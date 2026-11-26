#!/usr/bin/env bash
# Shared helper: start or stop the local test servers used by the benchmark.
#   servers.sh up    -> standalone Valkey on :6379 (core 0) and a 3-node cluster on :7001-7003 (cores 1-3)
#   servers.sh tls   -> one more standalone Valkey (core 0) with a plain port :6391 and a TLS port :6390,
#                       for TLS off vs on against the same server. Makes a throwaway test CA and server
#                       certificate in ./tls first, if they aren't there.
#   servers.sh down  -> remove them all
# Host networking, no persistence. IMAGE overrides the Valkey image.
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${IMAGE:-valkey/valkey:9.0-alpine}
NAMES=(rr-s rr-c7001 rr-c7002 rr-c7003 rr-tls)

down() { docker rm -f "${NAMES[@]}" >/dev/null 2>&1 || true; }

certs() {
  [ -f tls/server.crt ] && return 0
  mkdir -p tls
  openssl req -x509 -newkey rsa:2048 -nodes -keyout tls/ca.key -out tls/ca.crt -days 60 \
    -subj "/CN=benchmark test CA" 2>/dev/null
  openssl req -newkey rsa:2048 -nodes -keyout tls/server.key -out tls/server.csr -subj "/CN=localhost" 2>/dev/null
  printf 'subjectAltName=IP:127.0.0.1,DNS:localhost\n' > tls/san.ext
  openssl x509 -req -in tls/server.csr -CA tls/ca.crt -CAkey tls/ca.key -CAcreateserial \
    -out tls/server.crt -days 60 -extfile tls/san.ext 2>/dev/null
  chmod 644 tls/server.key # the server runs as uid 999 in the container; a test-only key
}

tls() {
  certs
  docker rm -f rr-tls >/dev/null 2>&1 || true
  docker run -d --name rr-tls --network host --cpuset-cpus 0,12 -v "$PWD/tls:/tls:ro" "$IMAGE" \
    valkey-server --port 6391 --tls-port 6390 --tls-cert-file /tls/server.crt --tls-key-file /tls/server.key \
    --tls-ca-cert-file /tls/ca.crt --tls-auth-clients no --save "" --appendonly no >/dev/null
  for _ in $(seq 1 20); do
    docker exec rr-tls valkey-cli --tls --cacert /tls/ca.crt -p 6390 ping 2>/dev/null | grep -q PONG && return 0
    sleep 0.5
  done
  echo "TLS server did not answer" >&2
  return 1
}

up() {
  down
  docker run -d --name rr-s --network host --cpuset-cpus 0,12 "$IMAGE" \
    valkey-server --port 6379 --save "" --appendonly no >/dev/null
  local core=1
  for port in 7001 7002 7003; do
    docker run -d --name "rr-c$port" --network host --cpuset-cpus "$core,$((core + 12))" "$IMAGE" \
      valkey-server --port "$port" --save "" --appendonly no --cluster-enabled yes \
      --cluster-config-file "nodes-$port.conf" >/dev/null
    core=$((core + 1))
  done
  sleep 2
  docker exec rr-c7001 valkey-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-yes >/dev/null
  for _ in $(seq 1 30); do
    docker exec rr-c7001 valkey-cli -p 7001 cluster info | grep -q 'cluster_state:ok' && return 0
    sleep 1
  done
  echo "cluster did not become ready" >&2
  return 1
}

case "${1:-}" in
  up) up ;;
  tls) tls ;;
  down) down ;;
  *) echo "usage: $0 up|tls|down" >&2; exit 2 ;;
esac
