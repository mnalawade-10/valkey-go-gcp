// Command fleet runs two tests:
//
//   - Cold start: K app "instances" (one client each) all start at the same moment; how long until
//     each gets its first reply.
//   - Fleet footprint: how many server connections and how much server memory each instance costs,
//     idle and under load. Connection counts don't depend on the machine.
//
// Each instance gets a client with its library's defaults for an app with -cpus vCPUs (the process
// runs with GOMAXPROCS=-cpus, which is what the libraries size their defaults from). Under load,
// every instance runs -load goroutines doing GETs on 64 keys spread over the cluster's slots.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	goredis "github.com/redis/go-redis/v9"
	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
	"github.com/valkey-io/valkey-go"
)

var (
	mode      = flag.String("mode", "standalone", "standalone or cluster")
	instances = flag.Int("instances", 50, "app instances to simulate")
	loadG     = flag.Int("load", 32, "goroutines per instance during the load phase")
	loadDur   = flag.Duration("dur", 3*time.Second, "duration of the load phase")
	cpus      = flag.Int("cpus", 4, "vCPUs of each simulated app instance (GOMAXPROCS)")
	libs      = flag.String("libs", "", "libraries to test (default: all for the mode)")
	reps      = flag.Int("reps", 3, "runs per library; the table shows the median of each number")
)

// median returns, field by field, the median of the runs.
func median(runs []result) result {
	pickD := func(f func(r result) time.Duration) time.Duration {
		v := make([]time.Duration, len(runs))
		for i, r := range runs {
			v[i] = f(r)
		}
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		return v[len(v)/2]
	}
	pickF := func(f func(r result) float64) float64 {
		v := make([]float64, len(runs))
		for i, r := range runs {
			v[i] = f(r)
		}
		sort.Float64s(v)
		return v[len(v)/2]
	}
	m := runs[0]
	m.ttfP50 = pickD(func(r result) time.Duration { return r.ttfP50 })
	m.ttfP99 = pickD(func(r result) time.Duration { return r.ttfP99 })
	m.ttfMax = pickD(func(r result) time.Duration { return r.ttfMax })
	m.coldTotal = pickD(func(r result) time.Duration { return r.coldTotal })
	m.idleConns = pickF(func(r result) float64 { return r.idleConns })
	m.busiestConns = pickF(func(r result) float64 { return r.busiestConns })
	m.totalConnsLoad = pickF(func(r result) float64 { return r.totalConnsLoad })
	m.peakMem = pickF(func(r result) float64 { return r.peakMem })
	m.opsPerSecond = pickF(func(r result) float64 { return r.opsPerSecond })
	var errs, ops int64
	for _, r := range runs {
		errs += r.errors
		ops += r.ops
	}
	m.errors, m.ops = errs, ops // totals over all runs
	return m
}

var (
	standaloneAddrs = []string{"127.0.0.1:6379"}
	clusterAddrs    = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}
	keys            = func() []string {
		k := make([]string, 64)
		for i := range k {
			k[i] = "fleet:" + strconv.Itoa(i)
		}
		return k
	}()
)

type instance struct {
	get   func(ctx context.Context, key string) error
	close func()
}

func addrs() []string {
	if *mode == "cluster" {
		return clusterAddrs
	}
	return standaloneAddrs
}

func newInstance(lib string) (instance, error) {
	a := addrs()
	switch lib {
	case "valkey-go", "valkey-go (ForceSingleClient)":
		opt := valkey.ClientOption{InitAddress: a}
		if lib != "valkey-go" {
			// Skips cluster auto-detection and the client-tracking handshake.
			opt.ForceSingleClient, opt.DisableCache = true, true
		}
		c, err := valkey.NewClient(opt)
		if err != nil {
			return instance{}, err
		}
		return instance{
			get:   func(ctx context.Context, k string) error { return c.Do(ctx, c.B().Get().Key(k).Build()).Error() },
			close: c.Close,
		}, nil
	case "go-redis", "go-redis-autopipeline":
		var c goredis.UniversalClient
		var cmd goredis.Cmdable
		var ap *goredis.AutoPipeliner
		var err error
		if *mode == "cluster" {
			cc := goredis.NewClusterClient(&goredis.ClusterOptions{Addrs: a})
			c, cmd = cc, cc
			if lib == "go-redis-autopipeline" {
				ap, err = cc.AutoPipeline()
			}
		} else {
			sc := goredis.NewClient(&goredis.Options{Addr: a[0]})
			c, cmd = sc, sc
			if lib == "go-redis-autopipeline" {
				ap, err = sc.AutoPipeline()
			}
		}
		if err != nil {
			c.Close()
			return instance{}, err
		}
		if ap != nil {
			cmd = ap
		}
		return instance{
			get: func(ctx context.Context, k string) error { return cmd.Get(ctx, k).Err() },
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
			c, err := glide.NewClusterClient(cfg)
			if err != nil {
				return instance{}, err
			}
			return instance{
				get:   func(ctx context.Context, k string) error { _, err := c.Get(ctx, k); return err },
				close: c.Close,
			}, nil
		}
		c, err := glide.NewClient(config.NewClientConfiguration().WithAddress(nodes[0]))
		if err != nil {
			return instance{}, err
		}
		return instance{
			get:   func(ctx context.Context, k string) error { _, err := c.Get(ctx, k); return err },
			close: c.Close,
		}, nil
	}
	return instance{}, fmt.Errorf("unknown library %q", lib)
}

// admin talks to every server node directly, one connection each, to read INFO.
type admin struct{ nodes []valkey.Client }

func newAdmin() (*admin, error) {
	var a admin
	for _, s := range addrs() {
		c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s}, ForceSingleClient: true, DisableCache: true, PipelineMultiplex: -1})
		if err != nil {
			return nil, err
		}
		a.nodes = append(a.nodes, c)
	}
	return &a, nil
}

// perNode returns an INFO field for every node.
func (a *admin) perNode(ctx context.Context, section, field string) ([]int64, error) {
	out := make([]int64, len(a.nodes))
	for i, c := range a.nodes {
		s, err := c.Do(ctx, c.B().Info().Section(section).Build()).ToString()
		if err != nil {
			return nil, err
		}
		found := false
		for _, line := range strings.Split(s, "\r\n") {
			if v, ok := strings.CutPrefix(line, field+":"); ok {
				if out[i], err = strconv.ParseInt(strings.TrimSpace(v), 10, 64); err != nil {
					return nil, err
				}
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("INFO %s has no %s", section, field)
		}
	}
	return out, nil
}

func sum(v []int64) (s int64) {
	for _, x := range v {
		s += x
	}
	return s
}

type result struct {
	lib                    string
	ttfP50, ttfP99, ttfMax time.Duration
	coldTotal              time.Duration
	idleConns              float64 // per instance, all nodes
	busiestConns           float64 // per instance, on the busiest node, under load
	totalConnsLoad         float64 // per instance, all nodes, under load
	peakMem                float64 // bytes per instance, all nodes, under load
	opsPerSecond           float64
	errors, ops            int64
}

func run(ctx context.Context, ad *admin, lib string, settled []int64) (result, error) {
	r := result{lib: lib}
	// Wait until connections from the previous run have closed.
	for deadline := time.Now().Add(10 * time.Second); time.Now().Before(deadline); time.Sleep(200 * time.Millisecond) {
		c, err := ad.perNode(ctx, "clients", "connected_clients")
		if err != nil {
			return r, err
		}
		if sum(c) <= sum(settled) {
			break
		}
	}
	baseConns, err := ad.perNode(ctx, "clients", "connected_clients")
	if err != nil {
		return r, err
	}
	baseMem, err := ad.perNode(ctx, "memory", "used_memory")
	if err != nil {
		return r, err
	}

	// Cold start: all instances start together.
	insts := make([]instance, *instances)
	ttf := make([]time.Duration, *instances)
	errs := make([]error, *instances)
	var wg sync.WaitGroup
	gate := make(chan struct{})
	for i := range insts {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-gate
			start := time.Now()
			in, err := newInstance(lib)
			if err == nil {
				err = in.get(ctx, keys[i%len(keys)])
			}
			ttf[i], insts[i], errs[i] = time.Since(start), in, err
		}(i)
	}
	t0 := time.Now()
	close(gate)
	wg.Wait()
	r.coldTotal = time.Since(t0)
	for i, e := range errs {
		if e != nil {
			return r, fmt.Errorf("instance %d: %w", i, e)
		}
	}
	defer func() {
		for _, in := range insts {
			if in.close != nil {
				in.close()
			}
		}
	}()
	sort.Slice(ttf, func(i, j int) bool { return ttf[i] < ttf[j] })
	r.ttfP50, r.ttfP99, r.ttfMax = ttf[len(ttf)/2], ttf[len(ttf)*99/100], ttf[len(ttf)-1]

	time.Sleep(time.Second)
	idle, err := ad.perNode(ctx, "clients", "connected_clients")
	if err != nil {
		return r, err
	}
	r.idleConns = float64(sum(idle)-sum(baseConns)) / float64(*instances)

	// Load: every instance runs -load goroutines doing GETs; sample the server meanwhile.
	var ops, failed atomic.Int64
	peakPerNode := make([]int64, len(ad.nodes))
	var peakTotal, peakMem int64
	var mu sync.Mutex
	stop := make(chan struct{})
	var samplers sync.WaitGroup
	samplers.Add(1)
	go func() {
		defer samplers.Done()
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-stop:
				return
			case <-t.C:
				c, err1 := ad.perNode(ctx, "clients", "connected_clients")
				m, err2 := ad.perNode(ctx, "memory", "used_memory")
				if err1 != nil || err2 != nil {
					continue
				}
				mu.Lock()
				for i := range c {
					if c[i] > peakPerNode[i] {
						peakPerNode[i] = c[i]
					}
				}
				peakTotal = max(peakTotal, sum(c))
				peakMem = max(peakMem, sum(m))
				mu.Unlock()
			}
		}
	}()
	deadline := time.Now().Add(*loadDur)
	start := time.Now()
	var load sync.WaitGroup
	for _, in := range insts {
		for g := 0; g < *loadG; g++ {
			load.Add(1)
			go func(in instance, g int) {
				defer load.Done()
				for n := g; time.Now().Before(deadline); n++ {
					if err := in.get(ctx, keys[n%len(keys)]); err != nil {
						failed.Add(1)
						continue
					}
					ops.Add(1)
				}
			}(in, g)
		}
	}
	load.Wait()
	elapsed := time.Since(start)
	close(stop)
	samplers.Wait()
	r.errors, r.ops = failed.Load(), ops.Load()
	busiest := 0.0
	for i := range peakPerNode {
		busiest = max(busiest, float64(peakPerNode[i]-baseConns[i]))
	}
	r.busiestConns = busiest / float64(*instances)
	r.totalConnsLoad = float64(peakTotal-sum(baseConns)) / float64(*instances)
	r.peakMem = float64(peakMem-sum(baseMem)) / float64(*instances)
	r.opsPerSecond = float64(r.ops) / elapsed.Seconds()
	return r, nil
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*cpus)
	ctx := context.Background()

	ad, err := newAdmin()
	if err != nil {
		fmt.Fprintln(os.Stderr, "admin:", err)
		os.Exit(1)
	}
	setter, err := valkey.NewClient(valkey.ClientOption{InitAddress: addrs(), DisableCache: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, "setter:", err)
		os.Exit(1)
	}
	for _, k := range keys {
		if err := setter.Do(ctx, setter.B().Set().Key(k).Value("v").Build()).Error(); err != nil {
			fmt.Fprintln(os.Stderr, "set:", err)
			os.Exit(1)
		}
	}
	setter.Close()

	list := *libs
	if list == "" {
		list = "valkey-go,go-redis,go-redis-autopipeline,glide"
		if *mode == "standalone" {
			list = "valkey-go,valkey-go (ForceSingleClient),go-redis,go-redis-autopipeline,glide"
		}
	}
	nodes := len(addrs())
	settled, err := ad.perNode(ctx, "clients", "connected_clients")
	if err != nil {
		fmt.Fprintln(os.Stderr, "connections:", err)
		os.Exit(1)
	}
	fmt.Printf("mode=%s nodes=%d instances=%d load=%d goroutines/instance for %s on %d keys, app vCPUs (GOMAXPROCS)=%d, median of %d runs\n\n",
		*mode, nodes, *instances, *loadG, *loadDur, len(keys), *cpus, *reps)
	fmt.Println("| Library | Connections per instance, idle | Connections per instance under load (busiest node) | Instances until 32,000 connections on a node | Server memory per instance, under load | First reply p50 / p99 / max | All instances up | Failed GETs under load | GETs/s, all instances |")
	fmt.Println("| :--- | :- | :- | :- | :- | :- | :- | :- | :- |")
	for _, lib := range strings.Split(list, ",") {
		var runs []result
		for i := 0; i < *reps; i++ {
			r, err := run(ctx, ad, lib, settled)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", lib, err)
				os.Exit(1)
			}
			runs = append(runs, r)
		}
		r := median(runs)
		limit := "n/a"
		if r.busiestConns > 0 {
			limit = fmt.Sprintf("%.0f", 32000/r.busiestConns)
		}
		failed := "0"
		if r.errors > 0 {
			failed = fmt.Sprintf("%d (%.3f%%)", r.errors, 100*float64(r.errors)/float64(r.errors+r.ops))
		}
		fmt.Printf("| %s | %.1f | %.1f | %s | %.0f KB | %s / %s / %s | %s | %s | %.0f |\n",
			lib, r.idleConns, r.busiestConns, limit, r.peakMem/1024,
			r.ttfP50.Round(100*time.Microsecond), r.ttfP99.Round(100*time.Microsecond), r.ttfMax.Round(100*time.Microsecond),
			r.coldTotal.Round(time.Millisecond), failed, r.opsPerSecond)
	}
}
