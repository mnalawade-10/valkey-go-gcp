#!/usr/bin/env bash
# Diagnostic for TLS off vs on: is the server the limit? Runs each client for a few seconds at a fixed
# number of requests in flight against servers.sh tls, and reads the Valkey server's CPU time from
# /proc before and after. 100% means the server's single main thread is saturated.
#   ./tls_server_cpu.sh [in-flight] [clients...]
set -uo pipefail
cd "$(dirname "$0")"
CONC=${1:-100}
shift || true
CLIENTS=("$@")
[ ${#CLIENTS[@]} -eq 0 ] && CLIENTS=(valkey-go valkey-go+tls go-redis+tls glide+tls valkey-go/1-conn+tls)
BIN=${BIN:-/tmp/latency-bin}
PID=$(pgrep -P "$(docker inspect -f '{{.State.Pid}}' rr-tls)" valkey-server) # the container's PID 1 is tini
TCK=$(getconf CLK_TCK)
ticks() { sed 's/.*) //' "/proc/$PID/stat" | awk '{print $12 + $13}'; } # utime + stime, after "pid (comm) "
printf '| Client | In flight | Requests/s | p99 (ms) | Server CPU (one core = 100%%) |\n| :--- | :- | :- | :- | :- |\n'
for c in "${CLIENTS[@]}"; do
  t0=$(ticks)
  s0=$(date +%s.%N)
  row=$(taskset -c 4-11,16-23 "$BIN" -no-preload -addr 127.0.0.1:6391 -tls-addr 127.0.0.1:6390 -tls-ca ../tls/ca.crt \
    -conc "$CONC" -dur 5s -warm 1s -clients "$c" 2>/dev/null | grep "^| $c |")
  s1=$(date +%s.%N)
  t1=$(ticks)
  rps=$(echo "$row" | awk -F'|' '{gsub(/ /,"",$4); print $4}')
  p99=$(echo "$row" | awk -F'|' '{gsub(/ /,"",$7); print $7}')
  cpu=$(awk -v a="$t0" -v b="$t1" -v s="$s0" -v e="$s1" -v k="$TCK" 'BEGIN {printf "%.0f%%", (b - a) / k / (e - s) * 100}')
  echo "| $c | $CONC | $rps | $p99 | $cpu |"
done
