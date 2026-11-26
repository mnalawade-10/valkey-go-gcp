#!/usr/bin/env bash
# Version sweep: valkey-go releases against the README-era rueidis v0.0.90, SET on one server.
# Each valkey-go version gets its own build (one module version per build); rueidis v0.0.90 runs in the
# same process as an anchor, so each version is compared with the README-era client under the same
# conditions. go-redis and GLIDE are excluded to keep it short.
set -euo pipefail
cd "$(dirname "$0")"
SRC=$PWD
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
VERSIONS=${VERSIONS:-v1.0.40 v1.0.50 v1.0.60 v1.0.70 v1.0.80}  # v1.0.35-v1.0.39 declare the old module path github.com/rueian/valkey-go
COUNT=${COUNT:-6}
BENCHTIME=${BENCHTIME:-50000x}
CLIENT_CPUS=${CLIENT_CPUS:-4-11,16-23}
OUT=${OUT:-results}
MUX=${MUX:-0}                        # PipelineMultiplex for both clients: 0 = default (4 connections), -1 = one
FILE=${FILE:-sweep_set.txt}
EXCL=goredis-v9.0.0-rc.2,goredis-v9.23.0,goredis-v9.23.0-autopipeline,glide-v2.6.1,rueidis-v0.0.90-csc,valkeygo-v1.0.80-csc,goredis-v9.23.0-csc
mkdir -p "$OUT"
: > "$OUT/$FILE"

cleanup() { docker rm -f rr-sweep >/dev/null 2>&1 || true; }
trap cleanup EXIT
cleanup
docker run -d --name rr-sweep --network host --cpuset-cpus 0,12 "$IMAGE" \
  valkey-server --port 6379 --save "" --appendonly no >/dev/null
sleep 1

for v in $VERSIONS; do
  dir=/tmp/rr-sweep-$v
  rm -rf "$dir" && mkdir -p "$dir"
  cp "$SRC"/go.mod "$SRC"/go.sum "$SRC"/*_test.go "$dir"/
  (cd "$dir" && go mod edit -require="github.com/valkey-io/valkey-go@$v" && go mod tidy && go test -c -o sweep.test .)
  echo "== valkey-go $v" | tee -a "$OUT/$FILE"
  taskset -c "$CLIENT_CPUS" "$dir/sweep.test" -test.run '^$' -test.bench "${BENCH:-^BenchmarkSingleClientSet\$}" -test.benchmem \
    -test.benchtime "$BENCHTIME" -test.count "$COUNT" -test.cpu 10 -exclude "$EXCL" -mux "$MUX" |
    sed "s/lib=valkeygo-v1.0.80/lib=valkeygo-$v/" | tee -a "$OUT/$FILE"
  rm -rf "$dir"
done
