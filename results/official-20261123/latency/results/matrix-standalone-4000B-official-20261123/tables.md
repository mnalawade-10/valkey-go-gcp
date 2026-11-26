mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13097 | 0.072 | 0.086 | 0.153 | 0.389 | 1.560 | 59.4 | 27 | 0 |
| go-redis | 10 | 45631 | 0.212 | 0.283 | 0.545 | 0.845 | 1.813 | 73.4 | 32 | 0 |
| go-redis | 100 | 48448 | 1.981 | 2.249 | 3.941 | 4.475 | 6.851 | 72.8 | 71 | 0 |
| go-redis | 1000 | 48818 | 20.367 | 21.567 | 23.599 | 24.831 | 26.367 | 84.0 | 333 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.141 | 0.211 | 0.385 | 0.829 | 1.574 | 1.125 | 1.329 | 130.5 | 278 | 0 |
| go-redis | 10000 | 10000 | 0.147 | 0.292 | 0.498 | 0.833 | 1.955 | 1.226 | 1.505 | 88.2 | 442 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8816 | 0.103 | 0.139 | 0.250 | 0.502 | 1.935 | 116.5 | 29 | 0 |
| go-redis-autopipeline | 10 | 45826 | 0.198 | 0.282 | 0.591 | 0.871 | 1.989 | 27.7 | 32 | 0 |
| go-redis-autopipeline | 100 | 128429 | 0.749 | 0.918 | 1.552 | 1.975 | 2.807 | 12.7 | 58 | 0 |
| go-redis-autopipeline | 1000 | 192627 | 5.255 | 5.895 | 7.459 | 8.983 | 10.551 | 9.8 | 310 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 0.895 | 1.524 | 2.671 | 2.917 | 3.967 | 1.114 | 3.365 | 129.4 | 282 | 0 |
| go-redis-autopipeline | 10000 | 9998 | 0.436 | 1.613 | 2.663 | 3.031 | 3.405 | 1.118 | 2.797 | 36.0 | 448 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7352 | 0.126 | 0.163 | 0.281 | 0.495 | 1.624 | 126.6 | 34 | 0 |
| glide | 10 | 27693 | 0.345 | 0.473 | 0.666 | 1.054 | 1.968 | 68.4 | 37 | 0 |
| glide | 100 | 29252 | 3.297 | 5.031 | 6.995 | 8.035 | 9.655 | 67.3 | 66 | 0 |
| glide | 1000 | 30205 | 33.023 | 36.831 | 40.095 | 43.839 | 45.983 | 66.7 | 316 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.225 | 0.302 | 0.527 | 1.152 | 1.990 | 0.901 | 1.254 | 178.4 | 278 | 0 |
| glide | 10000 | 10000 | 0.149 | 0.202 | 0.429 | 0.916 | 2.147 | 0.307 | 0.673 | 129.0 | 452 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14024 | 0.067 | 0.081 | 0.136 | 0.360 | 1.585 | 51.6 | 35 | 0 |
| valkey-go | 10 | 62019 | 0.155 | 0.218 | 0.338 | 0.566 | 1.739 | 49.7 | 39 | 0 |
| valkey-go | 100 | 224172 | 0.433 | 0.581 | 1.023 | 1.515 | 2.345 | 20.5 | 65 | 0 |
| valkey-go | 1000 | 259524 | 3.775 | 6.299 | 8.455 | 10.279 | 11.567 | 18.7 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.153 | 0.221 | 0.387 | 0.691 | 1.824 | 1.120 | 1.326 | 146.0 | 272 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.214 | 0.370 | 0.733 | 1.938 | 1.155 | 1.377 | 100.3 | 453 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8945 | 0.101 | 0.139 | 0.255 | 0.479 | 1.557 | 116.3 | 30 | 0 |
| go-redis-autopipeline | 10 | 46419 | 0.196 | 0.276 | 0.579 | 0.895 | 1.734 | 27.8 | 32 | 0 |
| go-redis-autopipeline | 100 | 131176 | 0.732 | 0.898 | 1.516 | 2.001 | 2.509 | 12.4 | 57 | 0 |
| go-redis-autopipeline | 1000 | 192592 | 5.259 | 5.935 | 7.471 | 9.167 | 9.927 | 9.6 | 308 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.279 | 1.520 | 2.663 | 2.849 | 3.721 | 1.111 | 3.331 | 122.7 | 276 | 0 |
| go-redis-autopipeline | 10000 | 9996 | 0.442 | 1.630 | 2.689 | 2.985 | 3.929 | 1.119 | 2.835 | 35.8 | 434 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7624 | 0.122 | 0.157 | 0.270 | 0.506 | 1.739 | 122.9 | 32 | 0 |
| glide | 10 | 28162 | 0.341 | 0.454 | 0.654 | 1.043 | 2.979 | 67.3 | 38 | 0 |
| glide | 100 | 31633 | 3.257 | 3.769 | 4.683 | 5.531 | 8.983 | 63.4 | 64 | 0 |
| glide | 1000 | 30840 | 32.319 | 36.095 | 40.479 | 43.615 | 46.047 | 65.5 | 314 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.223 | 0.308 | 0.554 | 1.262 | 1.874 | 0.902 | 1.238 | 175.9 | 272 | 0 |
| glide | 10000 | 10000 | 0.140 | 0.195 | 0.459 | 1.208 | 2.749 | 0.373 | 0.822 | 125.0 | 450 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14044 | 0.067 | 0.080 | 0.124 | 0.352 | 2.137 | 51.9 | 36 | 0 |
| valkey-go | 10 | 61479 | 0.156 | 0.221 | 0.340 | 0.563 | 1.759 | 49.7 | 41 | 0 |
| valkey-go | 100 | 219466 | 0.444 | 0.589 | 0.995 | 1.497 | 3.045 | 20.1 | 66 | 0 |
| valkey-go | 1000 | 255734 | 3.651 | 6.683 | 10.671 | 12.591 | 15.535 | 18.5 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.148 | 0.211 | 0.371 | 0.555 | 1.192 | 1.115 | 1.311 | 134.6 | 262 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.207 | 0.360 | 0.588 | 1.734 | 1.140 | 1.357 | 101.6 | 450 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12937 | 0.072 | 0.089 | 0.144 | 0.393 | 1.571 | 61.2 | 27 | 0 |
| go-redis | 10 | 46010 | 0.211 | 0.281 | 0.530 | 0.825 | 1.728 | 70.9 | 33 | 0 |
| go-redis | 100 | 48804 | 1.976 | 2.189 | 3.941 | 4.427 | 6.975 | 69.3 | 71 | 0 |
| go-redis | 1000 | 48257 | 20.559 | 22.143 | 24.495 | 26.767 | 29.695 | 81.5 | 332 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.143 | 0.208 | 0.380 | 0.690 | 1.706 | 1.117 | 1.320 | 130.1 | 276 | 0 |
| go-redis | 10000 | 9999 | 0.151 | 0.295 | 0.484 | 0.708 | 1.591 | 1.206 | 1.491 | 83.6 | 419 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7378 | 0.126 | 0.162 | 0.281 | 0.517 | 1.752 | 125.7 | 33 | 0 |
| glide | 10 | 28233 | 0.338 | 0.467 | 0.653 | 0.961 | 2.337 | 66.7 | 36 | 0 |
| glide | 100 | 31573 | 3.237 | 3.831 | 4.951 | 6.083 | 7.695 | 62.4 | 63 | 0 |
| glide | 1000 | 30063 | 32.767 | 39.103 | 45.983 | 46.911 | 51.007 | 66.9 | 320 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.216 | 0.294 | 0.520 | 1.287 | 2.119 | 0.896 | 1.221 | 166.9 | 280 | 0 |
| glide | 10000 | 10000 | 0.145 | 0.200 | 0.444 | 0.923 | 2.279 | 0.363 | 0.780 | 125.9 | 439 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 11005 | 0.070 | 0.147 | 0.219 | 0.439 | 2.020 | 66.4 | 37 | 0 |
| valkey-go | 10 | 62036 | 0.155 | 0.218 | 0.335 | 0.572 | 2.653 | 49.7 | 41 | 0 |
| valkey-go | 100 | 222367 | 0.434 | 0.589 | 1.038 | 1.544 | 4.335 | 20.3 | 66 | 0 |
| valkey-go | 1000 | 254405 | 3.381 | 7.535 | 10.503 | 12.111 | 13.863 | 18.6 | 320 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.153 | 0.220 | 0.397 | 0.694 | 1.571 | 1.116 | 1.332 | 135.8 | 274 | 0 |
| valkey-go | 10000 | 10000 | 0.106 | 0.220 | 0.379 | 0.686 | 2.283 | 1.150 | 1.375 | 96.5 | 453 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13237 | 0.071 | 0.087 | 0.145 | 0.395 | 1.510 | 59.9 | 29 | 0 |
| go-redis | 10 | 45416 | 0.213 | 0.284 | 0.549 | 0.878 | 2.199 | 72.7 | 34 | 0 |
| go-redis | 100 | 47228 | 2.018 | 2.415 | 4.015 | 4.843 | 7.103 | 72.5 | 70 | 0 |
| go-redis | 1000 | 48605 | 20.447 | 21.791 | 24.239 | 25.631 | 27.855 | 82.5 | 333 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.143 | 0.210 | 0.390 | 0.602 | 1.394 | 1.124 | 1.331 | 128.7 | 276 | 0 |
| go-redis | 10000 | 9998 | 0.148 | 0.292 | 0.496 | 0.795 | 2.307 | 1.221 | 1.503 | 85.9 | 442 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8610 | 0.107 | 0.141 | 0.247 | 0.503 | 1.586 | 119.2 | 29 | 0 |
| go-redis-autopipeline | 10 | 45715 | 0.199 | 0.282 | 0.576 | 0.930 | 1.785 | 28.1 | 31 | 0 |
| go-redis-autopipeline | 100 | 127387 | 0.751 | 0.945 | 1.564 | 2.129 | 2.587 | 12.7 | 58 | 0 |
| go-redis-autopipeline | 1000 | 193651 | 5.251 | 5.875 | 7.247 | 8.623 | 8.935 | 9.7 | 309 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 0.631 | 1.514 | 2.619 | 2.957 | 3.829 | 1.107 | 3.209 | 133.8 | 276 | 0 |
| go-redis-autopipeline | 10000 | 9998 | 0.439 | 1.616 | 2.647 | 3.019 | 3.949 | 1.122 | 2.809 | 36.2 | 438 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14031 | 0.067 | 0.080 | 0.128 | 0.352 | 2.199 | 51.9 | 37 | 0 |
| valkey-go | 10 | 61242 | 0.156 | 0.222 | 0.344 | 0.569 | 1.919 | 49.7 | 39 | 0 |
| valkey-go | 100 | 213159 | 0.453 | 0.625 | 1.102 | 1.572 | 3.345 | 20.9 | 67 | 0 |
| valkey-go | 1000 | 255623 | 3.603 | 6.751 | 8.879 | 10.279 | 11.767 | 18.9 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.155 | 0.223 | 0.376 | 0.549 | 1.599 | 1.120 | 1.329 | 153.0 | 268 | 0 |
| valkey-go | 10000 | 9999 | 0.103 | 0.209 | 0.373 | 0.629 | 2.040 | 1.147 | 1.364 | 101.5 | 450 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13058 | 0.072 | 0.086 | 0.148 | 0.375 | 1.565 | 59.2 | 27 | 0 |
| go-redis | 10 | 45309 | 0.214 | 0.285 | 0.537 | 0.828 | 1.704 | 72.8 | 32 | 0 |
| go-redis | 100 | 48433 | 1.986 | 2.237 | 3.951 | 4.387 | 6.131 | 70.1 | 69 | 0 |
| go-redis | 1000 | 47183 | 20.591 | 23.599 | 30.511 | 34.399 | 38.559 | 83.5 | 331 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.153 | 0.232 | 0.482 | 1.928 | 3.243 | 1.134 | 1.406 | 145.4 | 278 | 0 |
| go-redis | 10000 | 9999 | 0.149 | 0.300 | 0.499 | 0.802 | 1.654 | 1.216 | 1.506 | 88.2 | 421 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8798 | 0.103 | 0.138 | 0.250 | 0.505 | 1.831 | 118.0 | 27 | 0 |
| go-redis-autopipeline | 10 | 44248 | 0.205 | 0.292 | 0.610 | 1.026 | 3.305 | 28.8 | 30 | 0 |
| go-redis-autopipeline | 100 | 130210 | 0.736 | 0.909 | 1.560 | 1.969 | 3.081 | 12.9 | 57 | 0 |
| go-redis-autopipeline | 1000 | 197038 | 5.143 | 5.731 | 7.435 | 9.799 | 12.007 | 9.9 | 308 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 0.514 | 1.524 | 2.651 | 2.867 | 3.735 | 1.113 | 3.175 | 138.0 | 284 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.441 | 1.629 | 2.657 | 3.001 | 4.171 | 1.118 | 2.819 | 36.7 | 445 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7222 | 0.128 | 0.168 | 0.294 | 0.512 | 1.540 | 128.8 | 31 | 0 |
| glide | 10 | 28277 | 0.338 | 0.453 | 0.661 | 0.964 | 1.986 | 67.9 | 38 | 0 |
| glide | 100 | 31956 | 3.237 | 3.665 | 4.687 | 5.543 | 6.635 | 63.3 | 64 | 0 |
| glide | 1000 | 31808 | 31.311 | 35.551 | 39.167 | 42.655 | 44.063 | 65.3 | 315 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.230 | 0.351 | 0.567 | 1.047 | 2.059 | 0.935 | 1.337 | 186.8 | 264 | 0 |
| glide | 10000 | 10000 | 0.147 | 0.202 | 0.444 | 1.137 | 2.087 | 0.323 | 0.717 | 127.0 | 447 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12968 | 0.072 | 0.088 | 0.149 | 0.396 | 1.449 | 59.0 | 27 | 0 |
| go-redis | 10 | 44591 | 0.217 | 0.290 | 0.548 | 0.870 | 2.437 | 73.3 | 32 | 0 |
| go-redis | 100 | 48420 | 1.987 | 2.221 | 3.945 | 4.371 | 5.951 | 71.5 | 71 | 0 |
| go-redis | 1000 | 48053 | 20.575 | 22.447 | 26.047 | 28.223 | 31.503 | 84.2 | 331 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.145 | 0.210 | 0.412 | 0.752 | 1.455 | 1.123 | 1.327 | 135.7 | 278 | 0 |
| go-redis | 10000 | 9998 | 0.146 | 0.296 | 0.514 | 1.066 | 2.061 | 1.218 | 1.503 | 88.3 | 422 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8883 | 0.102 | 0.137 | 0.243 | 0.502 | 1.637 | 117.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 45629 | 0.199 | 0.281 | 0.583 | 0.890 | 1.543 | 28.0 | 31 | 0 |
| go-redis-autopipeline | 100 | 129040 | 0.745 | 0.912 | 1.539 | 1.957 | 2.761 | 12.8 | 56 | 0 |
| go-redis-autopipeline | 1000 | 192314 | 5.251 | 5.963 | 7.775 | 9.127 | 11.351 | 10.0 | 308 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 0.652 | 1.514 | 2.649 | 2.843 | 3.703 | 1.111 | 3.221 | 130.8 | 268 | 0 |
| go-redis-autopipeline | 10000 | 9996 | 0.431 | 1.610 | 2.637 | 2.959 | 3.151 | 1.114 | 2.783 | 35.6 | 440 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7156 | 0.129 | 0.167 | 0.305 | 0.501 | 1.663 | 128.0 | 34 | 0 |
| glide | 10 | 27189 | 0.352 | 0.476 | 0.691 | 0.998 | 2.023 | 70.8 | 37 | 0 |
| glide | 100 | 29639 | 3.401 | 4.583 | 5.155 | 6.035 | 7.243 | 66.8 | 64 | 0 |
| glide | 1000 | 29877 | 33.727 | 37.407 | 39.967 | 41.599 | 43.551 | 68.2 | 309 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.227 | 0.309 | 0.548 | 1.015 | 1.686 | 0.896 | 1.245 | 180.4 | 260 | 0 |
| glide | 10000 | 10000 | 0.142 | 0.197 | 0.453 | 1.375 | 2.615 | 0.330 | 0.717 | 126.7 | 446 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14115 | 0.067 | 0.080 | 0.127 | 0.352 | 1.614 | 51.4 | 35 | 0 |
| valkey-go | 10 | 60377 | 0.159 | 0.225 | 0.343 | 0.569 | 1.640 | 50.1 | 39 | 0 |
| valkey-go | 100 | 220679 | 0.439 | 0.591 | 1.047 | 1.564 | 2.717 | 20.5 | 65 | 0 |
| valkey-go | 1000 | 265401 | 3.355 | 7.055 | 9.631 | 11.503 | 14.143 | 18.5 | 316 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.154 | 0.224 | 0.395 | 1.094 | 1.614 | 1.118 | 1.336 | 140.3 | 274 | 0 |
| valkey-go | 10000 | 10000 | 0.102 | 0.211 | 0.388 | 1.104 | 2.427 | 1.144 | 1.365 | 103.4 | 391 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8939 | 0.102 | 0.136 | 0.240 | 0.495 | 1.571 | 115.0 | 27 | 0 |
| go-redis-autopipeline | 10 | 46918 | 0.193 | 0.276 | 0.561 | 0.939 | 1.728 | 27.5 | 29 | 0 |
| go-redis-autopipeline | 100 | 130864 | 0.734 | 0.901 | 1.527 | 2.053 | 2.721 | 12.7 | 57 | 0 |
| go-redis-autopipeline | 1000 | 194711 | 5.191 | 5.887 | 7.663 | 9.695 | 10.831 | 9.9 | 308 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.243 | 1.504 | 2.643 | 2.859 | 3.997 | 1.116 | 3.229 | 128.0 | 262 | 0 |
| go-redis-autopipeline | 10000 | 9999 | 0.430 | 1.607 | 2.629 | 2.945 | 3.397 | 1.116 | 2.777 | 34.7 | 443 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7091 | 0.129 | 0.171 | 0.319 | 0.546 | 1.765 | 130.0 | 34 | 0 |
| glide | 10 | 27681 | 0.345 | 0.472 | 0.672 | 1.012 | 1.784 | 68.0 | 37 | 0 |
| glide | 100 | 30298 | 3.351 | 4.319 | 5.103 | 5.751 | 7.799 | 64.9 | 65 | 0 |
| glide | 1000 | 29604 | 33.791 | 37.343 | 39.487 | 44.799 | 47.775 | 68.4 | 319 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.216 | 0.289 | 0.525 | 1.146 | 1.632 | 0.891 | 1.210 | 166.9 | 268 | 0 |
| glide | 10000 | 10000 | 0.146 | 0.212 | 0.467 | 1.103 | 2.547 | 0.328 | 0.763 | 126.6 | 450 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14162 | 0.067 | 0.078 | 0.119 | 0.344 | 1.453 | 51.3 | 37 | 0 |
| valkey-go | 10 | 61851 | 0.155 | 0.218 | 0.341 | 0.575 | 1.788 | 50.2 | 39 | 0 |
| valkey-go | 100 | 170711 | 0.509 | 0.927 | 1.446 | 2.257 | 3.499 | 24.7 | 67 | 0 |
| valkey-go | 1000 | 255863 | 3.599 | 6.843 | 9.199 | 11.031 | 12.799 | 18.9 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.151 | 0.215 | 0.397 | 0.679 | 1.338 | 1.118 | 1.321 | 141.6 | 262 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.216 | 0.401 | 0.712 | 1.854 | 1.145 | 1.370 | 98.4 | 452 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12931 | 0.072 | 0.087 | 0.152 | 0.404 | 1.725 | 60.7 | 27 | 0 |
| go-redis | 10 | 45164 | 0.214 | 0.285 | 0.552 | 0.861 | 1.921 | 73.1 | 31 | 0 |
| go-redis | 100 | 48450 | 1.986 | 2.233 | 3.947 | 4.503 | 6.191 | 70.2 | 69 | 0 |
| go-redis | 1000 | 47952 | 20.655 | 22.687 | 25.471 | 28.639 | 32.447 | 82.3 | 331 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.141 | 0.205 | 0.390 | 0.580 | 1.705 | 1.118 | 1.319 | 122.1 | 276 | 0 |
| go-redis | 10000 | 9998 | 0.146 | 0.293 | 0.505 | 0.812 | 1.717 | 1.210 | 1.499 | 85.8 | 442 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7356 | 0.125 | 0.163 | 0.293 | 0.496 | 1.596 | 124.5 | 34 | 0 |
| glide | 10 | 27142 | 0.353 | 0.475 | 0.687 | 0.993 | 2.137 | 69.5 | 36 | 0 |
| glide | 100 | 31667 | 3.241 | 3.749 | 4.823 | 5.387 | 6.163 | 63.2 | 63 | 0 |
| glide | 1000 | 30565 | 32.431 | 36.223 | 43.231 | 44.479 | 48.575 | 66.4 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.223 | 0.301 | 0.530 | 1.105 | 1.727 | 0.897 | 1.223 | 174.3 | 272 | 0 |
| glide | 10000 | 10000 | 0.142 | 0.192 | 0.427 | 0.901 | 2.119 | 0.339 | 0.740 | 124.0 | 451 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14219 | 0.067 | 0.079 | 0.123 | 0.376 | 1.575 | 51.5 | 37 | 0 |
| valkey-go | 10 | 62268 | 0.154 | 0.217 | 0.336 | 0.563 | 1.680 | 49.6 | 38 | 0 |
| valkey-go | 100 | 220000 | 0.441 | 0.592 | 1.043 | 1.549 | 2.581 | 20.5 | 66 | 0 |
| valkey-go | 1000 | 255248 | 3.689 | 6.783 | 9.559 | 10.855 | 12.559 | 18.8 | 319 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.149 | 0.209 | 0.386 | 0.567 | 1.698 | 1.121 | 1.313 | 141.0 | 258 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.212 | 0.377 | 0.635 | 1.711 | 1.150 | 1.371 | 99.6 | 454 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13137 | 0.072 | 0.085 | 0.145 | 0.390 | 1.561 | 58.3 | 27 | 0 |
| go-redis | 10 | 44828 | 0.216 | 0.288 | 0.541 | 0.842 | 1.992 | 71.3 | 32 | 0 |
| go-redis | 100 | 48226 | 1.990 | 2.273 | 3.955 | 4.523 | 6.303 | 70.2 | 69 | 0 |
| go-redis | 1000 | 48504 | 20.479 | 21.791 | 23.903 | 25.391 | 27.567 | 82.6 | 331 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.143 | 0.215 | 0.395 | 0.615 | 1.196 | 1.120 | 1.335 | 128.2 | 274 | 0 |
| go-redis | 10000 | 9998 | 0.143 | 0.290 | 0.501 | 0.883 | 1.834 | 1.218 | 1.503 | 89.4 | 442 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8900 | 0.102 | 0.137 | 0.240 | 0.487 | 1.638 | 116.6 | 27 | 0 |
| go-redis-autopipeline | 10 | 45705 | 0.199 | 0.286 | 0.588 | 0.915 | 2.179 | 27.9 | 30 | 0 |
| go-redis-autopipeline | 100 | 130077 | 0.738 | 0.907 | 1.517 | 2.027 | 3.063 | 12.6 | 57 | 0 |
| go-redis-autopipeline | 1000 | 192804 | 5.247 | 5.895 | 7.591 | 9.143 | 10.407 | 9.8 | 307 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.258 | 1.508 | 2.625 | 2.821 | 3.469 | 1.104 | 3.175 | 125.9 | 276 | 0 |
| go-redis-autopipeline | 10000 | 9999 | 0.432 | 1.615 | 2.671 | 3.013 | 4.295 | 1.120 | 2.819 | 36.0 | 427 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14192 | 0.067 | 0.080 | 0.125 | 0.341 | 1.584 | 52.4 | 36 | 0 |
| valkey-go | 10 | 63024 | 0.152 | 0.215 | 0.327 | 0.562 | 1.523 | 49.5 | 39 | 0 |
| valkey-go | 100 | 220904 | 0.440 | 0.588 | 1.026 | 1.510 | 2.497 | 20.6 | 67 | 0 |
| valkey-go | 1000 | 261304 | 3.679 | 6.379 | 8.783 | 11.479 | 14.135 | 18.6 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.150 | 0.217 | 0.404 | 0.831 | 1.196 | 1.123 | 1.321 | 143.7 | 270 | 0 |
| valkey-go | 10000 | 9999 | 0.103 | 0.212 | 0.346 | 0.545 | 1.849 | 1.147 | 1.366 | 96.8 | 448 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12267 | 0.072 | 0.101 | 0.200 | 0.399 | 2.981 | 63.0 | 27 | 0 |
| go-redis | 10 | 45462 | 0.213 | 0.283 | 0.541 | 0.838 | 1.798 | 71.9 | 31 | 0 |
| go-redis | 100 | 47754 | 2.015 | 2.257 | 4.003 | 4.471 | 6.383 | 71.5 | 71 | 0 |
| go-redis | 1000 | 48541 | 20.415 | 21.871 | 24.335 | 26.303 | 29.327 | 83.7 | 331 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.142 | 0.218 | 0.398 | 0.708 | 1.505 | 1.125 | 1.347 | 124.5 | 278 | 0 |
| go-redis | 10000 | 10000 | 0.147 | 0.290 | 0.493 | 0.765 | 1.817 | 1.214 | 1.494 | 85.1 | 430 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9097 | 0.101 | 0.132 | 0.234 | 0.488 | 1.789 | 114.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 45045 | 0.201 | 0.289 | 0.609 | 0.966 | 1.843 | 28.5 | 31 | 0 |
| go-redis-autopipeline | 100 | 130763 | 0.734 | 0.906 | 1.541 | 1.950 | 2.775 | 12.8 | 58 | 0 |
| go-redis-autopipeline | 1000 | 196154 | 5.163 | 5.839 | 7.671 | 9.799 | 13.535 | 10.0 | 308 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.245 | 1.517 | 2.643 | 2.821 | 3.207 | 1.114 | 3.241 | 122.7 | 284 | 0 |
| go-redis-autopipeline | 10000 | 9996 | 0.436 | 1.615 | 2.697 | 3.057 | 4.411 | 1.118 | 2.823 | 35.8 | 428 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7000 | 0.133 | 0.173 | 0.301 | 0.522 | 1.602 | 131.7 | 32 | 0 |
| glide | 10 | 27937 | 0.342 | 0.468 | 0.661 | 1.071 | 2.309 | 69.5 | 36 | 0 |
| glide | 100 | 31563 | 3.231 | 3.771 | 5.035 | 5.503 | 7.187 | 63.2 | 63 | 0 |
| glide | 1000 | 29994 | 33.119 | 37.087 | 42.559 | 47.295 | 66.431 | 67.0 | 319 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.223 | 0.307 | 0.534 | 1.135 | 1.537 | 0.891 | 1.227 | 176.5 | 278 | 0 |
| glide | 10000 | 10000 | 0.144 | 0.199 | 0.449 | 1.130 | 1.688 | 0.352 | 0.788 | 126.9 | 449 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13478 | 0.070 | 0.084 | 0.143 | 0.388 | 1.702 | 58.3 | 28 | 0 |
| go-redis | 10 | 45041 | 0.215 | 0.288 | 0.547 | 0.846 | 1.935 | 74.1 | 31 | 0 |
| go-redis | 100 | 47808 | 2.007 | 2.283 | 3.989 | 4.499 | 6.399 | 71.1 | 68 | 0 |
| go-redis | 1000 | 47904 | 20.687 | 22.127 | 24.399 | 26.063 | 28.447 | 83.2 | 329 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 1999 | 0.142 | 0.211 | 0.395 | 1.040 | 1.404 | 1.126 | 1.347 | 131.6 | 278 | 0 |
| go-redis | 10000 | 9998 | 0.146 | 0.289 | 0.514 | 1.168 | 1.850 | 1.219 | 1.508 | 86.9 | 434 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9006 | 0.101 | 0.134 | 0.242 | 0.529 | 1.639 | 115.2 | 29 | 0 |
| go-redis-autopipeline | 10 | 46564 | 0.196 | 0.273 | 0.572 | 0.941 | 1.741 | 27.5 | 31 | 0 |
| go-redis-autopipeline | 100 | 128843 | 0.745 | 0.925 | 1.537 | 1.973 | 3.899 | 12.9 | 56 | 0 |
| go-redis-autopipeline | 1000 | 194409 | 5.215 | 5.907 | 7.435 | 8.839 | 10.951 | 9.9 | 310 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.267 | 1.505 | 2.641 | 2.855 | 3.821 | 1.112 | 3.307 | 115.4 | 264 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.426 | 1.609 | 2.651 | 2.961 | 3.653 | 1.116 | 2.801 | 34.8 | 435 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7220 | 0.129 | 0.167 | 0.293 | 0.543 | 2.137 | 128.0 | 34 | 0 |
| glide | 10 | 29101 | 0.332 | 0.432 | 0.632 | 0.924 | 2.809 | 67.5 | 36 | 0 |
| glide | 100 | 32055 | 3.189 | 3.729 | 4.979 | 5.731 | 7.467 | 63.2 | 64 | 0 |
| glide | 1000 | 30807 | 32.511 | 36.095 | 38.783 | 40.479 | 41.887 | 66.5 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.221 | 0.298 | 0.529 | 1.150 | 2.267 | 0.900 | 1.232 | 175.3 | 260 | 0 |
| glide | 10000 | 10000 | 0.147 | 0.204 | 0.442 | 1.096 | 2.141 | 0.356 | 0.779 | 123.7 | 444 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14207 | 0.067 | 0.079 | 0.123 | 0.354 | 1.521 | 51.2 | 35 | 0 |
| valkey-go | 10 | 62047 | 0.154 | 0.219 | 0.347 | 0.573 | 1.685 | 50.0 | 39 | 0 |
| valkey-go | 100 | 222573 | 0.435 | 0.586 | 1.041 | 1.499 | 2.655 | 20.6 | 67 | 0 |
| valkey-go | 1000 | 256198 | 3.519 | 7.051 | 9.887 | 11.031 | 12.167 | 18.9 | 317 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.166 | 0.285 | 0.479 | 1.175 | 2.161 | 1.122 | 1.442 | 162.5 | 266 | 0 |
| valkey-go | 10000 | 10000 | 0.105 | 0.224 | 0.383 | 0.646 | 1.624 | 1.147 | 1.379 | 97.5 | 452 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 8758 | 0.103 | 0.142 | 0.254 | 0.538 | 1.757 | 117.4 | 29 | 0 |
| go-redis-autopipeline | 10 | 44577 | 0.203 | 0.294 | 0.620 | 0.956 | 1.700 | 28.6 | 32 | 0 |
| go-redis-autopipeline | 100 | 130732 | 0.731 | 0.913 | 1.538 | 2.024 | 3.151 | 12.8 | 57 | 0 |
| go-redis-autopipeline | 1000 | 195934 | 5.155 | 5.807 | 8.239 | 9.903 | 10.671 | 10.1 | 307 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.244 | 1.510 | 2.625 | 2.879 | 3.725 | 1.105 | 3.227 | 122.4 | 278 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.439 | 1.620 | 2.683 | 3.003 | 3.279 | 1.119 | 2.825 | 35.9 | 442 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7365 | 0.125 | 0.164 | 0.301 | 0.512 | 1.708 | 127.3 | 34 | 0 |
| glide | 10 | 27951 | 0.344 | 0.453 | 0.659 | 1.088 | 1.989 | 68.6 | 37 | 0 |
| glide | 100 | 30823 | 3.281 | 4.151 | 5.079 | 5.795 | 7.591 | 64.1 | 65 | 0 |
| glide | 1000 | 29536 | 33.119 | 39.231 | 46.271 | 47.615 | 49.407 | 67.9 | 315 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.227 | 0.308 | 0.485 | 1.002 | 1.684 | 0.893 | 1.212 | 180.0 | 270 | 0 |
| glide | 10000 | 10000 | 0.146 | 0.199 | 0.404 | 0.964 | 1.797 | 0.284 | 0.670 | 127.4 | 448 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14101 | 0.067 | 0.080 | 0.136 | 0.298 | 1.735 | 51.1 | 37 | 0 |
| valkey-go | 10 | 61769 | 0.155 | 0.221 | 0.347 | 0.574 | 1.724 | 50.2 | 38 | 0 |
| valkey-go | 100 | 220023 | 0.440 | 0.600 | 1.075 | 1.575 | 3.609 | 20.8 | 68 | 0 |
| valkey-go | 1000 | 257107 | 3.519 | 7.187 | 10.535 | 12.367 | 13.407 | 18.9 | 318 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.145 | 0.210 | 0.342 | 1.282 | 1.549 | 1.109 | 1.317 | 131.2 | 268 | 0 |
| valkey-go | 10000 | 9998 | 0.103 | 0.209 | 0.336 | 0.549 | 1.955 | 1.146 | 1.359 | 98.8 | 448 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13100 | 0.072 | 0.087 | 0.147 | 0.306 | 1.591 | 59.5 | 28 | 0 |
| go-redis | 10 | 45437 | 0.214 | 0.283 | 0.526 | 0.859 | 1.817 | 72.5 | 32 | 0 |
| go-redis | 100 | 48132 | 2.000 | 2.205 | 3.983 | 4.515 | 6.743 | 71.7 | 70 | 0 |
| go-redis | 1000 | 48389 | 20.527 | 21.823 | 23.807 | 26.463 | 28.975 | 83.2 | 330 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=4000 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.140 | 0.203 | 0.326 | 0.630 | 1.054 | 1.117 | 1.304 | 122.7 | 278 | 0 |
| go-redis | 10000 | 10000 | 0.149 | 0.295 | 0.484 | 0.913 | 2.149 | 1.214 | 1.497 | 85.7 | 422 | 0 |
