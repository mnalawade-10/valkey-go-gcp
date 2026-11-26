#!/usr/bin/env bash
# Reruns the benchmark behind valkey-go's README claim on this machine (see harness_test.go).
# - Starts a standalone Valkey on :6379 and a 3-node cluster on :7001-7003 in Docker (host network).
# - Pins the servers and the benchmark to different CPU cores.
# - Runs SET and GET on the single server, and SET on the cluster, and writes raw go test output to results/.
# Override with env vars, e.g. a quick smoke test: COUNT=1 BENCHTIME=2000x OUT=smoke ./run.sh
set -euo pipefail
cd "$(dirname "$0")"

IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
COUNT=${COUNT:-10}
BENCHTIME=${BENCHTIME:-100000x}
CPU=${CPU:-10}                         # GOMAXPROCS, as on the README's 10-core MacBook
CLIENT_CPUS=${CLIENT_CPUS:-4-11,16-23}  # cores 4-11 (both hyperthreads) for the benchmark
OUT=${OUT:-results}
EXCLUDE=${EXCLUDE:-}                   # e.g. EXCLUDE=glide-v2.6.1
mkdir -p "$OUT"

cleanup() { docker rm -f rr-valkey rr-valkey-7001 rr-valkey-7002 rr-valkey-7003 >/dev/null 2>&1 || true; }
trap cleanup EXIT
cleanup

# Standalone on core 0; cluster nodes on cores 1, 2 and 3 (vCPU n and n+12 share a core).
docker run -d --name rr-valkey --network host --cpuset-cpus 0,12 "$IMAGE" \
  valkey-server --port 6379 --save "" --appendonly no >/dev/null
core=1
for port in 7001 7002 7003; do
  docker run -d --name "rr-valkey-$port" --network host --cpuset-cpus "$core,$((core + 12))" "$IMAGE" \
    valkey-server --port "$port" --save "" --appendonly no --cluster-enabled yes \
    --cluster-config-file "nodes-$port.conf" >/dev/null
  core=$((core + 1))
done
sleep 2
docker exec rr-valkey-7001 valkey-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-yes >/dev/null
for _ in $(seq 1 30); do
  if docker exec rr-valkey-7001 valkey-cli -p 7001 cluster info | grep -q 'cluster_state:ok'; then break; fi
  sleep 1
done

go test -c -o readme-repro.test .
{
  date -u
  go version
  docker exec rr-valkey valkey-server --version
  echo "image: $IMAGE"
  lscpu | grep 'Model name'
  echo "vCPUs: $(nproc); benchmark on CPUs $CLIENT_CPUS with GOMAXPROCS=$CPU"
  echo "benchtime=$BENCHTIME count=$COUNT exclude=$EXCLUDE"
  # The server's own ceiling on this machine, from the same CPUs as the benchmark:
  # 10 connections without pipelining, then with 16 commands per round trip.
  for p in 1 16; do
    docker run --rm --network host --cpuset-cpus "$CLIENT_CPUS" "$IMAGE" \
      valkey-benchmark -h 127.0.0.1 -p 6379 -t set,get -n 200000 -c 10 -P "$p" -q 2>&1 |
      tr '\r' '\n' | grep 'requests per second' | sed "s/^/valkey-benchmark -c 10 -P $p: /"
  done
} > "$OUT/env.txt"

run() { # $1 output name, $2 benchmark regex
  taskset -c "$CLIENT_CPUS" ./readme-repro.test -test.run '^$' -test.bench "$2" -test.benchmem \
    -test.benchtime "$BENCHTIME" -test.count "$COUNT" -test.cpu "$CPU" -test.timeout 3h \
    -exclude "$EXCLUDE" | tee "$OUT/$1.txt"
}
run single_set '^BenchmarkSingleClientSet$'
run single_get '^BenchmarkSingleClientGet$'
run cluster_set '^BenchmarkClusterClientSet$'
