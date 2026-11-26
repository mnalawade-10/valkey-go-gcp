# valkey-go: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 60329 | 0 | 0 | 0.338 |
| 1 | 65132 | 0 | 0 | 0.284 |
| 2 | 65108 | 0 | 0 | 0.281 |
| 3 | 65940 | 0 | 0 | 0.275 |
| 4 | 65047 | 0 | 0 | 0.283 |
| 5 | 6425 | 0 | 0 | 0.397 |
| 6 | 56603 | 14287 | 0 | 0.245 |
| 7 | 58956 | 14669 | 0 | 0.245 |
| 8 | 58759 | 14598 | 0 | 0.248 |
| 9 | 58646 | 14755 | 0 | 0.246 |
| 10 | 59150 | 14716 | 0 | 0.247 |
| 11 | 56985 | 14364 | 0 | 0.260 |
| 12 | 57398 | 14260 | 0 | 0.259 |
| 13 | 57772 | 14510 | 0 | 0.251 |
| 14 | 56422 | 14264 | 0 | 0.259 |
| 15 | 40245 | 10018 | 0 | 0.479 |
| 16 | 34854 | 8565 | 0 | 0.463 |
| 17 | 57470 | 14333 | 0 | 0.255 |
| 18 | 57801 | 14264 | 0 | 0.251 |
| 19 | 59297 | 14766 | 0 | 0.249 |
| 20 | 1 | 0 | 0 | 0.099 |

First readonly error: READONLY You can't write against a read only replica.

| valkey-go | 10 | 54917 | 0.140 | 0.207 | 0.333 | 0.491 | 920.063 | 46.5 | 101 | 192372 |
roles after: :6400 slave, :6401 master
