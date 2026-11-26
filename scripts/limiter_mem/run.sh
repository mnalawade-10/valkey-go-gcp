#!/usr/bin/env bash
# Runs the rate-limiter memory meter against a fresh Valkey in Docker.
# Usage: ./run.sh                      (Valkey 9.0, 1,000,000 identifiers)
#        IMAGE=valkey/valkey:8.0-alpine N=1000000 ./run.sh
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
N=${N:-1000000}
CLIENT_CPUS=${CLIENT_CPUS:-4-11,16-23}
OUT=${OUT:-results}
mkdir -p "$OUT"
name="$(echo "$IMAGE" | tr '/:' '__')_n$N"

cleanup() { docker rm -f rr-limiter >/dev/null 2>&1 || true; }
trap cleanup EXIT
cleanup
docker run -d --name rr-limiter --network host --cpuset-cpus 0,12 "$IMAGE" \
  valkey-server --port 6379 --save "" --appendonly no >/dev/null
sleep 1

go build -o limiter_mem .
taskset -c "$CLIENT_CPUS" ./limiter_mem -n "$N" ${EXTRA:-} | tee "$OUT/$name.txt"
rm -f limiter_mem
