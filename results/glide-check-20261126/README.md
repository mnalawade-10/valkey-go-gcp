# GLIDE fairness check (Nov 26, 2026)

This checks that the comparison doesn't put GLIDE at a disadvantage. It ran on the same machine and server as the main run.

**1. Same settings as GLIDE's own benchmark.** Our latency tool creates GLIDE clients with the defaults (address only, plus TLS where tested), exactly like GLIDE's benchmark tool ([go/benchmarks](https://github.com/valkey-io/valkey-glide/tree/e2e42d128a7ffc1c0ba0f38db2afdc1bfa635813/go/benchmarks)). The README benchmark (L1) only raises GLIDE's request timeout to 5 s, so slow replies don't count as errors.

**2. GLIDE's own benchmark tool measures the same.** We built GLIDE's tool from that commit, unchanged except that its `go.mod` uses the released GLIDE v2.6.1 and go-redis v9.23.0 instead of a local checkout, and ran it twice per client, with its workload (80% GET, 20% SET, 100-byte values) and the same CPUs:

```bash
taskset -c 4-11,16-23 ./benchmarks -minimal -dataSize "[100]" -concurrentTasks "[1 10 100 1000]" -clients glide -host 127.0.0.1 -port 6379
taskset -c 4-11,16-23 ./benchmarks -minimal -dataSize "[100]" -concurrentTasks "[1 10 100 1000]" -clients go-redis -host 127.0.0.1 -port 6379
```

| Requests in flight | GLIDE, GLIDE's tool | GLIDE, our tool (Nov 23 run) | go-redis, GLIDE's tool | go-redis, our tool (Nov 23 run) |
| :--- | ---: | ---: | ---: | ---: |
| 100 | 35.0k, 35.6k | 36k | 54.2k, 53.9k | 53k |
| 1,000 | 34.6k, 34.6k | 35k | 51.7k, 52.7k | 53k |

Under load the two tools agree, and GLIDE reaches about two thirds of go-redis's throughput in both. At 1 and 10 in flight, GLIDE's tool reports lower numbers than ours, because its `-minimal` mode runs only 1,000 requests per task, with no warm-up. Files: `own_glide_*.log`, `own_go-redis_*.log`, and `our_tool_same_day.txt` (our tool again on Nov 26: GLIDE 33.7k and 32.3k, go-redis 49.6k and 49.8k at 100 and 1,000 in flight).

**3. Where GLIDE's limit comes from.** While GLIDE kept 100 requests in flight, one GLIDE thread ran at about a full core. The Go threads were mostly idle, and 16 vCPUs were available (`glide_thread_cpu.txt`). So GLIDE v2.6.1 stops at about 35,000 requests per second because of its own client, not the machine or our tools.
