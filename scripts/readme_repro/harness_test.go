// Package readmerepro reruns the benchmark behind valkey-go's README claim
// ("~14x throughput over go-redis", rueian/rueidis-benchmark commit fd05043,
// Dec 17 2022) with the README-era clients and today's clients side by side,
// on one machine and one server.
//
// It keeps the README harness: the same parallelism levels (b.SetParallelism),
// key size, value sizes, one key for the single-server benchmarks, 1,024 random
// keys for the cluster benchmark, and go-redis's pool sized to the largest
// goroutine count. Two deliberate changes, both applied to every client:
//   - an untimed warm-up, so go-redis's connection dials (one per goroutine)
//     and any lazy setup are not billed to the timed region;
//   - sub-benchmark names use key=value parts (lib=, par=, val=), so benchstat
//     can put the clients in columns: benchstat -col /lib results.txt
//
// Each case also reports the client process's CPU time per command (cpu-ns/op, from getrusage over the
// timed region; it includes Go's benchmark loop, the same for every client).
package readmerepro

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/valyala/fastrand"
)

// exclude skips clients by name, e.g. -exclude=glide-v2.6.1 (comma-separated).
var exclude = flag.String("exclude", "", "comma-separated lib names to skip")

func excluded(name string) bool {
	for _, n := range strings.Split(*exclude, ",") {
		if strings.TrimSpace(n) == name {
			return true
		}
	}
	return false
}

// Bench is one benchmark case: a client (TargetBuilder) at one parallelism and value size.
type Bench struct {
	Parallelism   int
	Keys          []string
	Val           string
	TargetBuilder TargetBuilder
}

// Target is a ready client: Do runs one command, Close releases the client.
type Target struct {
	Close func()
	Do    func(keys []string, value string) error
}

// TargetBuilder creates a Target for a Bench.
type TargetBuilder struct {
	Name string
	Make func(bench Bench) (Target, error)
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(fastrand.Uint32n(26) + 'a'))
	}
	return sb.String()
}

func compose(parallelisms, keySizes, valSizes []int, numKeys int, builders []TargetBuilder) []Bench {
	benchmarks := make([]Bench, 0, len(parallelisms)*len(keySizes)*len(valSizes)*len(builders))
	for _, p := range parallelisms {
		for _, k := range keySizes {
			for _, v := range valSizes {
				keys := make([]string, numKeys)
				for i := range keys {
					keys[i] = gen(k)
				}
				val := gen(v)
				// Clients take turns inside each (parallelism, value) group, as in the README harness.
				for _, builder := range builders {
					if excluded(builder.Name) {
						continue
					}
					benchmarks = append(benchmarks, Bench{Parallelism: p, Keys: keys, Val: val, TargetBuilder: builder})
				}
			}
		}
	}
	return benchmarks
}

// warmUp runs a few untimed commands from as many goroutines as the timed
// region will use, so connection pools are filled before the timer starts.
func warmUp(b *testing.B, bench Bench, target Target) {
	goroutines := 1
	if bench.Parallelism > 0 {
		goroutines = bench.Parallelism * runtime.GOMAXPROCS(0)
	}
	var wg sync.WaitGroup
	var failed atomic.Int64
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				if err := target.Do(bench.Keys, bench.Val); err != nil {
					failed.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	if n := failed.Load(); n > 0 {
		b.Fatalf("%s warm-up: %d commands failed", bench.TargetBuilder.Name, n)
	}
}

// cpuNow is the CPU time this process has used so far (user plus system).
func cpuNow() time.Duration {
	var ru syscall.Rusage
	_ = syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	return time.Duration(ru.Utime.Nano() + ru.Stime.Nano())
}

// RunBenchmark runs every case as a sub-benchmark named lib=<client>/par=<p>/val=<bytes>.
func RunBenchmark(b *testing.B, benchmarks []Bench) {
	for _, bench := range benchmarks {
		bench := bench
		name := fmt.Sprintf("lib=%s/par=%d/val=%d", bench.TargetBuilder.Name, bench.Parallelism, len(bench.Val))
		b.Run(name, func(b *testing.B) {
			target, err := bench.TargetBuilder.Make(bench)
			if err != nil {
				b.Fatalf("%s setup fail: %v", bench.TargetBuilder.Name, err)
			}
			defer target.Close()
			warmUp(b, bench, target)
			var failed atomic.Int64
			cpu0 := cpuNow()
			if bench.Parallelism == 0 {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if err := target.Do(bench.Keys, bench.Val); err != nil {
						failed.Add(1)
					}
				}
			} else {
				b.SetParallelism(bench.Parallelism)
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						if err := target.Do(bench.Keys, bench.Val); err != nil {
							failed.Add(1)
						}
					}
				})
			}
			b.StopTimer()
			cpu := cpuNow() - cpu0
			if sec := b.Elapsed().Seconds(); sec > 0 {
				b.ReportMetric(float64(b.N)/sec, "ops/s")
			}
			b.ReportMetric(float64(cpu.Nanoseconds())/float64(b.N), "cpu-ns/op")
			if n := failed.Load(); n > 0 {
				b.Errorf("%s: %d of %d commands failed", bench.TargetBuilder.Name, n, b.N)
			}
		})
	}
}

// pick returns the only key, or a random one when there are several (cluster benchmark).
func pick(keys []string) string {
	if len(keys) == 1 {
		return keys[0]
	}
	return keys[fastrand.Uint32n(uint32(len(keys)))]
}
