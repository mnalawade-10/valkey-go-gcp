package readmerepro

import (
	"context"
	"flag"
	"net"
	"runtime"
	"strconv"
	"testing"
	"time"

	goredisold "github.com/go-redis/redis/v9" // go-redis v9.0.0-rc.2, the version in the README chart
	goredis "github.com/redis/go-redis/v9"    // go-redis v9.23.0, the latest release
	"github.com/rueian/rueidis"               // rueidis v0.0.90, valkey-go's ancestor in the README chart
	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
	"github.com/valkey-io/valkey-glide/go/v2/models"
	"github.com/valkey-io/valkey-go" // valkey-go v1.0.80, the latest release
)

var (
	ctx            = context.Background()
	singleAddress  = "127.0.0.1:6379"
	clusterAddress = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}
	// mux sets ClientOption.PipelineMultiplex for rueidis and valkey-go (connections = 2^mux; -1 means one).
	// 0 keeps the library default: 4 connections per node on a single server, 1 on a cluster.
	mux = flag.Int("mux", 0, "PipelineMultiplex for rueidis and valkey-go (0 = library default)")
)

// Library labels used in the lib= part of the benchmark names.
const (
	libRueidisOld   = "rueidis-v0.0.90"
	libGoRedisOld   = "goredis-v9.0.0-rc.2"
	libValkeyGo     = "valkeygo-v1.0.80"
	libGoRedis      = "goredis-v9.23.0"
	libGoRedisAuto  = "goredis-v9.23.0-autopipeline"
	libGlide        = "glide-v2.6.1"
	cscSuffix       = "-csc"
	cscTTL          = 10 * time.Second
	glideCmdTimeout = 5 * time.Second
)

type glideCmd interface {
	Set(ctx context.Context, key string, value string) (string, error)
	Get(ctx context.Context, key string) (models.Result[string], error)
	FlushAll(ctx context.Context) (string, error)
	Close()
}

func glideClient(addrs []string) (glideCmd, error) {
	nodes := make([]*config.NodeAddress, 0, len(addrs))
	for _, a := range addrs {
		host, port, err := net.SplitHostPort(a)
		if err != nil {
			return nil, err
		}
		p, err := strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, &config.NodeAddress{Host: host, Port: p})
	}
	if len(nodes) == 1 {
		return glide.NewClient(config.NewClientConfiguration().WithAddress(nodes[0]).WithRequestTimeout(glideCmdTimeout))
	}
	cfg := config.NewClusterClientConfiguration().WithRequestTimeout(glideCmdTimeout)
	for _, n := range nodes {
		cfg = cfg.WithAddress(n)
	}
	return glide.NewClusterClient(cfg)
}

// builders returns the clients to compare. get selects GET (true) or SET (false);
// csc adds the client-side caching variants (single server, GET only);
// poolSize is go-redis's PoolSize, set to the largest goroutine count as in the README harness.
func builders(addrs []string, get, csc bool, poolSize int) []TargetBuilder {
	cluster := len(addrs) > 1
	prepare := func(bench Bench, set func(k, v string) error) error {
		if !get {
			return nil
		}
		seen := make(map[string]struct{}, len(bench.Keys))
		for _, k := range bench.Keys {
			if _, ok := seen[k]; ok {
				continue
			}
			seen[k] = struct{}{}
			if err := set(k, bench.Val); err != nil {
				return err
			}
		}
		return nil
	}

	rueidisOld := func(name string, cache bool) TargetBuilder {
		return TargetBuilder{Name: name, Make: func(bench Bench) (Target, error) {
			client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: addrs, PipelineMultiplex: *mux})
			if err != nil {
				return Target{}, err
			}
			if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
				return Target{}, err
			}
			if err := prepare(bench, func(k, v string) error {
				return client.Do(ctx, client.B().Set().Key(k).Value(v).Build()).Error()
			}); err != nil {
				return Target{}, err
			}
			do := func(keys []string, value string) error {
				return client.Do(ctx, client.B().Set().Key(pick(keys)).Value(value).Build()).Error()
			}
			if get && cache {
				do = func(keys []string, _ string) error {
					return client.DoCache(ctx, client.B().Get().Key(pick(keys)).Cache(), cscTTL).Error()
				}
			} else if get {
				do = func(keys []string, _ string) error {
					return client.Do(ctx, client.B().Get().Key(pick(keys)).Build()).Error()
				}
			}
			return Target{Close: client.Close, Do: do}, nil
		}}
	}

	valkeyGo := func(name string, cache bool) TargetBuilder {
		return TargetBuilder{Name: name, Make: func(bench Bench) (Target, error) {
			client, err := valkey.NewClient(valkey.ClientOption{InitAddress: addrs, PipelineMultiplex: *mux})
			if err != nil {
				return Target{}, err
			}
			if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
				return Target{}, err
			}
			if err := prepare(bench, func(k, v string) error {
				return client.Do(ctx, client.B().Set().Key(k).Value(v).Build()).Error()
			}); err != nil {
				return Target{}, err
			}
			do := func(keys []string, value string) error {
				return client.Do(ctx, client.B().Set().Key(pick(keys)).Value(value).Build()).Error()
			}
			if get && cache {
				do = func(keys []string, _ string) error {
					return client.DoCache(ctx, client.B().Get().Key(pick(keys)).Cache(), cscTTL).Error()
				}
			} else if get {
				do = func(keys []string, _ string) error {
					return client.Do(ctx, client.B().Get().Key(pick(keys)).Build()).Error()
				}
			}
			return Target{Close: client.Close, Do: do}, nil
		}}
	}

	goRedisOld := TargetBuilder{Name: libGoRedisOld, Make: func(bench Bench) (Target, error) {
		var client goredisold.UniversalClient
		if cluster {
			client = goredisold.NewClusterClient(&goredisold.ClusterOptions{Addrs: addrs, PoolSize: poolSize})
		} else {
			client = goredisold.NewClient(&goredisold.Options{Addr: addrs[0], PoolSize: poolSize})
		}
		if err := client.FlushAll(ctx).Err(); err != nil {
			return Target{}, err
		}
		if err := prepare(bench, func(k, v string) error { return client.Set(ctx, k, v, 0).Err() }); err != nil {
			return Target{}, err
		}
		do := func(keys []string, value string) error { return client.Set(ctx, pick(keys), value, 0).Err() }
		if get {
			do = func(keys []string, _ string) error { return client.Get(ctx, pick(keys)).Err() }
		}
		return Target{Close: func() { client.Close() }, Do: do}, nil
	}}

	// goRedis covers the plain client, the experimental AutoPipeline (blocking face,
	// default options) and the experimental client-side cache (standalone, RESP3).
	goRedis := func(name string, auto, cache bool) TargetBuilder {
		return TargetBuilder{Name: name, Make: func(bench Bench) (Target, error) {
			var client goredis.UniversalClient
			var ap *goredis.AutoPipeliner
			var err error
			if cluster {
				cc := goredis.NewClusterClient(&goredis.ClusterOptions{Addrs: addrs, PoolSize: poolSize})
				client = cc
				if auto {
					ap, err = cc.AutoPipeline()
				}
			} else {
				opts := &goredis.Options{Addr: addrs[0], PoolSize: poolSize}
				if cache {
					opts.Protocol = 3
					opts.ClientSideCacheConfig = &goredis.ClientSideCacheConfig{MaxEntries: 10_000}
				}
				c := goredis.NewClient(opts)
				client = c
				if auto {
					ap, err = c.AutoPipeline()
				}
			}
			if err != nil {
				client.Close()
				return Target{}, err
			}
			if err := client.FlushAll(ctx).Err(); err != nil {
				return Target{}, err
			}
			if err := prepare(bench, func(k, v string) error { return client.Set(ctx, k, v, 0).Err() }); err != nil {
				return Target{}, err
			}
			closeAll := func() {
				if ap != nil {
					ap.Close()
				}
				client.Close()
			}
			if ap != nil {
				do := func(keys []string, value string) error { return ap.Set(ctx, pick(keys), value, 0).Err() }
				if get {
					do = func(keys []string, _ string) error { return ap.Get(ctx, pick(keys)).Err() }
				}
				return Target{Close: closeAll, Do: do}, nil
			}
			do := func(keys []string, value string) error { return client.Set(ctx, pick(keys), value, 0).Err() }
			if get {
				do = func(keys []string, _ string) error { return client.Get(ctx, pick(keys)).Err() }
			}
			return Target{Close: closeAll, Do: do}, nil
		}}
	}

	glideBuilder := TargetBuilder{Name: libGlide, Make: func(bench Bench) (Target, error) {
		client, err := glideClient(addrs)
		if err != nil {
			return Target{}, err
		}
		if _, err := client.FlushAll(ctx); err != nil {
			return Target{}, err
		}
		if err := prepare(bench, func(k, v string) error { _, err := client.Set(ctx, k, v); return err }); err != nil {
			return Target{}, err
		}
		do := func(keys []string, value string) error { _, err := client.Set(ctx, pick(keys), value); return err }
		if get {
			do = func(keys []string, _ string) error { _, err := client.Get(ctx, pick(keys)); return err }
		}
		return Target{Close: client.Close, Do: do}, nil
	}}

	list := []TargetBuilder{
		rueidisOld(libRueidisOld, false),
		goRedisOld,
		valkeyGo(libValkeyGo, false),
		goRedis(libGoRedis, false, false),
		goRedis(libGoRedisAuto, true, false),
		glideBuilder,
	}
	if csc {
		list = append(list,
			rueidisOld(libRueidisOld+cscSuffix, true),
			valkeyGo(libValkeyGo+cscSuffix, true),
			goRedis(libGoRedis+cscSuffix, false, true),
		)
	}
	return list
}

// The README's single-server benchmarks: parallelism 1, 8 and 64, one 16-byte key,
// values of 64, 256 and 1,024 bytes, go-redis pool = 64 x NumCPU.
func BenchmarkSingleClientSet(b *testing.B) {
	ncpu := runtime.NumCPU()
	RunBenchmark(b, compose([]int{1, 8, 64}, []int{16}, []int{64, 256, 1024}, 1,
		builders([]string{singleAddress}, false, false, 64*ncpu)))
}

func BenchmarkSingleClientGet(b *testing.B) {
	ncpu := runtime.NumCPU()
	RunBenchmark(b, compose([]int{1, 8, 64}, []int{16}, []int{64, 256, 1024}, 1,
		builders([]string{singleAddress}, true, true, 64*ncpu)))
}

// Not in the README: the single-server SET with 1,024 random keys instead of one,
// to see how the clients spread real traffic over their connections.
func BenchmarkSingleClientSetManyKeys(b *testing.B) {
	ncpu := runtime.NumCPU()
	RunBenchmark(b, compose([]int{1, 8, 64}, []int{16}, []int{64, 256, 1024}, 1024,
		builders([]string{singleAddress}, false, false, 64*ncpu)))
}

// The README's cluster benchmark: parallelism 4, 16 and 96, 1,024 random 16-byte keys,
// values of 64, 256 and 1,024 bytes, go-redis pool = 96 x NumCPU.
func BenchmarkClusterClientSet(b *testing.B) {
	ncpu := runtime.NumCPU()
	RunBenchmark(b, compose([]int{4, 16, 96}, []int{16}, []int{64, 256, 1024}, 1024,
		builders(clusterAddress, false, false, 96*ncpu)))
}
