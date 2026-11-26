# valkey-go: failover 5s into the measured window
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=20s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
at 5s ran "docker exec rr-fo-a valkey-cli -p 6400 failover": OK (err=<nil>)

Timeline (second of the measured window):

| Second | OK | READONLY errors | Other errors | p99 (ms) |
| :- | :- | :- | :- | :- |
| 0 | 66236 | 0 | 0 | 0.272 |
| 1 | 65652 | 0 | 0 | 0.273 |
| 2 | 65147 | 0 | 0 | 0.277 |
| 3 | 65395 | 0 | 0 | 0.277 |
| 4 | 65634 | 0 | 0 | 0.268 |
| 5 | 6694 | 0 | 0 | 0.361 |
| 6 | 57513 | 14556 | 0 | 0.252 |
| 7 | 58974 | 14624 | 0 | 0.247 |
| 8 | 57961 | 14467 | 0 | 0.268 |
| 9 | 58991 | 14743 | 0 | 0.248 |
| 10 | 59985 | 14997 | 0 | 0.241 |
| 11 | 58903 | 14785 | 0 | 0.245 |
| 12 | 58434 | 14567 | 0 | 0.254 |
| 13 | 57832 | 14609 | 0 | 0.256 |
| 14 | 58834 | 14929 | 0 | 0.242 |
| 15 | 58461 | 14322 | 0 | 0.249 |
| 16 | 58836 | 14656 | 0 | 0.254 |
| 17 | 58430 | 14425 | 0 | 0.249 |
| 18 | 58386 | 14542 | 0 | 0.249 |
| 19 | 58861 | 14636 | 0 | 0.251 |
| 20 | 1 | 0 | 0 | 0.191 |

First readonly error: READONLY You can't write against a read only replica.

| valkey-go | 10 | 57758 | 0.136 | 0.194 | 0.261 | 0.380 | 902.143 | 44.7 | 100 | 204861 |
roles after: :6400 slave, :6401 master
