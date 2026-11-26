// Command latency measures latency (layer L2), and runs the valkey-go off/on variants (layer L3).
//
// Two modes:
//
// Fixed rate (default): every client gets the same schedule, request i is due at start + i/rate. Many
// workers (-workers) take the next request, wait until it's due and send it, so a slow reply never holds
// back the next request ("open loop", which avoids coordinated omission). Three HDR histograms per run:
//   - latency: from the call to the reply (what the table reports);
//   - schedule lag: from the due time to the call; it shows whether the generator kept the rate.
//     On a VM, a sleeping goroutine wakes up late at low rates, so this lag is the load generator's
//     own delay, not the client's;
//   - from due time: lag plus latency, the strict wrk2-style number.
//
// Fixed concurrency (-conc 1,10,100,1000): that many goroutines each send their next request as soon as
// the previous reply arrives, as in GLIDE's benchmark. It reports requests per second and latency.
// With -batch N, each request is a pipeline of N GETs (not MGET), GLIDE's "batch of 10 GETs".
//
// Workload: GLIDE's standard mix, 80% GET of existing keys and 20% SET, random keys.
// With -elephants (fixed-rate mode), it is the "elephants and mice" test: small GETs at the fixed
// rate while other goroutines on the same client SET large values; it reports the small GETs' latency.
//
// Both modes also report the client process's CPU time per request (getrusage; it includes the load
// generator's own small, equal overhead) and its resident memory after the run (/proc/self/statm). Run one
// client per process when comparing memory, since a process doesn't give memory back right away.
//
// valkey-go variants for the off/on matrix are named valkey-go/<variant>; see newValkeyGo. Add +tls to any
// client name (valkey-go+tls, go-redis+tls, glide+tls, valkey-go/1-conn+tls) to connect with TLS to -tls-addr,
// trusting the CA in -tls-ca; servers.sh tls starts a server with a plain and a TLS port for this.
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"math/rand/v2"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/HdrHistogram/hdrhistogram-go"
	goredis "github.com/redis/go-redis/v9"
	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
	"github.com/valkey-io/valkey-glide/go/v2/pipeline"
	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeyaside"
	"github.com/valkey-io/valkey-go/valkeycompat"
	"github.com/valkey-io/valkey-go/valkeyhook"
	"github.com/valkey-io/valkey-go/valkeylock"
	"github.com/valkey-io/valkey-go/valkeyotel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	mode        = flag.String("mode", "standalone", "standalone or cluster")
	clientsArg  = flag.String("clients", "valkey-go,go-redis,go-redis-autopipeline,glide", "clients to test (valkey-go/<variant> for the off/on matrix)")
	ratesArg    = flag.String("rates", "2000,10000,20000", "request rates (ops/s), fixed-rate mode")
	concArg     = flag.String("conc", "", "fixed-concurrency mode: requests kept in flight, e.g. 1,10,100,1000")
	batch       = flag.Int("batch", 0, "with -conc: each request is a pipeline of this many GETs (0 = single commands, GLIDE's 80/20 mix)")
	dur         = flag.Duration("dur", 15*time.Second, "measured time per client and rate or concurrency")
	warm        = flag.Duration("warm", 2*time.Second, "unmeasured warm-up before each measurement")
	size        = flag.Int("size", 100, "value size in bytes")
	numKeys     = flag.Int("keys", 100000, "keys in the data set")
	workers     = flag.Int("workers", 512, "concurrent workers available to send requests, fixed-rate mode")
	elephants   = flag.Int("elephants", 0, "large SETs per second sent alongside the mice (0 = none), fixed-rate mode")
	eleSize     = flag.Int("elephant-size", 1<<20, "size of each large value in bytes")
	mice        = flag.Bool("mice", false, "send only small GETs (the baseline for -elephants); implied by -elephants")
	noPreload   = flag.Bool("no-preload", false, "skip loading the keys (they're already there)")
	preloadOnly = flag.Bool("preload-only", false, "load the keys and exit")
	benchOut    = flag.String("benchfmt", "", "also append each result to this file in Go benchmark format, for benchstat")
	addrArg     = flag.String("addr", "", "standalone server address(es), comma-separated (default 127.0.0.1:6379)")
	tlsAddrArg  = flag.String("tls-addr", "127.0.0.1:6390", "server address(es) for +tls clients, comma-separated")
	tlsCA       = flag.String("tls-ca", "tls/ca.crt", "CA certificate (PEM) that +tls clients trust")
	hot         = flag.Float64("hot", 0, "with -conc: reads only; this share of GETs goes to -hot-keys hot keys, the rest to keys that are never there (the client-side cache test)")
	hotKeys     = flag.Int("hot-keys", 1000, "how many hot keys -hot reads from")
	timeline    = flag.Bool("timeline", false, "with -conc: also print requests, errors and p99 for each second (the failover drill)")
	atCmd       = flag.String("at-cmd", "", "with -conc: shell command to run -at into the measured window, e.g. a failover")
	at          = flag.Duration("at", 5*time.Second, "when to run -at-cmd, from the start of the measured window")
)

const cacheTTL = 10 * time.Second // valkey-go/docache: how long an entry may stay in the client-side cache

// cacheReads and cacheHits count valkey-go/docache GETs and how many the client-side cache answered.
var cacheReads, cacheHits atomic.Int64

var (
	standaloneAddrs = []string{"127.0.0.1:6379"}
	clusterAddrs    = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}
)

type client struct {
	get, set func(ctx context.Context, key, val string) error
	batchGet func(ctx context.Context, keys []string) error
	close    func()
}

func addrs() []string {
	if *mode == "cluster" {
		return clusterAddrs
	}
	if *addrArg != "" {
		return strings.Split(*addrArg, ",")
	}
	return standaloneAddrs
}

// clientAddrs is where a client connects: the TLS address(es) for +tls clients.
func clientAddrs(useTLS bool) []string {
	if useTLS {
		return strings.Split(*tlsAddrArg, ",")
	}
	return addrs()
}

// caPEM reads the CA certificate that +tls clients trust.
func caPEM() ([]byte, error) { return os.ReadFile(*tlsCA) }

// tlsConfig is the Go TLS setup for valkey-go and go-redis: trust the test CA, check the server's name.
func tlsConfig() (*tls.Config, error) {
	pem, err := caPEM()
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(pem) {
		return nil, fmt.Errorf("no certificate in %s", *tlsCA)
	}
	host, _, _ := net.SplitHostPort(clientAddrs(true)[0])
	return &tls.Config{RootCAs: pool, ServerName: host, MinVersion: tls.VersionTLS12}, nil
}

func ignoreNil(err error) error {
	if err == nil || valkey.IsValkeyNil(err) || err == goredis.Nil {
		return nil
	}
	return err
}

// newValkeyGo builds valkey-go with one feature changed from its default. Variants (ClientOption fields
// from valkey.go at v1.0.80):
//
//	(none)        the defaults
//	no-autopipe   DisableAutoPipelining: each request takes a connection from the pool
//	always-pipe   AlwaysPipelining: pipeline even requests that aren't concurrent
//	flush-<d>      MaxFlushDelay <d>, e.g. flush-20us: wait up to <d> after each write to batch more commands
//	1-conn        PipelineMultiplex -1: one connection instead of the default 4 (standalone)
//	no-retry      DisableRetry: don't retry read-only commands after a connection error
//	docache       GETs through the client-side cache (DoCache, 10 s TTL) instead of Do
//	compat        the go-redis-style API from valkeycompat instead of the native one
//	redirect      Standalone.EnableRedirect: follow the server's REDIRECT to the new primary after a failover
//	aside         GETs through valkeyaside's cache-aside Get (client-side cache, loader only for missing keys);
//	              SETs stay plain, so the server tells the cache to drop the key
//	lock          every request is one valkeylock cycle (TryWithContext, then release) on a random name:
//	              lock cycles per second, a module benchmark rather than an off/on of the same workload
//	hook          wrapped by valkeyhook with a hook that only passes each call on
//	otel          valkeyotel.NewClient with the global OpenTelemetry providers (no-ops unless the app sets them)
//	otel-sdk      valkeyotel.NewClient with the OpenTelemetry SDK: every span sampled and batched to an
//	              exporter that drops it, metrics kept by a reader nobody reads; the cost without the network
func newValkeyGo(variant string, useTLS bool) (client, error) {
	opt := valkey.ClientOption{InitAddress: clientAddrs(useTLS)}
	if useTLS {
		cfg, err := tlsConfig()
		if err != nil {
			return client{}, err
		}
		opt.TLSConfig = cfg
	}
	useCache, useCompat, useHook, otelMode := false, false, false, ""
	switch variant {
	case "":
	case "no-autopipe":
		opt.DisableAutoPipelining = true
	case "always-pipe":
		opt.AlwaysPipelining = true
	case "1-conn":
		opt.PipelineMultiplex = -1
	case "no-retry":
		opt.DisableRetry = true
	case "docache":
		useCache = true
	case "compat":
		useCompat = true
	case "hook":
		useHook = true
	case "redirect":
		opt.Standalone.EnableRedirect = true
	case "aside", "lock":
		return newModule(variant, opt)
	case "otel", "otel-sdk":
		otelMode = variant
	default:
		d, err := time.ParseDuration(strings.TrimPrefix(variant, "flush-"))
		if !strings.HasPrefix(variant, "flush-") || err != nil {
			return client{}, fmt.Errorf("unknown valkey-go variant %q", variant)
		}
		opt.MaxFlushDelay = d
	}
	var c valkey.Client
	var err error
	switch otelMode {
	case "otel":
		c, err = valkeyotel.NewClient(opt)
	case "otel-sdk":
		tp := sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithBatcher(discardExporter{}))
		mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewManualReader()))
		c, err = valkeyotel.NewClient(opt, valkeyotel.WithTracerProvider(tp), valkeyotel.WithMeterProvider(mp))
	default:
		c, err = valkey.NewClient(opt)
	}
	if err != nil {
		return client{}, err
	}
	if useHook {
		c = valkeyhook.WithHook(c, passHook{})
	}
	if useCompat {
		ad := valkeycompat.NewAdapter(c)
		return client{
			get: func(ctx context.Context, k, _ string) error { return ignoreNil(ad.Get(ctx, k).Err()) },
			set: func(ctx context.Context, k, v string) error { return ad.Set(ctx, k, v, 0).Err() },
			batchGet: func(ctx context.Context, keys []string) error {
				_, err := ad.Pipelined(ctx, func(p valkeycompat.Pipeliner) error {
					for _, k := range keys {
						p.Get(ctx, k)
					}
					return nil
				})
				return ignoreNil(err)
			},
			close: c.Close,
		}, nil
	}
	return client{
		get: func(ctx context.Context, k, _ string) error {
			if useCache {
				r := c.DoCache(ctx, c.B().Get().Key(k).Cache(), cacheTTL)
				cacheReads.Add(1)
				if r.IsCacheHit() {
					cacheHits.Add(1)
				}
				return ignoreNil(r.Error())
			}
			return ignoreNil(c.Do(ctx, c.B().Get().Key(k).Build()).Error())
		},
		set: func(ctx context.Context, k, v string) error {
			return c.Do(ctx, c.B().Set().Key(k).Value(v).Build()).Error()
		},
		batchGet: func(ctx context.Context, keys []string) error {
			var rs []valkey.ValkeyResult
			if useCache {
				cmds := make([]valkey.CacheableTTL, len(keys))
				for i, k := range keys {
					cmds[i] = valkey.CT(c.B().Get().Key(k).Cache(), cacheTTL)
				}
				rs = c.DoMultiCache(ctx, cmds...)
			} else {
				cmds := make(valkey.Commands, len(keys))
				for i, k := range keys {
					cmds[i] = c.B().Get().Key(k).Build()
				}
				rs = c.DoMulti(ctx, cmds...)
			}
			for _, r := range rs {
				if err := ignoreNil(r.Error()); err != nil {
					return err
				}
			}
			return nil
		},
		close: c.Close,
	}, nil
}

// newModule builds the valkeyaside and valkeylock variants (see newValkeyGo).
func newModule(variant string, opt valkey.ClientOption) (client, error) {
	if variant == "aside" {
		ac, err := valkeyaside.NewClient(valkeyaside.ClientOption{ClientOption: opt})
		if err != nil {
			return client{}, err
		}
		small := strings.Repeat("v", *size)
		load := func(context.Context, string) (string, error) { return small, nil }
		raw := ac.Client()
		return client{
			get: func(ctx context.Context, k, _ string) error { _, err := ac.Get(ctx, cacheTTL, k, load); return err },
			set: func(ctx context.Context, k, v string) error {
				return raw.Do(ctx, raw.B().Set().Key(k).Value(v).Build()).Error()
			},
			batchGet: func(ctx context.Context, keys []string) error { // no multi-key Get: one after another
				for _, k := range keys {
					if _, err := ac.Get(ctx, cacheTTL, k, load); err != nil {
						return err
					}
				}
				return nil
			},
			close: ac.Close,
		}, nil
	}
	lk, err := valkeylock.NewLocker(valkeylock.LockerOption{ClientOption: opt})
	if err != nil {
		return client{}, err
	}
	cycle := func(ctx context.Context) error {
		_, release, err := lk.TryWithContext(ctx, "lock:"+strconv.Itoa(rand.IntN(10_000_000)))
		if err != nil {
			return err
		}
		release()
		return nil
	}
	return client{
		get: func(ctx context.Context, _, _ string) error { return cycle(ctx) },
		set: func(ctx context.Context, _, _ string) error { return cycle(ctx) },
		batchGet: func(ctx context.Context, keys []string) error {
			for range keys {
				if err := cycle(ctx); err != nil {
					return err
				}
			}
			return nil
		},
		close: lk.Close,
	}, nil
}

// passHook is a valkeyhook.Hook that only passes each call on: the cost of the hook layer itself.
type passHook struct{}

func (passHook) Do(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResult {
	return c.Do(ctx, cmd)
}
func (passHook) DoMulti(c valkey.Client, ctx context.Context, multi ...valkey.Completed) []valkey.ValkeyResult {
	return c.DoMulti(ctx, multi...)
}
func (passHook) DoCache(c valkey.Client, ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) valkey.ValkeyResult {
	return c.DoCache(ctx, cmd, ttl)
}
func (passHook) DoMultiCache(c valkey.Client, ctx context.Context, multi ...valkey.CacheableTTL) []valkey.ValkeyResult {
	return c.DoMultiCache(ctx, multi...)
}
func (passHook) Receive(c valkey.Client, ctx context.Context, sub valkey.Completed, fn func(valkey.PubSubMessage)) error {
	return c.Receive(ctx, sub, fn)
}
func (passHook) DoStream(c valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream {
	return c.DoStream(ctx, cmd)
}
func (passHook) DoMultiStream(c valkey.Client, ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream {
	return c.DoMultiStream(ctx, multi...)
}

// discardExporter drops every span: valkey-go/otel-sdk measures tracing without the network export.
type discardExporter struct{}

func (discardExporter) ExportSpans(context.Context, []sdktrace.ReadOnlySpan) error { return nil }
func (discardExporter) Shutdown(context.Context) error                             { return nil }

func newClient(name string) (client, error) {
	name, useTLS := strings.CutSuffix(name, "+tls")
	base, variant, _ := strings.Cut(name, "/")
	if base == "valkey-go" {
		return newValkeyGo(variant, useTLS)
	}
	if variant != "" {
		return client{}, fmt.Errorf("variants are only for valkey-go: %q", name)
	}
	a := clientAddrs(useTLS)
	switch name {
	case "go-redis", "go-redis-autopipeline":
		var tc *tls.Config
		if useTLS {
			var err error
			if tc, err = tlsConfig(); err != nil {
				return client{}, err
			}
		}
		var c goredis.UniversalClient
		if *mode == "cluster" {
			c = goredis.NewClusterClient(&goredis.ClusterOptions{Addrs: a, TLSConfig: tc})
		} else {
			c = goredis.NewClient(&goredis.Options{Addr: a[0], TLSConfig: tc})
		}
		var cmd goredis.Cmdable = c
		var ap *goredis.AutoPipeliner
		if name == "go-redis-autopipeline" {
			var err error
			switch cc := c.(type) {
			case *goredis.Client:
				ap, err = cc.AutoPipeline()
			case *goredis.ClusterClient:
				ap, err = cc.AutoPipeline()
			}
			if err != nil {
				c.Close()
				return client{}, err
			}
			cmd = ap
		}
		return client{
			get: func(ctx context.Context, k, _ string) error { return ignoreNil(cmd.Get(ctx, k).Err()) },
			set: func(ctx context.Context, k, v string) error { return cmd.Set(ctx, k, v, 0).Err() },
			// A batch is an explicit pipeline, so go-redis and go-redis-autopipeline behave the same here.
			batchGet: func(ctx context.Context, keys []string) error {
				_, err := c.Pipelined(ctx, func(p goredis.Pipeliner) error {
					for _, k := range keys {
						p.Get(ctx, k)
					}
					return nil
				})
				return ignoreNil(err)
			},
			close: func() {
				if ap != nil {
					ap.Close()
				}
				c.Close()
			},
		}, nil
	case "glide":
		nodes := make([]*config.NodeAddress, 0, len(a))
		for _, s := range a {
			host, port, _ := net.SplitHostPort(s)
			p, _ := strconv.Atoi(port)
			nodes = append(nodes, &config.NodeAddress{Host: host, Port: p})
		}
		if *mode == "cluster" {
			cfg := config.NewClusterClientConfiguration()
			for _, n := range nodes {
				cfg = cfg.WithAddress(n)
			}
			if useTLS {
				pem, err := caPEM()
				if err != nil {
					return client{}, err
				}
				cfg = cfg.WithUseTLS(true).WithAdvancedConfiguration(config.NewAdvancedClusterClientConfiguration().
					WithTlsConfiguration(config.NewTlsConfiguration().WithRootCertificates(pem)))
			}
			c, err := glide.NewClusterClient(cfg)
			if err != nil {
				return client{}, err
			}
			return client{
				get: func(ctx context.Context, k, _ string) error { _, err := c.Get(ctx, k); return err },
				set: func(ctx context.Context, k, v string) error { _, err := c.Set(ctx, k, v); return err },
				batchGet: func(ctx context.Context, keys []string) error {
					b := pipeline.NewClusterBatch(false) // non-atomic: a pipeline, split by slot
					for _, k := range keys {
						b.Get(k)
					}
					_, err := c.Exec(ctx, *b, true)
					return err
				},
				close: c.Close,
			}, nil
		}
		cfg := config.NewClientConfiguration().WithAddress(nodes[0])
		if useTLS {
			pem, err := caPEM()
			if err != nil {
				return client{}, err
			}
			cfg = cfg.WithUseTLS(true).WithAdvancedConfiguration(config.NewAdvancedClientConfiguration().
				WithTlsConfiguration(config.NewTlsConfiguration().WithRootCertificates(pem)))
		}
		c, err := glide.NewClient(cfg)
		if err != nil {
			return client{}, err
		}
		return client{
			get: func(ctx context.Context, k, _ string) error { _, err := c.Get(ctx, k); return err },
			set: func(ctx context.Context, k, v string) error { _, err := c.Set(ctx, k, v); return err },
			batchGet: func(ctx context.Context, keys []string) error {
				b := pipeline.NewStandaloneBatch(false)
				for _, k := range keys {
					b.Get(k)
				}
				_, err := c.Exec(ctx, *b, true)
				return err
			},
			close: c.Close,
		}, nil
	}
	return client{}, fmt.Errorf("unknown client %q", name)
}

func keyName(i int) string { return "lat:" + strconv.Itoa(i) }

func preload(ctx context.Context) error {
	c, err := valkey.NewClient(valkey.ClientOption{InitAddress: addrs(), DisableCache: true})
	if err != nil {
		return err
	}
	defer c.Close()
	val := strings.Repeat("v", *size)
	const chunk = 1000
	for i := 0; i < *numKeys; i += chunk {
		cmds := make(valkey.Commands, 0, chunk)
		for j := i; j < i+chunk && j < *numKeys; j++ {
			cmds = append(cmds, c.B().Set().Key(keyName(j)).Value(val).Build())
		}
		for _, r := range c.DoMulti(ctx, cmds...) {
			if err := r.Error(); err != nil {
				return err
			}
		}
	}
	return nil
}

// cpuTime is the CPU time this process has used so far (user plus system).
func cpuTime() time.Duration {
	var ru syscall.Rusage
	_ = syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	return time.Duration(ru.Utime.Nano() + ru.Stime.Nano())
}

// rssMB is this process's resident memory now, in MiB.
func rssMB() float64 {
	b, err := os.ReadFile("/proc/self/statm")
	if err != nil {
		return 0
	}
	f := strings.Fields(string(b))
	if len(f) < 2 {
		return 0
	}
	pages, _ := strconv.ParseFloat(f[1], 64)
	return pages * float64(os.Getpagesize()) / (1 << 20)
}

type result struct {
	achieved    float64
	svc         *hdrhistogram.Histogram // from the call to the reply
	lag         *hdrhistogram.Histogram // fixed rate: from the due time to the call (the load generator's own delay)
	fromDue     *hdrhistogram.Histogram // fixed rate: from the due time to the reply
	errors      int64
	elephantsOK int64
	cpuPerReq   time.Duration // client process CPU per request, warm-up included
	hitPct      float64       // valkey-go/docache: share of GETs the client-side cache answered
	rss         float64
}

func newHist() *hdrhistogram.Histogram { return hdrhistogram.New(1, 60_000_000, 3) } // microseconds, up to 60 s

func merge(hs []*hdrhistogram.Histogram) *hdrhistogram.Histogram {
	m := newHist()
	for _, h := range hs {
		m.Merge(h)
	}
	return m
}

// runRate drives one client at one rate: warm-up, then the measured window.
func runRate(ctx context.Context, c client, rate int) result {
	small := strings.Repeat("v", *size)
	big := strings.Repeat("E", *eleSize)
	interval := float64(time.Second) / float64(rate)
	total := int64(float64(rate) * (warm.Seconds() + dur.Seconds()))
	warmN := int64(float64(rate) * warm.Seconds())
	var next atomic.Int64
	var errs atomic.Int64
	svcH := make([]*hdrhistogram.Histogram, *workers)
	lagH := make([]*hdrhistogram.Histogram, *workers)
	dueH := make([]*hdrhistogram.Histogram, *workers)
	cpu0 := cpuTime()
	start := time.Now()
	var wg sync.WaitGroup
	for w := 0; w < *workers; w++ {
		svcH[w], lagH[w], dueH[w] = newHist(), newHist(), newHist()
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			rng := rand.New(rand.NewPCG(uint64(w), uint64(rate)))
			for {
				i := next.Add(1) - 1
				if i >= total {
					return
				}
				due := start.Add(time.Duration(float64(i) * interval))
				if d := time.Until(due); d > 0 {
					time.Sleep(d)
				}
				k := keyName(rng.IntN(*numKeys))
				sent := time.Now()
				var err error
				if !*mice && *elephants == 0 && rng.IntN(5) == 0 {
					err = c.set(ctx, k, small)
				} else {
					err = c.get(ctx, k, "")
				}
				done := time.Now()
				if err != nil {
					errs.Add(1)
					continue
				}
				if i >= warmN {
					_ = svcH[w].RecordValue(done.Sub(sent).Microseconds())
					_ = lagH[w].RecordValue(max(sent.Sub(due), 0).Microseconds())
					_ = dueH[w].RecordValue(done.Sub(due).Microseconds())
				}
			}
		}(w)
	}
	// Elephants: large SETs at a fixed rate on the same client, on their own goroutines.
	var eleOK atomic.Int64
	eleDone := make(chan struct{})
	var ewg sync.WaitGroup
	if *elephants > 0 {
		eInterval := time.Second / time.Duration(*elephants)
		for e := 0; e < 8; e++ {
			ewg.Add(1)
			go func(e int) {
				defer ewg.Done()
				t := time.NewTicker(eInterval * 8)
				defer t.Stop()
				for n := 0; ; n++ {
					select {
					case <-eleDone:
						return
					case <-t.C:
						if err := c.set(ctx, "elephant:"+strconv.Itoa(e)+":"+strconv.Itoa(n%4), big); err == nil {
							eleOK.Add(1)
						}
					}
				}
			}(e)
		}
	}
	wg.Wait()
	close(eleDone)
	ewg.Wait()
	elapsed := time.Since(start)
	r := result{svc: merge(svcH), lag: merge(lagH), fromDue: merge(dueH), errors: errs.Load(), elephantsOK: eleOK.Load()}
	r.achieved = float64(r.svc.TotalCount()) / (elapsed.Seconds() - warm.Seconds())
	r.cpuPerReq = (cpuTime() - cpu0) / time.Duration(max(total, 1))
	r.rss = rssMB()
	return r
}

// secStats is one second of the -timeline: requests that succeeded, failed with READONLY (still talking
// to a node that is now a replica), or failed otherwise, and their latency.
type secStats struct {
	ok, readonly, other atomic.Int64
}

func classify(err error) string {
	if strings.Contains(err.Error(), "READONLY") {
		return "readonly"
	}
	return "other"
}

// runConc keeps n requests in flight on one client: warm-up, then the measured window.
func runConc(ctx context.Context, c client, n int) result {
	small := strings.Repeat("v", *size)
	var stop, measuring atomic.Bool
	var errs, all atomic.Int64
	hs := make([]*hdrhistogram.Histogram, n)
	secs := int(dur.Seconds()) + 2
	var tl []secStats
	var tlH [][]*hdrhistogram.Histogram // [worker][second]
	var t0 atomic.Pointer[time.Time]
	var firstErr sync.Map // error class -> first message, for the timeline
	if *timeline {
		tl = make([]secStats, secs)
		tlH = make([][]*hdrhistogram.Histogram, n)
		for w := range tlH {
			tlH[w] = make([]*hdrhistogram.Histogram, secs)
			for i := range tlH[w] {
				tlH[w][i] = newHist()
			}
		}
	}
	cpu0 := cpuTime()
	var wg sync.WaitGroup
	for w := 0; w < n; w++ {
		hs[w] = newHist()
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			rng := rand.New(rand.NewPCG(uint64(w), uint64(n)))
			keys := make([]string, *batch)
			for !stop.Load() {
				var err error
				sent := time.Now()
				if *hot > 0 { // reads only: a hot key, or a key that is never there
					k := "cold:" + strconv.Itoa(rng.IntN(10_000_000))
					if rng.Float64() < *hot {
						k = keyName(rng.IntN(*hotKeys))
					}
					err = c.get(ctx, k, "")
				} else if *batch > 0 {
					for j := range keys {
						keys[j] = keyName(rng.IntN(*numKeys))
					}
					err = c.batchGet(ctx, keys)
				} else {
					k := keyName(rng.IntN(*numKeys))
					if rng.IntN(5) == 0 {
						err = c.set(ctx, k, small)
					} else {
						err = c.get(ctx, k, "")
					}
				}
				d := time.Since(sent)
				all.Add(1)
				if start := t0.Load(); *timeline && start != nil && measuring.Load() {
					if i := int(time.Since(*start) / time.Second); i < secs {
						switch {
						case err == nil:
							tl[i].ok.Add(1)
							_ = tlH[w][i].RecordValue(d.Microseconds())
						case classify(err) == "readonly":
							tl[i].readonly.Add(1)
							firstErr.LoadOrStore("readonly", err.Error())
						default:
							tl[i].other.Add(1)
							firstErr.LoadOrStore("other", err.Error())
						}
					}
				}
				if err != nil {
					errs.Add(1)
					continue
				}
				if measuring.Load() {
					_ = hs[w].RecordValue(d.Microseconds())
				}
			}
		}(w)
	}
	time.Sleep(*warm)
	cacheReads.Store(0)
	cacheHits.Store(0)
	now := time.Now()
	t0.Store(&now)
	measuring.Store(true)
	startMeasure := time.Now()
	if *atCmd != "" {
		go func() {
			time.Sleep(*at)
			out, err := exec.Command("sh", "-c", *atCmd).CombinedOutput()
			fmt.Fprintf(os.Stderr, "at %s ran %q: %s (err=%v)\n", *at, *atCmd, strings.TrimSpace(string(out)), err)
		}()
	}
	time.Sleep(*dur)
	measuring.Store(false)
	elapsed := time.Since(startMeasure)
	stop.Store(true)
	wg.Wait()
	r := result{svc: merge(hs), errors: errs.Load()}
	if *timeline {
		fmt.Printf("\nTimeline (second of the measured window):\n\n| Second | OK | READONLY errors | Other errors | p99 (ms) |\n| :- | :- | :- | :- | :- |\n")
		for i := 0; i < secs-1; i++ {
			m := newHist()
			for w := range tlH {
				m.Merge(tlH[w][i])
			}
			fmt.Printf("| %d | %d | %d | %d | %s |\n", i, tl[i].ok.Load(), tl[i].readonly.Load(), tl[i].other.Load(), ms(m.ValueAtQuantile(99)))
		}
		firstErr.Range(func(k, v any) bool {
			fmt.Printf("\nFirst %s error: %v\n", k, v)
			return true
		})
		fmt.Println()
	}
	r.achieved = float64(r.svc.TotalCount()) / elapsed.Seconds()
	if n := cacheReads.Load(); n > 0 {
		r.hitPct = 100 * float64(cacheHits.Load()) / float64(n)
	}
	r.cpuPerReq = (cpuTime() - cpu0) / time.Duration(max(all.Load(), 1))
	r.rss = rssMB()
	return r
}

// load names the workload in benchmark lines.
func load() string {
	switch {
	case *concArg != "" && *hot > 0:
		return fmt.Sprintf("hot%g", *hot*100)
	case *concArg != "" && *batch > 0:
		return fmt.Sprintf("batch%d", *batch)
	case *elephants > 0:
		return "elephants"
	case *mice:
		return "mice"
	}
	return "mix"
}

// benchFile appends results in Go benchmark format, so benchstat can compare repeated runs. The client
// isn't part of the name: each client or variant writes its own file, and benchstat compares the files.
type benchFile struct{ f *os.File }

func openBench() benchFile {
	if *benchOut == "" {
		return benchFile{}
	}
	f, err := os.OpenFile(*benchOut, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "benchfmt:", err)
		os.Exit(1)
	}
	return benchFile{f}
}

func (b benchFile) write(name string, r result) {
	if b.f == nil {
		return
	}
	h := r.svc
	fmt.Fprintf(b.f, "BenchmarkLatency/mode=%s/size=%d/%s 1 %.0f req/s %d p50-us %d p99-us %d p999-us %.2f cpu-us/req %.1f rss-MiB %.1f cache-hit-pct %d errors\n",
		*mode, *size, name, r.achieved, h.ValueAtQuantile(50), h.ValueAtQuantile(99), h.ValueAtQuantile(99.9),
		float64(r.cpuPerReq.Nanoseconds())/1000, r.rss, r.hitPct, r.errors)
}

func (b benchFile) Close() {
	if b.f != nil {
		b.f.Close()
	}
}

func ms(us int64) string { return fmt.Sprintf("%.3f", float64(us)/1000) }

func ints(s, what string) []int {
	var out []int
	for _, f := range strings.Split(s, ",") {
		v, err := strconv.Atoi(strings.TrimSpace(f))
		if err != nil || v <= 0 {
			fmt.Fprintf(os.Stderr, "bad %s: %q\n", what, f)
			os.Exit(1)
		}
		out = append(out, v)
	}
	return out
}

func main() {
	flag.Parse()
	ctx := context.Background()
	if !*noPreload {
		if err := preload(ctx); err != nil {
			fmt.Fprintln(os.Stderr, "preload:", err)
			os.Exit(1)
		}
	}
	if *preloadOnly {
		return
	}
	bench := openBench()
	defer bench.Close()
	conc := *concArg != ""
	what := "GLIDE mix (80% GET, 20% SET)"
	switch {
	case conc && *hot > 0:
		what = fmt.Sprintf("GETs only, %g%% of them to %d hot keys (the cache hit test)", *hot*100, *hotKeys)
	case conc && *batch > 0:
		what = fmt.Sprintf("batches of %d GETs, each sent as one pipeline", *batch)
	case *elephants > 0:
		what = fmt.Sprintf("small GETs, plus %d SETs/s of %d KB on the same client", *elephants, *eleSize/1024)
	case *mice:
		what = "small GETs only"
	}
	if conc {
		fmt.Printf("mode=%s fixed concurrency, workload=%s value=%d B keys=%d warm-up=%s measured=%s per level\n",
			*mode, what, *size, *numKeys, *warm, *dur)
		fmt.Println("Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.")
		fmt.Println()
		fmt.Println("| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |")
		fmt.Println("| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |")
	} else {
		fmt.Printf("mode=%s fixed rate, workload=%s value=%d B keys=%d workers=%d warm-up=%s measured=%s per rate\n",
			*mode, what, *size, *numKeys, *workers, *warm, *dur)
		fmt.Println("Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).")
		fmt.Println()
		fmt.Println("| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |")
		fmt.Println("| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |")
	}
	for _, name := range strings.Split(*clientsArg, ",") {
		c, err := newClient(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
			os.Exit(1)
		}
		if conc {
			for _, n := range ints(*concArg, "concurrency") {
				r := runConc(ctx, c, n)
				h := r.svc
				fmt.Printf("| %s | %d | %.0f | %s | %s | %s | %s | %s | %.1f | %.0f | %d |\n", name, n, r.achieved,
					ms(h.ValueAtQuantile(50)), ms(h.ValueAtQuantile(90)), ms(h.ValueAtQuantile(99)),
					ms(h.ValueAtQuantile(99.9)), ms(h.Max()), float64(r.cpuPerReq.Nanoseconds())/1000, r.rss, r.errors)
				bench.write(fmt.Sprintf("load=%s/inflight=%d", load(), n), r)
				if r.hitPct > 0 {
					fmt.Fprintf(os.Stderr, "%s at %d in flight: client-side cache answered %.1f%% of GETs\n", name, n, r.hitPct)
				}
			}
		} else {
			for _, rate := range ints(*ratesArg, "rate") {
				r := runRate(ctx, c, rate)
				h := r.svc
				fmt.Printf("| %s | %d | %.0f | %s | %s | %s | %s | %s | %s | %s | %.1f | %.0f | %d |\n", name, rate, r.achieved,
					ms(h.ValueAtQuantile(50)), ms(h.ValueAtQuantile(90)), ms(h.ValueAtQuantile(99)),
					ms(h.ValueAtQuantile(99.9)), ms(h.Max()), ms(r.lag.ValueAtQuantile(99)), ms(r.fromDue.ValueAtQuantile(99)),
					float64(r.cpuPerReq.Nanoseconds())/1000, r.rss, r.errors)
				bench.write(fmt.Sprintf("load=%s/rate=%d", load(), rate), r)
				if *elephants > 0 {
					fmt.Fprintf(os.Stderr, "%s at %d/s: %d large SETs completed\n", name, rate, r.elephantsOK)
				}
			}
		}
		c.close()
	}
}
