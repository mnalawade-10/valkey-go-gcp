#!/usr/bin/env bash
# Failover drill, local version: what clients see during a graceful failover, as in maintenance.
#
# A primary on :6400 and its replica on :6401, in Docker. Each client keeps 10 requests in flight with
# GLIDE's 80/20 mix against :6400, the address the app knows. $AT into the measured window the primary hands
# over to the replica (the FAILOVER command), so :6400 becomes a replica. The timeline shows, second by
# second: requests that worked, READONLY errors (writes still going to the old primary, now a replica),
# other errors, and p99.
#
# On Memorystore the primary endpoint moves to the new primary; here the address stays on the old node, so
# a client only recovers if it follows the server's REDIRECT (valkey-go/redirect, Valkey 8+) or detects the
# role change itself (what valkey-go PR #142 adds).
#
#   ./failover_drill.sh                       # valkey-go, valkey-go/redirect, go-redis, glide
#   ./failover_drill.sh valkey-go go-redis    # chosen clients
set -uo pipefail
cd "$(dirname "$0")"
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
OUT=${OUT:-results/failover-drill}
DUR=${DUR:-20s}
AT=${AT:-5s}
CPUS=${CPUS:-4-11,16-23}
CLIENTS=("$@")
[ ${#CLIENTS[@]} -eq 0 ] && CLIENTS=(valkey-go valkey-go/redirect go-redis glide)

mkdir -p "$OUT"
go build -o "$OUT/latency.bin" . || exit 1

cli() { docker exec "$1" valkey-cli -p "$2" "${@:3}"; }
role() { cli "$1" "$2" role | head -1; }
synced() { cli "$1" "$2" info replication | grep -q 'master_link_status:up'; }
wait_for() { for _ in $(seq 1 60); do "$@" && return 0; sleep 0.5; done; echo "timed out: $*" >&2; return 1; }

docker rm -f rr-fo-a rr-fo-b >/dev/null 2>&1
docker run -d --name rr-fo-a --network host --cpuset-cpus 1,13 "$IMAGE" \
  valkey-server --port 6400 --save "" --appendonly no >/dev/null
docker run -d --name rr-fo-b --network host --cpuset-cpus 2,14 "$IMAGE" \
  valkey-server --port 6401 --save "" --appendonly no --replicaof 127.0.0.1 6400 >/dev/null
wait_for synced rr-fo-b 6401 || exit 1

a_is_primary() { [ "$(role rr-fo-a 6400)" = master ] && synced rr-fo-b 6401; }
for c in "${CLIENTS[@]}"; do
  if ! a_is_primary; then # hand back to :6400 after the previous client's drill
    cli rr-fo-b 6401 failover >/dev/null
    wait_for a_is_primary || exit 1
  fi
  file="$OUT/$(echo "$c" | tr / _).md"
  echo "# $c: failover $AT into the measured window" > "$file"
  taskset -c "$CPUS" "$OUT/latency.bin" -addr 127.0.0.1:6400 -conc 10 -dur "$DUR" -warm 1s -timeline \
    -at "$AT" -at-cmd "docker exec rr-fo-a valkey-cli -p 6400 failover" -clients "$c" >> "$file" 2>&1
  echo "roles after: :6400 $(role rr-fo-a 6400), :6401 $(role rr-fo-b 6401)" >> "$file"
  echo "$c done"
done
docker rm -f rr-fo-a rr-fo-b >/dev/null 2>&1
echo "done: $OUT"
