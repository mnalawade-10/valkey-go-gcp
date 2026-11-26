# valkey-go: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 62810 | 0 | 0 | 0.304 |
| 1 | 57352 | 0 | 0 | 0.370 |
| 2 | 63597 | 0 | 0 | 0.286 |
| 3 | 62621 | 0 | 0 | 0.313 |
| 4 | 61672 | 0 | 0 | 0.325 |
| 5 | 6276 | 0 | 0 | 0.387 |
| 6 | 54627 | 13944 | 0 | 0.263 |
| 7 | 58307 | 14516 | 0 | 0.255 |
| 8 | 58400 | 14637 | 0 | 0.250 |
| 9 | 58359 | 14612 | 0 | 0.249 |
| 10 | 58658 | 14429 | 0 | 0.253 |
| 11 | 56955 | 14481 | 0 | 0.261 |
| 12 | 56828 | 14014 | 0 | 0.255 |
| 13 | 57853 | 14621 | 0 | 0.252 |
| 14 | 57249 | 14471 | 0 | 0.265 |
| 15 | 57839 | 14400 | 0 | 0.257 |
| 16 | 58237 | 14422 | 0 | 0.252 |
| 17 | 57579 | 14247 | 0 | 0.251 |
| 18 | 57903 | 14289 | 0 | 0.253 |
| 19 | 57698 | 14424 | 0 | 0.255 |
| 20 | 6 | 0 | 0 | 0.174 |

First readonly error: READONLY You can't write against a read only replica.

| valkey-go | 10 | 56041 | 0.139 | 0.201 | 0.284 | 0.455 | 920.575 | 45.1 | 101 | 201509 |
roles after: :6400 slave, :6401 master
