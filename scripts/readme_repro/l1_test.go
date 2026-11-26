package readmerepro

import (
	"math/rand/v2"
	"testing"
)

// L1: the README benchmark with today's clients, on three key shapes:
// the README harness with today's clients out of the box, and three ways of choosing keys.
//   - one:     every command uses the same key (the README's choice)
//   - uniform: 1,024 keys, each equally likely
//   - zipf:    1,024 keys with a skewed popularity (Zipf s=1.1: the top key gets about 18% of commands)
// Run with -exclude=rueidis-v0.0.90,goredis-v9.0.0-rc.2 to keep only today's clients; go-redis gets its
// default pool here (PoolSize 0 = 10 x GOMAXPROCS).

// zipfKeys returns n samples over distinct keys with Zipf popularity, so a uniform pick from the
// result follows the Zipf distribution.
func zipfKeys(distinct []string, n int) []string {
	r := rand.New(rand.NewPCG(1, 2))
	z := rand.NewZipf(r, 1.1, 1, uint64(len(distinct)-1))
	out := make([]string, n)
	for i := range out {
		out[i] = distinct[z.Uint64()]
	}
	return out
}

func composeShape(shape string, parallelisms, valSizes []int, builders []TargetBuilder) []Bench {
	var benchmarks []Bench
	for _, p := range parallelisms {
		for _, v := range valSizes {
			var keys []string
			switch shape {
			case "one":
				keys = []string{gen(16)}
			case "uniform":
				keys = make([]string, 1024)
				for i := range keys {
					keys[i] = gen(16)
				}
			case "zipf":
				distinct := make([]string, 1024)
				for i := range distinct {
					distinct[i] = gen(16)
				}
				keys = zipfKeys(distinct, 1<<16)
			}
			val := gen(v)
			for _, builder := range builders {
				if excluded(builder.Name) {
					continue
				}
				benchmarks = append(benchmarks, Bench{Parallelism: p, Keys: keys, Val: val, TargetBuilder: builder})
			}
		}
	}
	return benchmarks
}

func BenchmarkL1(b *testing.B) {
	for _, topo := range []struct {
		name   string
		addrs  []string
		shapes []string
		pars   []int
	}{
		{"single", []string{singleAddress}, []string{"one", "uniform", "zipf"}, []int{1, 8, 64}},
		{"cluster", clusterAddress, []string{"uniform", "zipf"}, []int{4, 16, 96}},
	} {
		b.Run("topo="+topo.name, func(b *testing.B) {
			for _, op := range []string{"set", "get"} {
				b.Run("op="+op, func(b *testing.B) {
					for _, shape := range topo.shapes {
						b.Run("keys="+shape, func(b *testing.B) {
							RunBenchmark(b, composeShape(shape, topo.pars, []int{64},
								builders(topo.addrs, op == "get", false, 0)))
						})
					}
				})
			}
		})
	}
}
