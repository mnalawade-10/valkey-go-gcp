#!/usr/bin/env bash
# Runs the same workloads for every client and every valkey-go variant (the off/on
# matrix, L3), repeated and with the order rotated each round, so drift on the machine hits all of them
# equally. One process per run, so the memory numbers are clean.
#
# Workloads: GLIDE's mix (80% GET, 20% SET) at fixed concurrency ($CONC requests in flight), and batches
# of 10 GETs sent as one pipeline ($BATCH_CONC batches in flight).
#
# Output in $OUT: one Go-benchmark-format file per client or variant, the readable tables (tables.md),
# and benchstat comparisons: clients-benchstat.txt (the four clients) and variants-benchstat.txt (each
# valkey-go variant against plain valkey-go).
#
#   COUNT=10 ./run_matrix.sh                # official: 10 repeats, 10 s each
#   COUNT=3 DUR=5s WARM=1s ./run_matrix.sh  # quick check
#
# TLS off vs on (after ../servers.sh tls): plain and TLS against the same server, each client against its
# +tls version (tls-<client>-benchstat.txt):
#   CLIENTS=valkey-go,go-redis,glide VARIANTS=valkey-go+tls,go-redis+tls,glide+tls \
#   EXTRA="-addr 127.0.0.1:6391 -tls-addr 127.0.0.1:6390 -tls-ca $PWD/../tls/ca.crt" OUT=results/matrix-tls ./run_matrix.sh
#
# Client-side cache at fixed hit rates (reads only; -hot sets the share of GETs that go to 1,000 hot keys):
#   CLIENTS=valkey-go VARIANTS=valkey-go/docache LOADS="-conc 1,10,100 -hot 0.5|-conc 1,10,100 -hot 0.9|-conc 1,10,100 -hot 0.99" \
#   OUT=results/matrix-cache ./run_matrix.sh
#
# Needs the local servers from ../servers.sh up.
set -uo pipefail
cd "$(dirname "$0")"
COUNT=${COUNT:-10}
DUR=${DUR:-10s}
WARM=${WARM:-2s}
CONC=${CONC:-1,10,100,1000}
BATCH_CONC=${BATCH_CONC:-1,10,100}
SIZE=${SIZE:-100}
MODE=${MODE:-standalone}
CPUS=${CPUS:-4-11,16-23}
CLIENTS=${CLIENTS:-valkey-go,go-redis,go-redis-autopipeline,glide}
VARIANTS=${VARIANTS-valkey-go/no-autopipe,valkey-go/always-pipe,valkey-go/flush-20us,valkey-go/1-conn,valkey-go/no-retry,valkey-go/docache,valkey-go/compat}
OUT=${OUT:-results/matrix-$MODE-${SIZE}B}
BENCHSTAT=${BENCHSTAT:-benchstat}
EXTRA=${EXTRA:-} # more flags for every run, e.g. -addr
# Workloads, separated by "|": each is one run of the tool per client and round.
LOADS=${LOADS:-"-conc $CONC|-conc $BATCH_CONC -batch 10"}

mkdir -p "$OUT"
BIN=$PWD/$OUT/latency.bin
go build -o "$BIN" . || exit 1
{
  echo "date: $(date -u +%FT%TZ)"
  echo "count=$COUNT dur=$DUR warm=$WARM conc=$CONC batch_conc=$BATCH_CONC size=$SIZE mode=$MODE cpus=$CPUS"
  echo "clients=$CLIENTS"
  echo "variants=$VARIANTS"
  echo "extra=$EXTRA"
  echo "loads=$LOADS"
  go version
  grep -E "valkey-go|go-redis|valkey-glide" go.mod | xargs -L1 echo
} > "$OUT/env.txt"

# shellcheck disable=SC2086 # $EXTRA holds several flags on purpose
taskset -c "$CPUS" "$BIN" -preload-only -mode "$MODE" -size "$SIZE" $EXTRA || exit 1
IFS=, read -ra ALL <<< "$CLIENTS${VARIANTS:+,$VARIANTS}"
n=${#ALL[@]}
for r in $(seq "$COUNT"); do
  for i in $(seq 0 $((n - 1))); do
    name=${ALL[$(((i + r) % n))]}
    file="$PWD/$OUT/$(echo "$name" | tr / _).txt"
    IFS="|" read -ra LOADLIST <<< "$LOADS"
    for args in "${LOADLIST[@]}"; do
      # shellcheck disable=SC2086 # $args and $EXTRA hold several flags on purpose
      taskset -c "$CPUS" "$BIN" -no-preload -mode "$MODE" -size "$SIZE" -dur "$DUR" -warm "$WARM" \
        -clients "$name" -benchfmt "$file" $args $EXTRA >> "$OUT/tables.md" 2>> "$OUT/stderr.log"
    done
    echo "round $r/$COUNT: $name done"
  done
done

cd "$OUT" || exit 1
"$BENCHSTAT" $(for c in ${CLIENTS//,/ }; do echo "$(echo "$c" | tr / _).txt"; done) > clients-benchstat.txt 2>&1
if [ -n "$VARIANTS" ]; then
  "$BENCHSTAT" valkey-go.txt $(for v in ${VARIANTS//,/ }; do echo "$(echo "$v" | tr / _).txt"; done) > variants-benchstat.txt 2>&1
  for v in ${VARIANTS//,/ }; do # each +tls variant against the same client without TLS
    case "$v" in *+tls)
      base=$(echo "${v%+tls}" | tr / _)
      "$BENCHSTAT" "$base.txt" "$(echo "$v" | tr / _).txt" > "tls-$base-benchstat.txt" 2>&1 ;;
    esac
  done
fi
echo "errors reported: $(grep -h -o '[0-9]* errors$' ./*.txt | awk '{s+=$1} END {print s+0}')"
echo "done: $OUT"
