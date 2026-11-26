mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14783 | 0.065 | 0.075 | 0.111 | 0.225 | 1.627 | 51.9 | 30 | 0 |
| go-redis | 10 | 53201 | 0.187 | 0.242 | 0.324 | 0.563 | 1.683 | 59.6 | 35 | 0 |
| go-redis | 100 | 54059 | 1.779 | 1.946 | 3.545 | 4.087 | 6.927 | 66.1 | 70 | 0 |
| go-redis | 1000 | 54814 | 18.079 | 19.007 | 20.943 | 22.079 | 24.079 | 77.4 | 256 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12094 | 0.077 | 0.094 | 0.159 | 0.390 | 1.624 | 63.0 | 29 | 0 |
| go-redis | 10 | 38954 | 0.252 | 0.325 | 0.533 | 0.981 | 2.385 | 79.1 | 36 | 0 |
| go-redis | 100 | 40122 | 2.391 | 2.595 | 4.763 | 5.243 | 7.379 | 81.0 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10001 | 0.092 | 0.124 | 0.192 | 0.392 | 2.753 | 102.7 | 29 | 0 |
| go-redis-autopipeline | 10 | 59047 | 0.158 | 0.211 | 0.344 | 0.828 | 4.867 | 21.3 | 32 | 0 |
| go-redis-autopipeline | 100 | 211348 | 0.465 | 0.559 | 0.706 | 1.453 | 2.075 | 8.0 | 58 | 0 |
| go-redis-autopipeline | 1000 | 483877 | 2.069 | 2.663 | 3.265 | 4.967 | 8.015 | 6.5 | 307 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12014 | 0.077 | 0.095 | 0.162 | 0.367 | 1.490 | 63.9 | 28 | 0 |
| go-redis-autopipeline | 10 | 38332 | 0.254 | 0.334 | 0.548 | 0.937 | 1.961 | 79.4 | 36 | 0 |
| go-redis-autopipeline | 100 | 39751 | 2.391 | 2.853 | 4.771 | 5.755 | 7.407 | 81.6 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7406 | 0.131 | 0.166 | 0.254 | 0.451 | 1.729 | 120.2 | 32 | 0 |
| glide | 10 | 34031 | 0.282 | 0.373 | 0.525 | 0.857 | 1.965 | 54.8 | 38 | 0 |
| glide | 100 | 36251 | 2.855 | 3.179 | 4.001 | 4.687 | 6.883 | 55.1 | 65 | 0 |
| glide | 1000 | 36074 | 27.871 | 30.783 | 37.343 | 42.239 | 44.991 | 56.7 | 236 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6674 | 0.138 | 0.180 | 0.313 | 0.595 | 1.541 | 123.2 | 33 | 0 |
| glide | 10 | 26318 | 0.359 | 0.532 | 0.921 | 1.498 | 2.103 | 90.5 | 36 | 0 |
| glide | 100 | 29464 | 3.381 | 4.355 | 5.431 | 6.411 | 8.575 | 83.3 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12952 | 0.074 | 0.084 | 0.124 | 0.250 | 1.567 | 48.6 | 35 | 0 |
| valkey-go+tls | 10 | 36102 | 0.259 | 0.430 | 0.572 | 0.754 | 1.930 | 62.3 | 42 | 0 |
| valkey-go+tls | 100 | 37186 | 2.267 | 5.615 | 7.955 | 9.047 | 10.791 | 63.6 | 68 | 0 |
| valkey-go+tls | 1000 | 37339 | 26.575 | 41.695 | 51.519 | 54.463 | 55.231 | 64.2 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10769 | 0.089 | 0.104 | 0.157 | 0.314 | 1.553 | 61.3 | 37 | 0 |
| valkey-go+tls | 10 | 28310 | 0.332 | 0.541 | 0.714 | 0.901 | 1.887 | 80.8 | 40 | 0 |
| valkey-go+tls | 100 | 29229 | 3.049 | 6.719 | 10.119 | 12.119 | 16.671 | 81.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12191 | 0.078 | 0.093 | 0.133 | 0.259 | 1.573 | 54.6 | 28 | 0 |
| go-redis+tls | 10 | 37201 | 0.267 | 0.335 | 0.440 | 0.644 | 1.429 | 65.1 | 33 | 0 |
| go-redis+tls | 100 | 37567 | 2.547 | 2.835 | 5.067 | 5.775 | 8.767 | 64.3 | 71 | 0 |
| go-redis+tls | 1000 | 37497 | 26.383 | 28.863 | 33.119 | 35.999 | 39.231 | 74.7 | 242 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9882 | 0.095 | 0.116 | 0.192 | 0.440 | 1.656 | 72.8 | 29 | 0 |
| go-redis+tls | 10 | 29164 | 0.339 | 0.424 | 0.648 | 1.053 | 2.047 | 84.2 | 36 | 0 |
| go-redis+tls | 100 | 30004 | 3.183 | 3.633 | 6.315 | 7.323 | 9.055 | 79.9 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 8371 | 0.110 | 0.141 | 0.226 | 0.431 | 1.612 | 108.3 | 29 | 0 |
| go-redis-autopipeline+tls | 10 | 53705 | 0.175 | 0.230 | 0.378 | 0.825 | 1.602 | 21.1 | 32 | 0 |
| go-redis-autopipeline+tls | 100 | 197676 | 0.496 | 0.598 | 0.774 | 1.492 | 2.339 | 7.9 | 60 | 0 |
| go-redis-autopipeline+tls | 1000 | 455426 | 2.163 | 2.835 | 3.453 | 5.115 | 7.591 | 6.6 | 311 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 10035 | 0.094 | 0.111 | 0.184 | 0.416 | 1.727 | 72.1 | 30 | 0 |
| go-redis-autopipeline+tls | 10 | 29018 | 0.341 | 0.428 | 0.624 | 1.017 | 1.945 | 84.1 | 38 | 0 |
| go-redis-autopipeline+tls | 100 | 30419 | 3.165 | 3.425 | 6.283 | 6.795 | 8.471 | 80.5 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6809 | 0.137 | 0.175 | 0.290 | 0.483 | 1.605 | 123.1 | 37 | 0 |
| glide+tls | 10 | 30715 | 0.317 | 0.436 | 0.592 | 0.823 | 1.810 | 59.4 | 39 | 0 |
| glide+tls | 100 | 33325 | 2.973 | 3.373 | 4.459 | 5.119 | 5.623 | 56.3 | 66 | 0 |
| glide+tls | 1000 | 32649 | 30.303 | 33.855 | 40.895 | 45.375 | 45.887 | 60.7 | 238 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5886 | 0.156 | 0.203 | 0.361 | 0.627 | 2.089 | 129.1 | 33 | 0 |
| glide+tls | 10 | 23430 | 0.400 | 0.583 | 0.958 | 1.460 | 3.327 | 95.9 | 36 | 0 |
| glide+tls | 100 | 25979 | 3.837 | 4.527 | 5.595 | 6.331 | 7.431 | 87.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15467 | 0.062 | 0.073 | 0.102 | 0.226 | 1.444 | 47.8 | 34 | 0 |
| valkey-go/flush-100us | 10 | 61150 | 0.157 | 0.239 | 0.313 | 1.334 | 1.810 | 35.6 | 41 | 0 |
| valkey-go/flush-100us | 100 | 417155 | 0.231 | 0.329 | 0.455 | 0.657 | 2.105 | 8.7 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 977414 | 1.041 | 1.408 | 1.776 | 2.909 | 5.387 | 5.5 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13277 | 0.071 | 0.085 | 0.129 | 0.264 | 1.541 | 52.4 | 35 | 0 |
| valkey-go/flush-100us | 10 | 45497 | 0.212 | 0.317 | 0.428 | 0.696 | 2.115 | 53.5 | 40 | 0 |
| valkey-go/flush-100us | 100 | 109972 | 0.927 | 1.165 | 1.648 | 2.219 | 3.975 | 22.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12830 | 0.075 | 0.086 | 0.122 | 0.246 | 1.562 | 48.4 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 49828 | 0.195 | 0.287 | 0.365 | 0.506 | 1.776 | 44.5 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 229190 | 0.406 | 0.625 | 0.919 | 1.643 | 2.117 | 13.6 | 69 | 0 |
| valkey-go/flush-100us+tls | 1000 | 226338 | 4.455 | 6.419 | 8.035 | 9.367 | 10.487 | 14.3 | 277 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10702 | 0.089 | 0.106 | 0.163 | 0.327 | 2.761 | 61.4 | 36 | 0 |
| valkey-go/flush-100us+tls | 10 | 34985 | 0.278 | 0.385 | 0.515 | 0.763 | 1.927 | 66.8 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 45567 | 1.923 | 4.091 | 5.967 | 6.663 | 8.295 | 54.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15245 | 0.062 | 0.075 | 0.119 | 0.240 | 1.593 | 47.6 | 36 | 0 |
| valkey-go | 10 | 73384 | 0.132 | 0.186 | 0.250 | 0.351 | 1.513 | 43.3 | 41 | 0 |
| valkey-go | 100 | 333960 | 0.296 | 0.409 | 0.531 | 0.834 | 2.361 | 15.4 | 67 | 0 |
| valkey-go | 1000 | 917225 | 1.126 | 1.459 | 1.767 | 2.853 | 6.675 | 6.7 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13321 | 0.071 | 0.082 | 0.128 | 0.261 | 1.634 | 51.9 | 37 | 0 |
| valkey-go | 10 | 48875 | 0.200 | 0.276 | 0.387 | 0.638 | 1.954 | 59.9 | 39 | 0 |
| valkey-go | 100 | 106230 | 0.970 | 1.224 | 1.655 | 2.417 | 3.597 | 28.2 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9846 | 0.093 | 0.125 | 0.194 | 0.410 | 2.321 | 105.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 61674 | 0.152 | 0.200 | 0.336 | 0.714 | 1.673 | 20.6 | 35 | 0 |
| go-redis-autopipeline | 100 | 212760 | 0.460 | 0.559 | 0.727 | 1.486 | 1.995 | 7.9 | 60 | 0 |
| go-redis-autopipeline | 1000 | 420808 | 2.395 | 3.299 | 4.419 | 6.487 | 8.807 | 7.1 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11694 | 0.078 | 0.100 | 0.171 | 0.396 | 1.664 | 69.8 | 27 | 0 |
| go-redis-autopipeline | 10 | 38524 | 0.253 | 0.331 | 0.569 | 0.983 | 1.879 | 81.8 | 34 | 0 |
| go-redis-autopipeline | 100 | 40301 | 2.375 | 2.633 | 4.747 | 5.571 | 7.475 | 84.2 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7242 | 0.131 | 0.168 | 0.299 | 0.548 | 3.135 | 126.3 | 33 | 0 |
| glide | 10 | 32481 | 0.297 | 0.389 | 0.548 | 0.840 | 2.022 | 56.9 | 36 | 0 |
| glide | 100 | 35832 | 2.875 | 3.235 | 4.271 | 4.935 | 6.423 | 57.0 | 63 | 0 |
| glide | 1000 | 33593 | 29.407 | 33.599 | 39.263 | 42.463 | 44.479 | 60.6 | 232 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6627 | 0.137 | 0.185 | 0.354 | 0.637 | 1.802 | 123.7 | 32 | 0 |
| glide | 10 | 25219 | 0.375 | 0.558 | 0.962 | 1.475 | 2.011 | 94.1 | 34 | 0 |
| glide | 100 | 29186 | 3.421 | 4.375 | 5.383 | 6.379 | 7.987 | 84.1 | 64 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12601 | 0.075 | 0.089 | 0.132 | 0.275 | 1.572 | 49.7 | 35 | 0 |
| valkey-go+tls | 10 | 35715 | 0.263 | 0.434 | 0.575 | 0.713 | 1.638 | 63.5 | 41 | 0 |
| valkey-go+tls | 100 | 37457 | 2.299 | 5.395 | 7.891 | 9.239 | 10.783 | 62.9 | 68 | 0 |
| valkey-go+tls | 1000 | 37190 | 20.479 | 56.063 | 70.847 | 73.535 | 74.367 | 63.9 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10474 | 0.091 | 0.107 | 0.163 | 0.350 | 1.654 | 62.8 | 38 | 0 |
| valkey-go+tls | 10 | 28417 | 0.332 | 0.536 | 0.711 | 0.938 | 1.857 | 80.2 | 43 | 0 |
| valkey-go+tls | 100 | 29539 | 2.829 | 7.259 | 9.791 | 11.151 | 12.095 | 81.3 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12182 | 0.078 | 0.093 | 0.134 | 0.261 | 1.662 | 55.0 | 31 | 0 |
| go-redis+tls | 10 | 37481 | 0.265 | 0.332 | 0.441 | 0.694 | 1.707 | 66.2 | 33 | 0 |
| go-redis+tls | 100 | 38423 | 2.511 | 2.693 | 4.987 | 5.343 | 8.863 | 66.9 | 73 | 0 |
| go-redis+tls | 1000 | 37372 | 26.543 | 28.959 | 32.239 | 35.487 | 40.095 | 78.8 | 230 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9860 | 0.095 | 0.115 | 0.204 | 0.477 | 1.659 | 72.9 | 30 | 0 |
| go-redis+tls | 10 | 28997 | 0.341 | 0.429 | 0.631 | 0.984 | 1.916 | 85.6 | 38 | 0 |
| go-redis+tls | 100 | 29811 | 3.211 | 3.541 | 6.375 | 7.171 | 9.807 | 82.9 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 8078 | 0.114 | 0.148 | 0.248 | 0.464 | 1.543 | 111.5 | 29 | 0 |
| go-redis-autopipeline+tls | 10 | 52184 | 0.179 | 0.238 | 0.404 | 0.890 | 1.962 | 21.8 | 34 | 0 |
| go-redis-autopipeline+tls | 100 | 189493 | 0.516 | 0.629 | 0.818 | 1.759 | 2.311 | 8.2 | 63 | 0 |
| go-redis-autopipeline+tls | 1000 | 424287 | 2.419 | 3.035 | 3.687 | 5.587 | 7.267 | 7.0 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9858 | 0.095 | 0.116 | 0.191 | 0.409 | 1.566 | 73.3 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 29200 | 0.340 | 0.425 | 0.628 | 0.983 | 2.028 | 84.4 | 36 | 0 |
| go-redis-autopipeline+tls | 100 | 29570 | 3.241 | 3.639 | 6.419 | 7.327 | 9.447 | 83.1 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6732 | 0.139 | 0.176 | 0.294 | 0.502 | 1.764 | 125.7 | 37 | 0 |
| glide+tls | 10 | 29709 | 0.327 | 0.453 | 0.632 | 0.966 | 2.043 | 61.3 | 40 | 0 |
| glide+tls | 100 | 31771 | 3.087 | 3.965 | 4.687 | 5.339 | 6.291 | 59.7 | 66 | 0 |
| glide+tls | 1000 | 32268 | 29.999 | 36.447 | 42.559 | 45.023 | 46.943 | 61.1 | 233 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5847 | 0.155 | 0.211 | 0.387 | 0.649 | 1.675 | 131.6 | 35 | 0 |
| glide+tls | 10 | 23703 | 0.397 | 0.574 | 0.944 | 1.435 | 2.779 | 95.8 | 37 | 0 |
| glide+tls | 100 | 25446 | 3.923 | 4.743 | 5.811 | 6.563 | 7.555 | 89.9 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15291 | 0.062 | 0.074 | 0.107 | 0.230 | 1.568 | 48.5 | 37 | 0 |
| valkey-go/flush-100us | 10 | 59911 | 0.160 | 0.245 | 0.325 | 1.333 | 1.736 | 37.2 | 41 | 0 |
| valkey-go/flush-100us | 100 | 415952 | 0.231 | 0.332 | 0.462 | 0.717 | 2.039 | 8.7 | 69 | 0 |
| valkey-go/flush-100us | 1000 | 986215 | 1.035 | 1.387 | 1.713 | 2.667 | 5.483 | 5.5 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13296 | 0.071 | 0.083 | 0.126 | 0.261 | 1.544 | 52.5 | 37 | 0 |
| valkey-go/flush-100us | 10 | 44823 | 0.215 | 0.323 | 0.441 | 0.704 | 1.826 | 54.1 | 41 | 0 |
| valkey-go/flush-100us | 100 | 108544 | 0.938 | 1.185 | 1.675 | 2.291 | 3.241 | 22.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12619 | 0.075 | 0.088 | 0.125 | 0.244 | 1.549 | 49.9 | 33 | 0 |
| valkey-go/flush-100us+tls | 10 | 49125 | 0.198 | 0.289 | 0.370 | 0.519 | 1.861 | 44.8 | 41 | 0 |
| valkey-go/flush-100us+tls | 100 | 208031 | 0.442 | 0.720 | 1.018 | 1.338 | 2.381 | 14.3 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 218296 | 3.509 | 9.615 | 11.519 | 12.159 | 13.383 | 14.8 | 264 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10664 | 0.090 | 0.104 | 0.158 | 0.309 | 1.519 | 61.4 | 38 | 0 |
| valkey-go/flush-100us+tls | 10 | 35073 | 0.277 | 0.386 | 0.524 | 0.750 | 1.920 | 66.4 | 41 | 0 |
| valkey-go/flush-100us+tls | 100 | 45438 | 1.834 | 4.311 | 6.263 | 7.247 | 8.103 | 55.2 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15594 | 0.062 | 0.071 | 0.105 | 0.211 | 1.469 | 46.5 | 36 | 0 |
| valkey-go | 10 | 73272 | 0.132 | 0.187 | 0.254 | 0.366 | 1.645 | 43.6 | 41 | 0 |
| valkey-go | 100 | 331204 | 0.298 | 0.412 | 0.540 | 0.871 | 2.701 | 15.6 | 67 | 0 |
| valkey-go | 1000 | 912140 | 1.125 | 1.467 | 1.885 | 3.047 | 5.543 | 6.7 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13622 | 0.070 | 0.083 | 0.128 | 0.260 | 2.007 | 51.2 | 35 | 0 |
| valkey-go | 10 | 48575 | 0.201 | 0.277 | 0.384 | 0.645 | 1.622 | 60.1 | 39 | 0 |
| valkey-go | 100 | 105241 | 0.977 | 1.243 | 1.672 | 2.493 | 4.077 | 28.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14431 | 0.066 | 0.076 | 0.110 | 0.233 | 1.576 | 53.0 | 30 | 0 |
| go-redis | 10 | 53039 | 0.187 | 0.244 | 0.337 | 0.552 | 1.630 | 59.8 | 33 | 0 |
| go-redis | 100 | 54100 | 1.781 | 1.930 | 3.547 | 3.881 | 5.939 | 65.7 | 70 | 0 |
| go-redis | 1000 | 53151 | 18.511 | 20.655 | 23.711 | 26.159 | 30.303 | 78.2 | 240 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12028 | 0.077 | 0.095 | 0.157 | 0.351 | 1.562 | 62.9 | 27 | 0 |
| go-redis | 10 | 38813 | 0.252 | 0.327 | 0.545 | 0.919 | 1.845 | 78.9 | 34 | 0 |
| go-redis | 100 | 40040 | 2.395 | 2.629 | 4.763 | 5.223 | 7.067 | 81.5 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7767 | 0.121 | 0.152 | 0.235 | 0.420 | 1.996 | 118.7 | 35 | 0 |
| glide | 10 | 32625 | 0.297 | 0.390 | 0.542 | 0.777 | 1.783 | 55.9 | 38 | 0 |
| glide | 100 | 36978 | 2.799 | 3.165 | 4.085 | 4.707 | 5.319 | 54.6 | 67 | 0 |
| glide | 1000 | 35325 | 28.095 | 31.551 | 35.263 | 39.167 | 40.927 | 56.9 | 240 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6705 | 0.136 | 0.181 | 0.316 | 0.560 | 1.633 | 122.0 | 33 | 0 |
| glide | 10 | 26451 | 0.358 | 0.528 | 0.913 | 1.433 | 2.213 | 88.6 | 35 | 0 |
| glide | 100 | 31033 | 3.195 | 4.203 | 5.231 | 6.251 | 8.631 | 78.8 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12724 | 0.074 | 0.086 | 0.132 | 0.292 | 1.650 | 49.1 | 33 | 0 |
| valkey-go+tls | 10 | 35737 | 0.263 | 0.435 | 0.577 | 0.718 | 1.689 | 62.1 | 42 | 0 |
| valkey-go+tls | 100 | 37397 | 2.291 | 5.651 | 8.051 | 9.039 | 10.255 | 62.5 | 66 | 0 |
| valkey-go+tls | 1000 | 37265 | 26.303 | 39.295 | 48.287 | 51.295 | 52.127 | 64.0 | 195 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10411 | 0.092 | 0.108 | 0.165 | 0.341 | 1.555 | 63.7 | 38 | 0 |
| valkey-go+tls | 10 | 27896 | 0.338 | 0.548 | 0.723 | 0.914 | 1.858 | 81.9 | 42 | 0 |
| valkey-go+tls | 100 | 28511 | 3.009 | 6.943 | 11.647 | 17.407 | 19.583 | 82.8 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12146 | 0.078 | 0.094 | 0.137 | 0.281 | 1.478 | 54.9 | 29 | 0 |
| go-redis+tls | 10 | 37194 | 0.266 | 0.336 | 0.469 | 0.700 | 1.756 | 65.2 | 34 | 0 |
| go-redis+tls | 100 | 37995 | 2.537 | 2.747 | 5.039 | 5.491 | 7.051 | 65.5 | 74 | 0 |
| go-redis+tls | 1000 | 37050 | 26.639 | 29.071 | 32.207 | 33.887 | 36.895 | 76.9 | 232 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9988 | 0.094 | 0.112 | 0.189 | 0.443 | 1.640 | 73.1 | 30 | 0 |
| go-redis+tls | 10 | 29371 | 0.339 | 0.422 | 0.602 | 0.973 | 2.207 | 84.4 | 37 | 0 |
| go-redis+tls | 100 | 30276 | 3.171 | 3.475 | 6.287 | 7.031 | 9.143 | 81.9 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7951 | 0.117 | 0.151 | 0.255 | 0.459 | 1.608 | 113.3 | 31 | 0 |
| go-redis-autopipeline+tls | 10 | 52657 | 0.179 | 0.235 | 0.388 | 0.839 | 1.684 | 21.5 | 34 | 0 |
| go-redis-autopipeline+tls | 100 | 193565 | 0.505 | 0.617 | 0.814 | 1.573 | 2.155 | 8.0 | 60 | 0 |
| go-redis-autopipeline+tls | 1000 | 434656 | 2.243 | 3.009 | 3.627 | 5.055 | 7.875 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9957 | 0.095 | 0.113 | 0.186 | 0.422 | 1.689 | 72.6 | 31 | 0 |
| go-redis-autopipeline+tls | 10 | 29234 | 0.340 | 0.423 | 0.607 | 0.946 | 3.325 | 83.6 | 38 | 0 |
| go-redis-autopipeline+tls | 100 | 29762 | 3.205 | 3.681 | 6.359 | 7.327 | 9.583 | 81.0 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6595 | 0.144 | 0.179 | 0.299 | 0.506 | 1.652 | 128.1 | 37 | 0 |
| glide+tls | 10 | 30829 | 0.315 | 0.439 | 0.600 | 0.842 | 2.081 | 58.4 | 40 | 0 |
| glide+tls | 100 | 33535 | 2.943 | 3.359 | 4.631 | 5.347 | 5.971 | 57.0 | 66 | 0 |
| glide+tls | 1000 | 33058 | 29.775 | 33.279 | 46.079 | 48.927 | 49.279 | 60.6 | 236 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5972 | 0.154 | 0.202 | 0.359 | 0.639 | 1.682 | 129.1 | 36 | 0 |
| glide+tls | 10 | 23767 | 0.395 | 0.570 | 0.919 | 1.401 | 2.279 | 94.6 | 36 | 0 |
| glide+tls | 100 | 26071 | 3.823 | 4.599 | 5.655 | 6.263 | 7.303 | 88.2 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15949 | 0.061 | 0.070 | 0.100 | 0.213 | 1.566 | 46.0 | 36 | 0 |
| valkey-go/flush-100us | 10 | 60468 | 0.158 | 0.242 | 0.325 | 1.337 | 2.099 | 36.0 | 41 | 0 |
| valkey-go/flush-100us | 100 | 416298 | 0.232 | 0.331 | 0.451 | 0.681 | 1.987 | 8.7 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 990701 | 1.031 | 1.386 | 1.685 | 2.687 | 5.167 | 5.5 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13352 | 0.071 | 0.082 | 0.127 | 0.263 | 1.617 | 52.1 | 35 | 0 |
| valkey-go/flush-100us | 10 | 45273 | 0.213 | 0.319 | 0.431 | 0.669 | 1.807 | 54.1 | 39 | 0 |
| valkey-go/flush-100us | 100 | 108918 | 0.934 | 1.184 | 1.663 | 2.301 | 2.983 | 23.3 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12464 | 0.076 | 0.091 | 0.131 | 0.268 | 3.007 | 50.4 | 37 | 0 |
| valkey-go/flush-100us+tls | 10 | 49071 | 0.198 | 0.290 | 0.370 | 0.523 | 1.846 | 44.7 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 203370 | 0.447 | 0.740 | 1.120 | 1.693 | 3.261 | 14.6 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 220235 | 4.303 | 8.263 | 9.943 | 10.975 | 11.559 | 14.3 | 265 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10427 | 0.091 | 0.111 | 0.178 | 0.326 | 1.646 | 63.1 | 36 | 0 |
| valkey-go/flush-100us+tls | 10 | 35046 | 0.278 | 0.386 | 0.511 | 0.745 | 2.245 | 67.8 | 41 | 0 |
| valkey-go/flush-100us+tls | 100 | 45167 | 2.049 | 3.927 | 5.343 | 5.967 | 6.975 | 55.5 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15858 | 0.061 | 0.069 | 0.102 | 0.228 | 1.549 | 45.5 | 35 | 0 |
| valkey-go | 10 | 72706 | 0.133 | 0.189 | 0.254 | 0.368 | 1.598 | 43.8 | 39 | 0 |
| valkey-go | 100 | 328397 | 0.300 | 0.417 | 0.546 | 0.927 | 2.179 | 15.9 | 64 | 0 |
| valkey-go | 1000 | 888797 | 1.161 | 1.509 | 1.854 | 3.069 | 6.447 | 7.0 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12076 | 0.072 | 0.119 | 0.195 | 0.307 | 1.825 | 59.9 | 35 | 0 |
| valkey-go | 10 | 48149 | 0.202 | 0.281 | 0.400 | 0.672 | 1.796 | 66.2 | 39 | 0 |
| valkey-go | 100 | 105290 | 0.976 | 1.242 | 1.651 | 2.611 | 3.977 | 29.1 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13780 | 0.068 | 0.082 | 0.126 | 0.273 | 3.537 | 56.7 | 28 | 0 |
| go-redis | 10 | 51660 | 0.192 | 0.253 | 0.354 | 0.627 | 2.031 | 62.8 | 31 | 0 |
| go-redis | 100 | 53338 | 1.804 | 1.999 | 3.585 | 3.991 | 5.871 | 71.2 | 69 | 0 |
| go-redis | 1000 | 53311 | 18.575 | 20.031 | 22.303 | 25.183 | 28.367 | 81.2 | 292 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12081 | 0.077 | 0.097 | 0.161 | 0.385 | 1.808 | 63.1 | 28 | 0 |
| go-redis | 10 | 38594 | 0.253 | 0.331 | 0.548 | 0.911 | 1.896 | 79.7 | 35 | 0 |
| go-redis | 100 | 39675 | 2.405 | 2.737 | 4.795 | 5.503 | 8.063 | 83.1 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9648 | 0.094 | 0.127 | 0.210 | 0.395 | 1.563 | 107.2 | 30 | 0 |
| go-redis-autopipeline | 10 | 58701 | 0.158 | 0.213 | 0.367 | 0.796 | 2.205 | 21.9 | 33 | 0 |
| go-redis-autopipeline | 100 | 206153 | 0.476 | 0.576 | 0.732 | 1.481 | 2.105 | 8.4 | 57 | 0 |
| go-redis-autopipeline | 1000 | 459760 | 2.189 | 2.825 | 3.423 | 5.011 | 7.719 | 6.8 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12040 | 0.077 | 0.095 | 0.164 | 0.362 | 1.634 | 63.5 | 27 | 0 |
| go-redis-autopipeline | 10 | 38707 | 0.252 | 0.329 | 0.554 | 0.923 | 1.635 | 80.0 | 35 | 0 |
| go-redis-autopipeline | 100 | 39942 | 2.395 | 2.681 | 4.787 | 5.503 | 7.619 | 84.1 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12524 | 0.076 | 0.089 | 0.131 | 0.256 | 1.610 | 49.8 | 35 | 0 |
| valkey-go+tls | 10 | 35694 | 0.262 | 0.435 | 0.583 | 0.792 | 1.878 | 62.3 | 44 | 0 |
| valkey-go+tls | 100 | 37054 | 2.089 | 5.875 | 8.263 | 9.119 | 10.655 | 63.8 | 68 | 0 |
| valkey-go+tls | 1000 | 37265 | 18.783 | 62.783 | 78.975 | 81.407 | 82.239 | 64.9 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10731 | 0.089 | 0.105 | 0.162 | 0.340 | 1.698 | 60.5 | 37 | 0 |
| valkey-go+tls | 10 | 28664 | 0.329 | 0.531 | 0.699 | 0.870 | 1.870 | 81.9 | 40 | 0 |
| valkey-go+tls | 100 | 29548 | 2.987 | 6.671 | 9.671 | 12.039 | 14.279 | 81.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12274 | 0.078 | 0.092 | 0.135 | 0.294 | 1.649 | 56.1 | 28 | 0 |
| go-redis+tls | 10 | 37494 | 0.265 | 0.332 | 0.445 | 0.697 | 1.793 | 65.7 | 34 | 0 |
| go-redis+tls | 100 | 37528 | 2.555 | 2.801 | 5.075 | 5.655 | 8.199 | 65.0 | 72 | 0 |
| go-redis+tls | 1000 | 36772 | 26.927 | 29.615 | 32.399 | 34.719 | 38.751 | 77.1 | 234 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9910 | 0.094 | 0.115 | 0.199 | 0.426 | 1.613 | 73.0 | 28 | 0 |
| go-redis+tls | 10 | 29238 | 0.340 | 0.423 | 0.623 | 0.978 | 2.243 | 83.7 | 36 | 0 |
| go-redis+tls | 100 | 30202 | 3.171 | 3.503 | 6.283 | 6.995 | 9.703 | 83.3 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 8043 | 0.116 | 0.148 | 0.246 | 0.470 | 1.696 | 112.6 | 29 | 0 |
| go-redis-autopipeline+tls | 10 | 52676 | 0.179 | 0.233 | 0.385 | 0.857 | 1.649 | 21.8 | 35 | 0 |
| go-redis-autopipeline+tls | 100 | 190251 | 0.513 | 0.629 | 0.849 | 1.549 | 2.395 | 8.2 | 59 | 0 |
| go-redis-autopipeline+tls | 1000 | 433532 | 2.193 | 3.047 | 3.703 | 5.179 | 8.279 | 6.9 | 311 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9683 | 0.096 | 0.121 | 0.209 | 0.458 | 2.987 | 73.3 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 26850 | 0.356 | 0.493 | 0.793 | 1.107 | 2.191 | 92.8 | 37 | 0 |
| go-redis-autopipeline+tls | 100 | 27172 | 3.473 | 4.347 | 6.987 | 8.679 | 11.671 | 93.2 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6468 | 0.146 | 0.187 | 0.335 | 0.549 | 1.815 | 132.3 | 35 | 0 |
| glide+tls | 10 | 28265 | 0.351 | 0.470 | 0.680 | 1.010 | 1.899 | 63.1 | 38 | 0 |
| glide+tls | 100 | 31640 | 3.049 | 4.163 | 4.787 | 5.455 | 7.055 | 59.7 | 65 | 0 |
| glide+tls | 1000 | 32864 | 30.287 | 33.247 | 38.079 | 45.727 | 47.327 | 60.8 | 238 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5701 | 0.160 | 0.218 | 0.407 | 0.729 | 1.801 | 133.9 | 35 | 0 |
| glide+tls | 10 | 23208 | 0.401 | 0.591 | 0.973 | 1.513 | 2.655 | 97.0 | 38 | 0 |
| glide+tls | 100 | 25539 | 3.919 | 4.519 | 5.655 | 6.503 | 8.047 | 89.0 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15392 | 0.062 | 0.073 | 0.108 | 0.229 | 1.558 | 47.3 | 37 | 0 |
| valkey-go/flush-100us | 10 | 60871 | 0.157 | 0.241 | 0.317 | 1.328 | 1.746 | 36.2 | 42 | 0 |
| valkey-go/flush-100us | 100 | 414708 | 0.232 | 0.334 | 0.464 | 0.719 | 1.822 | 8.8 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 970932 | 1.050 | 1.420 | 1.828 | 2.787 | 6.315 | 5.7 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13131 | 0.072 | 0.087 | 0.134 | 0.284 | 1.642 | 53.8 | 37 | 0 |
| valkey-go/flush-100us | 10 | 44654 | 0.216 | 0.324 | 0.444 | 0.708 | 1.881 | 54.8 | 40 | 0 |
| valkey-go/flush-100us | 100 | 106529 | 0.956 | 1.210 | 1.726 | 2.289 | 3.525 | 24.0 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12605 | 0.075 | 0.089 | 0.129 | 0.248 | 1.581 | 49.9 | 33 | 0 |
| valkey-go/flush-100us+tls | 10 | 49713 | 0.195 | 0.289 | 0.372 | 0.520 | 1.879 | 44.9 | 41 | 0 |
| valkey-go/flush-100us+tls | 100 | 217001 | 0.419 | 0.684 | 1.059 | 1.533 | 2.905 | 14.1 | 66 | 0 |
| valkey-go/flush-100us+tls | 1000 | 218913 | 4.719 | 7.595 | 8.999 | 10.375 | 11.567 | 15.0 | 279 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10618 | 0.089 | 0.108 | 0.177 | 0.376 | 1.979 | 61.9 | 37 | 0 |
| valkey-go/flush-100us+tls | 10 | 34625 | 0.280 | 0.393 | 0.546 | 0.798 | 2.237 | 68.0 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 45476 | 1.978 | 3.985 | 5.591 | 6.183 | 10.351 | 55.1 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15272 | 0.062 | 0.074 | 0.107 | 0.214 | 1.550 | 48.4 | 34 | 0 |
| valkey-go | 10 | 71841 | 0.134 | 0.191 | 0.259 | 0.378 | 1.899 | 44.0 | 40 | 0 |
| valkey-go | 100 | 324678 | 0.303 | 0.422 | 0.562 | 0.901 | 2.239 | 15.9 | 66 | 0 |
| valkey-go | 1000 | 898258 | 1.144 | 1.491 | 1.840 | 3.291 | 5.511 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13048 | 0.072 | 0.088 | 0.147 | 0.322 | 1.920 | 53.3 | 37 | 0 |
| valkey-go | 10 | 48141 | 0.202 | 0.281 | 0.399 | 0.684 | 1.562 | 60.2 | 41 | 0 |
| valkey-go | 100 | 105685 | 0.973 | 1.242 | 1.673 | 2.413 | 3.787 | 29.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14127 | 0.067 | 0.080 | 0.125 | 0.295 | 1.629 | 53.9 | 30 | 0 |
| go-redis | 10 | 52119 | 0.190 | 0.249 | 0.341 | 0.619 | 2.065 | 61.6 | 32 | 0 |
| go-redis | 100 | 52620 | 1.823 | 2.059 | 3.629 | 4.061 | 5.623 | 69.9 | 70 | 0 |
| go-redis | 1000 | 52129 | 18.911 | 20.767 | 24.975 | 26.671 | 30.143 | 81.6 | 289 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11463 | 0.079 | 0.104 | 0.190 | 0.430 | 2.539 | 67.2 | 29 | 0 |
| go-redis | 10 | 37795 | 0.258 | 0.340 | 0.571 | 0.973 | 2.225 | 81.9 | 37 | 0 |
| go-redis | 100 | 39876 | 2.399 | 2.661 | 4.771 | 5.455 | 6.807 | 83.9 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9475 | 0.097 | 0.131 | 0.213 | 0.389 | 1.742 | 107.3 | 30 | 0 |
| go-redis-autopipeline | 10 | 59981 | 0.155 | 0.208 | 0.357 | 0.777 | 2.271 | 21.1 | 32 | 0 |
| go-redis-autopipeline | 100 | 202575 | 0.484 | 0.589 | 0.763 | 1.527 | 2.329 | 8.3 | 57 | 0 |
| go-redis-autopipeline | 1000 | 447516 | 2.205 | 2.931 | 3.505 | 5.447 | 8.295 | 6.9 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11767 | 0.078 | 0.100 | 0.181 | 0.399 | 1.673 | 66.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 38084 | 0.256 | 0.336 | 0.574 | 0.945 | 2.129 | 80.8 | 35 | 0 |
| go-redis-autopipeline | 100 | 40125 | 2.387 | 2.661 | 4.751 | 5.343 | 7.439 | 84.5 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7427 | 0.126 | 0.165 | 0.286 | 0.505 | 1.977 | 122.1 | 33 | 0 |
| glide | 10 | 33902 | 0.282 | 0.379 | 0.529 | 0.863 | 2.007 | 59.3 | 38 | 0 |
| glide | 100 | 35615 | 2.829 | 3.411 | 4.507 | 5.251 | 6.007 | 56.3 | 65 | 0 |
| glide | 1000 | 34749 | 28.495 | 32.767 | 37.471 | 57.663 | 61.247 | 58.7 | 223 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6668 | 0.137 | 0.184 | 0.332 | 0.596 | 1.697 | 123.0 | 33 | 0 |
| glide | 10 | 25722 | 0.369 | 0.540 | 0.940 | 1.503 | 2.681 | 91.5 | 36 | 0 |
| glide | 100 | 28775 | 3.463 | 4.459 | 5.523 | 6.451 | 8.335 | 85.0 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12347 | 0.077 | 0.092 | 0.138 | 0.273 | 1.431 | 55.0 | 29 | 0 |
| go-redis+tls | 10 | 36754 | 0.270 | 0.340 | 0.459 | 0.722 | 1.838 | 66.6 | 34 | 0 |
| go-redis+tls | 100 | 37430 | 2.571 | 2.839 | 5.111 | 5.635 | 7.359 | 65.7 | 71 | 0 |
| go-redis+tls | 1000 | 37570 | 26.399 | 28.943 | 33.151 | 36.223 | 40.543 | 78.0 | 248 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9951 | 0.094 | 0.114 | 0.196 | 0.430 | 1.574 | 73.6 | 28 | 0 |
| go-redis+tls | 10 | 29180 | 0.340 | 0.425 | 0.616 | 0.977 | 1.908 | 85.2 | 36 | 0 |
| go-redis+tls | 100 | 30026 | 3.195 | 3.493 | 6.347 | 7.071 | 9.247 | 82.6 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 8188 | 0.111 | 0.146 | 0.252 | 0.446 | 1.618 | 110.4 | 30 | 0 |
| go-redis-autopipeline+tls | 10 | 52246 | 0.179 | 0.236 | 0.401 | 0.786 | 1.711 | 21.9 | 33 | 0 |
| go-redis-autopipeline+tls | 100 | 192927 | 0.507 | 0.616 | 0.809 | 1.586 | 2.267 | 8.1 | 62 | 0 |
| go-redis-autopipeline+tls | 1000 | 414587 | 2.537 | 3.137 | 3.819 | 5.839 | 8.695 | 7.0 | 310 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 10017 | 0.094 | 0.113 | 0.190 | 0.424 | 1.590 | 72.0 | 27 | 0 |
| go-redis-autopipeline+tls | 10 | 29354 | 0.339 | 0.423 | 0.604 | 0.963 | 1.972 | 85.3 | 36 | 0 |
| go-redis-autopipeline+tls | 100 | 29759 | 3.225 | 3.567 | 6.391 | 7.095 | 10.103 | 83.0 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6475 | 0.148 | 0.183 | 0.308 | 0.526 | 1.640 | 133.1 | 35 | 0 |
| glide+tls | 10 | 30178 | 0.319 | 0.453 | 0.617 | 0.905 | 2.155 | 61.0 | 39 | 0 |
| glide+tls | 100 | 32627 | 3.039 | 3.495 | 4.663 | 5.327 | 6.339 | 59.8 | 67 | 0 |
| glide+tls | 1000 | 33223 | 29.951 | 32.799 | 35.615 | 38.111 | 38.687 | 61.3 | 230 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5806 | 0.157 | 0.210 | 0.383 | 0.700 | 1.710 | 130.5 | 33 | 0 |
| glide+tls | 10 | 23631 | 0.398 | 0.571 | 0.932 | 1.487 | 2.807 | 95.9 | 35 | 0 |
| glide+tls | 100 | 25796 | 3.873 | 4.483 | 5.635 | 6.739 | 8.391 | 88.8 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15839 | 0.060 | 0.072 | 0.108 | 0.231 | 1.731 | 45.9 | 34 | 0 |
| valkey-go/flush-100us | 10 | 60388 | 0.159 | 0.243 | 0.324 | 1.339 | 1.866 | 36.4 | 41 | 0 |
| valkey-go/flush-100us | 100 | 410171 | 0.234 | 0.338 | 0.472 | 0.695 | 1.861 | 8.9 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 961983 | 1.059 | 1.420 | 1.837 | 2.965 | 5.775 | 5.6 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13176 | 0.072 | 0.083 | 0.131 | 0.265 | 1.576 | 52.2 | 37 | 0 |
| valkey-go/flush-100us | 10 | 44671 | 0.215 | 0.325 | 0.443 | 0.713 | 1.527 | 55.6 | 40 | 0 |
| valkey-go/flush-100us | 100 | 107167 | 0.952 | 1.205 | 1.678 | 2.211 | 3.163 | 23.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12522 | 0.076 | 0.090 | 0.136 | 0.289 | 1.670 | 50.2 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 48910 | 0.198 | 0.292 | 0.386 | 0.558 | 1.820 | 45.4 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 224027 | 0.415 | 0.638 | 0.957 | 1.505 | 2.681 | 13.9 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 192478 | 3.597 | 11.335 | 18.751 | 27.727 | 29.663 | 15.7 | 291 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10545 | 0.090 | 0.108 | 0.177 | 0.344 | 1.665 | 62.4 | 37 | 0 |
| valkey-go/flush-100us+tls | 10 | 33996 | 0.285 | 0.400 | 0.561 | 0.833 | 2.125 | 69.1 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 32923 | 2.273 | 5.967 | 12.815 | 16.199 | 18.703 | 78.5 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15718 | 0.061 | 0.071 | 0.105 | 0.232 | 1.544 | 46.7 | 34 | 0 |
| valkey-go | 10 | 72884 | 0.132 | 0.188 | 0.257 | 0.371 | 1.632 | 43.7 | 41 | 0 |
| valkey-go | 100 | 327156 | 0.302 | 0.418 | 0.539 | 0.865 | 2.057 | 15.7 | 68 | 0 |
| valkey-go | 1000 | 904857 | 1.135 | 1.479 | 1.840 | 3.345 | 7.719 | 6.9 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13207 | 0.071 | 0.087 | 0.131 | 0.276 | 1.559 | 53.4 | 35 | 0 |
| valkey-go | 10 | 48844 | 0.200 | 0.276 | 0.390 | 0.666 | 1.682 | 59.7 | 39 | 0 |
| valkey-go | 100 | 104752 | 0.976 | 1.248 | 1.793 | 2.537 | 3.623 | 28.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14463 | 0.066 | 0.076 | 0.114 | 0.243 | 1.621 | 52.9 | 30 | 0 |
| go-redis | 10 | 52144 | 0.190 | 0.250 | 0.343 | 0.583 | 1.697 | 61.0 | 34 | 0 |
| go-redis | 100 | 53937 | 1.785 | 1.957 | 3.557 | 3.983 | 6.591 | 67.3 | 71 | 0 |
| go-redis | 1000 | 53988 | 18.415 | 19.535 | 21.663 | 23.071 | 25.679 | 78.8 | 257 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11512 | 0.078 | 0.106 | 0.186 | 0.404 | 2.145 | 66.6 | 29 | 0 |
| go-redis | 10 | 38325 | 0.255 | 0.331 | 0.554 | 0.937 | 1.862 | 80.6 | 36 | 0 |
| go-redis | 100 | 39852 | 2.403 | 2.665 | 4.791 | 5.411 | 7.331 | 84.0 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9681 | 0.094 | 0.129 | 0.212 | 0.415 | 1.868 | 106.4 | 27 | 0 |
| go-redis-autopipeline | 10 | 61894 | 0.151 | 0.201 | 0.344 | 0.811 | 2.021 | 20.6 | 31 | 0 |
| go-redis-autopipeline | 100 | 202346 | 0.483 | 0.592 | 0.777 | 1.664 | 2.399 | 8.2 | 58 | 0 |
| go-redis-autopipeline | 1000 | 445247 | 2.205 | 2.961 | 3.593 | 5.299 | 8.631 | 6.7 | 309 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12046 | 0.076 | 0.098 | 0.169 | 0.406 | 1.571 | 64.2 | 29 | 0 |
| go-redis-autopipeline | 10 | 38814 | 0.252 | 0.327 | 0.552 | 0.970 | 1.880 | 79.1 | 35 | 0 |
| go-redis-autopipeline | 100 | 39158 | 2.405 | 3.047 | 4.811 | 5.903 | 7.483 | 83.8 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 8024 | 0.115 | 0.151 | 0.270 | 0.477 | 3.053 | 115.9 | 33 | 0 |
| glide | 10 | 32441 | 0.293 | 0.405 | 0.575 | 0.883 | 1.711 | 56.6 | 36 | 0 |
| glide | 100 | 34139 | 2.875 | 4.179 | 4.639 | 5.207 | 6.059 | 58.0 | 65 | 0 |
| glide | 1000 | 32879 | 29.183 | 38.015 | 42.559 | 44.895 | 46.911 | 61.6 | 207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6473 | 0.139 | 0.194 | 0.368 | 0.656 | 1.919 | 127.1 | 33 | 0 |
| glide | 10 | 25623 | 0.368 | 0.550 | 0.996 | 1.576 | 2.395 | 93.9 | 35 | 0 |
| glide | 100 | 28884 | 3.445 | 4.463 | 5.555 | 6.551 | 8.615 | 85.1 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12470 | 0.075 | 0.090 | 0.143 | 0.310 | 1.583 | 50.2 | 33 | 0 |
| valkey-go+tls | 10 | 35649 | 0.263 | 0.434 | 0.579 | 0.724 | 1.792 | 63.7 | 42 | 0 |
| valkey-go+tls | 100 | 36960 | 2.169 | 5.811 | 8.035 | 8.727 | 9.239 | 63.6 | 64 | 0 |
| valkey-go+tls | 1000 | 36901 | 25.199 | 46.463 | 54.623 | 56.159 | 56.895 | 65.3 | 194 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10626 | 0.090 | 0.106 | 0.172 | 0.393 | 1.618 | 62.1 | 37 | 0 |
| valkey-go+tls | 10 | 28403 | 0.332 | 0.536 | 0.713 | 0.981 | 2.071 | 81.7 | 40 | 0 |
| valkey-go+tls | 100 | 29573 | 2.623 | 7.567 | 9.871 | 12.015 | 13.623 | 82.0 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7999 | 0.115 | 0.150 | 0.259 | 0.491 | 1.563 | 114.3 | 30 | 0 |
| go-redis-autopipeline+tls | 10 | 52455 | 0.178 | 0.238 | 0.411 | 0.832 | 1.871 | 21.7 | 35 | 0 |
| go-redis-autopipeline+tls | 100 | 189702 | 0.515 | 0.629 | 0.834 | 1.501 | 2.695 | 7.9 | 61 | 0 |
| go-redis-autopipeline+tls | 1000 | 401870 | 2.611 | 3.231 | 4.003 | 5.419 | 7.363 | 7.1 | 311 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9664 | 0.095 | 0.117 | 0.243 | 0.488 | 1.564 | 74.7 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 29033 | 0.342 | 0.426 | 0.629 | 0.978 | 2.491 | 84.0 | 37 | 0 |
| go-redis-autopipeline+tls | 100 | 29936 | 3.187 | 3.575 | 6.339 | 7.231 | 10.543 | 84.9 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6722 | 0.139 | 0.177 | 0.310 | 0.542 | 1.569 | 127.2 | 37 | 0 |
| glide+tls | 10 | 30124 | 0.319 | 0.456 | 0.618 | 1.017 | 1.927 | 60.6 | 40 | 0 |
| glide+tls | 100 | 32061 | 3.023 | 4.139 | 4.807 | 5.363 | 6.587 | 59.6 | 68 | 0 |
| glide+tls | 1000 | 33187 | 29.631 | 33.119 | 42.559 | 45.183 | 48.575 | 61.4 | 236 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5748 | 0.158 | 0.215 | 0.388 | 0.672 | 1.745 | 133.5 | 36 | 0 |
| glide+tls | 10 | 23313 | 0.401 | 0.578 | 0.958 | 1.506 | 2.431 | 96.4 | 38 | 0 |
| glide+tls | 100 | 25626 | 3.877 | 4.727 | 5.719 | 6.499 | 8.919 | 90.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15640 | 0.061 | 0.071 | 0.106 | 0.214 | 1.576 | 47.0 | 35 | 0 |
| valkey-go/flush-100us | 10 | 61107 | 0.157 | 0.240 | 0.315 | 1.318 | 1.867 | 36.1 | 41 | 0 |
| valkey-go/flush-100us | 100 | 412309 | 0.234 | 0.334 | 0.460 | 0.734 | 1.951 | 8.8 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 969463 | 1.049 | 1.409 | 1.850 | 2.623 | 5.443 | 5.5 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13319 | 0.071 | 0.084 | 0.136 | 0.288 | 1.555 | 52.5 | 36 | 0 |
| valkey-go/flush-100us | 10 | 44334 | 0.217 | 0.326 | 0.444 | 0.703 | 1.939 | 55.3 | 42 | 0 |
| valkey-go/flush-100us | 100 | 109429 | 0.930 | 1.174 | 1.641 | 2.191 | 3.017 | 23.6 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12551 | 0.075 | 0.089 | 0.133 | 0.269 | 1.631 | 50.0 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 49473 | 0.196 | 0.289 | 0.372 | 0.575 | 1.889 | 44.3 | 41 | 0 |
| valkey-go/flush-100us+tls | 100 | 206018 | 0.436 | 0.748 | 1.126 | 1.641 | 2.685 | 14.8 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 235019 | 4.067 | 7.151 | 10.055 | 15.071 | 17.695 | 13.9 | 270 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10645 | 0.090 | 0.106 | 0.162 | 0.325 | 1.644 | 62.5 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 34504 | 0.281 | 0.393 | 0.547 | 0.835 | 1.964 | 68.8 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 45136 | 2.008 | 4.123 | 5.507 | 6.115 | 6.731 | 56.3 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15524 | 0.062 | 0.072 | 0.105 | 0.237 | 1.537 | 46.9 | 36 | 0 |
| valkey-go | 10 | 72550 | 0.133 | 0.190 | 0.261 | 0.384 | 1.625 | 43.6 | 40 | 0 |
| valkey-go | 100 | 329875 | 0.300 | 0.414 | 0.529 | 0.860 | 2.327 | 15.7 | 67 | 0 |
| valkey-go | 1000 | 904474 | 1.132 | 1.489 | 1.908 | 3.081 | 6.623 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13236 | 0.071 | 0.085 | 0.131 | 0.285 | 2.859 | 52.2 | 35 | 0 |
| valkey-go | 10 | 47952 | 0.203 | 0.282 | 0.404 | 0.697 | 1.687 | 60.8 | 41 | 0 |
| valkey-go | 100 | 107382 | 0.960 | 1.214 | 1.601 | 2.357 | 3.727 | 28.7 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14074 | 0.067 | 0.082 | 0.128 | 0.272 | 1.585 | 54.5 | 28 | 0 |
| go-redis | 10 | 52722 | 0.188 | 0.246 | 0.337 | 0.589 | 1.641 | 60.7 | 32 | 0 |
| go-redis | 100 | 53525 | 1.798 | 1.965 | 3.589 | 3.977 | 6.199 | 66.6 | 72 | 0 |
| go-redis | 1000 | 52880 | 18.751 | 20.079 | 23.215 | 25.167 | 28.175 | 78.1 | 283 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12083 | 0.077 | 0.096 | 0.160 | 0.371 | 1.512 | 63.5 | 27 | 0 |
| go-redis | 10 | 38835 | 0.252 | 0.327 | 0.545 | 0.977 | 1.908 | 78.4 | 35 | 0 |
| go-redis | 100 | 39492 | 2.423 | 2.701 | 4.839 | 5.583 | 7.767 | 83.1 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9939 | 0.093 | 0.123 | 0.190 | 0.373 | 1.540 | 104.0 | 27 | 0 |
| go-redis-autopipeline | 10 | 61546 | 0.151 | 0.202 | 0.348 | 0.771 | 1.668 | 20.8 | 32 | 0 |
| go-redis-autopipeline | 100 | 178904 | 0.512 | 0.776 | 1.112 | 1.837 | 2.803 | 9.2 | 58 | 0 |
| go-redis-autopipeline | 1000 | 443250 | 2.243 | 2.953 | 4.123 | 6.599 | 10.103 | 6.5 | 309 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11809 | 0.078 | 0.101 | 0.179 | 0.402 | 1.560 | 64.6 | 30 | 0 |
| go-redis-autopipeline | 10 | 38930 | 0.251 | 0.327 | 0.550 | 0.947 | 1.754 | 78.9 | 38 | 0 |
| go-redis-autopipeline | 100 | 39478 | 2.419 | 2.783 | 4.835 | 5.679 | 7.531 | 82.8 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7673 | 0.120 | 0.158 | 0.267 | 0.462 | 2.321 | 119.8 | 34 | 0 |
| glide | 10 | 34183 | 0.284 | 0.365 | 0.516 | 0.775 | 2.079 | 54.8 | 37 | 0 |
| glide | 100 | 36258 | 2.783 | 3.251 | 4.483 | 5.375 | 7.327 | 55.0 | 62 | 0 |
| glide | 1000 | 36375 | 27.135 | 31.151 | 38.719 | 63.455 | 67.903 | 57.3 | 237 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6756 | 0.135 | 0.183 | 0.327 | 0.587 | 1.769 | 121.6 | 33 | 0 |
| glide | 10 | 25327 | 0.373 | 0.555 | 0.976 | 1.550 | 2.461 | 93.1 | 35 | 0 |
| glide | 100 | 29080 | 3.423 | 4.427 | 5.547 | 6.731 | 7.955 | 84.6 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12773 | 0.074 | 0.089 | 0.132 | 0.261 | 1.586 | 49.6 | 35 | 0 |
| valkey-go+tls | 10 | 35104 | 0.267 | 0.441 | 0.593 | 0.783 | 1.900 | 63.5 | 41 | 0 |
| valkey-go+tls | 100 | 37021 | 2.205 | 5.775 | 7.995 | 9.063 | 9.967 | 63.3 | 66 | 0 |
| valkey-go+tls | 1000 | 36891 | 26.703 | 55.103 | 72.191 | 74.175 | 75.647 | 64.1 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10925 | 0.088 | 0.104 | 0.160 | 0.296 | 1.549 | 59.9 | 37 | 0 |
| valkey-go+tls | 10 | 28267 | 0.333 | 0.538 | 0.722 | 0.986 | 1.832 | 81.3 | 41 | 0 |
| valkey-go+tls | 100 | 29678 | 2.953 | 6.759 | 9.263 | 10.487 | 11.927 | 80.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 11876 | 0.078 | 0.096 | 0.146 | 0.332 | 1.667 | 57.1 | 30 | 0 |
| go-redis+tls | 10 | 37312 | 0.266 | 0.334 | 0.447 | 0.704 | 1.808 | 65.6 | 33 | 0 |
| go-redis+tls | 100 | 37958 | 2.531 | 2.773 | 5.027 | 5.627 | 9.015 | 66.4 | 73 | 0 |
| go-redis+tls | 1000 | 37487 | 26.479 | 28.799 | 32.751 | 35.967 | 39.199 | 77.4 | 234 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 10015 | 0.094 | 0.113 | 0.196 | 0.487 | 1.650 | 71.6 | 28 | 0 |
| go-redis+tls | 10 | 29081 | 0.341 | 0.428 | 0.622 | 0.943 | 1.917 | 85.7 | 35 | 0 |
| go-redis+tls | 100 | 29693 | 3.207 | 3.673 | 6.363 | 7.347 | 10.639 | 83.4 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6780 | 0.136 | 0.177 | 0.305 | 0.535 | 1.715 | 124.7 | 35 | 0 |
| glide+tls | 10 | 29139 | 0.329 | 0.475 | 0.622 | 0.850 | 1.962 | 61.9 | 38 | 0 |
| glide+tls | 100 | 32871 | 3.009 | 3.415 | 4.687 | 5.247 | 6.095 | 58.4 | 66 | 0 |
| glide+tls | 1000 | 32180 | 30.431 | 34.687 | 42.687 | 44.671 | 48.063 | 61.3 | 238 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6048 | 0.151 | 0.200 | 0.354 | 0.655 | 1.678 | 128.2 | 35 | 0 |
| glide+tls | 10 | 23101 | 0.404 | 0.598 | 0.960 | 1.498 | 2.769 | 96.6 | 36 | 0 |
| glide+tls | 100 | 24609 | 4.017 | 5.003 | 6.347 | 7.535 | 11.111 | 92.7 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15682 | 0.061 | 0.071 | 0.104 | 0.210 | 1.444 | 46.2 | 34 | 0 |
| valkey-go/flush-100us | 10 | 61582 | 0.155 | 0.237 | 0.313 | 1.334 | 1.715 | 35.8 | 41 | 0 |
| valkey-go/flush-100us | 100 | 412784 | 0.233 | 0.337 | 0.467 | 0.731 | 2.030 | 8.8 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 964474 | 1.055 | 1.422 | 1.813 | 2.819 | 5.055 | 5.4 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13071 | 0.072 | 0.087 | 0.140 | 0.279 | 1.585 | 53.3 | 35 | 0 |
| valkey-go/flush-100us | 10 | 45293 | 0.213 | 0.320 | 0.429 | 0.666 | 1.520 | 54.4 | 40 | 0 |
| valkey-go/flush-100us | 100 | 108396 | 0.937 | 1.188 | 1.671 | 2.173 | 3.619 | 23.3 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12812 | 0.074 | 0.088 | 0.130 | 0.270 | 1.632 | 50.0 | 33 | 0 |
| valkey-go/flush-100us+tls | 10 | 49261 | 0.196 | 0.291 | 0.387 | 0.589 | 2.353 | 44.7 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 212150 | 0.425 | 0.705 | 1.139 | 1.609 | 2.587 | 14.0 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 197111 | 4.475 | 10.071 | 11.943 | 13.719 | 14.967 | 15.3 | 281 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10661 | 0.090 | 0.106 | 0.162 | 0.335 | 1.649 | 61.7 | 36 | 0 |
| valkey-go/flush-100us+tls | 10 | 34310 | 0.283 | 0.395 | 0.526 | 0.774 | 1.841 | 67.5 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 45483 | 1.899 | 4.291 | 6.347 | 7.111 | 8.751 | 54.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15767 | 0.061 | 0.070 | 0.106 | 0.230 | 1.542 | 45.8 | 35 | 0 |
| valkey-go | 10 | 72685 | 0.133 | 0.188 | 0.256 | 0.362 | 1.947 | 43.7 | 40 | 0 |
| valkey-go | 100 | 328392 | 0.301 | 0.417 | 0.542 | 0.891 | 2.285 | 15.7 | 65 | 0 |
| valkey-go | 1000 | 910303 | 1.132 | 1.472 | 1.782 | 2.983 | 6.391 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13184 | 0.071 | 0.086 | 0.137 | 0.279 | 1.579 | 52.6 | 37 | 0 |
| valkey-go | 10 | 47848 | 0.204 | 0.282 | 0.398 | 0.678 | 1.813 | 60.2 | 39 | 0 |
| valkey-go | 100 | 105199 | 0.980 | 1.238 | 1.630 | 2.539 | 4.167 | 28.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14508 | 0.066 | 0.079 | 0.116 | 0.258 | 1.573 | 54.1 | 28 | 0 |
| go-redis | 10 | 52103 | 0.191 | 0.248 | 0.340 | 0.581 | 1.601 | 59.7 | 33 | 0 |
| go-redis | 100 | 52295 | 1.824 | 2.075 | 3.647 | 4.319 | 6.183 | 68.4 | 70 | 0 |
| go-redis | 1000 | 52837 | 18.815 | 19.983 | 21.743 | 22.831 | 24.351 | 79.5 | 249 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11830 | 0.078 | 0.099 | 0.169 | 0.414 | 1.691 | 64.5 | 28 | 0 |
| go-redis | 10 | 38324 | 0.255 | 0.332 | 0.557 | 0.955 | 1.915 | 80.3 | 36 | 0 |
| go-redis | 100 | 40155 | 2.383 | 2.663 | 4.751 | 5.411 | 7.287 | 83.3 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9383 | 0.096 | 0.138 | 0.239 | 0.421 | 2.217 | 108.4 | 29 | 0 |
| go-redis-autopipeline | 10 | 57740 | 0.159 | 0.223 | 0.384 | 0.799 | 1.624 | 21.1 | 34 | 0 |
| go-redis-autopipeline | 100 | 202366 | 0.483 | 0.592 | 0.775 | 1.716 | 2.117 | 7.8 | 59 | 0 |
| go-redis-autopipeline | 1000 | 447167 | 2.241 | 2.937 | 3.663 | 5.499 | 10.295 | 6.8 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11762 | 0.078 | 0.100 | 0.172 | 0.381 | 1.607 | 65.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 38212 | 0.256 | 0.333 | 0.555 | 0.986 | 1.767 | 80.5 | 34 | 0 |
| go-redis-autopipeline | 100 | 39802 | 2.407 | 2.675 | 4.815 | 5.575 | 7.415 | 80.7 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7612 | 0.119 | 0.164 | 0.288 | 0.525 | 1.657 | 120.3 | 33 | 0 |
| glide | 10 | 32718 | 0.294 | 0.388 | 0.552 | 0.873 | 1.787 | 56.7 | 38 | 0 |
| glide | 100 | 35418 | 2.923 | 3.267 | 4.143 | 4.839 | 6.003 | 55.5 | 64 | 0 |
| glide | 1000 | 34656 | 28.303 | 32.719 | 39.711 | 44.255 | 51.391 | 58.8 | 208 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6659 | 0.136 | 0.185 | 0.337 | 0.684 | 1.616 | 122.7 | 32 | 0 |
| glide | 10 | 26295 | 0.358 | 0.533 | 0.962 | 1.554 | 2.663 | 90.7 | 33 | 0 |
| glide | 100 | 29020 | 3.441 | 4.383 | 5.419 | 6.387 | 8.759 | 84.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12560 | 0.075 | 0.089 | 0.134 | 0.262 | 1.646 | 50.6 | 35 | 0 |
| valkey-go+tls | 10 | 35733 | 0.262 | 0.433 | 0.575 | 0.791 | 2.063 | 63.4 | 42 | 0 |
| valkey-go+tls | 100 | 37673 | 2.279 | 5.367 | 7.495 | 8.575 | 10.183 | 63.1 | 66 | 0 |
| valkey-go+tls | 1000 | 37031 | 31.007 | 43.743 | 51.135 | 53.951 | 55.327 | 64.6 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10143 | 0.092 | 0.116 | 0.184 | 0.355 | 1.437 | 65.8 | 36 | 0 |
| valkey-go+tls | 10 | 24233 | 0.369 | 0.663 | 1.127 | 1.498 | 2.493 | 92.8 | 41 | 0 |
| valkey-go+tls | 100 | 29022 | 2.839 | 7.367 | 9.999 | 12.199 | 14.135 | 82.2 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 12365 | 0.077 | 0.089 | 0.134 | 0.289 | 1.658 | 53.8 | 30 | 0 |
| go-redis+tls | 10 | 37181 | 0.267 | 0.335 | 0.443 | 0.699 | 1.932 | 65.1 | 33 | 0 |
| go-redis+tls | 100 | 37972 | 2.535 | 2.771 | 5.051 | 5.659 | 8.367 | 63.7 | 73 | 0 |
| go-redis+tls | 1000 | 37583 | 26.527 | 28.799 | 31.103 | 32.863 | 35.679 | 75.6 | 252 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9909 | 0.095 | 0.114 | 0.197 | 0.415 | 1.546 | 73.2 | 28 | 0 |
| go-redis+tls | 10 | 29209 | 0.341 | 0.423 | 0.612 | 0.980 | 1.905 | 84.3 | 36 | 0 |
| go-redis+tls | 100 | 29444 | 3.259 | 3.579 | 6.459 | 7.179 | 8.959 | 81.3 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7915 | 0.116 | 0.152 | 0.262 | 0.471 | 1.793 | 115.7 | 33 | 0 |
| go-redis-autopipeline+tls | 10 | 52703 | 0.178 | 0.235 | 0.396 | 0.863 | 1.662 | 21.5 | 33 | 0 |
| go-redis-autopipeline+tls | 100 | 193126 | 0.506 | 0.620 | 0.809 | 1.610 | 2.157 | 7.9 | 59 | 0 |
| go-redis-autopipeline+tls | 1000 | 424509 | 2.359 | 3.097 | 3.767 | 5.331 | 7.455 | 6.9 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9859 | 0.095 | 0.114 | 0.192 | 0.432 | 1.480 | 73.5 | 30 | 0 |
| go-redis-autopipeline+tls | 10 | 29164 | 0.341 | 0.425 | 0.613 | 0.952 | 2.013 | 84.6 | 38 | 0 |
| go-redis-autopipeline+tls | 100 | 29339 | 3.245 | 3.777 | 6.459 | 7.571 | 9.823 | 80.9 | 81 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15009 | 0.063 | 0.075 | 0.106 | 0.227 | 1.511 | 49.6 | 34 | 0 |
| valkey-go/flush-100us | 10 | 60674 | 0.158 | 0.241 | 0.318 | 1.320 | 1.886 | 36.3 | 39 | 0 |
| valkey-go/flush-100us | 100 | 412556 | 0.234 | 0.335 | 0.463 | 0.697 | 1.726 | 8.8 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 975023 | 1.045 | 1.398 | 1.757 | 2.693 | 4.983 | 5.3 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13186 | 0.072 | 0.086 | 0.129 | 0.264 | 1.415 | 54.2 | 37 | 0 |
| valkey-go/flush-100us | 10 | 44677 | 0.215 | 0.324 | 0.441 | 0.722 | 2.105 | 55.4 | 40 | 0 |
| valkey-go/flush-100us | 100 | 107542 | 0.943 | 1.201 | 1.687 | 2.237 | 4.375 | 23.2 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12703 | 0.075 | 0.086 | 0.131 | 0.281 | 1.625 | 49.3 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 49152 | 0.197 | 0.291 | 0.381 | 0.543 | 1.896 | 45.3 | 43 | 0 |
| valkey-go/flush-100us+tls | 100 | 215972 | 0.419 | 0.676 | 1.170 | 1.860 | 2.981 | 14.5 | 68 | 0 |
| valkey-go/flush-100us+tls | 1000 | 216490 | 3.405 | 10.223 | 11.887 | 14.079 | 16.399 | 14.6 | 281 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10289 | 0.092 | 0.111 | 0.173 | 0.333 | 1.602 | 64.1 | 37 | 0 |
| valkey-go/flush-100us+tls | 10 | 34927 | 0.278 | 0.389 | 0.529 | 0.793 | 1.935 | 67.6 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 44660 | 2.028 | 4.037 | 5.335 | 7.887 | 9.135 | 55.7 | 70 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15645 | 0.061 | 0.071 | 0.107 | 0.230 | 1.491 | 46.3 | 35 | 0 |
| valkey-go | 10 | 72515 | 0.133 | 0.189 | 0.256 | 0.356 | 1.383 | 43.7 | 41 | 0 |
| valkey-go | 100 | 325564 | 0.304 | 0.420 | 0.548 | 0.938 | 2.167 | 16.0 | 67 | 0 |
| valkey-go | 1000 | 884574 | 1.162 | 1.526 | 1.859 | 3.057 | 7.423 | 7.2 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12795 | 0.073 | 0.090 | 0.154 | 0.356 | 2.663 | 54.4 | 35 | 0 |
| valkey-go | 10 | 46775 | 0.207 | 0.290 | 0.427 | 0.701 | 1.759 | 61.7 | 40 | 0 |
| valkey-go | 100 | 103950 | 0.987 | 1.261 | 1.706 | 2.565 | 3.775 | 29.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13802 | 0.068 | 0.082 | 0.127 | 0.283 | 1.542 | 55.8 | 28 | 0 |
| go-redis | 10 | 51281 | 0.192 | 0.257 | 0.361 | 0.622 | 1.880 | 64.0 | 32 | 0 |
| go-redis | 100 | 52200 | 1.839 | 2.067 | 3.653 | 4.111 | 5.763 | 73.8 | 68 | 0 |
| go-redis | 1000 | 45647 | 19.263 | 31.695 | 35.327 | 37.823 | 41.407 | 92.7 | 279 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11800 | 0.078 | 0.098 | 0.180 | 0.403 | 1.649 | 64.8 | 29 | 0 |
| go-redis | 10 | 38120 | 0.256 | 0.335 | 0.574 | 0.964 | 1.784 | 81.6 | 36 | 0 |
| go-redis | 100 | 39222 | 2.439 | 2.751 | 4.859 | 5.703 | 8.159 | 87.9 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9696 | 0.094 | 0.128 | 0.226 | 0.444 | 1.616 | 107.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 59903 | 0.155 | 0.210 | 0.371 | 0.805 | 3.037 | 21.2 | 33 | 0 |
| go-redis-autopipeline | 100 | 208699 | 0.469 | 0.575 | 0.749 | 1.506 | 2.183 | 7.9 | 58 | 0 |
| go-redis-autopipeline | 1000 | 453997 | 2.223 | 2.889 | 3.535 | 5.183 | 7.307 | 6.7 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11730 | 0.078 | 0.100 | 0.173 | 0.388 | 1.566 | 67.7 | 27 | 0 |
| go-redis-autopipeline | 10 | 38510 | 0.254 | 0.330 | 0.558 | 0.969 | 2.309 | 80.2 | 35 | 0 |
| go-redis-autopipeline | 100 | 40323 | 2.379 | 2.595 | 4.739 | 5.275 | 7.343 | 83.4 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7558 | 0.120 | 0.162 | 0.278 | 0.471 | 1.969 | 120.4 | 32 | 0 |
| glide | 10 | 33648 | 0.286 | 0.376 | 0.533 | 0.788 | 1.995 | 56.0 | 38 | 0 |
| glide | 100 | 35181 | 2.909 | 3.329 | 4.451 | 4.979 | 6.107 | 57.5 | 64 | 0 |
| glide | 1000 | 33533 | 29.519 | 33.727 | 39.615 | 42.463 | 43.455 | 59.9 | 227 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6413 | 0.142 | 0.193 | 0.346 | 0.642 | 1.771 | 128.3 | 32 | 0 |
| glide | 10 | 25025 | 0.378 | 0.558 | 0.970 | 1.549 | 2.373 | 94.1 | 36 | 0 |
| glide | 100 | 28629 | 3.497 | 4.431 | 5.439 | 6.247 | 7.847 | 85.2 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12669 | 0.075 | 0.088 | 0.134 | 0.283 | 1.582 | 51.4 | 35 | 0 |
| valkey-go+tls | 10 | 35077 | 0.266 | 0.443 | 0.602 | 0.797 | 1.970 | 64.0 | 42 | 0 |
| valkey-go+tls | 100 | 36961 | 2.377 | 5.307 | 7.367 | 8.599 | 9.487 | 64.4 | 66 | 0 |
| valkey-go+tls | 1000 | 36606 | 29.807 | 43.903 | 51.807 | 55.263 | 56.383 | 66.1 | 196 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10332 | 0.092 | 0.110 | 0.178 | 0.347 | 1.827 | 64.1 | 38 | 0 |
| valkey-go+tls | 10 | 28428 | 0.332 | 0.535 | 0.707 | 0.923 | 1.941 | 81.1 | 41 | 0 |
| valkey-go+tls | 100 | 29358 | 2.987 | 6.619 | 9.303 | 10.919 | 11.999 | 82.0 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 11479 | 0.079 | 0.106 | 0.185 | 0.300 | 1.584 | 60.0 | 30 | 0 |
| go-redis+tls | 10 | 35731 | 0.274 | 0.360 | 0.497 | 0.723 | 1.723 | 69.7 | 33 | 0 |
| go-redis+tls | 100 | 37019 | 2.595 | 2.853 | 5.171 | 5.751 | 7.715 | 66.2 | 73 | 0 |
| go-redis+tls | 1000 | 37268 | 26.703 | 29.071 | 31.711 | 33.823 | 36.927 | 77.1 | 254 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9934 | 0.094 | 0.113 | 0.196 | 0.432 | 1.549 | 74.0 | 30 | 0 |
| go-redis+tls | 10 | 28978 | 0.342 | 0.428 | 0.632 | 0.986 | 2.131 | 86.4 | 38 | 0 |
| go-redis+tls | 100 | 29474 | 3.231 | 3.723 | 6.415 | 7.579 | 9.495 | 84.7 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7951 | 0.117 | 0.150 | 0.259 | 0.456 | 1.729 | 113.5 | 30 | 0 |
| go-redis-autopipeline+tls | 10 | 51788 | 0.180 | 0.239 | 0.407 | 0.867 | 1.733 | 21.9 | 33 | 0 |
| go-redis-autopipeline+tls | 100 | 185759 | 0.524 | 0.647 | 0.856 | 1.572 | 1.989 | 8.3 | 60 | 0 |
| go-redis-autopipeline+tls | 1000 | 410977 | 2.399 | 3.201 | 3.853 | 5.911 | 8.711 | 7.2 | 310 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9758 | 0.095 | 0.115 | 0.201 | 0.449 | 1.553 | 75.3 | 29 | 0 |
| go-redis-autopipeline+tls | 10 | 28782 | 0.344 | 0.433 | 0.628 | 0.967 | 1.892 | 85.2 | 35 | 0 |
| go-redis-autopipeline+tls | 100 | 29273 | 3.255 | 3.767 | 6.455 | 7.531 | 9.911 | 84.2 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6601 | 0.140 | 0.184 | 0.349 | 0.584 | 1.737 | 128.2 | 37 | 0 |
| glide+tls | 10 | 25946 | 0.345 | 0.597 | 0.959 | 1.451 | 3.297 | 68.5 | 38 | 0 |
| glide+tls | 100 | 32221 | 3.023 | 3.581 | 4.887 | 11.015 | 13.543 | 64.2 | 67 | 0 |
| glide+tls | 1000 | 31725 | 31.055 | 35.167 | 40.575 | 44.735 | 47.199 | 62.3 | 224 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5646 | 0.163 | 0.220 | 0.401 | 0.724 | 1.692 | 136.2 | 35 | 0 |
| glide+tls | 10 | 23086 | 0.405 | 0.592 | 0.992 | 1.555 | 2.151 | 98.6 | 38 | 0 |
| glide+tls | 100 | 25701 | 3.849 | 4.655 | 5.843 | 6.755 | 9.239 | 90.0 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12576 | 0.075 | 0.093 | 0.149 | 0.294 | 1.628 | 50.6 | 35 | 0 |
| valkey-go/flush-100us+tls | 10 | 48471 | 0.199 | 0.295 | 0.393 | 0.629 | 1.884 | 45.3 | 44 | 0 |
| valkey-go/flush-100us+tls | 100 | 203087 | 0.444 | 0.737 | 1.267 | 1.846 | 2.927 | 15.1 | 69 | 0 |
| valkey-go/flush-100us+tls | 1000 | 218996 | 4.407 | 7.447 | 9.839 | 11.663 | 15.527 | 14.9 | 244 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10438 | 0.091 | 0.110 | 0.174 | 0.348 | 3.137 | 63.1 | 38 | 0 |
| valkey-go/flush-100us+tls | 10 | 32806 | 0.294 | 0.422 | 0.573 | 0.826 | 2.587 | 70.8 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 44030 | 1.974 | 4.363 | 5.939 | 6.759 | 8.011 | 57.8 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15442 | 0.062 | 0.073 | 0.108 | 0.226 | 1.774 | 47.5 | 37 | 0 |
| valkey-go | 10 | 74072 | 0.130 | 0.186 | 0.254 | 0.347 | 1.485 | 42.5 | 39 | 0 |
| valkey-go | 100 | 331407 | 0.299 | 0.412 | 0.535 | 0.813 | 1.952 | 15.4 | 66 | 0 |
| valkey-go | 1000 | 905561 | 1.134 | 1.479 | 1.848 | 2.975 | 6.827 | 6.5 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13410 | 0.071 | 0.081 | 0.127 | 0.259 | 1.821 | 51.6 | 35 | 0 |
| valkey-go | 10 | 48598 | 0.201 | 0.277 | 0.390 | 0.646 | 1.843 | 58.7 | 38 | 0 |
| valkey-go | 100 | 106404 | 0.965 | 1.227 | 1.700 | 2.385 | 3.513 | 28.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14530 | 0.066 | 0.076 | 0.116 | 0.250 | 1.533 | 53.5 | 28 | 0 |
| go-redis | 10 | 50861 | 0.191 | 0.260 | 0.403 | 0.828 | 6.651 | 63.1 | 34 | 0 |
| go-redis | 100 | 52978 | 1.815 | 2.008 | 3.609 | 4.039 | 6.027 | 67.7 | 72 | 0 |
| go-redis | 1000 | 52734 | 18.799 | 20.143 | 23.135 | 27.151 | 31.199 | 79.4 | 249 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11935 | 0.077 | 0.099 | 0.164 | 0.372 | 1.641 | 64.0 | 28 | 0 |
| go-redis | 10 | 38940 | 0.252 | 0.326 | 0.549 | 0.904 | 1.502 | 79.4 | 35 | 0 |
| go-redis | 100 | 39683 | 2.397 | 2.791 | 4.779 | 5.751 | 7.731 | 84.3 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9822 | 0.094 | 0.126 | 0.199 | 0.401 | 2.199 | 105.8 | 27 | 0 |
| go-redis-autopipeline | 10 | 59514 | 0.156 | 0.209 | 0.354 | 0.875 | 6.087 | 21.5 | 32 | 0 |
| go-redis-autopipeline | 100 | 191880 | 0.501 | 0.625 | 0.992 | 2.151 | 6.675 | 8.6 | 58 | 0 |
| go-redis-autopipeline | 1000 | 407983 | 2.403 | 3.315 | 4.531 | 6.011 | 9.727 | 7.2 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10153 | 0.081 | 0.128 | 0.254 | 1.516 | 2.067 | 72.6 | 27 | 0 |
| go-redis-autopipeline | 10 | 37118 | 0.259 | 0.350 | 0.619 | 1.282 | 2.079 | 82.2 | 36 | 0 |
| go-redis-autopipeline | 100 | 39539 | 2.415 | 2.711 | 4.803 | 5.375 | 8.495 | 85.5 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7362 | 0.129 | 0.164 | 0.280 | 0.463 | 1.892 | 124.7 | 33 | 0 |
| glide | 10 | 29414 | 0.312 | 0.453 | 0.810 | 1.794 | 3.035 | 61.1 | 37 | 0 |
| glide | 100 | 35627 | 2.875 | 3.293 | 4.371 | 4.935 | 5.863 | 57.2 | 64 | 0 |
| glide | 1000 | 34850 | 28.719 | 31.711 | 34.399 | 56.415 | 61.599 | 59.2 | 232 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6478 | 0.139 | 0.193 | 0.357 | 1.083 | 2.411 | 127.0 | 33 | 0 |
| glide | 10 | 24691 | 0.379 | 0.570 | 1.063 | 1.742 | 2.693 | 95.0 | 35 | 0 |
| glide | 100 | 27227 | 3.641 | 4.739 | 6.027 | 7.283 | 10.671 | 89.4 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12135 | 0.076 | 0.098 | 0.163 | 0.319 | 1.997 | 52.5 | 35 | 0 |
| valkey-go+tls | 10 | 34937 | 0.267 | 0.444 | 0.604 | 0.881 | 2.849 | 64.5 | 43 | 0 |
| valkey-go+tls | 100 | 36455 | 2.149 | 6.027 | 8.247 | 9.511 | 10.855 | 64.9 | 66 | 0 |
| valkey-go+tls | 1000 | 36140 | 25.695 | 47.583 | 57.023 | 63.359 | 64.415 | 65.3 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 10525 | 0.091 | 0.106 | 0.167 | 0.320 | 1.513 | 63.1 | 37 | 0 |
| valkey-go+tls | 10 | 28309 | 0.333 | 0.538 | 0.719 | 0.924 | 2.069 | 81.6 | 43 | 0 |
| valkey-go+tls | 100 | 28770 | 3.115 | 6.775 | 9.279 | 10.895 | 12.415 | 82.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 11582 | 0.079 | 0.101 | 0.174 | 0.469 | 2.583 | 58.6 | 30 | 0 |
| go-redis+tls | 10 | 37371 | 0.266 | 0.334 | 0.452 | 0.682 | 2.189 | 66.4 | 33 | 0 |
| go-redis+tls | 100 | 37800 | 2.547 | 2.779 | 5.051 | 5.519 | 7.815 | 66.2 | 73 | 0 |
| go-redis+tls | 1000 | 35557 | 27.599 | 31.391 | 37.247 | 40.191 | 44.767 | 80.4 | 230 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9430 | 0.096 | 0.121 | 0.227 | 1.118 | 1.854 | 76.9 | 30 | 0 |
| go-redis+tls | 10 | 29020 | 0.341 | 0.428 | 0.642 | 1.098 | 2.333 | 87.1 | 36 | 0 |
| go-redis+tls | 100 | 29232 | 3.245 | 3.863 | 6.443 | 7.547 | 9.807 | 83.9 | 79 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7647 | 0.119 | 0.161 | 0.276 | 1.365 | 4.647 | 117.1 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 49027 | 0.190 | 0.261 | 0.458 | 1.303 | 2.249 | 22.0 | 31 | 0 |
| go-redis-autopipeline+tls | 100 | 173006 | 0.558 | 0.698 | 1.058 | 1.986 | 3.679 | 8.3 | 58 | 0 |
| go-redis-autopipeline+tls | 1000 | 394832 | 2.465 | 3.327 | 4.651 | 6.179 | 8.727 | 7.2 | 311 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9177 | 0.097 | 0.129 | 0.252 | 1.273 | 2.085 | 77.3 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 28363 | 0.346 | 0.442 | 0.672 | 1.271 | 3.517 | 87.7 | 36 | 0 |
| go-redis-autopipeline+tls | 100 | 27153 | 3.383 | 4.667 | 6.831 | 8.791 | 11.471 | 91.9 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5593 | 0.158 | 0.234 | 0.465 | 1.594 | 3.167 | 144.8 | 34 | 0 |
| glide+tls | 10 | 29469 | 0.324 | 0.470 | 0.667 | 1.278 | 1.965 | 62.5 | 38 | 0 |
| glide+tls | 100 | 32665 | 3.013 | 3.631 | 4.787 | 5.363 | 6.339 | 59.3 | 64 | 0 |
| glide+tls | 1000 | 31754 | 31.327 | 34.783 | 36.927 | 38.751 | 40.031 | 62.5 | 229 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 4924 | 0.175 | 0.272 | 0.592 | 1.633 | 3.141 | 148.3 | 36 | 0 |
| glide+tls | 10 | 22606 | 0.406 | 0.618 | 1.105 | 1.894 | 2.663 | 100.4 | 38 | 0 |
| glide+tls | 100 | 25241 | 3.945 | 4.795 | 5.887 | 6.571 | 7.663 | 91.6 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15242 | 0.062 | 0.072 | 0.115 | 0.258 | 1.468 | 47.6 | 37 | 0 |
| valkey-go/flush-100us | 10 | 59897 | 0.159 | 0.244 | 0.332 | 1.373 | 3.613 | 36.6 | 43 | 0 |
| valkey-go/flush-100us | 100 | 400165 | 0.239 | 0.349 | 0.500 | 0.770 | 4.967 | 9.2 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 979512 | 1.040 | 1.399 | 1.704 | 2.817 | 5.219 | 5.7 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12324 | 0.073 | 0.097 | 0.170 | 0.610 | 1.970 | 56.3 | 35 | 0 |
| valkey-go/flush-100us | 10 | 41106 | 0.228 | 0.355 | 0.556 | 1.514 | 2.649 | 58.1 | 40 | 0 |
| valkey-go/flush-100us | 100 | 107548 | 0.945 | 1.196 | 1.688 | 2.237 | 3.561 | 23.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15396 | 0.062 | 0.074 | 0.111 | 0.243 | 1.342 | 47.8 | 35 | 0 |
| valkey-go | 10 | 71217 | 0.135 | 0.193 | 0.273 | 0.451 | 1.939 | 44.1 | 40 | 0 |
| valkey-go | 100 | 304297 | 0.316 | 0.455 | 0.722 | 1.680 | 2.691 | 16.4 | 67 | 0 |
| valkey-go | 1000 | 838922 | 1.194 | 1.631 | 2.425 | 3.703 | 6.387 | 7.2 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12555 | 0.072 | 0.094 | 0.183 | 0.388 | 2.961 | 55.1 | 37 | 0 |
| valkey-go | 10 | 47375 | 0.205 | 0.285 | 0.413 | 0.710 | 1.699 | 61.1 | 39 | 0 |
| valkey-go | 100 | 104893 | 0.980 | 1.251 | 1.623 | 2.359 | 3.449 | 29.3 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13808 | 0.067 | 0.082 | 0.139 | 0.535 | 1.945 | 55.4 | 30 | 0 |
| go-redis | 10 | 49324 | 0.195 | 0.272 | 0.417 | 1.209 | 1.950 | 65.1 | 34 | 0 |
| go-redis | 100 | 51658 | 1.819 | 2.315 | 3.651 | 4.571 | 6.211 | 69.0 | 70 | 0 |
| go-redis | 1000 | 49229 | 19.439 | 23.807 | 29.311 | 31.151 | 34.783 | 82.2 | 237 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11864 | 0.078 | 0.098 | 0.167 | 0.364 | 1.576 | 64.9 | 27 | 0 |
| go-redis | 10 | 38235 | 0.253 | 0.333 | 0.587 | 1.319 | 2.673 | 79.8 | 36 | 0 |
| go-redis | 100 | 37592 | 2.439 | 3.433 | 4.923 | 6.475 | 8.295 | 86.6 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9592 | 0.094 | 0.130 | 0.221 | 0.521 | 1.625 | 107.8 | 29 | 0 |
| go-redis-autopipeline | 10 | 60188 | 0.154 | 0.210 | 0.366 | 0.932 | 3.189 | 21.0 | 33 | 0 |
| go-redis-autopipeline | 100 | 202643 | 0.480 | 0.590 | 0.853 | 1.664 | 2.281 | 8.3 | 59 | 0 |
| go-redis-autopipeline | 1000 | 450767 | 2.233 | 2.897 | 3.827 | 5.595 | 8.535 | 6.9 | 307 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12025 | 0.077 | 0.097 | 0.173 | 0.400 | 1.533 | 63.5 | 27 | 0 |
| go-redis-autopipeline | 10 | 37184 | 0.258 | 0.351 | 0.634 | 1.312 | 2.357 | 81.9 | 35 | 0 |
| go-redis-autopipeline | 100 | 39977 | 2.385 | 2.751 | 4.759 | 5.479 | 7.679 | 84.2 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7609 | 0.126 | 0.155 | 0.256 | 0.437 | 1.139 | 122.8 | 33 | 0 |
| glide | 10 | 30763 | 0.307 | 0.422 | 0.650 | 1.458 | 2.229 | 59.3 | 38 | 0 |
| glide | 100 | 33432 | 2.987 | 3.871 | 4.827 | 5.671 | 6.699 | 59.5 | 65 | 0 |
| glide | 1000 | 34901 | 28.479 | 31.871 | 36.223 | 40.095 | 44.575 | 59.8 | 235 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6513 | 0.137 | 0.191 | 0.362 | 1.291 | 1.627 | 125.9 | 33 | 0 |
| glide | 10 | 25583 | 0.366 | 0.552 | 0.996 | 1.603 | 3.741 | 93.3 | 36 | 0 |
| glide | 100 | 28989 | 3.429 | 4.463 | 5.687 | 7.075 | 10.119 | 85.2 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 12068 | 0.076 | 0.100 | 0.169 | 0.370 | 3.003 | 53.5 | 33 | 0 |
| valkey-go+tls | 10 | 35537 | 0.264 | 0.436 | 0.581 | 0.736 | 1.783 | 63.3 | 42 | 0 |
| valkey-go+tls | 100 | 36815 | 2.277 | 5.619 | 8.055 | 9.527 | 11.791 | 65.6 | 64 | 0 |
| valkey-go+tls | 1000 | 35996 | 31.999 | 44.831 | 54.847 | 59.519 | 61.567 | 67.5 | 197 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go+tls | 1 | 9638 | 0.094 | 0.124 | 0.234 | 0.646 | 3.057 | 68.7 | 37 | 0 |
| valkey-go+tls | 10 | 27992 | 0.335 | 0.543 | 0.744 | 1.086 | 2.195 | 82.5 | 40 | 0 |
| valkey-go+tls | 100 | 29038 | 2.867 | 7.215 | 10.407 | 12.031 | 13.127 | 83.0 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 11876 | 0.077 | 0.099 | 0.170 | 0.582 | 3.263 | 57.2 | 28 | 0 |
| go-redis+tls | 10 | 36441 | 0.268 | 0.346 | 0.505 | 1.030 | 3.221 | 66.6 | 33 | 0 |
| go-redis+tls | 100 | 36619 | 2.593 | 3.099 | 5.159 | 6.183 | 8.191 | 68.9 | 71 | 0 |
| go-redis+tls | 1000 | 36406 | 27.151 | 30.207 | 34.591 | 37.535 | 40.831 | 79.6 | 252 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis+tls | 1 | 9739 | 0.095 | 0.119 | 0.222 | 0.626 | 1.684 | 75.2 | 27 | 0 |
| go-redis+tls | 10 | 27857 | 0.350 | 0.451 | 0.718 | 1.608 | 3.559 | 88.6 | 35 | 0 |
| go-redis+tls | 100 | 28313 | 3.291 | 4.299 | 6.551 | 8.471 | 10.911 | 88.0 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 7568 | 0.120 | 0.168 | 0.297 | 1.083 | 3.417 | 117.0 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 50874 | 0.181 | 0.252 | 0.439 | 1.338 | 3.363 | 21.9 | 32 | 0 |
| go-redis-autopipeline+tls | 100 | 169117 | 0.554 | 0.780 | 1.182 | 1.882 | 3.669 | 8.8 | 58 | 0 |
| go-redis-autopipeline+tls | 1000 | 393678 | 2.463 | 3.369 | 4.719 | 8.107 | 9.479 | 7.7 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline+tls | 1 | 9816 | 0.095 | 0.117 | 0.209 | 0.451 | 1.187 | 74.0 | 28 | 0 |
| go-redis-autopipeline+tls | 10 | 28569 | 0.346 | 0.438 | 0.648 | 0.969 | 2.437 | 87.1 | 37 | 0 |
| go-redis-autopipeline+tls | 100 | 29295 | 3.243 | 3.791 | 6.451 | 7.587 | 11.695 | 87.2 | 78 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 6147 | 0.143 | 0.214 | 0.403 | 1.502 | 5.755 | 135.5 | 34 | 0 |
| glide+tls | 10 | 27972 | 0.341 | 0.496 | 0.685 | 1.166 | 2.245 | 65.1 | 38 | 0 |
| glide+tls | 100 | 31593 | 3.119 | 3.759 | 4.979 | 6.151 | 10.111 | 61.6 | 64 | 0 |
| glide+tls | 1000 | 31500 | 31.423 | 35.167 | 40.319 | 45.503 | 46.943 | 63.2 | 222 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide+tls | 1 | 5525 | 0.163 | 0.231 | 0.427 | 1.087 | 1.751 | 140.2 | 34 | 0 |
| glide+tls | 10 | 22861 | 0.405 | 0.608 | 1.032 | 1.681 | 2.715 | 98.3 | 37 | 0 |
| glide+tls | 100 | 24780 | 3.973 | 5.043 | 6.335 | 7.787 | 10.607 | 92.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15402 | 0.062 | 0.074 | 0.109 | 0.244 | 1.501 | 48.9 | 36 | 0 |
| valkey-go/flush-100us | 10 | 59409 | 0.160 | 0.247 | 0.335 | 1.339 | 2.261 | 37.3 | 40 | 0 |
| valkey-go/flush-100us | 100 | 400679 | 0.237 | 0.349 | 0.511 | 1.071 | 2.335 | 9.0 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 916158 | 1.090 | 1.512 | 2.229 | 3.089 | 5.051 | 5.6 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12726 | 0.072 | 0.093 | 0.162 | 0.391 | 1.679 | 55.3 | 36 | 0 |
| valkey-go/flush-100us | 10 | 43950 | 0.217 | 0.330 | 0.473 | 1.048 | 1.868 | 56.0 | 41 | 0 |
| valkey-go/flush-100us | 100 | 102589 | 0.971 | 1.302 | 1.993 | 2.985 | 6.003 | 24.4 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 12043 | 0.076 | 0.098 | 0.169 | 0.402 | 1.659 | 53.8 | 33 | 0 |
| valkey-go/flush-100us+tls | 10 | 47074 | 0.202 | 0.304 | 0.443 | 1.322 | 4.035 | 46.5 | 42 | 0 |
| valkey-go/flush-100us+tls | 100 | 195804 | 0.446 | 0.795 | 1.407 | 2.275 | 3.029 | 15.4 | 69 | 0 |
| valkey-go/flush-100us+tls | 1000 | 208788 | 4.679 | 8.303 | 9.935 | 10.767 | 11.247 | 15.3 | 271 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us+tls | 1 | 10569 | 0.090 | 0.108 | 0.170 | 0.349 | 1.642 | 63.3 | 36 | 0 |
| valkey-go/flush-100us+tls | 10 | 33768 | 0.284 | 0.405 | 0.595 | 0.933 | 5.643 | 69.8 | 40 | 0 |
| valkey-go/flush-100us+tls | 100 | 43514 | 2.141 | 3.953 | 5.579 | 6.787 | 9.567 | 56.1 | 68 | 0 |
