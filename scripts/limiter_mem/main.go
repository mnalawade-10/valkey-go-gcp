// Command limiter_mem is the "rate-limiter memory meter".
//
// It checks the note that GCRA uses about half the memory of fixed window. The valkeylimiter
// README says GCRA keeps one key per identifier where fixed window keeps two, "cutting keyspace
// overhead in half"; this measures what that means in server memory. For each limiter it:
//  1. flushes the server and records used_memory;
//  2. calls Allow once for N distinct identifiers (each call creates that identifier's state);
//  3. records used_memory and DBSIZE again, giving bytes and keys per identifier;
//  4. times Allow on existing identifiers: one goroutine (latency) and many goroutines (throughput).
//
// Limiters: valkeylimiter fixed window (its default), valkeylimiter GCRA, valkeycompatrate (GCRA,
// drop-in for go-redis's redis_rate) and go-redis's redis_rate v10 itself.
// The limit is 10 per 24 hours, so no state expires while it's being measured.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	redisrate "github.com/go-redis/redis_rate/v10"
	goredis "github.com/redis/go-redis/v9"
	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeycompatrate"
	"github.com/valkey-io/valkey-go/valkeylimiter"
)

var (
	addr     = flag.String("addr", "127.0.0.1:6379", "Valkey address")
	n        = flag.Int("n", 1_000_000, "identifiers to create per limiter")
	fillConc = flag.Int("conc", 64, "goroutines used to create identifiers and for the throughput test")
	latCalls = flag.Int("lat", 20_000, "sequential Allow calls for the latency test")
	tputDur  = flag.Duration("tput", 5*time.Second, "duration of the throughput test")
	only     = flag.String("only", "", "comma-separated limiter names to run (default: all)")
)

const (
	limit  = 10
	window = 24 * time.Hour
)

type limiter struct {
	name  string
	keys  func(id string) []string // the keys one identifier's state lives in
	allow func(ctx context.Context, id string) (bool, error)
	close func()
}

func newLimiters() ([]limiter, error) {
	var out []limiter

	for _, alg := range []struct {
		name   string
		alg    valkeylimiter.Algorithm
		prefix string
		keys   func(id string) []string
	}{
		{"valkeylimiter fixed window", valkeylimiter.AlgorithmFixedWindow, valkeylimiter.PlaceholderPrefix, func(id string) []string {
			k := valkeylimiter.PlaceholderPrefix + ":{" + id + "}"
			return []string{k, k + ":ex"}
		}},
		{"valkeylimiter GCRA", valkeylimiter.AlgorithmGCRA, valkeylimiter.PlaceholderPrefix, func(id string) []string {
			return []string{valkeylimiter.PlaceholderPrefix + ":gcra:{" + id + "}"}
		}},
		// The same two algorithms with a short KeyPrefix, to separate the algorithm from key-name length.
		{"valkeylimiter fixed window (prefix rl)", valkeylimiter.AlgorithmFixedWindow, "rl", func(id string) []string {
			k := "rl:{" + id + "}"
			return []string{k, k + ":ex"}
		}},
		{"valkeylimiter GCRA (prefix rl)", valkeylimiter.AlgorithmGCRA, "rl", func(id string) []string {
			return []string{"rl:gcra:{" + id + "}"}
		}},
	} {
		rl, err := valkeylimiter.NewRateLimiter(valkeylimiter.RateLimiterOption{
			ClientOption: valkey.ClientOption{InitAddress: []string{*addr}},
			KeyPrefix:    alg.prefix,
			Limit:        limit,
			Window:       window,
			Algorithm:    alg.alg,
		})
		if err != nil {
			return nil, err
		}
		out = append(out, limiter{
			name: alg.name,
			keys: alg.keys,
			allow: func(ctx context.Context, id string) (bool, error) {
				res, err := rl.Allow(ctx, id)
				return res.Allowed, err
			},
			close: rl.Close,
		})
	}

	vc, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{*addr}})
	if err != nil {
		return nil, err
	}
	compat := valkeycompatrate.NewLimiterFromClient(vc)
	compatLimit := valkeycompatrate.Limit{Rate: limit, Burst: limit, Period: window}
	out = append(out, limiter{
		name: "valkeycompatrate (GCRA)",
		keys: func(id string) []string { return []string{"rate:" + id} },
		allow: func(ctx context.Context, id string) (bool, error) {
			res, err := compat.Allow(ctx, id, compatLimit)
			if err != nil {
				return false, err
			}
			return res.Allowed > 0, nil
		},
		close: vc.Close,
	})

	rdb := goredis.NewClient(&goredis.Options{Addr: *addr, PoolSize: *fillConc})
	rr := redisrate.NewLimiter(rdb)
	rrLimit := redisrate.Limit{Rate: limit, Burst: limit, Period: window}
	out = append(out, limiter{
		name: "go-redis redis_rate v10 (GCRA)",
		keys: func(id string) []string { return []string{"rate:" + id} },
		allow: func(ctx context.Context, id string) (bool, error) {
			res, err := rr.Allow(ctx, id, rrLimit)
			if err != nil {
				return false, err
			}
			return res.Allowed > 0, nil
		},
		close: func() { rdb.Close() },
	})
	return out, nil
}

func info(ctx context.Context, c valkey.Client, section, field string) (int64, error) {
	s, err := c.Do(ctx, c.B().Info().Section(section).Build()).ToString()
	if err != nil {
		return 0, err
	}
	for _, line := range strings.Split(s, "\r\n") {
		if v, ok := strings.CutPrefix(line, field+":"); ok {
			return strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		}
	}
	return 0, fmt.Errorf("INFO %s has no %s", section, field)
}

func usedMemory(ctx context.Context, c valkey.Client) (int64, error) {
	// Let lazy frees and background work settle, then read.
	time.Sleep(500 * time.Millisecond)
	return info(ctx, c, "memory", "used_memory")
}

func dbsize(ctx context.Context, c valkey.Client) (int64, error) {
	return c.Do(ctx, c.B().Dbsize().Build()).AsInt64()
}

func memoryUsage(ctx context.Context, c valkey.Client, key string) (int64, error) {
	v, err := c.Do(ctx, c.B().MemoryUsage().Key(key).Samples(0).Build()).AsInt64()
	if valkey.IsValkeyNil(err) {
		return 0, fmt.Errorf("key %q not found", key)
	}
	return v, err
}

func id(i int) string { return "user:" + strconv.Itoa(i) }

// fill calls Allow once for identifiers 0..n-1 from conc goroutines.
func fill(ctx context.Context, l limiter, count, conc int) error {
	var next atomic.Int64
	var failed, denied atomic.Int64
	var firstErr atomic.Value
	var wg sync.WaitGroup
	for g := 0; g < conc; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				i := int(next.Add(1)) - 1
				if i >= count {
					return
				}
				ok, err := l.allow(ctx, id(i))
				if err != nil {
					failed.Add(1)
					firstErr.CompareAndSwap(nil, err)
					continue
				}
				if !ok {
					denied.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	if f := failed.Load(); f > 0 {
		return fmt.Errorf("%d calls failed, first error: %v", f, firstErr.Load())
	}
	if d := denied.Load(); d > 0 {
		return fmt.Errorf("%d first calls were denied", d)
	}
	return nil
}

func percentile(sorted []time.Duration, p float64) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	i := int(float64(len(sorted)-1) * p)
	return sorted[i]
}

type result struct {
	name            string
	keysPerID       float64
	bytesPerID      float64
	sampleBytes     int64
	p50, p99        time.Duration
	opsPerSecond    float64
	identifiers     int
	usedBefore      int64
	usedAfter       int64
	keysBefore      int64
	keysAfter       int64
	sampleKeyBreaks string
}

func measure(ctx context.Context, admin valkey.Client, l limiter) (result, error) {
	r := result{name: l.name, identifiers: *n}
	if err := admin.Do(ctx, admin.B().Flushall().Sync().Build()).Error(); err != nil {
		return r, err
	}
	// One untimed round from as many goroutines as the fill uses: it loads the Lua script and
	// opens every connection the limiter will use, so connection buffers are the same in the
	// "before" and "after" readings. Then remove that state.
	var wg sync.WaitGroup
	warmErr := make(chan error, *fillConc)
	for g := 0; g < *fillConc; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				if _, err := l.allow(ctx, "warmup:"+strconv.Itoa(g)); err != nil {
					warmErr <- err
					return
				}
			}
		}(g)
	}
	wg.Wait()
	close(warmErr)
	if err := <-warmErr; err != nil {
		return r, fmt.Errorf("warm-up: %w", err)
	}
	for g := 0; g < *fillConc; g++ {
		for _, k := range l.keys("warmup:" + strconv.Itoa(g)) {
			if err := admin.Do(ctx, admin.B().Del().Key(k).Build()).Error(); err != nil {
				return r, err
			}
		}
	}
	var err error
	if r.usedBefore, err = usedMemory(ctx, admin); err != nil {
		return r, err
	}
	if r.keysBefore, err = dbsize(ctx, admin); err != nil {
		return r, err
	}

	if err := fill(ctx, l, *n, *fillConc); err != nil {
		return r, err
	}
	if r.usedAfter, err = usedMemory(ctx, admin); err != nil {
		return r, err
	}
	if r.keysAfter, err = dbsize(ctx, admin); err != nil {
		return r, err
	}
	r.bytesPerID = float64(r.usedAfter-r.usedBefore) / float64(*n)
	r.keysPerID = float64(r.keysAfter-r.keysBefore) / float64(*n)

	var parts []string
	for _, k := range l.keys(id(*n / 2)) {
		b, err := memoryUsage(ctx, admin, k)
		if err != nil {
			return r, err
		}
		val, _ := admin.Do(ctx, admin.B().Get().Key(k).Build()).ToString()
		ttl, _ := admin.Do(ctx, admin.B().Pttl().Key(k).Build()).AsInt64()
		enc, _ := admin.Do(ctx, admin.B().ObjectEncoding().Key(k).Build()).ToString()
		r.sampleBytes += b
		parts = append(parts, fmt.Sprintf("%s=%dB (value %q, %s, ttl %ds)", k, b, val, enc, ttl/1000))
	}
	r.sampleKeyBreaks = strings.Join(parts, "; ")

	// Latency: one goroutine, existing identifiers in random order.
	lat := make([]time.Duration, 0, *latCalls)
	for i := 0; i < *latCalls; i++ {
		target := id(rand.IntN(*n))
		start := time.Now()
		if _, err := l.allow(ctx, target); err != nil {
			return r, fmt.Errorf("latency test: %w", err)
		}
		lat = append(lat, time.Since(start))
	}
	sort.Slice(lat, func(i, j int) bool { return lat[i] < lat[j] })
	r.p50, r.p99 = percentile(lat, 0.50), percentile(lat, 0.99)

	// Throughput: many goroutines for a fixed time.
	var calls, failed atomic.Int64
	deadline := time.Now().Add(*tputDur)
	start := time.Now()
	for g := 0; g < *fillConc; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for time.Now().Before(deadline) {
				if _, err := l.allow(ctx, id(rand.IntN(*n))); err != nil {
					failed.Add(1)
					continue
				}
				calls.Add(1)
			}
		}()
	}
	wg.Wait()
	if f := failed.Load(); f > 0 {
		return r, fmt.Errorf("throughput test: %d calls failed", f)
	}
	r.opsPerSecond = float64(calls.Load()) / time.Since(start).Seconds()
	return r, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	admin, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{*addr}, DisableCache: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, "connect:", err)
		os.Exit(1)
	}
	defer admin.Close()
	version, _ := admin.Do(ctx, admin.B().Info().Section("server").Build()).ToString()
	for _, line := range strings.Split(version, "\r\n") {
		if strings.HasPrefix(line, "valkey_version:") || strings.HasPrefix(line, "redis_version:") {
			fmt.Println("server:", line)
		}
	}

	limiters, err := newLimiters()
	if err != nil {
		fmt.Fprintln(os.Stderr, "limiters:", err)
		os.Exit(1)
	}
	defer func() {
		for _, l := range limiters {
			l.close()
		}
	}()

	var results []result
	for _, l := range limiters {
		if *only != "" && !strings.Contains(","+*only+",", ","+l.name+",") {
			continue
		}
		fmt.Printf("measuring %s with %d identifiers...\n", l.name, *n)
		r, err := measure(ctx, admin, l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", l.name, err)
			os.Exit(1)
		}
		fmt.Printf("  used_memory %d -> %d, keys %d -> %d, sample: %s\n", r.usedBefore, r.usedAfter, r.keysBefore, r.keysAfter, r.sampleKeyBreaks)
		results = append(results, r)
	}
	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "no limiter matched -only")
		os.Exit(1)
	}

	base := results[0]
	fmt.Printf("\n| Limiter | Keys per identifier | Bytes per identifier (used_memory) | vs %s | MEMORY USAGE, one identifier | Allow p50 | Allow p99 | Allow/s (%d goroutines) |\n", base.name, *fillConc)
	fmt.Println("| :--- | :- | :- | :- | :- | :- | :- | :- |")
	for _, r := range results {
		fmt.Printf("| %s | %.2f | %.0f | %+.0f%% | %d B | %s | %s | %.0f |\n",
			r.name, r.keysPerID, r.bytesPerID, (r.bytesPerID/base.bytesPerID-1)*100, r.sampleBytes,
			r.p50.Round(time.Microsecond), r.p99.Round(time.Microsecond), r.opsPerSecond)
	}
	if err := admin.Do(ctx, admin.B().Flushall().Build()).Error(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Fprintln(os.Stderr, "final flush:", err)
	}
}
