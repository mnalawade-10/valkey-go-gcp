#!/usr/bin/env bash
# Same as run.sh, but runs only the limiters named in ONLY (comma-separated; names may contain spaces).
# Example: ONLY='valkeylimiter fixed window (prefix rl),valkeylimiter GCRA (prefix rl)' ./run_only.sh
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
N=${N:-1000000}
ONLY=${ONLY:?set ONLY to the limiter names}
CLIENT_CPUS=${CLIENT_CPUS:-4-11,16-23}
OUT=${OUT:-results}
TAG=${TAG:-only}
mkdir -p "$OUT"
name="$(echo "$IMAGE" | tr '/:' '__')_n${N}_${TAG}"

cleanup() { docker rm -f rr-limiter >/dev/null 2>&1 || true; }
trap cleanup EXIT
cleanup
docker run -d --name rr-limiter --network host --cpuset-cpus 0,12 "$IMAGE" \
  valkey-server --port 6379 --save "" --appendonly no >/dev/null
sleep 1

go build -o limiter_mem_only .
taskset -c "$CLIENT_CPUS" ./limiter_mem_only -n "$N" -only "$ONLY" | tee "$OUT/$name.txt"
rm -f limiter_mem_only
