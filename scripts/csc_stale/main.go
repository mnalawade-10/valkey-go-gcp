// Command csc_stale is the "stale-read meter".
//
// It measures client-side caching speed and correctness together. A writer keeps updating 100 hot
// keys with increasing version numbers; readers read them through each client's cache. A read is
// stale if it returns a version older than one whose write had already been acknowledged when the
// read started. For each client it reports how often that happens and for how long values stay stale,
// plus read speed and how many reads still reached the server.
//
// Clients (one at a time, single server):
//   - valkey-go without cache (baseline)
//   - valkey-go DoCache (server-assisted, opt-in tracking, 60 s client TTL)
//   - go-redis client-side cache (experimental, RESP3)
//   - GLIDE cache, time-based only (its default), 60 s TTL
//   - GLIDE cache, server-assisted (opt-in, CLIENT TRACKING ON BCAST), 60 s TTL
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"net"
	"os"
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
	addr       = flag.String("addr", "127.0.0.1:6379", "Valkey address (single server)")
	numKeys    = flag.Int("keys", 100, "hot keys")
	writeRate  = flag.Int("writes", 1000, "writes per second across all keys")
	readers    = flag.Int("readers", 8, "reader goroutines")
	duration   = flag.Duration("dur", 10*time.Second, "measurement time per client")
	clientsArg = flag.String("clients", "valkey-go (no cache),valkey-go DoCache,go-redis CSC,GLIDE cache (time-based),GLIDE cache (server-assisted)", "clients to test")
)

const (
	ringSize = 1 << 14
	cacheTTL = 60 * time.Second
)

type commit struct {
	ver atomic.Int64
	at  atomic.Int64 // unix nanos when the write of ver was acknowledged
}

type keyState struct {
	latest atomic.Int64
	ring   [ringSize]commit
}

func keyName(i int) string { return "csc:key:" + strconv.Itoa(i) }

type reader func(ctx context.Context, key string) (string, error)

type client struct {
	read  reader
	close func()
}

func newClient(name string) (client, error) {
	host, portS, _ := net.SplitHostPort(*addr)
	port, _ := strconv.Atoi(portS)
	switch name {
	case "valkey-go (no cache)", "valkey-go DoCache":
		c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{*addr}})
		if err != nil {
			return client{}, err
		}
		if name == "valkey-go (no cache)" {
			return client{read: func(ctx context.Context, k string) (string, error) {
				return c.Do(ctx, c.B().Get().Key(k).Build()).ToString()
			}, close: c.Close}, nil
		}
		return client{read: func(ctx context.Context, k string) (string, error) {
			return c.DoCache(ctx, c.B().Get().Key(k).Cache(), cacheTTL).ToString()
		}, close: c.Close}, nil
	case "go-redis CSC":
		c := goredis.NewClient(&goredis.Options{
			Addr:                  *addr,
			Protocol:              3,
			ClientSideCacheConfig: &goredis.ClientSideCacheConfig{MaxEntries: 10_000},
		})
		return client{read: func(ctx context.Context, k string) (string, error) {
			return c.Get(ctx, k).Result()
		}, close: func() { c.Close() }}, nil
	case "GLIDE cache (time-based)", "GLIDE cache (server-assisted)":
		cache, err := config.NewClientSideCache(10*1024, uint64(cacheTTL/time.Millisecond))
		if err != nil {
			return client{}, err
		}
		if name == "GLIDE cache (server-assisted)" {
			cache = cache.WithServerAssisted(true)
		}
		cache = cache.WithMetrics(true)
		c, err := glide.NewClient(config.NewClientConfiguration().
			WithAddress(&config.NodeAddress{Host: host, Port: port}).
			WithClientSideCache(cache))
		if err != nil {
			return client{}, err
		}
		return client{read: func(ctx context.Context, k string) (string, error) {
			r, err := c.Get(ctx, k)
			if err != nil {
				return "", err
			}
			if r.IsNil() {
				return "", fmt.Errorf("key %s missing", k)
			}
			return r.Value(), nil
		}, close: c.Close}, nil
	}
	return client{}, fmt.Errorf("unknown client %q", name)
}

func serverGets(ctx context.Context, admin valkey.Client) (int64, error) {
	s, err := admin.Do(ctx, admin.B().Info().Section("commandstats").Build()).ToString()
	if err != nil {
		return 0, err
	}
	for _, line := range strings.Split(s, "\r\n") {
		if rest, ok := strings.CutPrefix(line, "cmdstat_get:calls="); ok {
			n, _, _ := strings.Cut(rest, ",")
			return strconv.ParseInt(n, 10, 64)
		}
	}
	return 0, nil
}

type result struct {
	name                     string
	reads, stale, unknownAge int64
	serverGets               int64
	writes                   int64
	readP50, readP99         time.Duration
	ageP50, ageP99, ageMax   time.Duration
	elapsed                  time.Duration
}

func measure(ctx context.Context, admin valkey.Client, name string) (result, error) {
	r := result{name: name}
	if err := admin.Do(ctx, admin.B().Flushall().Sync().Build()).Error(); err != nil {
		return r, err
	}
	keys := make([]*keyState, *numKeys)
	for i := range keys {
		keys[i] = &keyState{}
		if err := admin.Do(ctx, admin.B().Set().Key(keyName(i)).Value("1").Build()).Error(); err != nil {
			return r, err
		}
		keys[i].ring[1].at.Store(time.Now().UnixNano())
		keys[i].ring[1].ver.Store(1)
		keys[i].latest.Store(1)
	}

	c, err := newClient(name)
	if err != nil {
		return r, err
	}
	defer c.close()
	// Fill the cache once.
	for i := range keys {
		if _, err := c.read(ctx, keyName(i)); err != nil {
			return r, fmt.Errorf("first read: %w", err)
		}
	}
	getsBefore, err := serverGets(ctx, admin)
	if err != nil {
		return r, err
	}

	writer, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{*addr}, DisableCache: true})
	if err != nil {
		return r, err
	}
	defer writer.Close()

	stop := make(chan struct{})
	var wg sync.WaitGroup
	var writes atomic.Int64
	var writeErr atomic.Value

	// Writer: round-robin over the keys at -writes per second.
	wg.Add(1)
	go func() {
		defer wg.Done()
		interval := time.Second / time.Duration(*writeRate)
		next := time.Now()
		i := 0
		for {
			select {
			case <-stop:
				return
			default:
			}
			if d := time.Until(next); d > 0 {
				time.Sleep(d)
			}
			next = next.Add(interval)
			ks := keys[i%len(keys)]
			ver := ks.latest.Load() + 1
			if err := writer.Do(ctx, writer.B().Set().Key(keyName(i%len(keys))).Value(strconv.FormatInt(ver, 10)).Build()).Error(); err != nil {
				writeErr.CompareAndSwap(nil, err)
				continue
			}
			e := &ks.ring[ver%ringSize]
			e.at.Store(time.Now().UnixNano())
			e.ver.Store(ver)
			ks.latest.Store(ver)
			writes.Add(1)
			i++
		}
	}()

	type rstats struct {
		reads, stale, unknown int64
		lat, ages             []time.Duration
		err                   error
	}
	stats := make([]rstats, *readers)
	start := time.Now()
	deadline := start.Add(*duration)
	for g := 0; g < *readers; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			st := &stats[g]
			rng := rand.New(rand.NewPCG(uint64(g), 42))
			for n := 0; ; n++ {
				now := time.Now()
				if now.After(deadline) {
					return
				}
				ki := rng.IntN(len(keys))
				ks := keys[ki]
				latest := ks.latest.Load()
				v, err := c.read(ctx, keyName(ki))
				took := time.Since(now)
				if err != nil {
					st.err = err
					return
				}
				ver, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					st.err = fmt.Errorf("bad value %q", v)
					return
				}
				st.reads++
				if n%16 == 0 {
					st.lat = append(st.lat, took)
				}
				if ver < latest {
					st.stale++
					e := &ks.ring[(ver+1)%ringSize]
					if e.ver.Load() == ver+1 {
						st.ages = append(st.ages, time.Duration(now.UnixNano()-e.at.Load()))
					} else {
						st.unknown++
					}
				}
			}
		}(g)
	}
	// Wait for the readers, then stop the writer.
	for time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	close(stop)
	wg.Wait()
	r.elapsed = time.Since(start)
	if e := writeErr.Load(); e != nil {
		return r, fmt.Errorf("writer: %v", e)
	}
	r.writes = writes.Load()

	var lat, ages []time.Duration
	for _, st := range stats {
		if st.err != nil {
			return r, fmt.Errorf("reader: %w", st.err)
		}
		r.reads += st.reads
		r.stale += st.stale
		r.unknownAge += st.unknown
		lat = append(lat, st.lat...)
		ages = append(ages, st.ages...)
	}
	getsAfter, err := serverGets(ctx, admin)
	if err != nil {
		return r, err
	}
	r.serverGets = getsAfter - getsBefore
	sort.Slice(lat, func(i, j int) bool { return lat[i] < lat[j] })
	sort.Slice(ages, func(i, j int) bool { return ages[i] < ages[j] })
	if len(lat) > 0 {
		r.readP50, r.readP99 = lat[len(lat)/2], lat[len(lat)*99/100]
	}
	if len(ages) > 0 {
		r.ageP50, r.ageP99, r.ageMax = ages[len(ages)/2], ages[len(ages)*99/100], ages[len(ages)-1]
	}
	return r, nil
}

func pct(a, b int64) string {
	if b == 0 {
		return "0%"
	}
	v := 100 * float64(a) / float64(b)
	if v > 0 && v < 0.01 {
		return "<0.01%"
	}
	return fmt.Sprintf("%.2f%%", v)
}

func main() {
	flag.Parse()
	ctx := context.Background()
	admin, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{*addr}, DisableCache: true, PipelineMultiplex: -1})
	if err != nil {
		fmt.Fprintln(os.Stderr, "admin:", err)
		os.Exit(1)
	}
	defer admin.Close()

	fmt.Printf("keys=%d writes=%d/s (each key every %s) readers=%d duration=%s cache TTL=%s\n\n",
		*numKeys, *writeRate, time.Duration(*numKeys)*time.Second/time.Duration(*writeRate), *readers, *duration, cacheTTL)
	fmt.Println("| Client | Reads/s | Read p50 / p99 | Reads that reached the server | Stale reads | Stale for: p50 / p99 / max |")
	fmt.Println("| :--- | :- | :- | :- | :- | :- |")
	for _, name := range strings.Split(*clientsArg, ",") {
		r, err := measure(ctx, admin, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
			os.Exit(1)
		}
		ages := "none"
		if r.stale > 0 {
			ages = fmt.Sprintf("%s / %s / %s", r.ageP50.Round(time.Microsecond), r.ageP99.Round(time.Microsecond), r.ageMax.Round(time.Millisecond))
			if r.unknownAge > 0 {
				ages += fmt.Sprintf(" (%d older than the ring)", r.unknownAge)
			}
		}
		fmt.Printf("| %s | %.0f | %s / %s | %s | %s | %s |\n", r.name, float64(r.reads)/r.elapsed.Seconds(),
			r.readP50.Round(100*time.Nanosecond), r.readP99.Round(100*time.Nanosecond),
			pct(r.serverGets, r.reads), pct(r.stale, r.reads), ages)
		fmt.Fprintf(os.Stderr, "%s: reads=%d stale=%d serverGETs=%d writes=%d\n", r.name, r.reads, r.stale, r.serverGets, r.writes)
	}
}
