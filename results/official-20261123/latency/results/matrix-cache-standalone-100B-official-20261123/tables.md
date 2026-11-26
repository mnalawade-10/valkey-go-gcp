mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18959 | 0.006 | 0.111 | 0.171 | 0.327 | 9.863 | 49.7 | 76 | 0 |
| valkey-go/docache | 10 | 115220 | 0.019 | 0.210 | 0.302 | 0.462 | 3.143 | 29.0 | 321 | 0 |
| valkey-go/docache | 100 | 267540 | 0.045 | 1.018 | 1.684 | 6.007 | 9.687 | 18.9 | 876 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 76097 | 0.001 | 0.086 | 0.140 | 0.222 | 2.947 | 12.4 | 64 | 0 |
| valkey-go/docache | 10 | 404456 | 0.001 | 0.104 | 0.301 | 0.425 | 3.261 | 8.3 | 233 | 0 |
| valkey-go/docache | 100 | 881924 | 0.001 | 0.312 | 1.448 | 2.081 | 9.703 | 7.5 | 524 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 458790 | 0.000 | 0.001 | 0.095 | 0.176 | 1.584 | 2.1 | 64 | 0 |
| valkey-go/docache | 10 | 2151625 | 0.001 | 0.001 | 0.124 | 0.257 | 4.279 | 2.4 | 180 | 0 |
| valkey-go/docache | 100 | 3160627 | 0.001 | 0.074 | 0.441 | 2.491 | 8.855 | 3.6 | 339 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15632 | 0.061 | 0.069 | 0.100 | 0.246 | 1.677 | 45.5 | 35 | 0 |
| valkey-go | 10 | 73885 | 0.130 | 0.186 | 0.253 | 0.386 | 1.584 | 43.2 | 41 | 0 |
| valkey-go | 100 | 352661 | 0.278 | 0.389 | 0.517 | 0.851 | 2.131 | 15.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15818 | 0.061 | 0.071 | 0.103 | 0.244 | 1.702 | 46.4 | 35 | 0 |
| valkey-go | 10 | 74973 | 0.129 | 0.183 | 0.250 | 0.381 | 1.248 | 42.8 | 41 | 0 |
| valkey-go | 100 | 345853 | 0.283 | 0.397 | 0.536 | 0.895 | 2.133 | 15.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15316 | 0.062 | 0.073 | 0.104 | 0.257 | 1.524 | 46.5 | 35 | 0 |
| valkey-go | 10 | 75326 | 0.128 | 0.182 | 0.247 | 0.478 | 1.505 | 43.5 | 40 | 0 |
| valkey-go | 100 | 347058 | 0.281 | 0.396 | 0.557 | 0.932 | 2.539 | 15.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15625 | 0.061 | 0.071 | 0.100 | 0.267 | 1.619 | 46.3 | 35 | 0 |
| valkey-go | 10 | 74765 | 0.129 | 0.183 | 0.246 | 0.476 | 1.327 | 43.1 | 40 | 0 |
| valkey-go | 100 | 346796 | 0.282 | 0.396 | 0.544 | 0.856 | 2.012 | 15.4 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15519 | 0.062 | 0.074 | 0.106 | 0.273 | 1.557 | 46.6 | 37 | 0 |
| valkey-go | 10 | 74905 | 0.129 | 0.182 | 0.248 | 0.473 | 1.619 | 43.3 | 40 | 0 |
| valkey-go | 100 | 350112 | 0.279 | 0.393 | 0.544 | 0.920 | 1.976 | 15.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15663 | 0.061 | 0.070 | 0.100 | 0.319 | 1.552 | 45.3 | 35 | 0 |
| valkey-go | 10 | 75096 | 0.128 | 0.183 | 0.252 | 0.485 | 2.067 | 43.5 | 40 | 0 |
| valkey-go | 100 | 343775 | 0.283 | 0.401 | 0.570 | 0.985 | 2.677 | 15.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18807 | 0.006 | 0.112 | 0.168 | 0.378 | 1.650 | 50.5 | 72 | 0 |
| valkey-go/docache | 10 | 89744 | 0.077 | 0.266 | 0.375 | 0.596 | 3.845 | 33.5 | 231 | 0 |
| valkey-go/docache | 100 | 197419 | 0.087 | 1.254 | 1.651 | 2.235 | 5.387 | 20.1 | 597 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 79090 | 0.001 | 0.079 | 0.137 | 0.219 | 1.645 | 12.0 | 64 | 0 |
| valkey-go/docache | 10 | 408795 | 0.001 | 0.099 | 0.300 | 0.441 | 6.103 | 8.2 | 237 | 0 |
| valkey-go/docache | 100 | 918454 | 0.001 | 0.255 | 1.405 | 2.025 | 10.575 | 7.3 | 571 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 472914 | 0.000 | 0.000 | 0.091 | 0.172 | 2.641 | 2.1 | 62 | 0 |
| valkey-go/docache | 10 | 2163552 | 0.001 | 0.001 | 0.123 | 0.255 | 4.247 | 2.4 | 174 | 0 |
| valkey-go/docache | 100 | 3162923 | 0.001 | 0.074 | 0.439 | 2.501 | 9.047 | 3.6 | 375 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18161 | 0.078 | 0.114 | 0.167 | 0.372 | 1.422 | 51.4 | 64 | 0 |
| valkey-go/docache | 10 | 87120 | 0.082 | 0.275 | 0.383 | 0.601 | 3.219 | 34.0 | 228 | 0 |
| valkey-go/docache | 100 | 191874 | 0.090 | 1.295 | 1.706 | 2.243 | 4.591 | 20.4 | 543 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 77031 | 0.001 | 0.083 | 0.141 | 0.243 | 1.524 | 12.3 | 64 | 0 |
| valkey-go/docache | 10 | 399393 | 0.001 | 0.101 | 0.309 | 0.467 | 3.723 | 8.5 | 233 | 0 |
| valkey-go/docache | 100 | 918769 | 0.001 | 0.245 | 1.411 | 2.057 | 6.743 | 7.5 | 573 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 452076 | 0.000 | 0.001 | 0.094 | 0.180 | 1.625 | 2.2 | 64 | 0 |
| valkey-go/docache | 10 | 2160719 | 0.001 | 0.001 | 0.124 | 0.261 | 4.839 | 2.4 | 172 | 0 |
| valkey-go/docache | 100 | 3199327 | 0.001 | 0.074 | 0.441 | 2.397 | 15.927 | 3.6 | 373 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15306 | 0.062 | 0.072 | 0.106 | 0.290 | 1.758 | 46.4 | 35 | 0 |
| valkey-go | 10 | 73748 | 0.131 | 0.186 | 0.251 | 0.473 | 1.615 | 43.5 | 39 | 0 |
| valkey-go | 100 | 348523 | 0.281 | 0.394 | 0.533 | 0.839 | 2.409 | 15.4 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15787 | 0.061 | 0.071 | 0.101 | 0.316 | 1.569 | 45.8 | 35 | 0 |
| valkey-go | 10 | 72633 | 0.132 | 0.190 | 0.261 | 0.492 | 1.863 | 44.0 | 39 | 0 |
| valkey-go | 100 | 344365 | 0.284 | 0.400 | 0.543 | 0.873 | 2.395 | 15.8 | 63 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15955 | 0.060 | 0.071 | 0.100 | 0.298 | 1.528 | 45.8 | 35 | 0 |
| valkey-go | 10 | 75030 | 0.128 | 0.182 | 0.246 | 0.484 | 1.560 | 43.7 | 39 | 0 |
| valkey-go | 100 | 341148 | 0.285 | 0.405 | 0.600 | 0.963 | 2.469 | 15.9 | 64 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14885 | 0.061 | 0.079 | 0.142 | 0.344 | 2.001 | 53.4 | 36 | 0 |
| valkey-go | 10 | 73716 | 0.130 | 0.186 | 0.253 | 0.502 | 1.769 | 43.8 | 41 | 0 |
| valkey-go | 100 | 350430 | 0.279 | 0.392 | 0.540 | 0.863 | 2.577 | 15.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15733 | 0.061 | 0.071 | 0.102 | 0.328 | 1.717 | 46.2 | 35 | 0 |
| valkey-go | 10 | 74530 | 0.129 | 0.183 | 0.246 | 0.487 | 1.752 | 43.5 | 41 | 0 |
| valkey-go | 100 | 348562 | 0.281 | 0.393 | 0.545 | 0.871 | 2.233 | 15.5 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15315 | 0.062 | 0.072 | 0.102 | 0.292 | 1.546 | 45.7 | 35 | 0 |
| valkey-go | 10 | 74236 | 0.130 | 0.185 | 0.249 | 0.483 | 1.501 | 43.7 | 41 | 0 |
| valkey-go | 100 | 347640 | 0.280 | 0.398 | 0.560 | 0.903 | 2.175 | 15.7 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18214 | 0.006 | 0.115 | 0.172 | 0.377 | 1.438 | 51.8 | 74 | 0 |
| valkey-go/docache | 10 | 88791 | 0.081 | 0.270 | 0.374 | 0.569 | 3.233 | 34.1 | 229 | 0 |
| valkey-go/docache | 100 | 189126 | 0.088 | 1.312 | 1.770 | 2.485 | 6.643 | 19.9 | 576 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 78151 | 0.001 | 0.080 | 0.137 | 0.224 | 2.297 | 12.2 | 66 | 0 |
| valkey-go/docache | 10 | 415062 | 0.001 | 0.098 | 0.292 | 0.424 | 3.447 | 8.4 | 251 | 0 |
| valkey-go/docache | 100 | 952404 | 0.001 | 0.230 | 1.346 | 1.986 | 6.075 | 7.2 | 554 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 476552 | 0.000 | 0.000 | 0.092 | 0.170 | 1.525 | 2.0 | 64 | 0 |
| valkey-go/docache | 10 | 2182240 | 0.000 | 0.001 | 0.122 | 0.256 | 5.655 | 2.4 | 189 | 0 |
| valkey-go/docache | 100 | 3203426 | 0.001 | 0.073 | 0.441 | 2.443 | 8.967 | 3.5 | 368 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 17633 | 0.080 | 0.118 | 0.169 | 0.394 | 1.609 | 53.3 | 64 | 0 |
| valkey-go/docache | 10 | 87433 | 0.081 | 0.274 | 0.384 | 0.581 | 7.007 | 33.2 | 226 | 0 |
| valkey-go/docache | 100 | 191188 | 0.086 | 1.295 | 1.664 | 2.153 | 4.539 | 20.3 | 545 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 76070 | 0.001 | 0.083 | 0.142 | 0.230 | 1.666 | 12.6 | 64 | 0 |
| valkey-go/docache | 10 | 409343 | 0.001 | 0.099 | 0.299 | 0.435 | 3.279 | 8.3 | 241 | 0 |
| valkey-go/docache | 100 | 913132 | 0.001 | 0.248 | 1.452 | 2.279 | 8.951 | 7.6 | 534 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 473334 | 0.000 | 0.000 | 0.092 | 0.169 | 1.511 | 2.0 | 62 | 0 |
| valkey-go/docache | 10 | 2160757 | 0.001 | 0.001 | 0.123 | 0.258 | 3.575 | 2.4 | 180 | 0 |
| valkey-go/docache | 100 | 2577592 | 0.001 | 0.069 | 0.496 | 4.783 | 11.567 | 3.7 | 314 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15618 | 0.061 | 0.071 | 0.101 | 0.294 | 1.527 | 45.1 | 35 | 0 |
| valkey-go | 10 | 74587 | 0.129 | 0.183 | 0.245 | 0.471 | 1.712 | 43.3 | 38 | 0 |
| valkey-go | 100 | 352336 | 0.277 | 0.391 | 0.536 | 0.811 | 2.161 | 15.2 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14733 | 0.064 | 0.079 | 0.117 | 0.366 | 3.009 | 49.3 | 34 | 0 |
| valkey-go | 10 | 74233 | 0.130 | 0.185 | 0.253 | 0.477 | 2.083 | 43.5 | 40 | 0 |
| valkey-go | 100 | 346637 | 0.282 | 0.396 | 0.547 | 0.944 | 2.030 | 15.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15564 | 0.061 | 0.073 | 0.104 | 0.281 | 1.445 | 46.5 | 37 | 0 |
| valkey-go | 10 | 73960 | 0.130 | 0.185 | 0.249 | 0.478 | 1.835 | 43.6 | 40 | 0 |
| valkey-go | 100 | 352440 | 0.277 | 0.390 | 0.532 | 0.861 | 2.073 | 15.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15393 | 0.061 | 0.073 | 0.118 | 0.348 | 1.994 | 46.4 | 37 | 0 |
| valkey-go | 10 | 74105 | 0.130 | 0.185 | 0.251 | 0.471 | 1.557 | 43.8 | 40 | 0 |
| valkey-go | 100 | 346396 | 0.281 | 0.400 | 0.566 | 0.900 | 2.255 | 15.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15748 | 0.061 | 0.070 | 0.106 | 0.324 | 1.633 | 44.9 | 35 | 0 |
| valkey-go | 10 | 69730 | 0.133 | 0.204 | 0.334 | 0.512 | 1.695 | 45.6 | 40 | 0 |
| valkey-go | 100 | 332026 | 0.292 | 0.423 | 0.596 | 0.889 | 2.449 | 15.9 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15543 | 0.061 | 0.073 | 0.106 | 0.313 | 1.582 | 47.2 | 35 | 0 |
| valkey-go | 10 | 74412 | 0.129 | 0.183 | 0.249 | 0.484 | 1.688 | 43.7 | 39 | 0 |
| valkey-go | 100 | 350174 | 0.280 | 0.392 | 0.535 | 0.922 | 2.171 | 15.6 | 64 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18014 | 0.006 | 0.117 | 0.179 | 0.375 | 1.829 | 51.3 | 70 | 0 |
| valkey-go/docache | 10 | 88367 | 0.078 | 0.271 | 0.384 | 0.599 | 3.185 | 33.7 | 220 | 0 |
| valkey-go/docache | 100 | 195130 | 0.086 | 1.271 | 1.656 | 2.231 | 4.783 | 20.6 | 568 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 77308 | 0.001 | 0.082 | 0.139 | 0.228 | 1.694 | 12.4 | 66 | 0 |
| valkey-go/docache | 10 | 413780 | 0.001 | 0.098 | 0.295 | 0.431 | 3.693 | 8.2 | 245 | 0 |
| valkey-go/docache | 100 | 922092 | 0.001 | 0.235 | 1.393 | 1.989 | 6.951 | 7.2 | 563 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 476424 | 0.000 | 0.000 | 0.090 | 0.169 | 1.638 | 2.1 | 62 | 0 |
| valkey-go/docache | 10 | 2176646 | 0.001 | 0.001 | 0.122 | 0.255 | 5.987 | 2.4 | 188 | 0 |
| valkey-go/docache | 100 | 3243846 | 0.001 | 0.072 | 0.431 | 2.433 | 9.127 | 3.6 | 354 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18016 | 0.078 | 0.116 | 0.169 | 0.378 | 1.433 | 51.4 | 68 | 0 |
| valkey-go/docache | 10 | 87772 | 0.079 | 0.273 | 0.380 | 0.579 | 4.979 | 34.0 | 227 | 0 |
| valkey-go/docache | 100 | 186785 | 0.087 | 1.331 | 1.727 | 2.329 | 4.283 | 20.0 | 570 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 76486 | 0.001 | 0.082 | 0.140 | 0.235 | 1.570 | 12.4 | 64 | 0 |
| valkey-go/docache | 10 | 407779 | 0.001 | 0.099 | 0.298 | 0.430 | 3.523 | 8.3 | 235 | 0 |
| valkey-go/docache | 100 | 886369 | 0.001 | 0.272 | 1.451 | 2.129 | 9.983 | 7.5 | 590 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 464464 | 0.000 | 0.001 | 0.092 | 0.174 | 1.571 | 2.1 | 62 | 0 |
| valkey-go/docache | 10 | 2153902 | 0.001 | 0.001 | 0.124 | 0.262 | 6.227 | 2.4 | 172 | 0 |
| valkey-go/docache | 100 | 3185779 | 0.001 | 0.074 | 0.441 | 2.459 | 9.855 | 3.6 | 364 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15939 | 0.061 | 0.069 | 0.098 | 0.274 | 1.471 | 44.9 | 34 | 0 |
| valkey-go | 10 | 75108 | 0.128 | 0.182 | 0.244 | 0.465 | 1.437 | 43.3 | 40 | 0 |
| valkey-go | 100 | 343566 | 0.283 | 0.404 | 0.578 | 0.962 | 2.547 | 15.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15787 | 0.061 | 0.070 | 0.102 | 0.290 | 1.475 | 46.1 | 35 | 0 |
| valkey-go | 10 | 75303 | 0.128 | 0.182 | 0.246 | 0.473 | 1.813 | 43.4 | 41 | 0 |
| valkey-go | 100 | 344006 | 0.284 | 0.399 | 0.558 | 0.905 | 2.587 | 15.7 | 68 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15538 | 0.062 | 0.071 | 0.105 | 0.312 | 1.957 | 46.4 | 35 | 0 |
| valkey-go | 10 | 73554 | 0.131 | 0.186 | 0.252 | 0.465 | 1.754 | 44.1 | 40 | 0 |
| valkey-go | 100 | 346279 | 0.281 | 0.398 | 0.562 | 0.962 | 2.289 | 15.9 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15360 | 0.062 | 0.075 | 0.110 | 0.312 | 1.566 | 47.8 | 37 | 0 |
| valkey-go | 10 | 73978 | 0.130 | 0.185 | 0.249 | 0.478 | 1.767 | 44.1 | 42 | 0 |
| valkey-go | 100 | 348285 | 0.281 | 0.395 | 0.537 | 0.871 | 2.349 | 15.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14951 | 0.063 | 0.076 | 0.116 | 0.337 | 1.637 | 48.4 | 35 | 0 |
| valkey-go | 10 | 73915 | 0.130 | 0.185 | 0.252 | 0.479 | 1.830 | 43.8 | 41 | 0 |
| valkey-go | 100 | 338357 | 0.287 | 0.411 | 0.579 | 0.934 | 2.305 | 15.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14971 | 0.063 | 0.076 | 0.116 | 0.352 | 1.561 | 47.6 | 35 | 0 |
| valkey-go | 10 | 74014 | 0.130 | 0.185 | 0.251 | 0.469 | 1.446 | 44.3 | 39 | 0 |
| valkey-go | 100 | 298569 | 0.309 | 0.519 | 0.749 | 1.145 | 3.245 | 17.3 | 64 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 17854 | 0.007 | 0.120 | 0.179 | 0.380 | 1.929 | 52.8 | 72 | 0 |
| valkey-go/docache | 10 | 86203 | 0.082 | 0.278 | 0.390 | 0.606 | 3.383 | 34.3 | 222 | 0 |
| valkey-go/docache | 100 | 189672 | 0.088 | 1.307 | 1.726 | 2.363 | 6.827 | 20.2 | 580 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 77370 | 0.001 | 0.081 | 0.140 | 0.221 | 1.704 | 12.4 | 66 | 0 |
| valkey-go/docache | 10 | 418925 | 0.001 | 0.098 | 0.289 | 0.421 | 3.195 | 8.3 | 243 | 0 |
| valkey-go/docache | 100 | 934073 | 0.001 | 0.234 | 1.381 | 2.028 | 6.151 | 7.3 | 557 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 483122 | 0.000 | 0.000 | 0.091 | 0.165 | 2.065 | 2.0 | 62 | 0 |
| valkey-go/docache | 10 | 2142703 | 0.001 | 0.001 | 0.124 | 0.259 | 7.263 | 2.4 | 172 | 0 |
| valkey-go/docache | 100 | 3241049 | 0.001 | 0.072 | 0.438 | 2.407 | 12.359 | 3.5 | 380 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 17799 | 0.081 | 0.116 | 0.174 | 0.380 | 1.529 | 52.0 | 66 | 0 |
| valkey-go/docache | 10 | 86257 | 0.078 | 0.278 | 0.405 | 0.605 | 3.679 | 34.2 | 226 | 0 |
| valkey-go/docache | 100 | 190045 | 0.087 | 1.302 | 1.717 | 2.445 | 5.415 | 20.5 | 549 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 78811 | 0.001 | 0.082 | 0.136 | 0.219 | 1.621 | 12.2 | 66 | 0 |
| valkey-go/docache | 10 | 407485 | 0.001 | 0.099 | 0.298 | 0.433 | 3.697 | 8.3 | 239 | 0 |
| valkey-go/docache | 100 | 904616 | 0.001 | 0.244 | 1.445 | 2.107 | 5.279 | 7.5 | 548 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 469300 | 0.000 | 0.001 | 0.093 | 0.170 | 1.376 | 2.1 | 64 | 0 |
| valkey-go/docache | 10 | 2143832 | 0.001 | 0.001 | 0.124 | 0.260 | 5.563 | 2.4 | 176 | 0 |
| valkey-go/docache | 100 | 3178294 | 0.001 | 0.075 | 0.443 | 2.421 | 11.759 | 3.6 | 362 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15323 | 0.062 | 0.074 | 0.106 | 0.299 | 1.536 | 47.8 | 35 | 0 |
| valkey-go | 10 | 74208 | 0.130 | 0.184 | 0.247 | 0.461 | 1.733 | 43.8 | 39 | 0 |
| valkey-go | 100 | 343903 | 0.285 | 0.400 | 0.540 | 0.836 | 2.225 | 15.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15588 | 0.062 | 0.070 | 0.100 | 0.293 | 1.825 | 46.1 | 35 | 0 |
| valkey-go | 10 | 73887 | 0.130 | 0.185 | 0.248 | 0.474 | 3.119 | 43.7 | 41 | 0 |
| valkey-go | 100 | 346402 | 0.283 | 0.395 | 0.541 | 0.897 | 3.459 | 15.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14986 | 0.063 | 0.075 | 0.106 | 0.322 | 1.509 | 49.4 | 35 | 0 |
| valkey-go | 10 | 74439 | 0.129 | 0.184 | 0.249 | 0.471 | 1.604 | 43.5 | 41 | 0 |
| valkey-go | 100 | 350478 | 0.279 | 0.391 | 0.543 | 0.944 | 2.205 | 15.5 | 65 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15472 | 0.062 | 0.071 | 0.100 | 0.283 | 1.796 | 45.5 | 34 | 0 |
| valkey-go | 10 | 74702 | 0.129 | 0.184 | 0.248 | 0.470 | 1.722 | 42.9 | 40 | 0 |
| valkey-go | 100 | 350905 | 0.280 | 0.390 | 0.531 | 0.822 | 2.131 | 15.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14745 | 0.063 | 0.077 | 0.113 | 0.367 | 4.719 | 49.2 | 37 | 0 |
| valkey-go | 10 | 73905 | 0.130 | 0.186 | 0.252 | 0.464 | 1.139 | 44.0 | 41 | 0 |
| valkey-go | 100 | 349004 | 0.281 | 0.394 | 0.541 | 0.893 | 2.223 | 15.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15818 | 0.061 | 0.069 | 0.099 | 0.285 | 1.493 | 45.6 | 37 | 0 |
| valkey-go | 10 | 73860 | 0.130 | 0.185 | 0.250 | 0.488 | 1.720 | 43.2 | 40 | 0 |
| valkey-go | 100 | 345720 | 0.283 | 0.397 | 0.551 | 0.938 | 2.547 | 15.6 | 66 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 50% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 18763 | 0.006 | 0.112 | 0.164 | 0.360 | 1.535 | 50.2 | 72 | 0 |
| valkey-go/docache | 10 | 87372 | 0.080 | 0.274 | 0.384 | 0.595 | 3.593 | 34.1 | 226 | 0 |
| valkey-go/docache | 100 | 189054 | 0.087 | 1.312 | 1.725 | 2.215 | 5.855 | 19.8 | 578 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 90% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 71606 | 0.001 | 0.080 | 0.182 | 0.303 | 1.853 | 13.3 | 62 | 0 |
| valkey-go/docache | 10 | 416260 | 0.001 | 0.098 | 0.293 | 0.425 | 6.051 | 8.3 | 237 | 0 |
| valkey-go/docache | 100 | 891701 | 0.001 | 0.250 | 1.460 | 2.183 | 11.975 | 7.4 | 598 | 0 |
mode=standalone fixed concurrency, workload=GETs only, 99% of them to 1000 hot keys (the cache hit test) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 462769 | 0.000 | 0.001 | 0.094 | 0.173 | 1.677 | 2.1 | 64 | 0 |
| valkey-go/docache | 10 | 2164926 | 0.001 | 0.001 | 0.123 | 0.256 | 3.445 | 2.4 | 176 | 0 |
| valkey-go/docache | 100 | 3213333 | 0.001 | 0.073 | 0.434 | 2.497 | 12.143 | 3.6 | 383 | 0 |
