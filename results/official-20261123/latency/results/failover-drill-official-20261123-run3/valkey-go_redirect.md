# valkey-go/redirect: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 64648 | 0 | 0 | 0.281 |
| 1 | 62353 | 0 | 0 | 0.298 |
| 2 | 64505 | 0 | 0 | 0.287 |
| 3 | 64303 | 0 | 0 | 0.286 |
| 4 | 63505 | 0 | 0 | 0.295 |
| 5 | 50115 | 0 | 0 | 0.319 |
| 6 | 64757 | 0 | 0 | 0.284 |
| 7 | 65215 | 0 | 0 | 0.275 |
| 8 | 63554 | 0 | 0 | 0.299 |
| 9 | 64007 | 0 | 0 | 0.279 |
| 10 | 64370 | 0 | 0 | 0.285 |
| 11 | 64560 | 0 | 0 | 0.287 |
| 12 | 65198 | 0 | 0 | 0.277 |
| 13 | 64247 | 0 | 0 | 0.278 |
| 14 | 65183 | 0 | 0 | 0.273 |
| 15 | 64844 | 0 | 0 | 0.279 |
| 16 | 65264 | 0 | 0 | 0.269 |
| 17 | 64125 | 0 | 0 | 0.274 |
| 18 | 62066 | 0 | 0 | 0.316 |
| 19 | 63879 | 0 | 0 | 0.283 |
| 20 | 5 | 0 | 0 | 0.174 |

| valkey-go/redirect | 10 | 63535 | 0.152 | 0.212 | 0.287 | 0.415 | 186.623 | 45.3 | 107 | 0 |
roles after: :6400 slave, :6401 master
