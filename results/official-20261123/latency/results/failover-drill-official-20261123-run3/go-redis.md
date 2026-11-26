# go-redis: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 46924 | 0 | 0 | 0.375 |
| 1 | 47273 | 0 | 0 | 0.355 |
| 2 | 47399 | 0 | 0 | 0.357 |
| 3 | 46795 | 0 | 0 | 0.363 |
| 4 | 47543 | 0 | 0 | 0.351 |
| 5 | 4711 | 2 | 0 | 0.457 |
| 6 | 351 | 90 | 0 | 0.813 |
| 7 | 394 | 96 | 0 | 0.759 |
| 8 | 515 | 98 | 0 | 0.807 |
| 9 | 459 | 99 | 0 | 0.796 |
| 10 | 326 | 94 | 0 | 0.787 |
| 11 | 354 | 95 | 0 | 0.770 |
| 12 | 367 | 95 | 0 | 0.740 |
| 13 | 316 | 95 | 0 | 0.824 |
| 14 | 307 | 94 | 0 | 0.764 |
| 15 | 323 | 98 | 0 | 0.776 |
| 16 | 294 | 98 | 0 | 0.713 |
| 17 | 395 | 96 | 0 | 0.786 |
| 18 | 319 | 100 | 0 | 0.818 |
| 19 | 381 | 101 | 0 | 0.742 |
| 20 | 0 | 0 | 0 | 0.000 |

First readonly error: READONLY You can't write against a read only replica.

| go-redis | 10 | 12287 | 0.209 | 0.275 | 0.399 | 0.711 | 1.925 | 75.3 | 93 | 1361 |
roles after: :6400 slave, :6401 master
