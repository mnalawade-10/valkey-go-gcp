# valkey-go, go-redis and GLIDE: benchmark scripts and raw results (Nov 2026)

These are the scripts and the raw results behind the benchmark comparison in valkey-go's README ("Benchmark Comparison with go-redis and GLIDE"). Everything was measured on one machine on Nov 23, 2026.

## Contents

| Path | What it is |
| :--- | :--- |
| `scripts/` | Every script and test tool used for the run. `run_all.sh` runs everything; `summarize_official.py` prints the tables |
| `results/official-20261123/` | The raw output of the run, its log (`run_all-official-20261123.log`) and every table (`official_summary.md`) |
| `results/glide-check-20261126/` | A check that the comparison is fair to GLIDE: GLIDE's own benchmark tool, run on the same machine, measures the same GLIDE throughput, and GLIDE's limit is inside its own client |
| `SHA256SUMS` | A checksum of every file |

The README's throughput table comes from the L1 part (`results/official-20261123/readme_repro/results_l1_official-20261123/`). Its latency table comes from the fixed-rate part (`results/official-20261123/latency/results/rates-standalone-100B-official-20261123/`).

## The machine and software

| Item | Value |
| :--- | :--- |
| Machine | A 24-vCPU VM: AMD EPYC 7B12, 12 cores × 2 threads, Linux 7.1.6 |
| CPU split | Single server on CPUs 0,12; cluster nodes on 1,13 / 2,14 / 3,15; the benchmark processes on 4–11,16–23 (`docker --cpuset-cpus`, `taskset`) |
| Server | `valkey/valkey:9.0-alpine`, Valkey 9.0.8 (`valkey/valkey@sha256:139d21c97b1a956e42e29b1e4ede4752a6ff52d1a0db9ef0ff4954a570f23976`), host networking, no persistence |
| Clients | valkey-go v1.0.80, go-redis v9.23.0, GLIDE for Go v2.6.1 (pinned in each tool's `go.mod`) |
| Tools | Go 1.27.3, Docker 29.8.0, memtier_benchmark 2.6.1 (`redislabs/memtier_benchmark:2.6.1`), benchstat from golang.org/x/perf v0.0.0-20261117185222-4febdb082b2c |

**Client settings:** default settings everywhere, except that in the README benchmark (L1) GLIDE's request timeout is 5 seconds, so slow replies don't count as errors.

## One-time setup

```bash
docker pull valkey/valkey:9.0-alpine
docker pull redislabs/memtier_benchmark:2.6.1
# benchstat compares repeated runs; the scripts take it from PATH (or from $BENCHSTAT):
go install golang.org/x/perf/cmd/benchstat@v0.0.0-20261117185222-4febdb082b2c
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Run it

```bash
cd scripts
IMAGE=valkey/valkey@sha256:139d21c97b1a956e42e29b1e4ede4752a6ff52d1a0db9ef0ff4954a570f23976 GOTOOLCHAIN=go1.27.3 ./run_all.sh
python3 summarize_official.py official-<UTC date of the run> > summary.md
```

The whole run takes about 6 hours; `PARTS="l1 rates" ./run_all.sh` runs only the parts behind the README tables. Keep the machine otherwise idle while it runs.

| Part | What it runs | Raw output |
| :--- | :--- | :--- |
| l0 | `memtier_l0/run.sh`: the server's own capacity, 3 runs of 60 s per setting | `memtier_l0/results-<tag>-run1..3` |
| sweep | `readme_repro/sweep.sh`: valkey-go releases v1.0.40 to v1.0.80, 10 runs each | `readme_repro/results-<tag>` |
| limiter | `limiter_mem/run.sh`: rate-limiter memory, 1,000,000 identifiers | `limiter_mem/results-<tag>` |
| l1 | The README benchmark (rueidis-benchmark harness) with today's clients, 10 runs of 50,000 commands per case | `readme_repro/results_l1_<tag>` |
| matrix | `latency/run_matrix.sh`: GLIDE's workload (80% GET, 20% SET) at 1–1,000 requests in flight, batches of 10, and valkey-go's settings turned off and on; 10 runs of 5 s | `latency/results/matrix-standalone-100B-<tag>` |
| rates | The same tool at fixed request rates (2,000, 10,000 and 20,000 per second), and with 1 MB writes mixed in | `latency/results/rates-standalone-100B-<tag>` |
| kb4 | 4 KB values | `latency/results/matrix-standalone-4000B-<tag>` |
| cache | valkey-go's client-side cache at 50%, 90% and 99% hits | `latency/results/matrix-cache-standalone-100B-<tag>` |
| cluster | A 3-node cluster | `latency/results/matrix-cluster-100B-<tag>` |
| tls | TLS off vs on against the same server | `latency/results/matrix-tls-standalone-100B-<tag>` |
| stale | Stale reads from each client's client-side cache, 3 runs | `csc_stale/results-<tag>` |
| fleet | Startup time, and connections and server memory per app instance, 5 runs | `fleet/results-<tag>` |
| failover | A local failover drill, 3 runs | `latency/results/failover-drill-<tag>-run1..3` |

## Will you get the same numbers?

You should get the same picture, but not the same absolute numbers:
- **The same:** which client is faster where, and the results that don't depend on the machine, such as connections per app instance and memory per rate-limiter identifier.
- **Different:** absolute throughput and latency. On the same kind of machine, expect most results within about 5–10%. On a different machine every number moves. go-redis without auto-pipelining stays near the server's limit for unpipelined commands, so ratios involving it change the most.

To get as close as possible:
1. **Use a 24-vCPU machine where CPUs n and n+12 are the two threads of one core.** `cat /sys/devices/system/cpu/cpu0/topology/thread_siblings_list` must print `0,12`. The scripts assume this layout to keep the server and the clients on separate cores.
2. **Pin the server image and Go version** as in the run command above.
3. **Compare with benchstat, file by file,** against `results/official-20261123`:
   ```bash
   benchstat results/official-20261123/latency/results/matrix-standalone-100B-official-20261123/valkey-go.txt \
             scripts/latency/results/matrix-standalone-100B-<your tag>/valkey-go.txt
   ```
   Lines marked `~` show no statistically significant difference.

## Notes

- The raw files are as recorded, except that local paths in them were replaced with `<scripts>`, and the summary's title was renamed. For publication, the scripts' comments were made general, and benchstat is now taken from `PATH`; the measurement code is the code that produced these results.
- During the run, an unrelated low-priority job ran for about a minute on the CPUs of one cluster node, while the fleet test's cluster part was running. That part was rerun afterwards, and its numbers matched within run-to-run noise (`fleet/results-official-20261123/fleet_cluster_recheck.txt`).
- The valkeylock variant reported failed lock attempts at 1,000 requests in flight in later rounds. That's still being checked; it doesn't affect the client comparison.
