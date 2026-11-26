# valkey-go/redirect: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 65251 | 0 | 0 | 0.283 |
| 1 | 64877 | 0 | 0 | 0.284 |
| 2 | 66118 | 0 | 0 | 0.274 |
| 3 | 65200 | 0 | 0 | 0.282 |
| 4 | 66121 | 0 | 0 | 0.275 |
| 5 | 52896 | 0 | 0 | 0.298 |
| 6 | 64987 | 0 | 0 | 0.279 |
| 7 | 63914 | 0 | 0 | 0.292 |
| 8 | 66223 | 0 | 0 | 0.275 |
| 9 | 63997 | 0 | 0 | 0.288 |
| 10 | 65021 | 0 | 0 | 0.279 |
| 11 | 65495 | 0 | 0 | 0.272 |
| 12 | 65709 | 0 | 0 | 0.272 |
| 13 | 64478 | 0 | 0 | 0.282 |
| 14 | 64556 | 0 | 0 | 0.279 |
| 15 | 64451 | 0 | 0 | 0.282 |
| 16 | 65335 | 0 | 0 | 0.280 |
| 17 | 64134 | 0 | 0 | 0.278 |
| 18 | 65384 | 0 | 0 | 0.271 |
| 19 | 65029 | 0 | 0 | 0.275 |
| 20 | 5 | 0 | 0 | 0.180 |

| valkey-go/redirect | 10 | 64459 | 0.150 | 0.209 | 0.280 | 0.415 | 159.743 | 44.4 | 106 | 0 |
roles after: :6400 slave, :6401 master
