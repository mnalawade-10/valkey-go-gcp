#!/usr/bin/env bash
# The full benchmark run. Every layer runs one after another, so no two tests
# share the machine. About 6.5 hours. Each part writes to its tool's results directory with $TAG in the
# name; the log records when each part started and finished.
#
#   ./run_all.sh                          # everything
#   PARTS="matrix cache" ./run_all.sh     # chosen parts
#
# Parts: l0 (memtier), sweep (release sweep), limiter (GCRA memory), l1 (README benchmark with today's
# clients), matrix (clients and valkey-go variants), rates (fixed rates, elephants and mice), kb4 (4 KB
# values), cache (fixed hit rates), cluster, tls, stale (stale-read meter), fleet (cold start, footprint),
# failover (drill).
#
# Don't edit this file, run_matrix.sh or the tools while it runs: bash reads scripts as it goes, and
# run_matrix.sh builds the latency tool at the start of each matrix. Avoid other heavy work on the machine.
set -uo pipefail
cd "$(dirname "$0")"
TAG=${TAG:-official-$(date -u +%Y%m%d)}
PARTS=${PARTS:-"l0 sweep limiter l1 matrix rates kb4 cache cluster tls stale fleet failover"}
COUNT=${COUNT:-10}
CPUS=4-11,16-23
LOG=$PWD/run_all-$TAG.log
log() { echo "$(date -u +%FT%TZ) $*" | tee -a "$LOG"; }
want() { [[ " $PARTS " == *" $1 "* ]]; }
filter() { grep -v glide_logger; }
matrix() { (cd latency && env COUNT="$COUNT" DUR=5s WARM=1s "$@" ./run_matrix.sh) >> "$LOG" 2>&1; }
VARIANTS_ALL=valkey-go/no-autopipe,valkey-go/always-pipe,valkey-go/flush-20us,valkey-go/flush-100us,valkey-go/1-conn,valkey-go/no-retry,valkey-go/docache,valkey-go/compat,valkey-go/hook,valkey-go/otel,valkey-go/otel-sdk,valkey-go/aside,valkey-go/lock

log "start $TAG, parts: $PARTS, count $COUNT"

# These tools start their own server on :6379, so the shared servers go down first.
./servers.sh down
if want l0; then
  log "L0 memtier: 3 runs, 60 s per setting"
  for r in 1 2 3; do (cd memtier_l0 && TEST_TIME=60 OUT="results-$TAG-run$r" ./run.sh) >> "$LOG" 2>&1; done
  log "L0 done"
fi
if want sweep; then
  log "release sweep: $COUNT runs per release"
  (cd readme_repro && COUNT="$COUNT" OUT="results-$TAG" ./sweep.sh &&
    python3 summarize_sweep.py "results-$TAG/sweep_set.txt" > "results-$TAG/sweep_summary.md") >> "$LOG" 2>&1
  log "sweep done"
fi
if want limiter; then
  log "rate-limiter memory: Valkey 9.0, 1,000,000 identifiers"
  (cd limiter_mem && OUT="results-$TAG" ./run.sh) >> "$LOG" 2>&1
  log "limiter done"
fi

./servers.sh up || { log "servers.sh up failed"; exit 1; }
if want l1; then
  log "L1: README benchmark with today's clients, three key shapes, $COUNT runs"
  (cd readme_repro && mkdir -p "results_l1_$TAG" && go test -c -o readme-repro.test . &&
    taskset -c "$CPUS" ./readme-repro.test -test.run '^$' -test.bench '^BenchmarkL1$' -test.benchmem \
      -test.benchtime 50000x -test.count "$COUNT" -test.cpu 10 -test.timeout 6h \
      -exclude rueidis-v0.0.90,goredis-v9.0.0-rc.2 2> "results_l1_$TAG/stderr.txt" | filter > "results_l1_$TAG/l1.txt"
    rm -f readme-repro.test
    "${BENCHSTAT:-benchstat}" -col /lib "results_l1_$TAG/l1.txt" > "results_l1_$TAG/benchstat.txt" 2>&1) >> "$LOG" 2>&1
  log "L1 done"
fi
if want matrix; then
  log "matrix: 4 clients and 13 valkey-go variants, fixed concurrency and batches of 10"
  matrix OUT="results/matrix-standalone-100B-$TAG" VARIANTS="$VARIANTS_ALL"
  log "matrix done"
fi
if want rates; then
  log "fixed rates, elephants and mice: 4 clients"
  matrix OUT="results/rates-standalone-100B-$TAG" VARIANTS= LOADS="-rates 2000,10000,20000|-rates 5000 -mice|-rates 5000 -elephants 50"
  log "rates done"
fi
if want kb4; then
  log "4 KB values: 4 clients, fixed concurrency and fixed rates"
  matrix OUT="results/matrix-standalone-4000B-$TAG" SIZE=4000 VARIANTS= LOADS="-conc 1,10,100,1000|-rates 2000,10000"
  log "kb4 done"
fi
if want cache; then
  log "client-side cache at 50%, 90% and 99% hits"
  matrix OUT="results/matrix-cache-standalone-100B-$TAG" CLIENTS=valkey-go VARIANTS=valkey-go/docache \
    LOADS="-conc 1,10,100 -hot 0.5|-conc 1,10,100 -hot 0.9|-conc 1,10,100 -hot 0.99"
  log "cache done"
fi
if want cluster; then
  log "3-node cluster: 4 clients, fixed concurrency and batches of 10 (baseline for #208 and #209)"
  matrix OUT="results/matrix-cluster-100B-$TAG" MODE=cluster VARIANTS=
  log "cluster done"
fi
if want tls; then
  log "TLS off vs on: 4 clients, and valkey-go with a 100 us write delay"
  ./servers.sh tls >> "$LOG" 2>&1
  matrix OUT="results/matrix-tls-standalone-100B-$TAG" \
    VARIANTS=valkey-go+tls,go-redis+tls,go-redis-autopipeline+tls,glide+tls,valkey-go/flush-100us,valkey-go/flush-100us+tls \
    EXTRA="-addr 127.0.0.1:6391 -tls-addr 127.0.0.1:6390 -tls-ca $PWD/tls/ca.crt"
  docker rm -f rr-tls >/dev/null 2>&1
  log "tls done"
fi
if want stale; then
  log "stale-read meter: 3 runs"
  (cd csc_stale && mkdir -p "results-$TAG" && go build -o csc_stale . &&
    for r in 1 2 3; do taskset -c "$CPUS" ./csc_stale -dur 10s 2>> "results-$TAG/stderr.txt" | filter > "results-$TAG/run$r.txt"; done
    rm -f csc_stale) >> "$LOG" 2>&1
  log "stale done"
fi
if want fleet; then
  log "cold start and fleet footprint: median of 5 runs, standalone and cluster"
  (cd fleet && mkdir -p "results-$TAG" && go build -o fleet . &&
    taskset -c "$CPUS" ./fleet -mode standalone -reps 5 2> "results-$TAG/fleet_standalone_stderr.txt" | filter > "results-$TAG/fleet_standalone.txt" &&
    taskset -c "$CPUS" ./fleet -mode cluster -reps 5 2> "results-$TAG/fleet_cluster_stderr.txt" | filter > "results-$TAG/fleet_cluster.txt"
    rm -f fleet) >> "$LOG" 2>&1
  log "fleet done"
fi
if want failover; then
  log "failover drill: 3 runs per client"
  for r in 1 2 3; do (cd latency && DUR=20s AT=5s OUT="results/failover-drill-$TAG-run$r" ./failover_drill.sh) >> "$LOG" 2>&1; done
  log "failover done"
fi
log "all done"
