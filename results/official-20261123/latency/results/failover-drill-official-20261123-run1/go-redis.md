# go-redis: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 46334 | 0 | 0 | 0.387 |
| 1 | 46475 | 0 | 0 | 0.377 |
| 2 | 46897 | 0 | 0 | 0.356 |
| 3 | 45975 | 0 | 0 | 0.390 |
| 4 | 45990 | 0 | 0 | 0.396 |
| 5 | 4530 | 4 | 0 | 0.527 |
| 6 | 404 | 93 | 0 | 0.773 |
| 7 | 392 | 97 | 0 | 0.896 |
| 8 | 427 | 98 | 0 | 0.789 |
| 9 | 348 | 96 | 0 | 0.812 |
| 10 | 387 | 97 | 0 | 0.792 |
| 11 | 352 | 96 | 0 | 0.770 |
| 12 | 433 | 97 | 0 | 0.745 |
| 13 | 324 | 91 | 0 | 0.780 |
| 14 | 428 | 99 | 0 | 0.777 |
| 15 | 386 | 97 | 0 | 0.865 |
| 16 | 346 | 94 | 0 | 0.756 |
| 17 | 501 | 100 | 0 | 0.700 |
| 18 | 399 | 97 | 0 | 0.792 |
| 19 | 426 | 100 | 0 | 0.872 |
| 20 | 0 | 0 | 0 | 0.000 |

First readonly error: READONLY You can't write against a read only replica.

| go-redis | 10 | 12087 | 0.213 | 0.284 | 0.426 | 0.732 | 2.221 | 76.9 | 95 | 1366 |
roles after: :6400 slave, :6401 master
