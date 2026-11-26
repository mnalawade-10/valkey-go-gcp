# glide: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)
[2m2026-11-23T10:03:35.961676Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m pipeline - Response wait took 250.138207ms (threshold=125ms, response_timeout=250ms)
[2m2026-11-23T10:03:35.962030Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m timeout_watchdog - Timeout: cmd=SET node=unknown cause=ServerUnresponsive { node: "unknown" } phase=Sent elapsed=250.315457ms configured=250ms pending=3 inflight=7→9 STABLE (system was already saturated at submission) p99=355µs suggested_timeout=250ms rss=95MB

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 36110 | 0 | 0 | 0.479 |
| 1 | 35008 | 0 | 0 | 0.534 |
| 2 | 35711 | 0 | 0 | 0.470 |
| 3 | 33513 | 0 | 0 | 0.548 |
| 4 | 33512 | 0 | 0 | 0.561 |
| 5 | 4213 | 0 | 30 | 0.513 |
| 6 | 22512 | 0 | 5562 | 0.639 |
| 7 | 26340 | 0 | 6676 | 0.549 |
| 8 | 25962 | 0 | 6643 | 0.550 |
| 9 | 27264 | 0 | 6916 | 0.526 |
| 10 | 26671 | 0 | 6679 | 0.539 |
| 11 | 27962 | 0 | 7159 | 0.507 |
| 12 | 26528 | 0 | 6736 | 0.540 |
| 13 | 26988 | 0 | 6700 | 0.569 |
| 14 | 28474 | 0 | 7095 | 0.521 |
| 15 | 29220 | 0 | 7267 | 0.490 |
| 16 | 27818 | 0 | 7022 | 0.497 |
| 17 | 28909 | 0 | 7141 | 0.468 |
| 18 | 29813 | 0 | 7547 | 0.466 |
| 19 | 29226 | 0 | 7225 | 0.469 |
| 20 | 5 | 0 | 1 | 0.465 |

First other error: timed out

| glide | 10 | 28088 | 0.277 | 0.371 | 0.523 | 0.839 | 180.223 | 55.1 | 98 | 96402 |
roles after: :6400 slave, :6401 master
