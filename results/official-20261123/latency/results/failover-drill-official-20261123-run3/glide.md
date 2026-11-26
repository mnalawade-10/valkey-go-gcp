# glide: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)
[2m2026-11-23T10:05:14.619947Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m timeout_watchdog - Timeout: cmd=SET node=unknown cause=ServerUnresponsive { node: "unknown" } phase=Sent elapsed=250.092767ms configured=250ms pending=7 inflight=6→10 STABLE (system was already saturated at submission) p99=373µs suggested_timeout=250ms rss=91MB
[2m2026-11-23T10:05:15.311470Z[0m [33m WARN[0m [2mglide_logger[0m[2m:[0m pipeline - Response wait took 187.825515ms (threshold=125ms, response_timeout=250ms)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 35268 | 0 | 0 | 0.484 |
| 1 | 34016 | 0 | 0 | 0.516 |
| 2 | 31802 | 0 | 0 | 0.549 |
| 3 | 33926 | 0 | 0 | 0.519 |
| 4 | 34694 | 0 | 0 | 0.528 |
| 5 | 3627 | 0 | 30 | 0.575 |
| 6 | 24396 | 0 | 5988 | 0.552 |
| 7 | 26855 | 0 | 6819 | 0.554 |
| 8 | 25219 | 0 | 6438 | 0.614 |
| 9 | 27706 | 0 | 7033 | 0.524 |
| 10 | 26971 | 0 | 6753 | 0.538 |
| 11 | 28729 | 0 | 7358 | 0.495 |
| 12 | 27107 | 0 | 6877 | 0.546 |
| 13 | 26798 | 0 | 6721 | 0.520 |
| 14 | 25078 | 0 | 6203 | 0.585 |
| 15 | 24198 | 0 | 6036 | 0.588 |
| 16 | 25688 | 0 | 6366 | 0.568 |
| 17 | 24611 | 0 | 6184 | 0.593 |
| 18 | 25509 | 0 | 6290 | 0.559 |
| 19 | 24140 | 0 | 6175 | 0.578 |
| 20 | 0 | 0 | 0 | 0.000 |

First other error: timed out

| glide | 10 | 26817 | 0.292 | 0.385 | 0.549 | 0.837 | 189.567 | 57.1 | 98 | 91273 |
roles after: :6400 slave, :6401 master
