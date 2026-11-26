# glide: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)
[2m2026-11-23T10:01:57.243547Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m timeout_watchdog - Timeout: cmd=SET node=unknown cause=ServerUnresponsive { node: "unknown" } phase=Sent elapsed=250.128837ms configured=250ms pending=4 inflight=2→10 STABLE (system was already saturated at submission) p99=400µs suggested_timeout=250ms rss=95MB
[2m2026-11-23T10:01:57.494802Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m pipeline - Response wait took 250.203427ms (threshold=125ms, response_timeout=250ms)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 34064 | 0 | 0 | 0.535 |
| 1 | 36273 | 0 | 0 | 0.485 |
| 2 | 33129 | 0 | 0 | 0.558 |
| 3 | 33269 | 0 | 0 | 0.548 |
| 4 | 35162 | 0 | 0 | 0.510 |
| 5 | 3426 | 0 | 30 | 0.615 |
| 6 | 25678 | 0 | 6334 | 0.550 |
| 7 | 27582 | 0 | 6980 | 0.514 |
| 8 | 27817 | 0 | 7096 | 0.501 |
| 9 | 27082 | 0 | 6899 | 0.533 |
| 10 | 28105 | 0 | 6984 | 0.486 |
| 11 | 27867 | 0 | 7202 | 0.493 |
| 12 | 26949 | 0 | 6785 | 0.526 |
| 13 | 27798 | 0 | 6940 | 0.489 |
| 14 | 26893 | 0 | 6723 | 0.525 |
| 15 | 26721 | 0 | 6611 | 0.544 |
| 16 | 24451 | 0 | 6073 | 0.576 |
| 17 | 25453 | 0 | 6370 | 0.579 |
| 18 | 24826 | 0 | 6224 | 0.652 |
| 19 | 24540 | 0 | 6243 | 0.633 |
| 20 | 2 | 0 | 0 | 0.254 |

First other error: timed out

| glide | 10 | 27354 | 0.284 | 0.381 | 0.541 | 0.900 | 173.951 | 56.0 | 98 | 93495 |
roles after: :6400 slave, :6401 master
