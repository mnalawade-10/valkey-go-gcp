# go-redis: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 47234 | 0 | 0 | 0.363 |
| 1 | 47729 | 0 | 0 | 0.359 |
| 2 | 47857 | 0 | 0 | 0.350 |
| 3 | 47471 | 0 | 0 | 0.361 |
| 4 | 42937 | 0 | 0 | 0.427 |
| 5 | 4698 | 3 | 0 | 0.499 |
| 6 | 395 | 96 | 0 | 0.703 |
| 7 | 352 | 96 | 0 | 0.691 |
| 8 | 447 | 96 | 0 | 0.817 |
| 9 | 403 | 99 | 0 | 0.758 |
| 10 | 432 | 97 | 0 | 0.724 |
| 11 | 382 | 94 | 0 | 0.923 |
| 12 | 410 | 93 | 0 | 0.706 |
| 13 | 373 | 92 | 0 | 0.783 |
| 14 | 467 | 101 | 0 | 0.775 |
| 15 | 387 | 96 | 0 | 0.899 |
| 16 | 467 | 99 | 0 | 0.752 |
| 17 | 290 | 95 | 0 | 0.870 |
| 18 | 305 | 93 | 0 | 0.851 |
| 19 | 301 | 94 | 0 | 0.797 |
| 20 | 0 | 0 | 0 | 0.000 |

First readonly error: READONLY You can't write against a read only replica.

| go-redis | 10 | 12167 | 0.210 | 0.282 | 0.419 | 0.735 | 3.381 | 74.9 | 93 | 1354 |
roles after: :6400 slave, :6401 master
