#!/usr/bin/env bash
# L0 engine baseline with memtier_benchmark.
# Starts a standalone Valkey (:6379) and a 3-node cluster (:7001-7003) in Docker, pinned to their own
# CPU cores, then runs memtier from a container pinned to the client cores:
# GLIDE's mix (20% SET, 80% GET), 100 B and 4,000 B values, without and with pipelining (1 and 16).
# Writes memtier's JSON per run to results/ and prints a summary table (summarize.py).
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
MEMTIER=${MEMTIER:-redislabs/memtier_benchmark:2.6.1}
CLIENT_CPUS=${CLIENT_CPUS:-4-11,16-23}
TEST_TIME=${TEST_TIME:-30}
KEYS=${KEYS:-100000}
OUT=${OUT:-results}
mkdir -p "$OUT"

cleanup() { docker rm -f rr-l0 rr-l0-7001 rr-l0-7002 rr-l0-7003 >/dev/null 2>&1 || true; }
trap cleanup EXIT
cleanup
docker run -d --name rr-l0 --network host --cpuset-cpus 0,12 "$IMAGE" \
  valkey-server --port 6379 --save "" --appendonly no >/dev/null
core=1
for port in 7001 7002 7003; do
  docker run -d --name "rr-l0-$port" --network host --cpuset-cpus "$core,$((core + 12))" "$IMAGE" \
    valkey-server --port "$port" --save "" --appendonly no --cluster-enabled yes \
    --cluster-config-file "nodes-$port.conf" >/dev/null
  core=$((core + 1))
done
sleep 2
docker exec rr-l0-7001 valkey-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-yes >/dev/null
for _ in $(seq 1 30); do
  docker exec rr-l0-7001 valkey-cli -p 7001 cluster info | grep -q 'cluster_state:ok' && break
  sleep 1
done
docker exec rr-l0 valkey-server --version > "$OUT/env.txt"
echo "memtier: $MEMTIER, client CPUs $CLIENT_CPUS, test time ${TEST_TIME}s, keys $KEYS" >> "$OUT/env.txt"

memtier() { # extra memtier args...
  docker run --rm --network host --cpuset-cpus "$CLIENT_CPUS" -v "$PWD/$OUT:/out" "$MEMTIER" "$@"
}

for topo in standalone cluster; do
  if [ "$topo" = standalone ]; then target=(-s 127.0.0.1 -p 6379); else target=(-s 127.0.0.1 -p 7001 --cluster-mode); fi
  for size in 100 4000; do
    # Load every key once so GETs hit, then run the mix.
    memtier "${target[@]}" --ratio=1:0 --key-pattern=P:P --key-maximum="$KEYS" -d "$size" \
      --requests=allkeys -t 4 -c 10 --hide-histogram >/dev/null 2>&1
    for pipe in 1 16; do
      name="${topo}_d${size}_p${pipe}"
      echo "== $name"
      memtier "${target[@]}" --ratio=1:4 --key-pattern=R:R --key-maximum="$KEYS" -d "$size" \
        -t 4 -c 25 --pipeline="$pipe" --test-time="$TEST_TIME" --distinct-client-seed \
        --print-percentiles=50,99,99.9 --hide-histogram --json-out-file="/out/$name.json" 2>&1 |
        grep -E "^(Type|Totals|Sets|Gets) " || true
    done
  done
done
python3 summarize.py "$OUT" | tee "$OUT/summary.md"
