# valkey-go/redirect: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 64448 | 0 | 0 | 0.277 |
| 1 | 64836 | 0 | 0 | 0.282 |
| 2 | 68046 | 0 | 0 | 0.263 |
| 3 | 65772 | 0 | 0 | 0.277 |
| 4 | 64286 | 0 | 0 | 0.282 |
| 5 | 52183 | 0 | 0 | 0.317 |
| 6 | 64410 | 0 | 0 | 0.278 |
| 7 | 63169 | 0 | 0 | 0.292 |
| 8 | 64130 | 0 | 0 | 0.284 |
| 9 | 63825 | 0 | 0 | 0.284 |
| 10 | 63540 | 0 | 0 | 0.283 |
| 11 | 63660 | 0 | 0 | 0.295 |
| 12 | 64396 | 0 | 0 | 0.288 |
| 13 | 64009 | 0 | 0 | 0.284 |
| 14 | 63220 | 0 | 0 | 0.284 |
| 15 | 64006 | 0 | 0 | 0.285 |
| 16 | 64467 | 0 | 0 | 0.288 |
| 17 | 62517 | 0 | 0 | 0.296 |
| 18 | 65142 | 0 | 0 | 0.282 |
| 19 | 62323 | 0 | 0 | 0.304 |
| 20 | 9 | 0 | 0 | 0.235 |

| valkey-go/redirect | 10 | 63619 | 0.152 | 0.212 | 0.286 | 0.423 | 175.615 | 44.7 | 111 | 0 |
roles after: :6400 slave, :6401 master
