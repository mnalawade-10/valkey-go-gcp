mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14300 | 0.067 | 0.077 | 0.112 | 0.270 | 1.623 | 53.8 | 27 | 0 |
| go-redis | 10 | 52804 | 0.188 | 0.245 | 0.337 | 0.585 | 1.516 | 61.4 | 33 | 0 |
| go-redis | 100 | 53341 | 1.798 | 1.997 | 3.583 | 4.029 | 5.739 | 66.7 | 68 | 0 |
| go-redis | 1000 | 44432 | 19.039 | 31.743 | 38.655 | 42.495 | 47.359 | 92.2 | 271 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11927 | 0.078 | 0.097 | 0.161 | 0.369 | 1.627 | 63.8 | 27 | 0 |
| go-redis | 10 | 38623 | 0.254 | 0.327 | 0.551 | 0.913 | 1.654 | 80.3 | 34 | 0 |
| go-redis | 100 | 40086 | 2.381 | 2.659 | 4.755 | 5.555 | 7.455 | 82.8 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9896 | 0.093 | 0.124 | 0.186 | 0.351 | 1.677 | 103.3 | 29 | 0 |
| go-redis-autopipeline | 10 | 59927 | 0.156 | 0.207 | 0.349 | 0.796 | 1.712 | 21.2 | 30 | 0 |
| go-redis-autopipeline | 100 | 211600 | 0.464 | 0.560 | 0.714 | 1.537 | 2.919 | 8.1 | 57 | 0 |
| go-redis-autopipeline | 1000 | 464910 | 2.225 | 2.755 | 3.299 | 5.167 | 7.303 | 6.8 | 310 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11960 | 0.078 | 0.097 | 0.161 | 0.360 | 1.567 | 63.9 | 28 | 0 |
| go-redis-autopipeline | 10 | 38559 | 0.254 | 0.330 | 0.558 | 0.941 | 1.684 | 80.9 | 35 | 0 |
| go-redis-autopipeline | 100 | 39958 | 2.393 | 2.649 | 4.775 | 5.451 | 7.355 | 83.0 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7695 | 0.121 | 0.155 | 0.260 | 0.472 | 1.746 | 119.7 | 33 | 0 |
| glide | 10 | 33438 | 0.286 | 0.384 | 0.547 | 0.917 | 2.189 | 56.1 | 38 | 0 |
| glide | 100 | 34431 | 2.923 | 3.523 | 4.395 | 4.935 | 5.747 | 57.4 | 64 | 0 |
| glide | 1000 | 35435 | 28.127 | 31.455 | 36.735 | 41.471 | 41.951 | 58.2 | 234 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6839 | 0.134 | 0.176 | 0.318 | 0.593 | 1.585 | 120.6 | 32 | 0 |
| glide | 10 | 25656 | 0.370 | 0.542 | 0.953 | 1.445 | 2.243 | 92.0 | 33 | 0 |
| glide | 100 | 28883 | 3.433 | 4.423 | 5.483 | 6.643 | 10.151 | 85.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15200 | 0.062 | 0.075 | 0.110 | 0.267 | 1.537 | 48.2 | 28 | 0 |
| valkey-go/no-autopipe | 10 | 52616 | 0.188 | 0.244 | 0.333 | 0.469 | 1.658 | 51.1 | 53 | 0 |
| valkey-go/no-autopipe | 100 | 54013 | 1.790 | 1.910 | 3.569 | 3.833 | 6.399 | 54.7 | 178 | 0 |
| valkey-go/no-autopipe | 1000 | 49292 | 20.047 | 22.143 | 26.143 | 42.271 | 52.511 | 55.2 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 13038 | 0.072 | 0.086 | 0.135 | 0.311 | 1.507 | 53.5 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 39323 | 0.253 | 0.316 | 0.429 | 0.667 | 1.678 | 64.4 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 39827 | 2.409 | 2.625 | 4.803 | 5.319 | 6.863 | 66.4 | 261 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12079 | 0.078 | 0.097 | 0.144 | 0.277 | 1.587 | 71.7 | 34 | 0 |
| valkey-go/always-pipe | 10 | 72917 | 0.133 | 0.187 | 0.253 | 0.375 | 1.892 | 43.8 | 40 | 0 |
| valkey-go/always-pipe | 100 | 327998 | 0.302 | 0.415 | 0.535 | 0.909 | 2.197 | 15.8 | 65 | 0 |
| valkey-go/always-pipe | 1000 | 908018 | 1.128 | 1.483 | 1.883 | 3.067 | 6.591 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9448 | 0.102 | 0.124 | 0.194 | 0.359 | 2.735 | 93.8 | 37 | 0 |
| valkey-go/always-pipe | 10 | 46474 | 0.209 | 0.292 | 0.423 | 0.705 | 2.243 | 61.0 | 39 | 0 |
| valkey-go/always-pipe | 100 | 103981 | 0.988 | 1.257 | 1.671 | 2.613 | 3.927 | 29.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 14976 | 0.063 | 0.075 | 0.115 | 0.284 | 1.677 | 47.3 | 37 | 0 |
| valkey-go/flush-20us | 10 | 70242 | 0.137 | 0.196 | 0.270 | 0.426 | 1.381 | 46.4 | 40 | 0 |
| valkey-go/flush-20us | 100 | 336650 | 0.295 | 0.404 | 0.525 | 0.905 | 3.019 | 14.3 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 883498 | 1.149 | 1.540 | 2.077 | 3.263 | 5.723 | 6.6 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12380 | 0.075 | 0.093 | 0.155 | 0.334 | 1.486 | 54.7 | 37 | 0 |
| valkey-go/flush-20us | 10 | 45964 | 0.211 | 0.296 | 0.429 | 0.710 | 1.894 | 63.8 | 41 | 0 |
| valkey-go/flush-20us | 100 | 104949 | 0.978 | 1.244 | 1.726 | 2.435 | 3.717 | 27.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15207 | 0.063 | 0.075 | 0.109 | 0.233 | 1.511 | 46.9 | 35 | 0 |
| valkey-go/flush-100us | 10 | 59588 | 0.161 | 0.246 | 0.333 | 1.317 | 1.879 | 37.4 | 41 | 0 |
| valkey-go/flush-100us | 100 | 402959 | 0.239 | 0.340 | 0.477 | 0.784 | 2.193 | 9.2 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 966970 | 1.051 | 1.417 | 1.804 | 2.901 | 5.039 | 5.7 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12911 | 0.073 | 0.088 | 0.140 | 0.304 | 2.453 | 53.6 | 34 | 0 |
| valkey-go/flush-100us | 10 | 44344 | 0.217 | 0.326 | 0.444 | 0.693 | 1.966 | 56.1 | 40 | 0 |
| valkey-go/flush-100us | 100 | 107611 | 0.944 | 1.198 | 1.692 | 2.247 | 3.413 | 23.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15479 | 0.062 | 0.072 | 0.107 | 0.257 | 1.705 | 45.6 | 27 | 0 |
| valkey-go/1-conn | 10 | 80697 | 0.122 | 0.161 | 0.219 | 0.371 | 1.688 | 23.1 | 30 | 0 |
| valkey-go/1-conn | 100 | 432963 | 0.223 | 0.323 | 0.423 | 0.742 | 1.805 | 7.0 | 58 | 0 |
| valkey-go/1-conn | 1000 | 1047842 | 0.940 | 1.058 | 1.358 | 2.595 | 4.551 | 5.3 | 309 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 13003 | 0.072 | 0.087 | 0.146 | 0.342 | 1.482 | 53.9 | 27 | 0 |
| valkey-go/1-conn | 10 | 50288 | 0.197 | 0.261 | 0.424 | 0.684 | 1.607 | 40.3 | 31 | 0 |
| valkey-go/1-conn | 100 | 121451 | 0.794 | 0.942 | 1.356 | 2.133 | 3.607 | 18.6 | 58 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15238 | 0.062 | 0.074 | 0.110 | 0.268 | 1.565 | 46.9 | 33 | 0 |
| valkey-go/no-retry | 10 | 72876 | 0.133 | 0.187 | 0.254 | 0.369 | 1.524 | 43.7 | 40 | 0 |
| valkey-go/no-retry | 100 | 325132 | 0.305 | 0.420 | 0.540 | 0.895 | 1.971 | 15.9 | 65 | 0 |
| valkey-go/no-retry | 1000 | 904633 | 1.138 | 1.483 | 1.840 | 3.037 | 5.599 | 7.0 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12834 | 0.073 | 0.089 | 0.133 | 0.299 | 2.137 | 54.1 | 37 | 0 |
| valkey-go/no-retry | 10 | 47925 | 0.203 | 0.280 | 0.399 | 0.673 | 1.826 | 61.2 | 39 | 0 |
| valkey-go/no-retry | 100 | 105435 | 0.975 | 1.246 | 1.644 | 2.271 | 3.511 | 29.0 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12643 | 0.089 | 0.116 | 0.178 | 0.343 | 1.636 | 74.3 | 68 | 0 |
| valkey-go/docache | 10 | 143028 | 0.002 | 0.205 | 0.292 | 0.405 | 3.141 | 24.5 | 153 | 0 |
| valkey-go/docache | 100 | 484161 | 0.010 | 0.635 | 0.955 | 1.385 | 5.407 | 12.9 | 182 | 0 |
| valkey-go/docache | 1000 | 621989 | 0.017 | 5.135 | 6.279 | 9.119 | 18.799 | 11.9 | 416 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 11491 | 0.032 | 0.191 | 0.273 | 0.441 | 4.111 | 148.7 | 155 | 0 |
| valkey-go/docache | 10 | 332515 | 0.022 | 0.034 | 0.196 | 0.358 | 4.851 | 35.3 | 193 | 0 |
| valkey-go/docache | 100 | 566963 | 0.103 | 0.371 | 1.099 | 3.163 | 25.247 | 28.5 | 259 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 15204 | 0.063 | 0.075 | 0.112 | 0.261 | 1.604 | 46.2 | 37 | 0 |
| valkey-go/compat | 10 | 69758 | 0.138 | 0.197 | 0.269 | 0.421 | 1.793 | 44.3 | 42 | 0 |
| valkey-go/compat | 100 | 305185 | 0.325 | 0.453 | 0.584 | 0.983 | 2.737 | 16.3 | 66 | 0 |
| valkey-go/compat | 1000 | 881250 | 1.163 | 1.533 | 1.927 | 3.901 | 8.019 | 7.0 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12114 | 0.076 | 0.096 | 0.161 | 0.371 | 3.023 | 59.1 | 36 | 0 |
| valkey-go/compat | 10 | 45509 | 0.209 | 0.301 | 0.476 | 0.837 | 2.313 | 67.3 | 41 | 0 |
| valkey-go/compat | 100 | 102756 | 0.991 | 1.282 | 1.921 | 2.857 | 4.359 | 34.4 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15299 | 0.062 | 0.072 | 0.104 | 0.262 | 1.556 | 46.1 | 35 | 0 |
| valkey-go/hook | 10 | 70859 | 0.136 | 0.194 | 0.268 | 0.397 | 1.604 | 44.4 | 39 | 0 |
| valkey-go/hook | 100 | 329430 | 0.300 | 0.415 | 0.541 | 0.834 | 2.449 | 15.7 | 66 | 0 |
| valkey-go/hook | 1000 | 922990 | 1.114 | 1.458 | 1.815 | 2.809 | 5.747 | 6.8 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12830 | 0.073 | 0.089 | 0.137 | 0.300 | 1.575 | 53.3 | 37 | 0 |
| valkey-go/hook | 10 | 47553 | 0.205 | 0.284 | 0.401 | 0.657 | 2.739 | 61.3 | 42 | 0 |
| valkey-go/hook | 100 | 85609 | 1.100 | 1.785 | 2.611 | 3.399 | 5.591 | 32.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14675 | 0.064 | 0.078 | 0.113 | 0.270 | 1.655 | 50.2 | 37 | 0 |
| valkey-go/otel | 10 | 70757 | 0.136 | 0.193 | 0.269 | 0.484 | 1.351 | 47.9 | 40 | 0 |
| valkey-go/otel | 100 | 315104 | 0.310 | 0.437 | 0.631 | 1.376 | 3.071 | 18.3 | 66 | 0 |
| valkey-go/otel | 1000 | 874441 | 1.159 | 1.564 | 2.239 | 4.419 | 8.087 | 8.7 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12392 | 0.075 | 0.095 | 0.150 | 0.304 | 1.782 | 56.4 | 37 | 0 |
| valkey-go/otel | 10 | 47206 | 0.205 | 0.285 | 0.419 | 0.728 | 1.645 | 63.8 | 42 | 0 |
| valkey-go/otel | 100 | 103254 | 0.989 | 1.274 | 1.812 | 2.759 | 4.643 | 32.0 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 11398 | 0.082 | 0.105 | 0.158 | 0.332 | 1.740 | 94.1 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 66422 | 0.141 | 0.206 | 0.340 | 0.679 | 2.139 | 59.8 | 41 | 0 |
| valkey-go/otel-sdk | 100 | 296465 | 0.307 | 0.474 | 1.087 | 2.205 | 3.583 | 26.5 | 67 | 0 |
| valkey-go/otel-sdk | 1000 | 795828 | 1.150 | 1.677 | 4.607 | 7.671 | 20.271 | 17.0 | 324 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9572 | 0.100 | 0.121 | 0.190 | 0.416 | 1.702 | 108.2 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 45640 | 0.208 | 0.296 | 0.521 | 0.901 | 1.894 | 76.2 | 41 | 0 |
| valkey-go/otel-sdk | 100 | 102135 | 0.994 | 1.312 | 1.972 | 3.291 | 4.855 | 42.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11749 | 0.096 | 0.121 | 0.197 | 0.388 | 1.677 | 83.0 | 74 | 0 |
| valkey-go/aside | 10 | 136283 | 0.012 | 0.199 | 0.319 | 0.491 | 4.459 | 21.2 | 169 | 0 |
| valkey-go/aside | 100 | 194511 | 0.025 | 1.489 | 1.861 | 2.805 | 5.727 | 17.3 | 196 | 0 |
| valkey-go/aside | 1000 | 193976 | 0.026 | 14.615 | 15.743 | 25.215 | 28.479 | 17.3 | 440 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1348 | 0.732 | 1.027 | 1.372 | 2.423 | 2.643 | 739.3 | 82 | 0 |
| valkey-go/aside | 10 | 79120 | 0.093 | 0.247 | 0.568 | 1.137 | 6.683 | 62.4 | 211 | 0 |
| valkey-go/aside | 100 | 51423 | 0.453 | 7.543 | 13.455 | 18.927 | 33.663 | 99.0 | 230 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2322 | 0.402 | 0.507 | 0.885 | 1.872 | 2.589 | 523.4 | 28 | 0 |
| valkey-go/lock | 10 | 14221 | 0.674 | 0.853 | 1.446 | 2.077 | 2.803 | 189.2 | 33 | 0 |
| valkey-go/lock | 100 | 31536 | 3.093 | 3.561 | 4.947 | 6.043 | 7.807 | 100.6 | 62 | 2 |
| valkey-go/lock | 1000 | 30889 | 31.839 | 35.455 | 40.223 | 50.687 | 61.183 | 117.2 | 333 | 12 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 241 | 4.071 | 4.531 | 5.799 | 7.603 | 8.751 | 5172.5 | 27 | 0 |
| valkey-go/lock | 10 | 1421 | 6.927 | 7.855 | 8.951 | 10.671 | 11.935 | 1885.2 | 32 | 0 |
| valkey-go/lock | 100 | 3063 | 32.383 | 34.655 | 37.535 | 39.263 | 42.879 | 1000.2 | 63 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15029 | 0.063 | 0.075 | 0.112 | 0.270 | 1.529 | 46.4 | 37 | 0 |
| valkey-go | 10 | 70146 | 0.138 | 0.194 | 0.261 | 0.394 | 1.506 | 44.1 | 40 | 0 |
| valkey-go | 100 | 312700 | 0.317 | 0.436 | 0.560 | 0.853 | 2.157 | 15.9 | 67 | 0 |
| valkey-go | 1000 | 901323 | 1.142 | 1.483 | 1.901 | 2.839 | 5.959 | 6.8 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13322 | 0.072 | 0.081 | 0.126 | 0.359 | 1.592 | 52.3 | 35 | 0 |
| valkey-go | 10 | 46455 | 0.210 | 0.289 | 0.421 | 0.739 | 2.083 | 60.8 | 39 | 0 |
| valkey-go | 100 | 104731 | 0.985 | 1.243 | 1.636 | 2.307 | 3.415 | 27.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9721 | 0.094 | 0.126 | 0.194 | 0.445 | 1.513 | 105.1 | 30 | 0 |
| go-redis-autopipeline | 10 | 60474 | 0.154 | 0.204 | 0.365 | 0.778 | 1.606 | 21.1 | 32 | 0 |
| go-redis-autopipeline | 100 | 204923 | 0.476 | 0.580 | 0.803 | 1.561 | 2.115 | 8.1 | 60 | 0 |
| go-redis-autopipeline | 1000 | 472415 | 2.133 | 2.747 | 3.431 | 4.883 | 7.023 | 6.6 | 307 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10554 | 0.079 | 0.135 | 0.266 | 0.506 | 2.997 | 70.1 | 28 | 0 |
| go-redis-autopipeline | 10 | 37870 | 0.258 | 0.336 | 0.580 | 0.934 | 1.619 | 80.5 | 35 | 0 |
| go-redis-autopipeline | 100 | 39465 | 2.405 | 2.847 | 4.803 | 5.671 | 7.271 | 83.3 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7498 | 0.126 | 0.158 | 0.270 | 0.506 | 1.672 | 121.9 | 33 | 0 |
| glide | 10 | 34489 | 0.280 | 0.363 | 0.514 | 0.778 | 1.859 | 54.3 | 36 | 0 |
| glide | 100 | 35944 | 2.851 | 3.265 | 4.327 | 4.867 | 6.139 | 55.8 | 63 | 0 |
| glide | 1000 | 33212 | 29.951 | 32.687 | 38.847 | 43.199 | 56.575 | 59.4 | 229 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6696 | 0.136 | 0.179 | 0.349 | 0.691 | 1.711 | 122.8 | 33 | 0 |
| glide | 10 | 25972 | 0.364 | 0.537 | 0.930 | 1.515 | 2.325 | 91.4 | 35 | 0 |
| glide | 100 | 29311 | 3.415 | 4.363 | 5.399 | 6.347 | 8.287 | 84.5 | 64 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15528 | 0.062 | 0.071 | 0.104 | 0.294 | 1.610 | 46.0 | 28 | 0 |
| valkey-go/no-autopipe | 10 | 52831 | 0.187 | 0.241 | 0.329 | 0.544 | 1.843 | 50.5 | 51 | 0 |
| valkey-go/no-autopipe | 100 | 53468 | 1.804 | 1.965 | 3.597 | 3.963 | 6.051 | 55.2 | 174 | 0 |
| valkey-go/no-autopipe | 1000 | 47496 | 20.991 | 22.543 | 25.279 | 42.495 | 49.983 | 51.7 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12995 | 0.072 | 0.086 | 0.135 | 0.391 | 1.417 | 53.7 | 28 | 0 |
| valkey-go/no-autopipe | 10 | 38175 | 0.259 | 0.328 | 0.469 | 0.751 | 1.854 | 64.8 | 52 | 0 |
| valkey-go/no-autopipe | 100 | 39406 | 2.441 | 2.659 | 4.867 | 5.323 | 6.915 | 65.7 | 255 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12536 | 0.074 | 0.091 | 0.138 | 0.382 | 1.763 | 68.4 | 35 | 0 |
| valkey-go/always-pipe | 10 | 72234 | 0.134 | 0.189 | 0.252 | 0.471 | 1.550 | 44.1 | 38 | 0 |
| valkey-go/always-pipe | 100 | 322567 | 0.306 | 0.422 | 0.555 | 0.829 | 2.251 | 15.8 | 64 | 0 |
| valkey-go/always-pipe | 1000 | 910262 | 1.127 | 1.471 | 1.885 | 2.881 | 5.331 | 6.7 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9476 | 0.102 | 0.122 | 0.192 | 0.426 | 1.718 | 95.5 | 36 | 0 |
| valkey-go/always-pipe | 10 | 46992 | 0.207 | 0.287 | 0.414 | 0.654 | 1.530 | 60.4 | 39 | 0 |
| valkey-go/always-pipe | 100 | 105455 | 0.974 | 1.235 | 1.704 | 2.459 | 3.363 | 28.2 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15568 | 0.062 | 0.069 | 0.102 | 0.282 | 1.789 | 45.4 | 34 | 0 |
| valkey-go/flush-20us | 10 | 71919 | 0.134 | 0.189 | 0.254 | 0.464 | 1.657 | 45.3 | 42 | 0 |
| valkey-go/flush-20us | 100 | 344909 | 0.288 | 0.389 | 0.511 | 0.827 | 1.980 | 13.7 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 939067 | 1.086 | 1.433 | 1.845 | 3.003 | 6.087 | 6.3 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12811 | 0.073 | 0.088 | 0.132 | 0.382 | 1.439 | 53.2 | 37 | 0 |
| valkey-go/flush-20us | 10 | 47073 | 0.207 | 0.286 | 0.419 | 0.674 | 1.868 | 61.2 | 41 | 0 |
| valkey-go/flush-20us | 100 | 103577 | 0.980 | 1.274 | 1.871 | 2.453 | 3.739 | 26.8 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15791 | 0.061 | 0.070 | 0.100 | 0.274 | 1.551 | 45.0 | 37 | 0 |
| valkey-go/flush-100us | 10 | 61050 | 0.157 | 0.239 | 0.315 | 1.334 | 1.785 | 35.9 | 40 | 0 |
| valkey-go/flush-100us | 100 | 404608 | 0.238 | 0.339 | 0.481 | 0.714 | 1.735 | 8.9 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 978825 | 1.041 | 1.400 | 1.740 | 2.697 | 4.787 | 5.5 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12976 | 0.073 | 0.088 | 0.131 | 0.378 | 1.513 | 53.5 | 37 | 0 |
| valkey-go/flush-100us | 10 | 43734 | 0.220 | 0.328 | 0.458 | 0.710 | 1.885 | 54.8 | 39 | 0 |
| valkey-go/flush-100us | 100 | 107832 | 0.945 | 1.188 | 1.675 | 2.187 | 3.475 | 22.8 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15340 | 0.062 | 0.073 | 0.104 | 0.291 | 1.625 | 47.0 | 28 | 0 |
| valkey-go/1-conn | 10 | 83371 | 0.118 | 0.155 | 0.210 | 0.425 | 1.685 | 22.6 | 33 | 0 |
| valkey-go/1-conn | 100 | 403010 | 0.235 | 0.352 | 0.564 | 0.869 | 2.027 | 7.4 | 59 | 0 |
| valkey-go/1-conn | 1000 | 1037541 | 0.952 | 1.062 | 1.340 | 2.613 | 4.543 | 5.3 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12799 | 0.072 | 0.091 | 0.151 | 0.399 | 1.607 | 55.4 | 28 | 0 |
| valkey-go/1-conn | 10 | 49090 | 0.202 | 0.269 | 0.443 | 0.697 | 1.537 | 40.9 | 30 | 0 |
| valkey-go/1-conn | 100 | 119891 | 0.805 | 0.942 | 1.369 | 1.968 | 2.819 | 18.3 | 56 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15676 | 0.062 | 0.070 | 0.103 | 0.268 | 1.429 | 45.4 | 36 | 0 |
| valkey-go/no-retry | 10 | 72278 | 0.134 | 0.188 | 0.252 | 0.465 | 1.812 | 43.5 | 42 | 0 |
| valkey-go/no-retry | 100 | 324651 | 0.304 | 0.419 | 0.552 | 0.828 | 2.375 | 15.7 | 67 | 0 |
| valkey-go/no-retry | 1000 | 910154 | 1.129 | 1.467 | 1.837 | 2.921 | 6.563 | 6.7 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12911 | 0.073 | 0.087 | 0.135 | 0.388 | 1.635 | 53.1 | 36 | 0 |
| valkey-go/no-retry | 10 | 46920 | 0.208 | 0.286 | 0.421 | 0.712 | 1.843 | 60.2 | 41 | 0 |
| valkey-go/no-retry | 100 | 104006 | 0.989 | 1.259 | 1.677 | 2.587 | 3.537 | 28.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12285 | 0.091 | 0.118 | 0.194 | 0.435 | 1.547 | 75.9 | 66 | 0 |
| valkey-go/docache | 10 | 135943 | 0.002 | 0.216 | 0.312 | 0.496 | 3.175 | 24.9 | 151 | 0 |
| valkey-go/docache | 100 | 457299 | 0.010 | 0.681 | 0.971 | 1.425 | 4.057 | 13.0 | 180 | 0 |
| valkey-go/docache | 1000 | 534066 | 0.017 | 5.971 | 7.967 | 12.479 | 19.679 | 12.1 | 372 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 9864 | 0.058 | 0.202 | 0.296 | 0.532 | 1.622 | 168.9 | 140 | 0 |
| valkey-go/docache | 10 | 339463 | 0.021 | 0.033 | 0.195 | 0.367 | 4.883 | 34.9 | 195 | 0 |
| valkey-go/docache | 100 | 552845 | 0.105 | 0.388 | 1.154 | 3.197 | 25.055 | 29.0 | 256 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 15109 | 0.063 | 0.073 | 0.106 | 0.317 | 1.624 | 46.6 | 35 | 0 |
| valkey-go/compat | 10 | 69504 | 0.139 | 0.197 | 0.269 | 0.485 | 1.579 | 43.9 | 41 | 0 |
| valkey-go/compat | 100 | 307625 | 0.322 | 0.444 | 0.585 | 0.941 | 2.569 | 16.4 | 67 | 0 |
| valkey-go/compat | 1000 | 890080 | 1.153 | 1.508 | 1.895 | 3.723 | 7.939 | 7.1 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12667 | 0.074 | 0.090 | 0.139 | 0.388 | 1.523 | 57.3 | 37 | 0 |
| valkey-go/compat | 10 | 45055 | 0.213 | 0.299 | 0.499 | 0.817 | 1.770 | 68.3 | 42 | 0 |
| valkey-go/compat | 100 | 102294 | 0.992 | 1.289 | 1.953 | 2.909 | 4.235 | 35.1 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15452 | 0.062 | 0.073 | 0.103 | 0.293 | 2.715 | 46.5 | 34 | 0 |
| valkey-go/hook | 10 | 72777 | 0.133 | 0.187 | 0.254 | 0.486 | 1.494 | 43.6 | 40 | 0 |
| valkey-go/hook | 100 | 315615 | 0.313 | 0.434 | 0.580 | 0.915 | 2.313 | 15.9 | 67 | 0 |
| valkey-go/hook | 1000 | 893854 | 1.146 | 1.506 | 1.927 | 2.889 | 6.591 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12969 | 0.073 | 0.085 | 0.130 | 0.394 | 1.784 | 51.9 | 37 | 0 |
| valkey-go/hook | 10 | 46933 | 0.207 | 0.287 | 0.413 | 0.693 | 1.712 | 61.2 | 39 | 0 |
| valkey-go/hook | 100 | 105717 | 0.975 | 1.235 | 1.650 | 2.293 | 3.425 | 28.6 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 15033 | 0.064 | 0.074 | 0.105 | 0.326 | 1.506 | 48.3 | 37 | 0 |
| valkey-go/otel | 10 | 70829 | 0.135 | 0.192 | 0.266 | 0.519 | 1.846 | 47.9 | 39 | 0 |
| valkey-go/otel | 100 | 309662 | 0.315 | 0.445 | 0.665 | 1.468 | 2.783 | 18.5 | 65 | 0 |
| valkey-go/otel | 1000 | 877781 | 1.157 | 1.546 | 2.129 | 4.815 | 8.423 | 8.8 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12529 | 0.075 | 0.091 | 0.136 | 0.403 | 1.752 | 56.1 | 37 | 0 |
| valkey-go/otel | 10 | 45951 | 0.210 | 0.294 | 0.457 | 0.748 | 2.115 | 65.7 | 40 | 0 |
| valkey-go/otel | 100 | 102320 | 0.998 | 1.285 | 1.808 | 2.723 | 3.927 | 31.9 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9138 | 0.094 | 0.157 | 0.250 | 0.492 | 1.805 | 117.8 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 66486 | 0.141 | 0.204 | 0.341 | 0.697 | 2.022 | 61.4 | 44 | 0 |
| valkey-go/otel-sdk | 100 | 292675 | 0.310 | 0.480 | 1.118 | 2.311 | 4.089 | 27.0 | 68 | 0 |
| valkey-go/otel-sdk | 1000 | 781927 | 1.163 | 1.715 | 4.775 | 7.859 | 16.607 | 17.3 | 326 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9598 | 0.098 | 0.120 | 0.194 | 0.442 | 1.476 | 108.8 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 45221 | 0.210 | 0.300 | 0.536 | 0.869 | 1.913 | 75.7 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 101846 | 0.998 | 1.312 | 2.009 | 3.349 | 5.091 | 42.3 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11585 | 0.098 | 0.121 | 0.195 | 0.429 | 1.549 | 83.7 | 76 | 0 |
| valkey-go/aside | 10 | 136339 | 0.012 | 0.199 | 0.318 | 0.511 | 4.131 | 21.3 | 169 | 0 |
| valkey-go/aside | 100 | 193781 | 0.025 | 1.490 | 1.846 | 3.547 | 13.199 | 17.2 | 196 | 0 |
| valkey-go/aside | 1000 | 192554 | 0.026 | 14.711 | 15.959 | 19.919 | 26.127 | 17.4 | 440 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1321 | 0.750 | 1.024 | 1.411 | 2.553 | 4.923 | 740.4 | 80 | 0 |
| valkey-go/aside | 10 | 84446 | 0.085 | 0.235 | 0.547 | 1.092 | 8.059 | 60.3 | 209 | 0 |
| valkey-go/aside | 100 | 53152 | 0.461 | 7.379 | 13.039 | 19.375 | 35.967 | 96.8 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2472 | 0.390 | 0.461 | 0.714 | 1.245 | 2.031 | 507.8 | 27 | 0 |
| valkey-go/lock | 10 | 14578 | 0.660 | 0.833 | 1.341 | 1.947 | 2.463 | 181.3 | 33 | 0 |
| valkey-go/lock | 100 | 30704 | 3.177 | 3.653 | 4.931 | 5.855 | 7.043 | 98.7 | 61 | 2 |
| valkey-go/lock | 1000 | 30227 | 32.703 | 35.295 | 40.735 | 52.383 | 58.719 | 116.7 | 333 | 12 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 244 | 4.023 | 4.503 | 5.967 | 7.287 | 7.391 | 5088.9 | 26 | 0 |
| valkey-go/lock | 10 | 1384 | 7.095 | 8.083 | 9.495 | 11.231 | 11.815 | 1875.7 | 34 | 0 |
| valkey-go/lock | 100 | 1809 | 57.023 | 59.999 | 63.775 | 64.735 | 65.855 | 1251.5 | 62 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15884 | 0.061 | 0.069 | 0.099 | 0.275 | 1.597 | 45.6 | 33 | 0 |
| valkey-go | 10 | 72362 | 0.134 | 0.189 | 0.253 | 0.467 | 1.758 | 43.4 | 40 | 0 |
| valkey-go | 100 | 317209 | 0.312 | 0.433 | 0.578 | 0.930 | 2.053 | 16.0 | 67 | 0 |
| valkey-go | 1000 | 797671 | 1.288 | 1.690 | 2.235 | 3.919 | 8.527 | 7.0 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12043 | 0.075 | 0.101 | 0.175 | 0.443 | 3.143 | 57.0 | 37 | 0 |
| valkey-go | 10 | 46963 | 0.207 | 0.286 | 0.423 | 0.692 | 1.629 | 61.7 | 41 | 0 |
| valkey-go | 100 | 95787 | 1.081 | 1.363 | 1.847 | 2.475 | 3.579 | 28.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13794 | 0.069 | 0.081 | 0.119 | 0.364 | 1.689 | 55.5 | 28 | 0 |
| go-redis | 10 | 52072 | 0.191 | 0.248 | 0.343 | 0.612 | 1.803 | 61.6 | 32 | 0 |
| go-redis | 100 | 52756 | 1.814 | 2.053 | 3.615 | 4.031 | 5.971 | 66.2 | 70 | 0 |
| go-redis | 1000 | 52999 | 18.767 | 20.159 | 21.903 | 23.407 | 26.335 | 78.1 | 254 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11800 | 0.078 | 0.099 | 0.166 | 0.414 | 1.914 | 64.9 | 29 | 0 |
| go-redis | 10 | 38380 | 0.255 | 0.329 | 0.576 | 0.979 | 1.745 | 80.0 | 37 | 0 |
| go-redis | 100 | 39316 | 2.407 | 2.917 | 4.799 | 5.791 | 7.347 | 83.2 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7548 | 0.124 | 0.157 | 0.269 | 0.499 | 1.775 | 122.9 | 33 | 0 |
| glide | 10 | 34037 | 0.279 | 0.382 | 0.540 | 0.778 | 2.515 | 55.4 | 35 | 0 |
| glide | 100 | 35060 | 2.907 | 3.377 | 4.451 | 5.043 | 6.115 | 56.5 | 64 | 0 |
| glide | 1000 | 34341 | 28.927 | 32.271 | 36.223 | 41.119 | 42.591 | 58.6 | 221 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6634 | 0.137 | 0.182 | 0.344 | 0.625 | 1.592 | 123.5 | 31 | 0 |
| glide | 10 | 25196 | 0.374 | 0.557 | 0.983 | 1.500 | 2.185 | 93.7 | 34 | 0 |
| glide | 100 | 29327 | 3.399 | 4.319 | 5.319 | 6.139 | 9.863 | 84.4 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15374 | 0.062 | 0.073 | 0.103 | 0.340 | 1.784 | 46.8 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 52937 | 0.187 | 0.241 | 0.325 | 0.543 | 1.412 | 50.7 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 52996 | 1.820 | 1.977 | 3.627 | 3.977 | 5.279 | 55.7 | 178 | 0 |
| valkey-go/no-autopipe | 1000 | 48097 | 20.527 | 22.351 | 24.687 | 41.983 | 50.015 | 52.1 | 1209 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12843 | 0.073 | 0.088 | 0.133 | 0.389 | 1.630 | 54.5 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 39280 | 0.254 | 0.316 | 0.427 | 0.668 | 1.676 | 63.7 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 39343 | 2.435 | 2.693 | 4.851 | 5.463 | 8.119 | 67.3 | 261 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12402 | 0.075 | 0.093 | 0.142 | 0.375 | 1.540 | 70.1 | 35 | 0 |
| valkey-go/always-pipe | 10 | 72741 | 0.133 | 0.188 | 0.255 | 0.473 | 1.017 | 43.5 | 40 | 0 |
| valkey-go/always-pipe | 100 | 321402 | 0.307 | 0.425 | 0.562 | 0.870 | 2.151 | 15.9 | 68 | 0 |
| valkey-go/always-pipe | 1000 | 878301 | 1.175 | 1.522 | 1.894 | 3.085 | 5.331 | 6.8 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9357 | 0.102 | 0.125 | 0.197 | 0.448 | 1.960 | 96.6 | 36 | 0 |
| valkey-go/always-pipe | 10 | 47006 | 0.207 | 0.287 | 0.423 | 0.700 | 1.825 | 61.9 | 41 | 0 |
| valkey-go/always-pipe | 100 | 101632 | 1.015 | 1.280 | 1.661 | 2.505 | 3.401 | 28.3 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15825 | 0.061 | 0.069 | 0.098 | 0.249 | 1.618 | 45.8 | 34 | 0 |
| valkey-go/flush-20us | 10 | 73129 | 0.132 | 0.186 | 0.253 | 0.472 | 1.715 | 44.6 | 40 | 0 |
| valkey-go/flush-20us | 100 | 338223 | 0.293 | 0.399 | 0.537 | 0.846 | 1.902 | 13.9 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 909157 | 1.128 | 1.483 | 1.867 | 3.095 | 6.279 | 6.5 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12936 | 0.073 | 0.087 | 0.130 | 0.374 | 1.831 | 52.8 | 37 | 0 |
| valkey-go/flush-20us | 10 | 46972 | 0.207 | 0.287 | 0.423 | 0.683 | 1.541 | 61.5 | 41 | 0 |
| valkey-go/flush-20us | 100 | 102014 | 1.010 | 1.278 | 1.748 | 2.527 | 3.955 | 27.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15572 | 0.062 | 0.069 | 0.099 | 0.292 | 1.496 | 45.7 | 37 | 0 |
| valkey-go/flush-100us | 10 | 60348 | 0.159 | 0.243 | 0.320 | 1.301 | 2.501 | 36.2 | 41 | 0 |
| valkey-go/flush-100us | 100 | 406032 | 0.237 | 0.338 | 0.479 | 0.758 | 1.969 | 8.9 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 953743 | 1.059 | 1.442 | 1.967 | 2.863 | 6.643 | 5.4 | 312 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12839 | 0.073 | 0.089 | 0.132 | 0.393 | 1.514 | 54.6 | 35 | 0 |
| valkey-go/flush-100us | 10 | 43426 | 0.221 | 0.332 | 0.471 | 0.830 | 3.607 | 56.0 | 39 | 0 |
| valkey-go/flush-100us | 100 | 101963 | 0.993 | 1.279 | 1.799 | 2.387 | 3.363 | 23.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15094 | 0.063 | 0.074 | 0.108 | 0.350 | 1.474 | 46.6 | 27 | 0 |
| valkey-go/1-conn | 10 | 81322 | 0.121 | 0.159 | 0.217 | 0.465 | 1.799 | 23.1 | 31 | 0 |
| valkey-go/1-conn | 100 | 417527 | 0.232 | 0.335 | 0.433 | 0.815 | 2.217 | 7.2 | 57 | 0 |
| valkey-go/1-conn | 1000 | 1016917 | 0.967 | 1.079 | 1.463 | 2.525 | 4.763 | 5.3 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12882 | 0.073 | 0.088 | 0.146 | 0.408 | 1.576 | 55.6 | 27 | 0 |
| valkey-go/1-conn | 10 | 45414 | 0.211 | 0.298 | 0.534 | 1.254 | 6.711 | 42.9 | 31 | 0 |
| valkey-go/1-conn | 100 | 113253 | 0.849 | 1.029 | 1.503 | 2.249 | 4.323 | 19.0 | 57 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15715 | 0.061 | 0.072 | 0.113 | 0.339 | 1.568 | 46.5 | 35 | 0 |
| valkey-go/no-retry | 10 | 72211 | 0.134 | 0.189 | 0.256 | 0.487 | 1.797 | 43.8 | 40 | 0 |
| valkey-go/no-retry | 100 | 313908 | 0.314 | 0.436 | 0.581 | 0.972 | 2.125 | 16.1 | 67 | 0 |
| valkey-go/no-retry | 1000 | 801784 | 1.227 | 1.774 | 2.643 | 3.955 | 7.495 | 7.4 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 13029 | 0.072 | 0.086 | 0.133 | 0.398 | 1.484 | 55.3 | 37 | 0 |
| valkey-go/no-retry | 10 | 46223 | 0.210 | 0.291 | 0.443 | 0.747 | 1.613 | 61.0 | 41 | 0 |
| valkey-go/no-retry | 100 | 100121 | 1.024 | 1.312 | 1.798 | 2.701 | 3.811 | 29.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12064 | 0.092 | 0.122 | 0.198 | 0.435 | 1.583 | 76.8 | 68 | 0 |
| valkey-go/docache | 10 | 140034 | 0.002 | 0.210 | 0.296 | 0.468 | 3.525 | 24.6 | 147 | 0 |
| valkey-go/docache | 100 | 460451 | 0.010 | 0.672 | 0.974 | 1.434 | 4.927 | 13.1 | 180 | 0 |
| valkey-go/docache | 1000 | 513310 | 0.017 | 6.255 | 8.027 | 11.543 | 18.095 | 12.1 | 416 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 10204 | 0.042 | 0.202 | 0.293 | 0.525 | 1.672 | 164.3 | 138 | 0 |
| valkey-go/docache | 10 | 337908 | 0.021 | 0.033 | 0.195 | 0.364 | 6.799 | 34.9 | 194 | 0 |
| valkey-go/docache | 100 | 543488 | 0.111 | 0.394 | 1.179 | 3.647 | 27.919 | 29.7 | 254 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14897 | 0.064 | 0.076 | 0.112 | 0.325 | 1.724 | 46.8 | 37 | 0 |
| valkey-go/compat | 10 | 68598 | 0.141 | 0.199 | 0.269 | 0.489 | 1.745 | 44.9 | 42 | 0 |
| valkey-go/compat | 100 | 309237 | 0.321 | 0.440 | 0.572 | 1.017 | 2.317 | 16.5 | 66 | 0 |
| valkey-go/compat | 1000 | 877134 | 1.171 | 1.531 | 1.920 | 3.531 | 7.155 | 7.2 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12509 | 0.075 | 0.092 | 0.144 | 0.407 | 1.547 | 57.5 | 38 | 0 |
| valkey-go/compat | 10 | 45517 | 0.211 | 0.295 | 0.498 | 0.839 | 1.845 | 68.3 | 42 | 0 |
| valkey-go/compat | 100 | 100721 | 1.012 | 1.310 | 1.950 | 2.991 | 4.275 | 35.0 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15307 | 0.062 | 0.074 | 0.105 | 0.273 | 1.740 | 46.6 | 35 | 0 |
| valkey-go/hook | 10 | 72363 | 0.133 | 0.189 | 0.256 | 0.475 | 1.057 | 43.7 | 39 | 0 |
| valkey-go/hook | 100 | 322042 | 0.307 | 0.424 | 0.562 | 0.934 | 2.017 | 16.0 | 67 | 0 |
| valkey-go/hook | 1000 | 893945 | 1.150 | 1.503 | 1.866 | 2.763 | 6.983 | 7.0 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12886 | 0.073 | 0.088 | 0.134 | 0.392 | 1.552 | 52.8 | 37 | 0 |
| valkey-go/hook | 10 | 46846 | 0.209 | 0.286 | 0.408 | 0.684 | 1.806 | 60.8 | 41 | 0 |
| valkey-go/hook | 100 | 103874 | 0.991 | 1.255 | 1.667 | 2.345 | 3.379 | 28.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14995 | 0.064 | 0.073 | 0.108 | 0.358 | 1.675 | 48.5 | 37 | 0 |
| valkey-go/otel | 10 | 70293 | 0.137 | 0.194 | 0.266 | 0.531 | 1.656 | 48.0 | 40 | 0 |
| valkey-go/otel | 100 | 313125 | 0.313 | 0.438 | 0.655 | 1.368 | 3.127 | 18.2 | 66 | 0 |
| valkey-go/otel | 1000 | 867521 | 1.171 | 1.566 | 2.273 | 4.575 | 8.115 | 8.7 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12689 | 0.074 | 0.089 | 0.134 | 0.392 | 1.561 | 55.6 | 36 | 0 |
| valkey-go/otel | 10 | 45642 | 0.212 | 0.296 | 0.455 | 0.760 | 2.427 | 65.8 | 40 | 0 |
| valkey-go/otel | 100 | 95709 | 1.035 | 1.505 | 2.012 | 2.775 | 4.483 | 31.8 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10724 | 0.089 | 0.109 | 0.163 | 0.409 | 1.661 | 102.4 | 38 | 0 |
| valkey-go/otel-sdk | 10 | 66823 | 0.141 | 0.204 | 0.340 | 0.663 | 1.621 | 60.0 | 41 | 0 |
| valkey-go/otel-sdk | 100 | 297884 | 0.305 | 0.470 | 1.095 | 2.325 | 4.379 | 26.9 | 69 | 0 |
| valkey-go/otel-sdk | 1000 | 798288 | 1.144 | 1.658 | 4.767 | 7.739 | 15.519 | 17.0 | 326 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9731 | 0.097 | 0.119 | 0.187 | 0.429 | 1.461 | 105.9 | 39 | 0 |
| valkey-go/otel-sdk | 10 | 45656 | 0.208 | 0.295 | 0.535 | 0.867 | 1.824 | 75.4 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 99529 | 1.017 | 1.348 | 2.023 | 3.219 | 4.771 | 41.5 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11680 | 0.097 | 0.121 | 0.193 | 0.431 | 1.407 | 83.6 | 74 | 0 |
| valkey-go/aside | 10 | 136505 | 0.012 | 0.198 | 0.316 | 0.506 | 3.695 | 21.2 | 170 | 0 |
| valkey-go/aside | 100 | 185075 | 0.026 | 1.556 | 2.201 | 3.183 | 11.887 | 19.2 | 196 | 0 |
| valkey-go/aside | 1000 | 188492 | 0.027 | 14.943 | 16.159 | 23.295 | 26.159 | 17.8 | 440 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1350 | 0.725 | 1.038 | 1.372 | 2.385 | 3.181 | 740.1 | 82 | 0 |
| valkey-go/aside | 10 | 85056 | 0.084 | 0.233 | 0.547 | 1.113 | 6.631 | 60.3 | 209 | 0 |
| valkey-go/aside | 100 | 52033 | 0.394 | 7.707 | 13.775 | 19.487 | 38.623 | 94.6 | 236 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2390 | 0.403 | 0.480 | 0.722 | 1.204 | 2.008 | 519.3 | 27 | 0 |
| valkey-go/lock | 10 | 14335 | 0.670 | 0.844 | 1.392 | 1.932 | 2.553 | 186.5 | 33 | 0 |
| valkey-go/lock | 100 | 27596 | 3.597 | 4.151 | 5.695 | 6.839 | 7.867 | 97.1 | 63 | 2 |
| valkey-go/lock | 1000 | 15530 | 53.535 | 74.175 | 82.431 | 87.615 | 102.591 | 143.3 | 336 | 7819 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 238 | 4.127 | 4.583 | 5.787 | 7.127 | 8.287 | 5067.7 | 29 | 0 |
| valkey-go/lock | 10 | 1216 | 8.099 | 9.167 | 10.559 | 12.871 | 13.567 | 2054.8 | 35 | 0 |
| valkey-go/lock | 100 | 1717 | 57.535 | 61.407 | 67.135 | 69.631 | 71.423 | 1314.3 | 64 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15328 | 0.062 | 0.071 | 0.101 | 0.295 | 1.596 | 46.3 | 36 | 0 |
| valkey-go | 10 | 72022 | 0.134 | 0.190 | 0.261 | 0.483 | 1.689 | 43.5 | 41 | 0 |
| valkey-go | 100 | 324006 | 0.304 | 0.421 | 0.572 | 0.929 | 2.439 | 15.6 | 68 | 0 |
| valkey-go | 1000 | 811707 | 1.273 | 1.651 | 2.101 | 3.151 | 6.327 | 6.8 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12744 | 0.074 | 0.089 | 0.127 | 0.389 | 1.542 | 53.9 | 35 | 0 |
| valkey-go | 10 | 47502 | 0.205 | 0.283 | 0.410 | 0.690 | 1.662 | 60.8 | 38 | 0 |
| valkey-go | 100 | 95606 | 1.082 | 1.368 | 1.790 | 2.575 | 3.591 | 28.0 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14113 | 0.067 | 0.081 | 0.113 | 0.348 | 2.183 | 54.9 | 28 | 0 |
| go-redis | 10 | 53074 | 0.187 | 0.243 | 0.332 | 0.610 | 1.553 | 60.0 | 33 | 0 |
| go-redis | 100 | 52999 | 1.816 | 2.002 | 3.625 | 4.019 | 7.159 | 65.3 | 70 | 0 |
| go-redis | 1000 | 53251 | 18.591 | 19.871 | 22.655 | 26.719 | 30.255 | 75.7 | 252 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11969 | 0.077 | 0.097 | 0.164 | 0.431 | 1.523 | 62.6 | 29 | 0 |
| go-redis | 10 | 38555 | 0.254 | 0.327 | 0.552 | 0.908 | 1.792 | 77.9 | 36 | 0 |
| go-redis | 100 | 40136 | 2.381 | 2.667 | 4.747 | 5.371 | 7.507 | 80.9 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10157 | 0.091 | 0.120 | 0.183 | 0.443 | 2.807 | 100.1 | 29 | 0 |
| go-redis-autopipeline | 10 | 61385 | 0.153 | 0.201 | 0.351 | 0.702 | 1.662 | 20.5 | 32 | 0 |
| go-redis-autopipeline | 100 | 206952 | 0.472 | 0.577 | 0.792 | 1.386 | 2.171 | 7.8 | 57 | 0 |
| go-redis-autopipeline | 1000 | 477441 | 2.049 | 2.733 | 3.433 | 5.043 | 6.487 | 6.2 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12017 | 0.077 | 0.095 | 0.160 | 0.428 | 1.594 | 62.9 | 27 | 0 |
| go-redis-autopipeline | 10 | 38597 | 0.254 | 0.326 | 0.560 | 0.950 | 2.909 | 78.8 | 34 | 0 |
| go-redis-autopipeline | 100 | 40044 | 2.387 | 2.687 | 4.759 | 5.295 | 8.083 | 80.8 | 72 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15398 | 0.062 | 0.073 | 0.104 | 0.297 | 1.629 | 45.9 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 52924 | 0.187 | 0.240 | 0.316 | 0.548 | 1.820 | 49.6 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 54030 | 1.787 | 1.932 | 3.563 | 3.901 | 4.831 | 54.3 | 177 | 0 |
| valkey-go/no-autopipe | 1000 | 48172 | 20.575 | 22.303 | 25.151 | 41.631 | 50.271 | 50.3 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12921 | 0.073 | 0.086 | 0.134 | 0.374 | 1.655 | 53.8 | 28 | 0 |
| valkey-go/no-autopipe | 10 | 39066 | 0.255 | 0.317 | 0.431 | 0.668 | 1.813 | 63.6 | 53 | 0 |
| valkey-go/no-autopipe | 100 | 39898 | 2.403 | 2.645 | 4.791 | 5.327 | 8.503 | 65.6 | 259 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9709 | 0.081 | 0.167 | 0.248 | 0.455 | 2.789 | 88.3 | 35 | 0 |
| valkey-go/always-pipe | 10 | 72811 | 0.133 | 0.187 | 0.254 | 0.471 | 1.722 | 43.6 | 41 | 0 |
| valkey-go/always-pipe | 100 | 325920 | 0.303 | 0.418 | 0.558 | 0.886 | 1.966 | 15.6 | 67 | 0 |
| valkey-go/always-pipe | 1000 | 865543 | 1.186 | 1.548 | 2.107 | 3.301 | 5.675 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9694 | 0.099 | 0.120 | 0.184 | 0.431 | 1.594 | 91.8 | 37 | 0 |
| valkey-go/always-pipe | 10 | 46483 | 0.209 | 0.290 | 0.422 | 0.695 | 3.823 | 61.4 | 39 | 0 |
| valkey-go/always-pipe | 100 | 100611 | 1.023 | 1.295 | 1.785 | 2.503 | 3.237 | 28.1 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15257 | 0.062 | 0.074 | 0.103 | 0.284 | 1.665 | 46.5 | 35 | 0 |
| valkey-go/flush-20us | 10 | 72110 | 0.134 | 0.188 | 0.253 | 0.487 | 1.822 | 44.6 | 41 | 0 |
| valkey-go/flush-20us | 100 | 347049 | 0.287 | 0.387 | 0.504 | 0.807 | 2.019 | 13.5 | 68 | 0 |
| valkey-go/flush-20us | 1000 | 914726 | 1.121 | 1.474 | 1.813 | 2.993 | 5.235 | 6.3 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12997 | 0.073 | 0.088 | 0.130 | 0.391 | 1.618 | 52.8 | 34 | 0 |
| valkey-go/flush-20us | 10 | 47113 | 0.207 | 0.286 | 0.410 | 0.672 | 1.916 | 62.4 | 38 | 0 |
| valkey-go/flush-20us | 100 | 101474 | 1.012 | 1.284 | 1.758 | 2.427 | 4.463 | 26.9 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15559 | 0.062 | 0.071 | 0.100 | 0.304 | 1.984 | 45.4 | 35 | 0 |
| valkey-go/flush-100us | 10 | 61136 | 0.157 | 0.239 | 0.312 | 1.307 | 1.987 | 35.2 | 39 | 0 |
| valkey-go/flush-100us | 100 | 407546 | 0.237 | 0.336 | 0.473 | 0.707 | 1.677 | 8.8 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 936365 | 1.092 | 1.469 | 1.894 | 2.823 | 5.151 | 5.3 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13118 | 0.072 | 0.087 | 0.136 | 0.370 | 1.282 | 52.4 | 35 | 0 |
| valkey-go/flush-100us | 10 | 43973 | 0.219 | 0.328 | 0.456 | 0.721 | 4.727 | 54.0 | 39 | 0 |
| valkey-go/flush-100us | 100 | 104440 | 0.974 | 1.224 | 1.733 | 2.275 | 3.631 | 22.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15562 | 0.062 | 0.071 | 0.099 | 0.270 | 1.645 | 45.7 | 26 | 0 |
| valkey-go/1-conn | 10 | 81311 | 0.122 | 0.159 | 0.213 | 0.438 | 1.679 | 22.7 | 29 | 0 |
| valkey-go/1-conn | 100 | 419820 | 0.230 | 0.334 | 0.438 | 0.802 | 3.305 | 7.0 | 55 | 0 |
| valkey-go/1-conn | 1000 | 1040186 | 0.948 | 1.057 | 1.353 | 2.569 | 4.979 | 5.1 | 305 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12946 | 0.072 | 0.087 | 0.134 | 0.404 | 1.570 | 53.9 | 26 | 0 |
| valkey-go/1-conn | 10 | 49825 | 0.199 | 0.263 | 0.428 | 0.652 | 1.642 | 39.5 | 30 | 0 |
| valkey-go/1-conn | 100 | 117712 | 0.820 | 0.950 | 1.388 | 2.016 | 3.039 | 18.0 | 56 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15722 | 0.061 | 0.068 | 0.099 | 0.255 | 1.555 | 45.2 | 35 | 0 |
| valkey-go/no-retry | 10 | 72816 | 0.133 | 0.187 | 0.249 | 0.479 | 1.678 | 42.8 | 43 | 0 |
| valkey-go/no-retry | 100 | 326901 | 0.302 | 0.416 | 0.559 | 0.884 | 2.125 | 15.5 | 67 | 0 |
| valkey-go/no-retry | 1000 | 904251 | 1.134 | 1.483 | 1.913 | 3.023 | 6.131 | 6.6 | 312 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 13308 | 0.071 | 0.084 | 0.123 | 0.383 | 1.606 | 51.5 | 37 | 0 |
| valkey-go/no-retry | 10 | 47232 | 0.206 | 0.285 | 0.412 | 0.667 | 1.805 | 59.8 | 41 | 0 |
| valkey-go/no-retry | 100 | 103840 | 0.995 | 1.251 | 1.655 | 2.389 | 3.295 | 26.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 13012 | 0.087 | 0.112 | 0.167 | 0.412 | 1.545 | 71.5 | 72 | 0 |
| valkey-go/docache | 10 | 142448 | 0.002 | 0.207 | 0.293 | 0.447 | 3.195 | 24.0 | 151 | 0 |
| valkey-go/docache | 100 | 472712 | 0.009 | 0.658 | 0.933 | 1.337 | 3.931 | 12.4 | 182 | 0 |
| valkey-go/docache | 1000 | 531162 | 0.016 | 6.115 | 7.919 | 10.903 | 17.135 | 10.9 | 397 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 13137 | 0.027 | 0.182 | 0.256 | 0.464 | 1.581 | 128.8 | 159 | 0 |
| valkey-go/docache | 10 | 302559 | 0.023 | 0.041 | 0.207 | 0.442 | 3.279 | 37.7 | 193 | 0 |
| valkey-go/docache | 100 | 548381 | 0.108 | 0.381 | 1.169 | 3.581 | 28.687 | 29.4 | 244 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 15079 | 0.063 | 0.074 | 0.109 | 0.340 | 1.531 | 45.8 | 35 | 0 |
| valkey-go/compat | 10 | 69404 | 0.140 | 0.196 | 0.264 | 0.477 | 1.479 | 43.8 | 40 | 0 |
| valkey-go/compat | 100 | 314208 | 0.314 | 0.433 | 0.575 | 0.977 | 2.161 | 15.8 | 67 | 0 |
| valkey-go/compat | 1000 | 894478 | 1.149 | 1.484 | 1.834 | 3.705 | 7.031 | 6.4 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12666 | 0.074 | 0.089 | 0.138 | 0.401 | 1.582 | 55.1 | 37 | 0 |
| valkey-go/compat | 10 | 45744 | 0.210 | 0.295 | 0.481 | 0.806 | 1.631 | 65.9 | 43 | 0 |
| valkey-go/compat | 100 | 102440 | 0.998 | 1.282 | 1.870 | 2.775 | 3.887 | 33.0 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15287 | 0.062 | 0.073 | 0.103 | 0.290 | 1.666 | 46.3 | 35 | 0 |
| valkey-go/hook | 10 | 71155 | 0.135 | 0.194 | 0.268 | 0.472 | 0.948 | 43.9 | 41 | 0 |
| valkey-go/hook | 100 | 324806 | 0.304 | 0.418 | 0.563 | 0.870 | 2.117 | 15.4 | 66 | 0 |
| valkey-go/hook | 1000 | 902432 | 1.139 | 1.469 | 1.827 | 3.083 | 6.895 | 6.0 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12947 | 0.073 | 0.087 | 0.131 | 0.392 | 2.299 | 52.6 | 37 | 0 |
| valkey-go/hook | 10 | 46959 | 0.207 | 0.286 | 0.413 | 0.697 | 1.783 | 60.5 | 41 | 0 |
| valkey-go/hook | 100 | 102982 | 1.002 | 1.261 | 1.690 | 2.297 | 3.199 | 27.3 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14608 | 0.064 | 0.078 | 0.109 | 0.333 | 1.454 | 50.6 | 37 | 0 |
| valkey-go/otel | 10 | 70518 | 0.136 | 0.194 | 0.268 | 0.513 | 1.897 | 47.6 | 42 | 0 |
| valkey-go/otel | 100 | 311184 | 0.314 | 0.442 | 0.649 | 1.360 | 2.485 | 17.8 | 66 | 0 |
| valkey-go/otel | 1000 | 872504 | 1.164 | 1.553 | 2.205 | 4.175 | 8.447 | 8.0 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12766 | 0.074 | 0.091 | 0.139 | 0.397 | 1.735 | 55.2 | 35 | 0 |
| valkey-go/otel | 10 | 45974 | 0.210 | 0.293 | 0.460 | 0.780 | 1.957 | 64.3 | 39 | 0 |
| valkey-go/otel | 100 | 103245 | 0.996 | 1.268 | 1.718 | 2.459 | 4.099 | 30.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10848 | 0.087 | 0.110 | 0.172 | 0.419 | 2.307 | 99.9 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 66510 | 0.142 | 0.205 | 0.334 | 0.647 | 2.593 | 58.9 | 40 | 0 |
| valkey-go/otel-sdk | 100 | 296137 | 0.309 | 0.473 | 1.067 | 2.125 | 3.875 | 25.9 | 69 | 0 |
| valkey-go/otel-sdk | 1000 | 779466 | 1.171 | 1.751 | 4.547 | 8.007 | 17.519 | 16.5 | 325 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9463 | 0.100 | 0.124 | 0.197 | 0.457 | 1.650 | 107.1 | 38 | 0 |
| valkey-go/otel-sdk | 10 | 45205 | 0.211 | 0.299 | 0.522 | 0.842 | 2.135 | 74.5 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 100398 | 1.011 | 1.331 | 1.984 | 3.261 | 4.891 | 39.7 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11784 | 0.096 | 0.120 | 0.189 | 0.429 | 1.395 | 81.9 | 76 | 0 |
| valkey-go/aside | 10 | 138970 | 0.011 | 0.196 | 0.310 | 0.507 | 3.741 | 20.5 | 169 | 0 |
| valkey-go/aside | 100 | 201548 | 0.024 | 1.445 | 1.808 | 3.497 | 10.815 | 16.4 | 198 | 0 |
| valkey-go/aside | 1000 | 206672 | 0.025 | 13.943 | 15.559 | 19.327 | 23.183 | 16.2 | 436 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1378 | 0.714 | 1.001 | 1.327 | 2.291 | 3.295 | 717.2 | 82 | 0 |
| valkey-go/aside | 10 | 96250 | 0.072 | 0.212 | 0.493 | 0.913 | 4.911 | 54.3 | 209 | 0 |
| valkey-go/aside | 100 | 58204 | 0.413 | 7.007 | 12.191 | 18.191 | 34.623 | 90.3 | 236 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2457 | 0.392 | 0.466 | 0.711 | 1.231 | 1.942 | 499.7 | 29 | 0 |
| valkey-go/lock | 10 | 14135 | 0.678 | 0.870 | 1.396 | 1.982 | 3.239 | 183.3 | 32 | 0 |
| valkey-go/lock | 100 | 27531 | 3.621 | 4.127 | 5.175 | 5.907 | 7.339 | 91.6 | 62 | 0 |
| valkey-go/lock | 1000 | 15009 | 53.855 | 74.367 | 82.687 | 92.351 | 126.079 | 139.9 | 333 | 8719 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 238 | 4.135 | 4.607 | 5.707 | 7.943 | 7.967 | 5038.4 | 27 | 0 |
| valkey-go/lock | 10 | 1246 | 7.915 | 8.911 | 10.111 | 12.279 | 14.151 | 1991.7 | 32 | 0 |
| valkey-go/lock | 100 | 1720 | 57.951 | 60.159 | 64.991 | 66.431 | 67.263 | 1213.0 | 62 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15509 | 0.062 | 0.070 | 0.099 | 0.298 | 1.495 | 45.1 | 37 | 0 |
| valkey-go | 10 | 72837 | 0.133 | 0.187 | 0.251 | 0.469 | 0.834 | 43.2 | 39 | 0 |
| valkey-go | 100 | 323895 | 0.304 | 0.422 | 0.574 | 0.854 | 2.709 | 15.4 | 66 | 0 |
| valkey-go | 1000 | 814650 | 1.265 | 1.630 | 2.159 | 3.099 | 5.387 | 6.1 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13073 | 0.073 | 0.083 | 0.125 | 0.365 | 1.515 | 52.4 | 35 | 0 |
| valkey-go | 10 | 47433 | 0.206 | 0.283 | 0.398 | 0.629 | 1.314 | 59.2 | 40 | 0 |
| valkey-go | 100 | 95832 | 1.084 | 1.361 | 1.740 | 2.527 | 3.851 | 26.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13922 | 0.068 | 0.080 | 0.112 | 0.356 | 1.846 | 53.9 | 30 | 0 |
| go-redis | 10 | 52941 | 0.187 | 0.243 | 0.334 | 0.585 | 2.131 | 58.9 | 33 | 0 |
| go-redis | 100 | 52647 | 1.824 | 2.004 | 3.645 | 4.089 | 6.383 | 62.7 | 67 | 0 |
| go-redis | 1000 | 52985 | 18.703 | 20.095 | 22.143 | 24.591 | 28.111 | 74.0 | 284 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11931 | 0.078 | 0.097 | 0.156 | 0.423 | 1.589 | 62.2 | 27 | 0 |
| go-redis | 10 | 38432 | 0.254 | 0.331 | 0.555 | 0.885 | 2.233 | 78.2 | 34 | 0 |
| go-redis | 100 | 40028 | 2.393 | 2.697 | 4.763 | 5.267 | 6.539 | 78.6 | 73 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10031 | 0.092 | 0.123 | 0.179 | 0.438 | 1.666 | 100.3 | 28 | 0 |
| go-redis-autopipeline | 10 | 62918 | 0.149 | 0.196 | 0.345 | 0.732 | 1.715 | 19.7 | 31 | 0 |
| go-redis-autopipeline | 100 | 210843 | 0.462 | 0.573 | 0.799 | 1.514 | 3.539 | 7.3 | 59 | 0 |
| go-redis-autopipeline | 1000 | 467932 | 2.139 | 2.803 | 3.597 | 5.023 | 7.727 | 6.0 | 312 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11893 | 0.078 | 0.097 | 0.158 | 0.433 | 1.617 | 62.5 | 30 | 0 |
| go-redis-autopipeline | 10 | 38312 | 0.255 | 0.329 | 0.548 | 0.942 | 2.631 | 77.4 | 38 | 0 |
| go-redis-autopipeline | 100 | 40327 | 2.381 | 2.619 | 4.755 | 5.215 | 6.815 | 77.6 | 76 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7885 | 0.118 | 0.150 | 0.244 | 0.498 | 1.686 | 118.8 | 32 | 0 |
| glide | 10 | 34087 | 0.280 | 0.379 | 0.527 | 0.736 | 2.059 | 53.5 | 35 | 0 |
| glide | 100 | 36062 | 2.853 | 3.249 | 4.263 | 4.899 | 6.855 | 54.7 | 64 | 0 |
| glide | 1000 | 35607 | 28.143 | 31.295 | 33.759 | 35.967 | 36.415 | 56.5 | 237 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6768 | 0.135 | 0.178 | 0.322 | 0.610 | 1.962 | 120.0 | 33 | 0 |
| glide | 10 | 26681 | 0.354 | 0.531 | 0.892 | 1.475 | 2.379 | 86.9 | 36 | 0 |
| glide | 100 | 28844 | 3.461 | 4.391 | 5.423 | 6.315 | 8.163 | 81.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12537 | 0.074 | 0.092 | 0.137 | 0.368 | 1.590 | 68.4 | 37 | 0 |
| valkey-go/always-pipe | 10 | 72215 | 0.134 | 0.190 | 0.257 | 0.461 | 1.300 | 43.6 | 41 | 0 |
| valkey-go/always-pipe | 100 | 328285 | 0.301 | 0.415 | 0.548 | 0.842 | 2.014 | 15.3 | 68 | 0 |
| valkey-go/always-pipe | 1000 | 884064 | 1.162 | 1.503 | 1.854 | 2.969 | 5.091 | 6.2 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 10169 | 0.093 | 0.114 | 0.169 | 0.422 | 1.585 | 86.5 | 36 | 0 |
| valkey-go/always-pipe | 10 | 47585 | 0.205 | 0.281 | 0.406 | 0.656 | 1.569 | 59.5 | 39 | 0 |
| valkey-go/always-pipe | 100 | 101205 | 1.023 | 1.284 | 1.659 | 2.341 | 3.263 | 26.6 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15398 | 0.062 | 0.074 | 0.105 | 0.288 | 1.503 | 46.9 | 37 | 0 |
| valkey-go/flush-20us | 10 | 62512 | 0.143 | 0.247 | 0.401 | 0.588 | 1.893 | 50.0 | 40 | 0 |
| valkey-go/flush-20us | 100 | 343726 | 0.287 | 0.394 | 0.549 | 0.850 | 1.900 | 13.4 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 901997 | 1.137 | 1.489 | 1.894 | 2.713 | 6.067 | 6.1 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 13065 | 0.072 | 0.086 | 0.132 | 0.380 | 1.552 | 53.3 | 35 | 0 |
| valkey-go/flush-20us | 10 | 47348 | 0.206 | 0.284 | 0.405 | 0.666 | 1.733 | 61.6 | 39 | 0 |
| valkey-go/flush-20us | 100 | 103594 | 0.997 | 1.253 | 1.660 | 2.435 | 3.207 | 25.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15742 | 0.061 | 0.071 | 0.099 | 0.258 | 1.573 | 45.0 | 36 | 0 |
| valkey-go/flush-100us | 10 | 60729 | 0.157 | 0.240 | 0.318 | 1.332 | 7.099 | 36.0 | 41 | 0 |
| valkey-go/flush-100us | 100 | 395521 | 0.242 | 0.346 | 0.510 | 1.168 | 3.759 | 8.9 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 924919 | 1.089 | 1.475 | 2.101 | 3.383 | 7.671 | 5.3 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12772 | 0.073 | 0.091 | 0.145 | 0.407 | 3.443 | 53.4 | 35 | 0 |
| valkey-go/flush-100us | 10 | 42726 | 0.224 | 0.338 | 0.492 | 0.860 | 3.579 | 57.2 | 39 | 0 |
| valkey-go/flush-100us | 100 | 102565 | 0.976 | 1.271 | 1.890 | 2.973 | 7.127 | 22.5 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 14856 | 0.062 | 0.076 | 0.125 | 0.411 | 8.231 | 47.5 | 27 | 0 |
| valkey-go/1-conn | 10 | 79978 | 0.123 | 0.162 | 0.224 | 0.471 | 1.607 | 23.2 | 32 | 0 |
| valkey-go/1-conn | 100 | 421131 | 0.230 | 0.333 | 0.437 | 0.892 | 1.713 | 7.1 | 56 | 0 |
| valkey-go/1-conn | 1000 | 1010928 | 0.963 | 1.097 | 1.675 | 2.465 | 4.303 | 5.2 | 305 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12705 | 0.073 | 0.091 | 0.156 | 0.421 | 1.466 | 54.2 | 28 | 0 |
| valkey-go/1-conn | 10 | 48296 | 0.204 | 0.274 | 0.454 | 0.798 | 3.643 | 41.0 | 30 | 0 |
| valkey-go/1-conn | 100 | 116302 | 0.833 | 0.957 | 1.381 | 2.079 | 3.127 | 18.1 | 57 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15276 | 0.062 | 0.072 | 0.103 | 0.326 | 1.435 | 46.0 | 35 | 0 |
| valkey-go/no-retry | 10 | 71650 | 0.134 | 0.192 | 0.267 | 0.504 | 1.722 | 43.7 | 41 | 0 |
| valkey-go/no-retry | 100 | 307860 | 0.314 | 0.449 | 0.689 | 1.523 | 3.231 | 15.9 | 68 | 0 |
| valkey-go/no-retry | 1000 | 827041 | 1.203 | 1.662 | 2.525 | 3.439 | 6.347 | 6.6 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12439 | 0.074 | 0.094 | 0.165 | 0.441 | 1.595 | 54.2 | 35 | 0 |
| valkey-go/no-retry | 10 | 46391 | 0.209 | 0.290 | 0.438 | 0.715 | 3.545 | 60.5 | 39 | 0 |
| valkey-go/no-retry | 100 | 97339 | 1.031 | 1.373 | 2.119 | 2.873 | 4.355 | 28.0 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 11068 | 0.096 | 0.136 | 0.212 | 1.203 | 4.995 | 82.4 | 66 | 0 |
| valkey-go/docache | 10 | 135006 | 0.002 | 0.216 | 0.319 | 0.554 | 3.191 | 24.7 | 149 | 0 |
| valkey-go/docache | 100 | 461728 | 0.010 | 0.670 | 0.986 | 1.549 | 6.539 | 12.8 | 180 | 0 |
| valkey-go/docache | 1000 | 521237 | 0.017 | 6.215 | 8.179 | 10.199 | 19.087 | 11.3 | 416 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 10763 | 0.034 | 0.196 | 0.280 | 0.496 | 2.895 | 155.1 | 152 | 0 |
| valkey-go/docache | 10 | 341769 | 0.021 | 0.032 | 0.195 | 0.376 | 4.049 | 34.3 | 194 | 0 |
| valkey-go/docache | 100 | 567441 | 0.105 | 0.361 | 1.117 | 3.305 | 26.751 | 28.3 | 251 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14659 | 0.064 | 0.077 | 0.118 | 0.370 | 1.655 | 47.4 | 35 | 0 |
| valkey-go/compat | 10 | 69274 | 0.139 | 0.198 | 0.271 | 0.514 | 2.437 | 44.1 | 40 | 0 |
| valkey-go/compat | 100 | 306552 | 0.322 | 0.443 | 0.615 | 1.120 | 2.315 | 16.1 | 66 | 0 |
| valkey-go/compat | 1000 | 865088 | 1.170 | 1.554 | 2.159 | 3.661 | 7.727 | 6.7 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12203 | 0.076 | 0.095 | 0.163 | 0.418 | 1.527 | 57.2 | 38 | 0 |
| valkey-go/compat | 10 | 45320 | 0.212 | 0.298 | 0.496 | 0.794 | 2.151 | 68.4 | 42 | 0 |
| valkey-go/compat | 100 | 86171 | 1.077 | 1.746 | 3.581 | 5.199 | 8.471 | 37.9 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 14425 | 0.063 | 0.080 | 0.139 | 0.649 | 2.089 | 48.2 | 34 | 0 |
| valkey-go/hook | 10 | 69647 | 0.138 | 0.197 | 0.274 | 0.537 | 2.495 | 44.7 | 39 | 0 |
| valkey-go/hook | 100 | 312808 | 0.314 | 0.439 | 0.615 | 1.029 | 2.403 | 16.3 | 67 | 0 |
| valkey-go/hook | 1000 | 868636 | 1.179 | 1.550 | 2.012 | 2.907 | 6.155 | 7.0 | 312 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12297 | 0.075 | 0.096 | 0.154 | 0.409 | 3.355 | 56.5 | 36 | 0 |
| valkey-go/hook | 10 | 44323 | 0.216 | 0.316 | 0.462 | 0.724 | 1.890 | 61.9 | 39 | 0 |
| valkey-go/hook | 100 | 103226 | 1.000 | 1.257 | 1.652 | 2.411 | 3.567 | 27.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14017 | 0.066 | 0.082 | 0.136 | 0.481 | 1.795 | 51.6 | 37 | 0 |
| valkey-go/otel | 10 | 69932 | 0.137 | 0.195 | 0.270 | 0.517 | 1.558 | 48.0 | 42 | 0 |
| valkey-go/otel | 100 | 305262 | 0.316 | 0.451 | 0.727 | 1.700 | 7.331 | 18.1 | 67 | 0 |
| valkey-go/otel | 1000 | 844764 | 1.181 | 1.617 | 2.553 | 4.379 | 8.519 | 8.3 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 11748 | 0.076 | 0.103 | 0.190 | 0.597 | 2.111 | 57.6 | 38 | 0 |
| valkey-go/otel | 10 | 43806 | 0.215 | 0.316 | 0.532 | 1.374 | 4.655 | 66.4 | 39 | 0 |
| valkey-go/otel | 100 | 100590 | 1.014 | 1.310 | 1.854 | 2.965 | 4.971 | 32.2 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10680 | 0.088 | 0.113 | 0.176 | 0.440 | 2.919 | 101.1 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 64531 | 0.145 | 0.213 | 0.356 | 0.707 | 1.949 | 60.4 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 279252 | 0.319 | 0.512 | 1.179 | 2.687 | 13.023 | 26.7 | 69 | 0 |
| valkey-go/otel-sdk | 1000 | 747603 | 1.197 | 1.963 | 4.559 | 7.515 | 15.279 | 16.7 | 324 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9374 | 0.103 | 0.123 | 0.200 | 0.479 | 1.592 | 111.0 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 44447 | 0.211 | 0.305 | 0.561 | 1.077 | 8.043 | 76.1 | 43 | 0 |
| valkey-go/otel-sdk | 100 | 100297 | 1.011 | 1.337 | 2.030 | 3.243 | 4.899 | 40.9 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11425 | 0.098 | 0.124 | 0.200 | 0.456 | 1.572 | 84.0 | 74 | 0 |
| valkey-go/aside | 10 | 132765 | 0.011 | 0.203 | 0.334 | 0.741 | 4.887 | 21.2 | 169 | 0 |
| valkey-go/aside | 100 | 201801 | 0.024 | 1.434 | 1.795 | 3.021 | 9.207 | 16.7 | 196 | 0 |
| valkey-go/aside | 1000 | 197913 | 0.026 | 14.367 | 15.935 | 21.407 | 29.663 | 16.9 | 419 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1109 | 0.864 | 1.222 | 2.277 | 3.713 | 4.459 | 845.9 | 78 | 0 |
| valkey-go/aside | 10 | 79922 | 0.086 | 0.248 | 0.646 | 1.549 | 9.279 | 61.8 | 211 | 0 |
| valkey-go/aside | 100 | 54525 | 0.617 | 6.515 | 12.135 | 18.607 | 33.247 | 99.3 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2268 | 0.411 | 0.523 | 0.950 | 2.008 | 2.327 | 536.8 | 28 | 0 |
| valkey-go/lock | 10 | 13907 | 0.682 | 0.874 | 1.623 | 2.529 | 3.417 | 189.3 | 33 | 0 |
| valkey-go/lock | 100 | 26321 | 3.713 | 4.575 | 5.951 | 7.019 | 8.591 | 102.5 | 62 | 0 |
| valkey-go/lock | 1000 | 15498 | 54.463 | 74.687 | 83.199 | 88.127 | 94.591 | 144.9 | 324 | 6528 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 229 | 4.195 | 5.027 | 6.751 | 8.431 | 8.631 | 5296.5 | 27 | 0 |
| valkey-go/lock | 10 | 1141 | 8.375 | 10.743 | 13.439 | 14.783 | 15.703 | 2116.5 | 33 | 1 |
| valkey-go/lock | 100 | 1444 | 62.207 | 96.895 | 103.295 | 104.575 | 105.535 | 1424.2 | 62 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 14772 | 0.063 | 0.080 | 0.124 | 0.367 | 1.554 | 48.9 | 34 | 0 |
| valkey-go | 10 | 69687 | 0.136 | 0.199 | 0.287 | 0.577 | 4.479 | 44.6 | 40 | 0 |
| valkey-go | 100 | 320631 | 0.308 | 0.425 | 0.569 | 0.898 | 2.043 | 15.9 | 67 | 0 |
| valkey-go | 1000 | 800460 | 1.271 | 1.669 | 2.425 | 4.203 | 8.271 | 6.6 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 11369 | 0.075 | 0.126 | 0.197 | 0.447 | 1.791 | 60.9 | 35 | 0 |
| valkey-go | 10 | 45011 | 0.212 | 0.300 | 0.494 | 1.438 | 6.063 | 65.3 | 39 | 0 |
| valkey-go | 100 | 94949 | 1.086 | 1.381 | 1.861 | 2.511 | 4.063 | 28.8 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 13422 | 0.067 | 0.082 | 0.159 | 1.036 | 3.487 | 54.9 | 28 | 0 |
| go-redis | 10 | 46059 | 0.203 | 0.300 | 0.530 | 1.577 | 3.013 | 67.3 | 32 | 0 |
| go-redis | 100 | 52211 | 1.816 | 2.199 | 3.625 | 4.347 | 8.215 | 68.1 | 70 | 0 |
| go-redis | 1000 | 51718 | 19.055 | 21.151 | 23.919 | 30.063 | 33.151 | 78.6 | 240 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11747 | 0.078 | 0.100 | 0.171 | 0.467 | 1.595 | 64.2 | 30 | 0 |
| go-redis | 10 | 37479 | 0.258 | 0.339 | 0.622 | 1.266 | 2.133 | 81.2 | 37 | 0 |
| go-redis | 100 | 37537 | 2.451 | 3.459 | 4.979 | 6.619 | 9.511 | 82.4 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9455 | 0.095 | 0.130 | 0.212 | 0.725 | 1.639 | 107.5 | 28 | 0 |
| go-redis-autopipeline | 10 | 59001 | 0.158 | 0.210 | 0.373 | 0.937 | 2.107 | 21.3 | 33 | 0 |
| go-redis-autopipeline | 100 | 202495 | 0.479 | 0.586 | 0.859 | 1.850 | 2.833 | 8.2 | 57 | 0 |
| go-redis-autopipeline | 1000 | 434920 | 2.203 | 3.061 | 4.071 | 6.071 | 7.907 | 6.8 | 307 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11327 | 0.079 | 0.104 | 0.190 | 0.629 | 1.731 | 67.4 | 27 | 0 |
| go-redis-autopipeline | 10 | 37457 | 0.257 | 0.339 | 0.630 | 1.450 | 2.223 | 82.3 | 35 | 0 |
| go-redis-autopipeline | 100 | 37024 | 2.445 | 3.633 | 4.987 | 7.115 | 9.823 | 83.7 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7187 | 0.129 | 0.175 | 0.287 | 0.749 | 1.965 | 127.7 | 32 | 0 |
| glide | 10 | 33230 | 0.289 | 0.380 | 0.549 | 0.891 | 1.968 | 56.1 | 37 | 0 |
| glide | 100 | 35767 | 2.833 | 3.343 | 4.419 | 5.143 | 5.543 | 56.9 | 64 | 0 |
| glide | 1000 | 35902 | 28.031 | 30.719 | 33.247 | 37.343 | 56.543 | 58.0 | 220 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 5968 | 0.146 | 0.216 | 0.426 | 1.621 | 2.231 | 130.7 | 33 | 0 |
| glide | 10 | 25271 | 0.374 | 0.554 | 0.982 | 1.618 | 2.493 | 94.6 | 35 | 0 |
| glide | 100 | 28942 | 3.453 | 4.395 | 5.391 | 6.255 | 8.663 | 85.2 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15586 | 0.062 | 0.071 | 0.100 | 0.273 | 1.641 | 45.9 | 28 | 0 |
| valkey-go/no-autopipe | 10 | 53306 | 0.187 | 0.239 | 0.314 | 0.539 | 1.565 | 51.5 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 53101 | 1.804 | 2.023 | 3.597 | 4.093 | 6.191 | 56.5 | 177 | 0 |
| valkey-go/no-autopipe | 1000 | 48979 | 20.143 | 21.599 | 24.671 | 40.927 | 46.975 | 53.1 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12810 | 0.073 | 0.091 | 0.141 | 0.381 | 1.495 | 54.5 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 39145 | 0.255 | 0.317 | 0.417 | 0.651 | 1.803 | 63.9 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 39794 | 2.415 | 2.617 | 4.815 | 5.211 | 6.955 | 66.8 | 261 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15488 | 0.062 | 0.071 | 0.099 | 0.253 | 1.414 | 46.9 | 36 | 0 |
| valkey-go/flush-20us | 10 | 72579 | 0.133 | 0.188 | 0.255 | 0.475 | 1.811 | 45.5 | 40 | 0 |
| valkey-go/flush-20us | 100 | 346103 | 0.288 | 0.389 | 0.509 | 0.805 | 2.429 | 13.7 | 66 | 0 |
| valkey-go/flush-20us | 1000 | 914473 | 1.123 | 1.471 | 1.799 | 2.629 | 4.967 | 6.3 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12973 | 0.073 | 0.088 | 0.129 | 0.370 | 1.703 | 52.9 | 35 | 0 |
| valkey-go/flush-20us | 10 | 46502 | 0.208 | 0.291 | 0.433 | 0.724 | 1.887 | 61.5 | 40 | 0 |
| valkey-go/flush-20us | 100 | 101597 | 1.015 | 1.277 | 1.723 | 2.531 | 4.159 | 27.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15511 | 0.062 | 0.070 | 0.100 | 0.259 | 1.736 | 45.4 | 35 | 0 |
| valkey-go/flush-100us | 10 | 60604 | 0.158 | 0.242 | 0.316 | 1.310 | 1.566 | 36.1 | 41 | 0 |
| valkey-go/flush-100us | 100 | 402295 | 0.239 | 0.340 | 0.479 | 0.742 | 2.207 | 9.1 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 891752 | 1.114 | 1.566 | 2.247 | 3.307 | 6.363 | 5.9 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12800 | 0.074 | 0.088 | 0.126 | 0.366 | 1.559 | 53.8 | 35 | 0 |
| valkey-go/flush-100us | 10 | 43633 | 0.221 | 0.329 | 0.455 | 0.699 | 1.710 | 55.5 | 39 | 0 |
| valkey-go/flush-100us | 100 | 103742 | 0.984 | 1.234 | 1.743 | 2.281 | 3.225 | 23.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15293 | 0.062 | 0.074 | 0.100 | 0.270 | 1.405 | 47.1 | 27 | 0 |
| valkey-go/1-conn | 10 | 80573 | 0.123 | 0.161 | 0.216 | 0.436 | 1.642 | 23.3 | 31 | 0 |
| valkey-go/1-conn | 100 | 416715 | 0.234 | 0.335 | 0.431 | 0.715 | 1.629 | 7.3 | 59 | 0 |
| valkey-go/1-conn | 1000 | 1020158 | 0.966 | 1.079 | 1.389 | 2.697 | 4.351 | 5.3 | 303 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12932 | 0.072 | 0.088 | 0.152 | 0.397 | 1.457 | 55.1 | 28 | 0 |
| valkey-go/1-conn | 10 | 46873 | 0.209 | 0.282 | 0.479 | 1.170 | 2.151 | 41.9 | 30 | 0 |
| valkey-go/1-conn | 100 | 113817 | 0.835 | 1.040 | 1.539 | 2.249 | 4.135 | 18.9 | 57 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 14490 | 0.063 | 0.078 | 0.131 | 0.507 | 2.849 | 48.7 | 35 | 0 |
| valkey-go/no-retry | 10 | 72245 | 0.133 | 0.190 | 0.259 | 0.502 | 1.690 | 44.0 | 40 | 0 |
| valkey-go/no-retry | 100 | 317394 | 0.308 | 0.433 | 0.612 | 1.135 | 2.231 | 15.9 | 67 | 0 |
| valkey-go/no-retry | 1000 | 880187 | 1.160 | 1.523 | 2.029 | 3.317 | 6.971 | 6.9 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12887 | 0.073 | 0.088 | 0.131 | 0.377 | 1.644 | 53.3 | 37 | 0 |
| valkey-go/no-retry | 10 | 46983 | 0.207 | 0.286 | 0.417 | 0.678 | 1.977 | 60.1 | 39 | 0 |
| valkey-go/no-retry | 100 | 100378 | 1.018 | 1.314 | 1.839 | 2.505 | 4.491 | 28.6 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12330 | 0.090 | 0.118 | 0.185 | 0.445 | 3.009 | 76.0 | 66 | 0 |
| valkey-go/docache | 10 | 140010 | 0.002 | 0.210 | 0.297 | 0.477 | 3.755 | 24.7 | 151 | 0 |
| valkey-go/docache | 100 | 459927 | 0.010 | 0.673 | 0.961 | 1.521 | 4.703 | 13.0 | 182 | 0 |
| valkey-go/docache | 1000 | 524955 | 0.017 | 6.131 | 7.787 | 9.975 | 16.927 | 11.8 | 420 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 9974 | 0.057 | 0.202 | 0.296 | 0.516 | 2.095 | 169.1 | 140 | 0 |
| valkey-go/docache | 10 | 340333 | 0.021 | 0.033 | 0.195 | 0.376 | 4.915 | 34.5 | 196 | 0 |
| valkey-go/docache | 100 | 530286 | 0.112 | 0.378 | 1.277 | 3.875 | 23.791 | 29.8 | 264 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14300 | 0.065 | 0.080 | 0.129 | 0.409 | 1.957 | 48.9 | 35 | 0 |
| valkey-go/compat | 10 | 68133 | 0.140 | 0.200 | 0.282 | 0.624 | 2.759 | 44.9 | 41 | 0 |
| valkey-go/compat | 100 | 291115 | 0.334 | 0.470 | 0.720 | 1.907 | 3.007 | 16.7 | 66 | 0 |
| valkey-go/compat | 1000 | 810340 | 1.219 | 1.707 | 2.803 | 4.607 | 8.295 | 7.4 | 312 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12338 | 0.075 | 0.093 | 0.151 | 0.426 | 1.578 | 57.5 | 37 | 0 |
| valkey-go/compat | 10 | 44861 | 0.212 | 0.302 | 0.519 | 1.000 | 2.267 | 69.1 | 41 | 0 |
| valkey-go/compat | 100 | 98140 | 1.029 | 1.362 | 2.097 | 2.911 | 4.147 | 35.5 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 14858 | 0.063 | 0.076 | 0.113 | 0.342 | 1.890 | 48.7 | 35 | 0 |
| valkey-go/hook | 10 | 70423 | 0.135 | 0.195 | 0.274 | 0.602 | 3.103 | 44.2 | 41 | 0 |
| valkey-go/hook | 100 | 308737 | 0.317 | 0.444 | 0.634 | 1.386 | 2.517 | 16.3 | 66 | 0 |
| valkey-go/hook | 1000 | 831904 | 1.199 | 1.645 | 2.601 | 4.275 | 10.735 | 7.2 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12732 | 0.074 | 0.089 | 0.139 | 0.393 | 1.643 | 54.2 | 35 | 0 |
| valkey-go/hook | 10 | 45611 | 0.213 | 0.296 | 0.442 | 0.742 | 2.379 | 62.9 | 39 | 0 |
| valkey-go/hook | 100 | 103557 | 0.992 | 1.261 | 1.715 | 2.541 | 3.839 | 28.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 15016 | 0.064 | 0.075 | 0.108 | 0.313 | 1.959 | 48.8 | 35 | 0 |
| valkey-go/otel | 10 | 66032 | 0.140 | 0.215 | 0.367 | 0.591 | 1.792 | 50.3 | 39 | 0 |
| valkey-go/otel | 100 | 312055 | 0.313 | 0.440 | 0.663 | 1.352 | 2.443 | 18.2 | 65 | 0 |
| valkey-go/otel | 1000 | 869015 | 1.167 | 1.560 | 2.293 | 4.423 | 8.887 | 8.6 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12325 | 0.076 | 0.093 | 0.134 | 0.387 | 1.553 | 58.6 | 35 | 0 |
| valkey-go/otel | 10 | 46230 | 0.209 | 0.292 | 0.458 | 0.774 | 2.489 | 65.1 | 40 | 0 |
| valkey-go/otel | 100 | 101801 | 1.004 | 1.294 | 1.791 | 2.609 | 3.869 | 31.9 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 11108 | 0.086 | 0.106 | 0.156 | 0.399 | 1.601 | 98.1 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 66575 | 0.141 | 0.204 | 0.347 | 0.679 | 2.291 | 61.0 | 39 | 0 |
| valkey-go/otel-sdk | 100 | 295379 | 0.307 | 0.474 | 1.096 | 2.347 | 5.223 | 27.0 | 67 | 0 |
| valkey-go/otel-sdk | 1000 | 790692 | 1.150 | 1.687 | 4.835 | 8.059 | 20.463 | 17.1 | 326 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9802 | 0.096 | 0.119 | 0.183 | 0.440 | 1.525 | 104.2 | 38 | 0 |
| valkey-go/otel-sdk | 10 | 44834 | 0.211 | 0.303 | 0.548 | 0.909 | 2.101 | 76.5 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 101732 | 1.000 | 1.312 | 1.974 | 3.295 | 4.523 | 42.0 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11632 | 0.098 | 0.121 | 0.186 | 0.427 | 1.535 | 83.4 | 74 | 0 |
| valkey-go/aside | 10 | 136538 | 0.012 | 0.199 | 0.318 | 0.501 | 6.119 | 21.2 | 170 | 0 |
| valkey-go/aside | 100 | 193519 | 0.025 | 1.487 | 1.844 | 3.527 | 11.415 | 17.4 | 196 | 0 |
| valkey-go/aside | 1000 | 192498 | 0.027 | 14.695 | 15.767 | 18.655 | 21.759 | 17.5 | 440 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1388 | 0.708 | 1.004 | 1.327 | 2.297 | 2.777 | 716.2 | 86 | 0 |
| valkey-go/aside | 10 | 84192 | 0.085 | 0.236 | 0.553 | 1.120 | 12.735 | 60.7 | 211 | 0 |
| valkey-go/aside | 100 | 51051 | 0.392 | 7.791 | 13.983 | 19.647 | 35.327 | 94.7 | 234 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2401 | 0.400 | 0.478 | 0.726 | 1.388 | 2.024 | 512.2 | 27 | 0 |
| valkey-go/lock | 10 | 14384 | 0.669 | 0.838 | 1.366 | 2.093 | 2.847 | 183.9 | 32 | 0 |
| valkey-go/lock | 100 | 27181 | 3.657 | 4.115 | 5.367 | 6.383 | 8.079 | 100.0 | 62 | 3 |
| valkey-go/lock | 1000 | 13929 | 55.455 | 75.455 | 83.711 | 88.127 | 98.367 | 152.8 | 332 | 10624 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 237 | 4.139 | 4.567 | 6.003 | 7.639 | 8.059 | 5128.3 | 28 | 0 |
| valkey-go/lock | 10 | 1206 | 8.143 | 9.287 | 10.839 | 12.383 | 12.959 | 2093.3 | 34 | 0 |
| valkey-go/lock | 100 | 1696 | 58.463 | 60.895 | 82.047 | 85.311 | 85.503 | 1361.8 | 64 | 2 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15241 | 0.063 | 0.073 | 0.103 | 0.307 | 1.506 | 46.2 | 35 | 0 |
| valkey-go | 10 | 71841 | 0.135 | 0.189 | 0.251 | 0.476 | 1.594 | 44.1 | 40 | 0 |
| valkey-go | 100 | 324735 | 0.304 | 0.420 | 0.558 | 0.896 | 2.411 | 16.0 | 65 | 0 |
| valkey-go | 1000 | 822634 | 1.256 | 1.634 | 2.025 | 3.121 | 6.095 | 6.9 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12943 | 0.072 | 0.089 | 0.146 | 0.390 | 1.503 | 53.3 | 35 | 0 |
| valkey-go | 10 | 46835 | 0.208 | 0.287 | 0.415 | 0.686 | 1.952 | 60.2 | 39 | 0 |
| valkey-go | 100 | 94944 | 1.087 | 1.381 | 1.829 | 2.587 | 3.793 | 28.8 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14336 | 0.067 | 0.078 | 0.107 | 0.359 | 1.553 | 52.8 | 28 | 0 |
| go-redis | 10 | 52416 | 0.189 | 0.247 | 0.342 | 0.600 | 1.898 | 60.3 | 32 | 0 |
| go-redis | 100 | 52853 | 1.816 | 2.030 | 3.623 | 4.053 | 5.331 | 66.6 | 70 | 0 |
| go-redis | 1000 | 53032 | 18.687 | 20.095 | 22.063 | 23.759 | 25.919 | 77.8 | 285 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 12071 | 0.077 | 0.094 | 0.156 | 0.422 | 1.637 | 62.7 | 29 | 0 |
| go-redis | 10 | 38131 | 0.256 | 0.332 | 0.580 | 0.968 | 2.051 | 80.2 | 34 | 0 |
| go-redis | 100 | 33256 | 2.443 | 4.419 | 6.755 | 9.135 | 15.455 | 94.8 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9385 | 0.099 | 0.133 | 0.207 | 0.447 | 1.987 | 109.9 | 30 | 0 |
| go-redis-autopipeline | 10 | 59312 | 0.157 | 0.209 | 0.356 | 0.769 | 2.253 | 21.3 | 32 | 0 |
| go-redis-autopipeline | 100 | 204160 | 0.477 | 0.585 | 0.826 | 1.472 | 2.197 | 8.0 | 60 | 0 |
| go-redis-autopipeline | 1000 | 446304 | 2.159 | 2.981 | 3.707 | 5.299 | 6.839 | 6.8 | 309 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11826 | 0.078 | 0.097 | 0.175 | 0.452 | 2.163 | 63.6 | 28 | 0 |
| go-redis-autopipeline | 10 | 38267 | 0.255 | 0.333 | 0.590 | 0.957 | 2.049 | 80.1 | 35 | 0 |
| go-redis-autopipeline | 100 | 39370 | 2.423 | 2.787 | 4.823 | 5.455 | 6.991 | 84.7 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7261 | 0.129 | 0.164 | 0.284 | 0.514 | 1.911 | 124.6 | 33 | 0 |
| glide | 10 | 32733 | 0.294 | 0.384 | 0.552 | 0.853 | 2.085 | 55.9 | 35 | 0 |
| glide | 100 | 36457 | 2.839 | 3.183 | 4.027 | 4.803 | 5.951 | 55.6 | 64 | 0 |
| glide | 1000 | 33872 | 28.927 | 33.631 | 39.807 | 44.127 | 54.463 | 59.5 | 225 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6658 | 0.137 | 0.181 | 0.339 | 0.594 | 1.728 | 122.7 | 34 | 0 |
| glide | 10 | 25183 | 0.376 | 0.554 | 0.948 | 1.494 | 2.879 | 93.2 | 36 | 0 |
| glide | 100 | 28958 | 3.447 | 4.395 | 5.403 | 6.303 | 9.255 | 85.2 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15769 | 0.061 | 0.072 | 0.102 | 0.268 | 1.505 | 45.7 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 51445 | 0.191 | 0.251 | 0.362 | 0.574 | 1.748 | 50.7 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 53206 | 1.814 | 1.975 | 3.613 | 3.975 | 5.923 | 55.0 | 176 | 0 |
| valkey-go/no-autopipe | 1000 | 48172 | 20.143 | 22.847 | 31.711 | 41.567 | 62.591 | 52.2 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12874 | 0.073 | 0.087 | 0.135 | 0.395 | 1.522 | 54.8 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 39018 | 0.255 | 0.318 | 0.431 | 0.677 | 1.502 | 64.3 | 52 | 0 |
| valkey-go/no-autopipe | 100 | 39755 | 2.415 | 2.647 | 4.815 | 5.251 | 7.023 | 66.1 | 259 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12324 | 0.075 | 0.094 | 0.143 | 0.371 | 1.514 | 70.7 | 37 | 0 |
| valkey-go/always-pipe | 10 | 72451 | 0.133 | 0.188 | 0.253 | 0.481 | 2.057 | 43.5 | 41 | 0 |
| valkey-go/always-pipe | 100 | 306257 | 0.321 | 0.447 | 0.602 | 1.196 | 3.429 | 16.2 | 68 | 0 |
| valkey-go/always-pipe | 1000 | 876809 | 1.167 | 1.540 | 1.983 | 2.913 | 6.167 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9540 | 0.100 | 0.123 | 0.190 | 0.436 | 1.742 | 93.1 | 37 | 0 |
| valkey-go/always-pipe | 10 | 47083 | 0.207 | 0.286 | 0.412 | 0.698 | 1.782 | 62.0 | 42 | 0 |
| valkey-go/always-pipe | 100 | 102195 | 1.010 | 1.278 | 1.717 | 2.511 | 3.557 | 28.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15473 | 0.062 | 0.072 | 0.103 | 0.320 | 1.457 | 46.7 | 35 | 0 |
| valkey-go/flush-100us | 10 | 60675 | 0.158 | 0.241 | 0.315 | 1.308 | 1.716 | 36.5 | 41 | 0 |
| valkey-go/flush-100us | 100 | 403858 | 0.238 | 0.339 | 0.485 | 0.773 | 2.259 | 9.0 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 959348 | 1.067 | 1.430 | 1.749 | 2.685 | 4.771 | 5.5 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12976 | 0.073 | 0.088 | 0.125 | 0.378 | 1.699 | 53.1 | 37 | 0 |
| valkey-go/flush-100us | 10 | 43463 | 0.221 | 0.330 | 0.460 | 0.728 | 1.948 | 55.9 | 39 | 0 |
| valkey-go/flush-100us | 100 | 104141 | 0.978 | 1.233 | 1.740 | 2.379 | 4.679 | 23.3 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15664 | 0.061 | 0.073 | 0.100 | 0.258 | 1.487 | 46.2 | 26 | 0 |
| valkey-go/1-conn | 10 | 80910 | 0.122 | 0.160 | 0.215 | 0.439 | 1.763 | 23.2 | 29 | 0 |
| valkey-go/1-conn | 100 | 417638 | 0.232 | 0.335 | 0.438 | 0.750 | 2.455 | 7.3 | 56 | 0 |
| valkey-go/1-conn | 1000 | 1012853 | 0.967 | 1.092 | 1.508 | 2.749 | 4.451 | 5.4 | 301 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 13228 | 0.071 | 0.084 | 0.146 | 0.386 | 1.496 | 53.0 | 26 | 0 |
| valkey-go/1-conn | 10 | 48381 | 0.204 | 0.274 | 0.457 | 0.707 | 1.834 | 41.5 | 30 | 0 |
| valkey-go/1-conn | 100 | 116874 | 0.827 | 0.973 | 1.373 | 2.059 | 2.823 | 18.4 | 55 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15278 | 0.063 | 0.074 | 0.104 | 0.314 | 1.563 | 47.1 | 34 | 0 |
| valkey-go/no-retry | 10 | 72183 | 0.134 | 0.189 | 0.259 | 0.488 | 1.540 | 43.3 | 41 | 0 |
| valkey-go/no-retry | 100 | 322736 | 0.305 | 0.424 | 0.577 | 0.941 | 2.169 | 15.7 | 66 | 0 |
| valkey-go/no-retry | 1000 | 893631 | 1.153 | 1.500 | 1.870 | 3.041 | 6.527 | 6.8 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12944 | 0.073 | 0.088 | 0.132 | 0.389 | 2.655 | 52.8 | 37 | 0 |
| valkey-go/no-retry | 10 | 47243 | 0.206 | 0.283 | 0.407 | 0.683 | 1.695 | 60.0 | 41 | 0 |
| valkey-go/no-retry | 100 | 102499 | 1.003 | 1.273 | 1.756 | 2.507 | 3.499 | 28.2 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12922 | 0.088 | 0.113 | 0.169 | 0.409 | 1.501 | 72.6 | 68 | 0 |
| valkey-go/docache | 10 | 140496 | 0.002 | 0.210 | 0.301 | 0.462 | 3.291 | 24.3 | 153 | 0 |
| valkey-go/docache | 100 | 470431 | 0.010 | 0.660 | 0.935 | 1.346 | 4.499 | 12.9 | 182 | 0 |
| valkey-go/docache | 1000 | 527550 | 0.017 | 6.119 | 7.759 | 10.319 | 18.655 | 11.8 | 415 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12185 | 0.029 | 0.188 | 0.264 | 0.485 | 2.507 | 138.9 | 155 | 0 |
| valkey-go/docache | 10 | 337456 | 0.021 | 0.033 | 0.195 | 0.361 | 6.395 | 34.3 | 195 | 0 |
| valkey-go/docache | 100 | 556967 | 0.113 | 0.360 | 1.126 | 3.269 | 25.423 | 29.0 | 266 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 15125 | 0.063 | 0.074 | 0.105 | 0.285 | 1.566 | 45.5 | 37 | 0 |
| valkey-go/compat | 10 | 69594 | 0.139 | 0.197 | 0.267 | 0.476 | 1.988 | 44.1 | 41 | 0 |
| valkey-go/compat | 100 | 302564 | 0.328 | 0.451 | 0.600 | 1.082 | 2.571 | 16.5 | 67 | 0 |
| valkey-go/compat | 1000 | 882200 | 1.167 | 1.521 | 1.874 | 3.917 | 8.163 | 7.2 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12652 | 0.075 | 0.089 | 0.140 | 0.403 | 2.993 | 55.1 | 37 | 0 |
| valkey-go/compat | 10 | 45675 | 0.210 | 0.296 | 0.501 | 0.814 | 1.631 | 67.8 | 42 | 0 |
| valkey-go/compat | 100 | 101996 | 1.000 | 1.293 | 1.898 | 2.905 | 5.395 | 34.7 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15823 | 0.061 | 0.069 | 0.099 | 0.272 | 1.617 | 45.5 | 35 | 0 |
| valkey-go/hook | 10 | 73364 | 0.132 | 0.186 | 0.251 | 0.473 | 1.688 | 43.5 | 41 | 0 |
| valkey-go/hook | 100 | 324359 | 0.305 | 0.420 | 0.555 | 0.928 | 2.235 | 15.9 | 67 | 0 |
| valkey-go/hook | 1000 | 887025 | 1.153 | 1.516 | 1.989 | 3.255 | 5.887 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12967 | 0.073 | 0.086 | 0.128 | 0.387 | 1.572 | 53.3 | 35 | 0 |
| valkey-go/hook | 10 | 46854 | 0.208 | 0.287 | 0.421 | 0.689 | 2.018 | 61.3 | 40 | 0 |
| valkey-go/hook | 100 | 104261 | 0.985 | 1.253 | 1.675 | 2.423 | 3.623 | 28.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14949 | 0.064 | 0.075 | 0.112 | 0.330 | 1.476 | 48.3 | 37 | 0 |
| valkey-go/otel | 10 | 71434 | 0.134 | 0.190 | 0.261 | 0.525 | 1.497 | 47.7 | 42 | 0 |
| valkey-go/otel | 100 | 314331 | 0.311 | 0.437 | 0.657 | 1.300 | 2.631 | 18.3 | 66 | 0 |
| valkey-go/otel | 1000 | 882185 | 1.149 | 1.539 | 2.193 | 4.523 | 7.911 | 8.7 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12556 | 0.075 | 0.091 | 0.137 | 0.396 | 1.558 | 55.7 | 37 | 0 |
| valkey-go/otel | 10 | 45910 | 0.210 | 0.295 | 0.456 | 0.790 | 1.785 | 66.4 | 41 | 0 |
| valkey-go/otel | 100 | 102429 | 0.999 | 1.286 | 1.773 | 2.645 | 4.659 | 32.1 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10761 | 0.088 | 0.112 | 0.163 | 0.409 | 1.706 | 102.2 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 66526 | 0.141 | 0.205 | 0.345 | 0.686 | 1.945 | 60.4 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 229803 | 0.358 | 0.742 | 1.337 | 2.731 | 5.043 | 31.9 | 68 | 0 |
| valkey-go/otel-sdk | 1000 | 790717 | 1.154 | 1.682 | 4.867 | 7.855 | 14.127 | 17.1 | 325 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9451 | 0.102 | 0.122 | 0.191 | 0.450 | 1.691 | 110.1 | 40 | 0 |
| valkey-go/otel-sdk | 10 | 45018 | 0.211 | 0.301 | 0.540 | 0.888 | 1.903 | 77.4 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 101526 | 0.999 | 1.316 | 2.014 | 3.417 | 4.655 | 42.1 | 71 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11642 | 0.097 | 0.123 | 0.193 | 0.438 | 1.580 | 83.1 | 76 | 0 |
| valkey-go/aside | 10 | 136920 | 0.012 | 0.199 | 0.314 | 0.502 | 3.413 | 21.2 | 171 | 0 |
| valkey-go/aside | 100 | 194935 | 0.025 | 1.492 | 1.856 | 2.749 | 11.031 | 17.2 | 196 | 0 |
| valkey-go/aside | 1000 | 190340 | 0.026 | 14.855 | 16.231 | 24.863 | 28.095 | 17.8 | 393 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1275 | 0.776 | 1.083 | 1.415 | 2.421 | 2.955 | 777.4 | 80 | 0 |
| valkey-go/aside | 10 | 87044 | 0.082 | 0.229 | 0.537 | 1.050 | 7.487 | 59.3 | 209 | 0 |
| valkey-go/aside | 100 | 52624 | 0.424 | 7.607 | 13.255 | 19.391 | 33.599 | 96.0 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2412 | 0.400 | 0.473 | 0.724 | 1.273 | 1.992 | 507.4 | 28 | 0 |
| valkey-go/lock | 10 | 14382 | 0.670 | 0.838 | 1.365 | 2.007 | 2.557 | 182.7 | 33 | 0 |
| valkey-go/lock | 100 | 27159 | 3.661 | 4.151 | 5.563 | 6.783 | 8.663 | 98.4 | 62 | 0 |
| valkey-go/lock | 1000 | 14496 | 54.975 | 75.455 | 83.647 | 88.191 | 97.279 | 152.3 | 335 | 8806 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 237 | 4.139 | 4.607 | 5.999 | 7.411 | 7.847 | 5111.9 | 29 | 0 |
| valkey-go/lock | 10 | 1217 | 8.099 | 9.175 | 10.407 | 11.695 | 13.455 | 2082.3 | 33 | 0 |
| valkey-go/lock | 100 | 1712 | 58.079 | 60.351 | 62.847 | 64.191 | 65.151 | 1331.1 | 64 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15505 | 0.062 | 0.072 | 0.102 | 0.269 | 1.466 | 46.0 | 35 | 0 |
| valkey-go | 10 | 73030 | 0.132 | 0.186 | 0.250 | 0.481 | 1.761 | 43.7 | 39 | 0 |
| valkey-go | 100 | 319224 | 0.308 | 0.428 | 0.587 | 0.976 | 2.797 | 16.0 | 66 | 0 |
| valkey-go | 1000 | 817994 | 1.263 | 1.641 | 2.067 | 2.985 | 5.767 | 6.8 | 318 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12980 | 0.073 | 0.085 | 0.127 | 0.377 | 1.627 | 53.7 | 35 | 0 |
| valkey-go | 10 | 46501 | 0.209 | 0.290 | 0.417 | 0.700 | 2.002 | 62.0 | 39 | 0 |
| valkey-go | 100 | 94946 | 1.088 | 1.380 | 1.843 | 2.533 | 4.515 | 28.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14159 | 0.067 | 0.078 | 0.111 | 0.368 | 1.453 | 53.3 | 27 | 0 |
| go-redis | 10 | 52391 | 0.189 | 0.245 | 0.336 | 0.601 | 1.617 | 59.6 | 33 | 0 |
| go-redis | 100 | 52685 | 1.815 | 2.057 | 3.617 | 4.195 | 6.179 | 65.4 | 70 | 0 |
| go-redis | 1000 | 52538 | 18.831 | 20.335 | 23.519 | 28.255 | 31.999 | 77.8 | 281 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11874 | 0.078 | 0.098 | 0.162 | 0.424 | 1.505 | 62.7 | 29 | 0 |
| go-redis | 10 | 37802 | 0.258 | 0.338 | 0.591 | 0.975 | 1.946 | 81.5 | 37 | 0 |
| go-redis | 100 | 39185 | 2.431 | 2.815 | 4.859 | 5.603 | 8.035 | 81.5 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10182 | 0.091 | 0.120 | 0.189 | 0.426 | 1.589 | 101.3 | 29 | 0 |
| go-redis-autopipeline | 10 | 60921 | 0.153 | 0.203 | 0.365 | 0.749 | 1.603 | 20.9 | 30 | 0 |
| go-redis-autopipeline | 100 | 201095 | 0.484 | 0.595 | 0.830 | 1.543 | 2.411 | 8.1 | 59 | 0 |
| go-redis-autopipeline | 1000 | 456995 | 2.147 | 2.823 | 3.619 | 5.655 | 7.063 | 6.4 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11997 | 0.078 | 0.095 | 0.162 | 0.424 | 1.661 | 62.2 | 28 | 0 |
| go-redis-autopipeline | 10 | 38446 | 0.254 | 0.329 | 0.569 | 0.931 | 1.732 | 79.9 | 34 | 0 |
| go-redis-autopipeline | 100 | 39703 | 2.409 | 2.707 | 4.803 | 5.367 | 7.863 | 81.8 | 72 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7833 | 0.117 | 0.151 | 0.253 | 0.502 | 1.640 | 116.1 | 33 | 0 |
| glide | 10 | 33554 | 0.288 | 0.376 | 0.535 | 0.797 | 2.034 | 55.1 | 38 | 0 |
| glide | 100 | 36420 | 2.811 | 3.213 | 4.295 | 5.023 | 6.483 | 55.6 | 64 | 0 |
| glide | 1000 | 35044 | 28.287 | 31.407 | 37.951 | 42.079 | 42.783 | 58.1 | 231 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6745 | 0.135 | 0.179 | 0.325 | 0.599 | 1.883 | 121.5 | 33 | 0 |
| glide | 10 | 26657 | 0.353 | 0.530 | 0.920 | 1.512 | 2.567 | 91.1 | 36 | 0 |
| glide | 100 | 29139 | 3.433 | 4.371 | 5.375 | 6.243 | 7.923 | 84.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15431 | 0.062 | 0.072 | 0.107 | 0.305 | 1.568 | 47.0 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 52891 | 0.187 | 0.241 | 0.323 | 0.557 | 1.655 | 50.7 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 53239 | 1.806 | 1.983 | 3.601 | 4.007 | 5.803 | 55.1 | 179 | 0 |
| valkey-go/no-autopipe | 1000 | 49972 | 19.615 | 21.391 | 24.127 | 39.935 | 47.583 | 52.1 | 1209 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 13179 | 0.072 | 0.084 | 0.134 | 0.381 | 1.821 | 53.5 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 39264 | 0.254 | 0.316 | 0.421 | 0.662 | 1.785 | 64.2 | 52 | 0 |
| valkey-go/no-autopipe | 100 | 39422 | 2.429 | 2.703 | 4.835 | 5.419 | 7.571 | 66.7 | 262 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12360 | 0.075 | 0.094 | 0.140 | 0.384 | 1.540 | 69.2 | 37 | 0 |
| valkey-go/always-pipe | 10 | 71197 | 0.135 | 0.193 | 0.261 | 0.482 | 1.179 | 43.8 | 41 | 0 |
| valkey-go/always-pipe | 100 | 320836 | 0.307 | 0.427 | 0.576 | 0.883 | 2.831 | 15.9 | 68 | 0 |
| valkey-go/always-pipe | 1000 | 875759 | 1.176 | 1.534 | 1.911 | 2.939 | 5.347 | 6.9 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9452 | 0.102 | 0.123 | 0.190 | 0.438 | 1.723 | 93.8 | 36 | 0 |
| valkey-go/always-pipe | 10 | 46732 | 0.208 | 0.289 | 0.425 | 0.673 | 1.889 | 60.8 | 39 | 0 |
| valkey-go/always-pipe | 100 | 102717 | 1.004 | 1.273 | 1.641 | 2.389 | 3.343 | 28.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 14975 | 0.063 | 0.075 | 0.117 | 0.350 | 1.602 | 47.5 | 36 | 0 |
| valkey-go/flush-20us | 10 | 71116 | 0.135 | 0.192 | 0.261 | 0.471 | 1.487 | 45.2 | 42 | 0 |
| valkey-go/flush-20us | 100 | 342156 | 0.291 | 0.392 | 0.514 | 0.849 | 2.267 | 13.8 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 907439 | 1.125 | 1.492 | 1.923 | 2.995 | 5.975 | 6.4 | 314 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12987 | 0.073 | 0.086 | 0.129 | 0.392 | 1.577 | 53.1 | 35 | 0 |
| valkey-go/flush-20us | 10 | 46922 | 0.207 | 0.288 | 0.417 | 0.678 | 1.422 | 61.4 | 40 | 0 |
| valkey-go/flush-20us | 100 | 103226 | 1.002 | 1.264 | 1.658 | 2.451 | 3.661 | 27.6 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15299 | 0.062 | 0.073 | 0.104 | 0.270 | 1.422 | 46.8 | 29 | 0 |
| valkey-go/1-conn | 10 | 80955 | 0.122 | 0.160 | 0.214 | 0.453 | 1.897 | 23.0 | 32 | 0 |
| valkey-go/1-conn | 100 | 421065 | 0.229 | 0.333 | 0.439 | 0.820 | 2.187 | 7.2 | 57 | 0 |
| valkey-go/1-conn | 1000 | 1011210 | 0.972 | 1.090 | 1.442 | 2.601 | 4.343 | 5.4 | 307 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12879 | 0.073 | 0.087 | 0.144 | 0.400 | 1.515 | 54.0 | 28 | 0 |
| valkey-go/1-conn | 10 | 48530 | 0.204 | 0.272 | 0.459 | 0.722 | 1.591 | 41.3 | 31 | 0 |
| valkey-go/1-conn | 100 | 117183 | 0.823 | 0.970 | 1.414 | 2.109 | 3.369 | 18.7 | 57 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15133 | 0.062 | 0.074 | 0.103 | 0.291 | 1.552 | 48.2 | 37 | 0 |
| valkey-go/no-retry | 10 | 72993 | 0.132 | 0.187 | 0.251 | 0.468 | 1.795 | 43.3 | 41 | 0 |
| valkey-go/no-retry | 100 | 323918 | 0.305 | 0.422 | 0.557 | 0.887 | 2.069 | 15.8 | 68 | 0 |
| valkey-go/no-retry | 1000 | 894457 | 1.149 | 1.501 | 1.856 | 3.181 | 6.195 | 7.0 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 13080 | 0.072 | 0.086 | 0.131 | 0.383 | 1.578 | 52.5 | 37 | 0 |
| valkey-go/no-retry | 10 | 38869 | 0.227 | 0.419 | 0.629 | 0.953 | 1.763 | 70.4 | 41 | 0 |
| valkey-go/no-retry | 100 | 100784 | 1.020 | 1.303 | 1.768 | 2.461 | 3.543 | 28.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12241 | 0.090 | 0.120 | 0.194 | 0.425 | 3.081 | 75.9 | 64 | 0 |
| valkey-go/docache | 10 | 139852 | 0.002 | 0.210 | 0.298 | 0.494 | 3.483 | 24.4 | 151 | 0 |
| valkey-go/docache | 100 | 461757 | 0.010 | 0.672 | 0.963 | 1.398 | 3.967 | 12.9 | 182 | 0 |
| valkey-go/docache | 1000 | 520871 | 0.017 | 6.183 | 7.883 | 10.935 | 19.471 | 11.9 | 372 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 10733 | 0.035 | 0.197 | 0.284 | 0.494 | 3.435 | 155.3 | 142 | 0 |
| valkey-go/docache | 10 | 336677 | 0.021 | 0.033 | 0.198 | 0.367 | 3.183 | 35.2 | 194 | 0 |
| valkey-go/docache | 100 | 560174 | 0.089 | 0.402 | 1.197 | 3.469 | 22.783 | 28.9 | 245 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14878 | 0.064 | 0.075 | 0.109 | 0.364 | 1.513 | 47.8 | 35 | 0 |
| valkey-go/compat | 10 | 68637 | 0.141 | 0.200 | 0.272 | 0.491 | 1.725 | 44.7 | 40 | 0 |
| valkey-go/compat | 100 | 305548 | 0.323 | 0.448 | 0.600 | 1.043 | 2.437 | 16.5 | 65 | 0 |
| valkey-go/compat | 1000 | 885653 | 1.159 | 1.515 | 1.901 | 4.263 | 8.751 | 7.2 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12312 | 0.076 | 0.093 | 0.149 | 0.413 | 1.700 | 57.8 | 38 | 0 |
| valkey-go/compat | 10 | 44739 | 0.214 | 0.304 | 0.507 | 0.838 | 2.245 | 69.3 | 42 | 0 |
| valkey-go/compat | 100 | 101652 | 1.005 | 1.296 | 1.897 | 2.863 | 4.039 | 34.9 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15411 | 0.062 | 0.073 | 0.104 | 0.323 | 1.503 | 46.6 | 35 | 0 |
| valkey-go/hook | 10 | 71376 | 0.135 | 0.192 | 0.265 | 0.496 | 1.683 | 44.2 | 38 | 0 |
| valkey-go/hook | 100 | 316376 | 0.312 | 0.434 | 0.585 | 0.917 | 2.043 | 16.0 | 65 | 0 |
| valkey-go/hook | 1000 | 901767 | 1.141 | 1.485 | 1.855 | 2.819 | 5.943 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12886 | 0.073 | 0.088 | 0.131 | 0.385 | 1.525 | 53.5 | 36 | 0 |
| valkey-go/hook | 10 | 46707 | 0.208 | 0.288 | 0.435 | 0.702 | 2.044 | 62.8 | 41 | 0 |
| valkey-go/hook | 100 | 101537 | 1.007 | 1.286 | 1.790 | 2.613 | 4.691 | 28.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14678 | 0.065 | 0.077 | 0.111 | 0.337 | 1.366 | 50.4 | 37 | 0 |
| valkey-go/otel | 10 | 70973 | 0.135 | 0.192 | 0.263 | 0.504 | 2.075 | 47.7 | 41 | 0 |
| valkey-go/otel | 100 | 313155 | 0.313 | 0.439 | 0.648 | 1.344 | 3.067 | 18.3 | 67 | 0 |
| valkey-go/otel | 1000 | 848276 | 1.187 | 1.614 | 2.385 | 4.623 | 7.527 | 8.6 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12472 | 0.075 | 0.090 | 0.137 | 0.405 | 2.987 | 56.2 | 38 | 0 |
| valkey-go/otel | 10 | 45926 | 0.210 | 0.293 | 0.459 | 0.787 | 1.870 | 65.3 | 42 | 0 |
| valkey-go/otel | 100 | 103584 | 0.988 | 1.266 | 1.748 | 2.603 | 3.699 | 31.4 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 11102 | 0.085 | 0.107 | 0.156 | 0.403 | 1.485 | 97.4 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 62757 | 0.146 | 0.221 | 0.446 | 0.885 | 3.357 | 61.6 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 289224 | 0.315 | 0.486 | 1.119 | 2.339 | 5.603 | 27.2 | 69 | 0 |
| valkey-go/otel-sdk | 1000 | 759063 | 1.190 | 1.797 | 5.027 | 8.367 | 15.911 | 17.5 | 328 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9181 | 0.105 | 0.127 | 0.207 | 0.467 | 1.674 | 113.5 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 43812 | 0.216 | 0.311 | 0.563 | 0.921 | 2.015 | 77.4 | 41 | 0 |
| valkey-go/otel-sdk | 100 | 89262 | 1.105 | 1.582 | 2.461 | 3.779 | 6.663 | 45.0 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 8776 | 0.114 | 0.176 | 0.318 | 0.532 | 1.600 | 109.2 | 62 | 0 |
| valkey-go/aside | 10 | 131361 | 0.013 | 0.205 | 0.331 | 0.551 | 4.139 | 21.9 | 169 | 0 |
| valkey-go/aside | 100 | 185963 | 0.026 | 1.558 | 1.950 | 3.323 | 9.599 | 17.9 | 196 | 0 |
| valkey-go/aside | 1000 | 189070 | 0.027 | 15.095 | 16.751 | 20.591 | 26.111 | 17.7 | 440 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1257 | 0.783 | 1.099 | 1.485 | 2.595 | 3.693 | 791.3 | 80 | 0 |
| valkey-go/aside | 10 | 82413 | 0.087 | 0.239 | 0.561 | 1.074 | 9.231 | 61.5 | 211 | 0 |
| valkey-go/aside | 100 | 50672 | 0.442 | 7.759 | 13.615 | 19.567 | 35.999 | 98.2 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2341 | 0.408 | 0.498 | 0.779 | 1.455 | 2.185 | 517.1 | 30 | 0 |
| valkey-go/lock | 10 | 14197 | 0.676 | 0.851 | 1.437 | 2.139 | 3.635 | 187.2 | 33 | 0 |
| valkey-go/lock | 100 | 26569 | 3.745 | 4.239 | 5.535 | 6.419 | 7.831 | 102.1 | 63 | 3 |
| valkey-go/lock | 1000 | 13951 | 56.223 | 76.159 | 84.863 | 90.495 | 96.127 | 156.1 | 334 | 9712 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 236 | 4.147 | 4.623 | 6.047 | 8.479 | 9.111 | 5160.6 | 28 | 0 |
| valkey-go/lock | 10 | 1189 | 8.247 | 9.375 | 11.655 | 13.823 | 14.543 | 2089.4 | 34 | 0 |
| valkey-go/lock | 100 | 1682 | 59.007 | 62.111 | 63.935 | 64.895 | 65.503 | 1361.7 | 62 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15272 | 0.062 | 0.072 | 0.107 | 0.354 | 1.556 | 46.3 | 34 | 0 |
| valkey-go | 10 | 71794 | 0.135 | 0.190 | 0.256 | 0.500 | 1.561 | 43.9 | 41 | 0 |
| valkey-go | 100 | 320160 | 0.309 | 0.428 | 0.567 | 0.921 | 2.291 | 16.0 | 66 | 0 |
| valkey-go | 1000 | 777199 | 1.325 | 1.747 | 2.293 | 3.395 | 7.015 | 7.1 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12880 | 0.073 | 0.090 | 0.139 | 0.403 | 1.613 | 53.5 | 37 | 0 |
| valkey-go | 10 | 46971 | 0.207 | 0.287 | 0.421 | 0.691 | 1.968 | 60.5 | 38 | 0 |
| valkey-go | 100 | 93862 | 1.103 | 1.398 | 1.833 | 2.671 | 3.575 | 28.7 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14197 | 0.067 | 0.079 | 0.116 | 0.367 | 1.619 | 53.1 | 28 | 0 |
| go-redis | 10 | 52493 | 0.189 | 0.247 | 0.336 | 0.597 | 1.683 | 61.4 | 33 | 0 |
| go-redis | 100 | 53017 | 1.808 | 2.053 | 3.597 | 4.083 | 6.135 | 68.8 | 68 | 0 |
| go-redis | 1000 | 52560 | 18.911 | 20.191 | 21.935 | 23.135 | 25.055 | 78.6 | 289 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11965 | 0.077 | 0.096 | 0.172 | 0.435 | 1.649 | 62.4 | 29 | 0 |
| go-redis | 10 | 37871 | 0.258 | 0.337 | 0.581 | 0.939 | 1.624 | 80.7 | 36 | 0 |
| go-redis | 100 | 38902 | 2.441 | 2.797 | 4.871 | 5.827 | 8.663 | 82.4 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9881 | 0.093 | 0.125 | 0.197 | 0.435 | 1.589 | 104.1 | 28 | 0 |
| go-redis-autopipeline | 10 | 59439 | 0.156 | 0.209 | 0.377 | 0.782 | 2.329 | 21.2 | 37 | 0 |
| go-redis-autopipeline | 100 | 197306 | 0.492 | 0.612 | 0.859 | 1.575 | 2.119 | 8.1 | 57 | 0 |
| go-redis-autopipeline | 1000 | 440322 | 2.253 | 2.975 | 3.693 | 5.555 | 7.231 | 6.7 | 309 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11916 | 0.078 | 0.096 | 0.170 | 0.430 | 1.583 | 63.5 | 28 | 0 |
| go-redis-autopipeline | 10 | 38206 | 0.256 | 0.331 | 0.579 | 0.966 | 2.179 | 80.9 | 37 | 0 |
| go-redis-autopipeline | 100 | 39013 | 2.443 | 2.801 | 4.875 | 5.555 | 7.407 | 84.2 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7756 | 0.117 | 0.154 | 0.272 | 0.500 | 1.878 | 118.0 | 34 | 0 |
| glide | 10 | 33648 | 0.283 | 0.385 | 0.539 | 0.792 | 2.014 | 55.7 | 37 | 0 |
| glide | 100 | 34540 | 2.921 | 3.583 | 4.479 | 5.007 | 6.159 | 57.2 | 66 | 0 |
| glide | 1000 | 34821 | 28.655 | 31.663 | 37.247 | 41.375 | 45.183 | 58.1 | 209 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6564 | 0.140 | 0.186 | 0.350 | 0.638 | 1.878 | 124.8 | 33 | 0 |
| glide | 10 | 25902 | 0.363 | 0.546 | 0.956 | 1.474 | 2.379 | 93.1 | 34 | 0 |
| glide | 100 | 29595 | 3.359 | 4.295 | 5.315 | 6.367 | 9.303 | 84.5 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15597 | 0.062 | 0.071 | 0.102 | 0.305 | 1.570 | 47.1 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 52028 | 0.188 | 0.246 | 0.384 | 0.578 | 1.757 | 51.7 | 51 | 0 |
| valkey-go/no-autopipe | 100 | 52226 | 1.826 | 2.127 | 3.645 | 4.247 | 5.795 | 58.9 | 176 | 0 |
| valkey-go/no-autopipe | 1000 | 48809 | 20.095 | 22.111 | 23.807 | 41.055 | 47.071 | 52.5 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 13003 | 0.072 | 0.086 | 0.142 | 0.394 | 1.524 | 53.4 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 38860 | 0.256 | 0.320 | 0.452 | 0.737 | 1.682 | 65.4 | 53 | 0 |
| valkey-go/no-autopipe | 100 | 38954 | 2.447 | 2.773 | 4.875 | 5.523 | 7.671 | 70.7 | 258 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12272 | 0.075 | 0.095 | 0.153 | 0.390 | 1.555 | 70.2 | 35 | 0 |
| valkey-go/always-pipe | 10 | 72167 | 0.134 | 0.189 | 0.255 | 0.477 | 1.910 | 44.3 | 40 | 0 |
| valkey-go/always-pipe | 100 | 321823 | 0.307 | 0.425 | 0.578 | 0.871 | 2.341 | 15.9 | 66 | 0 |
| valkey-go/always-pipe | 1000 | 875539 | 1.176 | 1.534 | 1.892 | 3.011 | 5.731 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9125 | 0.104 | 0.132 | 0.208 | 0.440 | 1.542 | 99.6 | 37 | 0 |
| valkey-go/always-pipe | 10 | 46526 | 0.209 | 0.290 | 0.423 | 0.705 | 1.758 | 61.9 | 40 | 0 |
| valkey-go/always-pipe | 100 | 101691 | 1.014 | 1.286 | 1.684 | 2.457 | 3.439 | 28.7 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15325 | 0.062 | 0.074 | 0.105 | 0.273 | 1.365 | 47.1 | 36 | 0 |
| valkey-go/flush-20us | 10 | 71968 | 0.134 | 0.190 | 0.260 | 0.482 | 1.752 | 45.4 | 43 | 0 |
| valkey-go/flush-20us | 100 | 343487 | 0.289 | 0.393 | 0.514 | 0.847 | 2.012 | 13.8 | 68 | 0 |
| valkey-go/flush-20us | 1000 | 904581 | 1.131 | 1.495 | 1.858 | 2.685 | 5.495 | 6.4 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12938 | 0.073 | 0.087 | 0.130 | 0.391 | 1.643 | 53.1 | 37 | 0 |
| valkey-go/flush-20us | 10 | 47076 | 0.206 | 0.287 | 0.421 | 0.677 | 1.643 | 61.5 | 39 | 0 |
| valkey-go/flush-20us | 100 | 103071 | 1.005 | 1.266 | 1.678 | 2.325 | 4.115 | 27.0 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15260 | 0.062 | 0.074 | 0.109 | 0.321 | 1.481 | 47.2 | 34 | 0 |
| valkey-go/flush-100us | 10 | 60584 | 0.158 | 0.242 | 0.316 | 1.298 | 1.933 | 36.2 | 39 | 0 |
| valkey-go/flush-100us | 100 | 407512 | 0.236 | 0.335 | 0.471 | 0.769 | 2.375 | 8.7 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 959975 | 1.065 | 1.419 | 1.752 | 2.611 | 5.751 | 5.3 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 13057 | 0.072 | 0.088 | 0.138 | 0.375 | 1.583 | 52.2 | 37 | 0 |
| valkey-go/flush-100us | 10 | 43344 | 0.221 | 0.334 | 0.467 | 0.717 | 1.871 | 55.4 | 41 | 0 |
| valkey-go/flush-100us | 100 | 103090 | 0.985 | 1.253 | 1.761 | 2.319 | 3.315 | 23.1 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15585 | 0.062 | 0.070 | 0.105 | 0.314 | 1.632 | 46.1 | 35 | 0 |
| valkey-go/no-retry | 10 | 72595 | 0.133 | 0.188 | 0.255 | 0.482 | 1.414 | 43.5 | 40 | 0 |
| valkey-go/no-retry | 100 | 324873 | 0.304 | 0.420 | 0.556 | 0.838 | 2.203 | 15.8 | 66 | 0 |
| valkey-go/no-retry | 1000 | 883631 | 1.165 | 1.518 | 1.885 | 3.249 | 6.963 | 6.9 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12969 | 0.073 | 0.087 | 0.131 | 0.369 | 1.445 | 52.1 | 37 | 0 |
| valkey-go/no-retry | 10 | 46360 | 0.210 | 0.291 | 0.414 | 0.695 | 1.712 | 59.7 | 40 | 0 |
| valkey-go/no-retry | 100 | 102882 | 1.002 | 1.265 | 1.715 | 2.331 | 3.391 | 28.5 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12669 | 0.089 | 0.114 | 0.174 | 0.413 | 1.556 | 74.3 | 74 | 0 |
| valkey-go/docache | 10 | 140609 | 0.002 | 0.209 | 0.293 | 0.461 | 3.857 | 24.4 | 151 | 0 |
| valkey-go/docache | 100 | 460075 | 0.010 | 0.676 | 0.955 | 1.466 | 5.027 | 13.1 | 180 | 0 |
| valkey-go/docache | 1000 | 516162 | 0.017 | 6.195 | 8.223 | 12.247 | 19.471 | 11.8 | 372 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 10022 | 0.044 | 0.202 | 0.296 | 0.505 | 3.231 | 165.6 | 136 | 0 |
| valkey-go/docache | 10 | 340958 | 0.021 | 0.033 | 0.191 | 0.360 | 4.635 | 34.7 | 197 | 0 |
| valkey-go/docache | 100 | 450686 | 0.145 | 0.459 | 1.281 | 3.995 | 28.399 | 34.8 | 240 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14956 | 0.064 | 0.075 | 0.107 | 0.295 | 2.301 | 45.6 | 35 | 0 |
| valkey-go/compat | 10 | 68259 | 0.141 | 0.202 | 0.282 | 0.494 | 2.051 | 44.5 | 40 | 0 |
| valkey-go/compat | 100 | 308590 | 0.321 | 0.442 | 0.583 | 1.014 | 2.607 | 16.4 | 68 | 0 |
| valkey-go/compat | 1000 | 874722 | 1.171 | 1.531 | 1.952 | 4.057 | 7.931 | 7.1 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12389 | 0.075 | 0.093 | 0.146 | 0.382 | 1.590 | 56.5 | 37 | 0 |
| valkey-go/compat | 10 | 45419 | 0.212 | 0.295 | 0.503 | 0.837 | 1.694 | 68.1 | 39 | 0 |
| valkey-go/compat | 100 | 100766 | 1.012 | 1.311 | 1.961 | 2.887 | 3.777 | 35.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15581 | 0.061 | 0.073 | 0.105 | 0.283 | 1.496 | 45.6 | 34 | 0 |
| valkey-go/hook | 10 | 71816 | 0.135 | 0.190 | 0.255 | 0.488 | 1.337 | 44.0 | 40 | 0 |
| valkey-go/hook | 100 | 319260 | 0.308 | 0.430 | 0.591 | 0.969 | 2.257 | 16.1 | 67 | 0 |
| valkey-go/hook | 1000 | 880760 | 1.163 | 1.531 | 1.940 | 3.079 | 5.691 | 7.1 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12919 | 0.073 | 0.087 | 0.131 | 0.387 | 1.628 | 53.3 | 37 | 0 |
| valkey-go/hook | 10 | 46922 | 0.207 | 0.287 | 0.420 | 0.675 | 1.759 | 59.8 | 39 | 0 |
| valkey-go/hook | 100 | 102540 | 1.006 | 1.276 | 1.688 | 2.413 | 4.071 | 28.4 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14694 | 0.064 | 0.076 | 0.113 | 0.340 | 1.687 | 50.0 | 35 | 0 |
| valkey-go/otel | 10 | 70148 | 0.137 | 0.194 | 0.269 | 0.511 | 1.378 | 47.4 | 40 | 0 |
| valkey-go/otel | 100 | 306358 | 0.319 | 0.450 | 0.677 | 1.417 | 3.057 | 18.5 | 66 | 0 |
| valkey-go/otel | 1000 | 858785 | 1.183 | 1.584 | 2.313 | 4.439 | 8.583 | 8.6 | 317 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12460 | 0.075 | 0.092 | 0.141 | 0.393 | 1.541 | 56.3 | 37 | 0 |
| valkey-go/otel | 10 | 44344 | 0.216 | 0.309 | 0.481 | 0.773 | 1.725 | 66.3 | 42 | 0 |
| valkey-go/otel | 100 | 101130 | 1.014 | 1.305 | 1.779 | 2.643 | 3.669 | 32.6 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10572 | 0.089 | 0.114 | 0.186 | 0.425 | 1.662 | 102.6 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 65420 | 0.143 | 0.209 | 0.362 | 0.730 | 2.441 | 61.2 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 288820 | 0.314 | 0.487 | 1.151 | 2.403 | 4.319 | 27.4 | 70 | 0 |
| valkey-go/otel-sdk | 1000 | 762992 | 1.188 | 1.780 | 4.871 | 8.351 | 14.551 | 17.5 | 327 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9412 | 0.100 | 0.125 | 0.207 | 0.453 | 1.921 | 110.4 | 38 | 0 |
| valkey-go/otel-sdk | 10 | 43860 | 0.216 | 0.311 | 0.559 | 0.876 | 1.820 | 77.2 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 99619 | 1.016 | 1.351 | 2.035 | 3.361 | 4.567 | 42.6 | 69 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11656 | 0.097 | 0.121 | 0.196 | 0.438 | 2.415 | 82.9 | 78 | 0 |
| valkey-go/aside | 10 | 134228 | 0.012 | 0.202 | 0.322 | 0.515 | 4.719 | 21.2 | 169 | 0 |
| valkey-go/aside | 100 | 187441 | 0.026 | 1.545 | 1.934 | 3.059 | 10.263 | 17.8 | 198 | 0 |
| valkey-go/aside | 1000 | 187646 | 0.027 | 15.079 | 16.479 | 24.415 | 30.687 | 17.9 | 397 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1324 | 0.747 | 1.041 | 1.376 | 2.581 | 3.723 | 744.7 | 82 | 0 |
| valkey-go/aside | 10 | 83360 | 0.085 | 0.239 | 0.565 | 1.210 | 7.139 | 60.7 | 209 | 0 |
| valkey-go/aside | 100 | 50339 | 0.421 | 7.859 | 13.807 | 19.791 | 34.879 | 96.9 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2330 | 0.409 | 0.503 | 0.785 | 1.475 | 1.949 | 519.1 | 27 | 0 |
| valkey-go/lock | 10 | 13937 | 0.686 | 0.879 | 1.464 | 2.147 | 2.773 | 184.8 | 33 | 0 |
| valkey-go/lock | 100 | 26219 | 3.753 | 4.411 | 5.903 | 7.503 | 8.527 | 99.4 | 62 | 0 |
| valkey-go/lock | 1000 | 14298 | 55.487 | 75.391 | 83.391 | 88.703 | 97.471 | 153.1 | 335 | 9753 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 233 | 4.203 | 4.727 | 5.923 | 7.775 | 9.279 | 5187.9 | 27 | 0 |
| valkey-go/lock | 10 | 1205 | 8.139 | 9.303 | 11.207 | 13.279 | 14.559 | 2085.3 | 33 | 0 |
| valkey-go/lock | 100 | 1703 | 58.495 | 61.087 | 63.359 | 64.991 | 66.495 | 1336.5 | 65 | 2 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15170 | 0.062 | 0.075 | 0.108 | 0.318 | 1.477 | 47.4 | 35 | 0 |
| valkey-go | 10 | 72401 | 0.134 | 0.189 | 0.255 | 0.461 | 1.587 | 43.4 | 40 | 0 |
| valkey-go | 100 | 322387 | 0.306 | 0.427 | 0.571 | 0.846 | 2.097 | 15.7 | 67 | 0 |
| valkey-go | 1000 | 796139 | 1.295 | 1.692 | 2.115 | 3.389 | 8.743 | 6.9 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12861 | 0.073 | 0.091 | 0.152 | 0.393 | 1.641 | 53.6 | 35 | 0 |
| valkey-go | 10 | 47097 | 0.206 | 0.286 | 0.427 | 0.727 | 1.795 | 60.6 | 40 | 0 |
| valkey-go | 100 | 94032 | 1.103 | 1.395 | 1.831 | 2.599 | 3.745 | 28.6 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14159 | 0.067 | 0.080 | 0.115 | 0.357 | 1.625 | 53.2 | 30 | 0 |
| go-redis | 10 | 52392 | 0.189 | 0.246 | 0.336 | 0.622 | 1.719 | 60.3 | 32 | 0 |
| go-redis | 100 | 52931 | 1.818 | 1.994 | 3.619 | 4.019 | 5.531 | 66.9 | 72 | 0 |
| go-redis | 1000 | 52771 | 18.783 | 20.111 | 22.303 | 24.623 | 27.775 | 78.6 | 239 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11818 | 0.078 | 0.099 | 0.170 | 0.426 | 2.921 | 64.5 | 29 | 0 |
| go-redis | 10 | 38279 | 0.255 | 0.330 | 0.574 | 0.935 | 1.939 | 79.8 | 34 | 0 |
| go-redis | 100 | 39641 | 2.409 | 2.733 | 4.799 | 5.471 | 7.583 | 84.0 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 9639 | 0.095 | 0.126 | 0.207 | 0.449 | 1.619 | 107.1 | 28 | 0 |
| go-redis-autopipeline | 10 | 61047 | 0.153 | 0.203 | 0.360 | 0.772 | 1.610 | 20.8 | 32 | 0 |
| go-redis-autopipeline | 100 | 203070 | 0.480 | 0.587 | 0.814 | 1.403 | 2.055 | 8.1 | 58 | 0 |
| go-redis-autopipeline | 1000 | 459103 | 2.163 | 2.791 | 3.659 | 5.099 | 6.639 | 6.5 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 11691 | 0.079 | 0.099 | 0.175 | 0.435 | 1.678 | 64.1 | 27 | 0 |
| go-redis-autopipeline | 10 | 37619 | 0.260 | 0.338 | 0.581 | 0.974 | 2.423 | 81.3 | 35 | 0 |
| go-redis-autopipeline | 100 | 39426 | 2.415 | 2.767 | 4.819 | 5.699 | 7.375 | 83.7 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7316 | 0.132 | 0.162 | 0.276 | 0.498 | 1.694 | 124.5 | 32 | 0 |
| glide | 10 | 31805 | 0.300 | 0.405 | 0.582 | 0.853 | 1.959 | 57.7 | 35 | 0 |
| glide | 100 | 34055 | 3.023 | 3.493 | 4.375 | 5.083 | 6.375 | 58.3 | 62 | 0 |
| glide | 1000 | 33824 | 29.359 | 32.447 | 36.671 | 39.359 | 40.479 | 59.1 | 235 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6602 | 0.138 | 0.186 | 0.351 | 0.612 | 1.758 | 124.2 | 35 | 0 |
| glide | 10 | 25542 | 0.370 | 0.549 | 0.961 | 1.514 | 2.023 | 93.3 | 37 | 0 |
| glide | 100 | 28116 | 3.543 | 4.551 | 5.663 | 6.583 | 8.647 | 87.1 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15151 | 0.063 | 0.074 | 0.103 | 0.310 | 1.549 | 49.0 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 52741 | 0.188 | 0.241 | 0.318 | 0.564 | 1.560 | 51.3 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 53524 | 1.800 | 1.955 | 3.589 | 3.937 | 5.767 | 55.5 | 177 | 0 |
| valkey-go/no-autopipe | 1000 | 47661 | 20.767 | 22.847 | 26.383 | 41.983 | 60.095 | 51.6 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12937 | 0.073 | 0.086 | 0.136 | 0.390 | 1.792 | 54.2 | 29 | 0 |
| valkey-go/no-autopipe | 10 | 38881 | 0.255 | 0.320 | 0.443 | 0.702 | 1.938 | 64.4 | 55 | 0 |
| valkey-go/no-autopipe | 100 | 39597 | 2.417 | 2.695 | 4.815 | 5.459 | 7.155 | 67.0 | 261 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12134 | 0.076 | 0.097 | 0.150 | 0.380 | 1.713 | 71.5 | 33 | 0 |
| valkey-go/always-pipe | 10 | 72622 | 0.133 | 0.188 | 0.253 | 0.472 | 1.739 | 43.8 | 40 | 0 |
| valkey-go/always-pipe | 100 | 283933 | 0.323 | 0.547 | 0.874 | 1.220 | 2.633 | 17.1 | 66 | 0 |
| valkey-go/always-pipe | 1000 | 866428 | 1.189 | 1.548 | 1.954 | 3.097 | 6.775 | 7.1 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9553 | 0.100 | 0.121 | 0.192 | 0.440 | 1.559 | 94.3 | 37 | 0 |
| valkey-go/always-pipe | 10 | 46525 | 0.209 | 0.289 | 0.430 | 0.710 | 1.589 | 62.0 | 41 | 0 |
| valkey-go/always-pipe | 100 | 101092 | 1.021 | 1.290 | 1.727 | 2.501 | 3.815 | 28.8 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15240 | 0.062 | 0.075 | 0.103 | 0.305 | 1.606 | 47.6 | 35 | 0 |
| valkey-go/flush-20us | 10 | 71887 | 0.134 | 0.190 | 0.259 | 0.488 | 1.667 | 45.0 | 41 | 0 |
| valkey-go/flush-20us | 100 | 342946 | 0.290 | 0.392 | 0.519 | 0.838 | 2.311 | 13.8 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 903671 | 1.131 | 1.491 | 1.886 | 2.909 | 6.351 | 6.4 | 311 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12805 | 0.073 | 0.088 | 0.135 | 0.395 | 1.623 | 54.0 | 37 | 0 |
| valkey-go/flush-20us | 10 | 47301 | 0.205 | 0.285 | 0.433 | 0.708 | 1.908 | 61.3 | 41 | 0 |
| valkey-go/flush-20us | 100 | 102890 | 1.003 | 1.266 | 1.690 | 2.493 | 3.789 | 26.9 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15522 | 0.062 | 0.072 | 0.101 | 0.326 | 1.522 | 46.0 | 35 | 0 |
| valkey-go/flush-100us | 10 | 60811 | 0.158 | 0.240 | 0.316 | 1.318 | 2.002 | 35.7 | 41 | 0 |
| valkey-go/flush-100us | 100 | 403253 | 0.238 | 0.340 | 0.491 | 0.796 | 2.040 | 8.9 | 68 | 0 |
| valkey-go/flush-100us | 1000 | 956962 | 1.066 | 1.435 | 1.796 | 2.829 | 5.575 | 5.5 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12892 | 0.073 | 0.089 | 0.132 | 0.367 | 1.484 | 53.5 | 37 | 0 |
| valkey-go/flush-100us | 10 | 43514 | 0.221 | 0.330 | 0.463 | 0.733 | 1.909 | 56.1 | 41 | 0 |
| valkey-go/flush-100us | 100 | 104177 | 0.978 | 1.235 | 1.740 | 2.303 | 3.385 | 23.3 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15712 | 0.061 | 0.070 | 0.101 | 0.279 | 1.565 | 45.1 | 26 | 0 |
| valkey-go/1-conn | 10 | 81137 | 0.122 | 0.159 | 0.216 | 0.426 | 1.669 | 23.2 | 29 | 0 |
| valkey-go/1-conn | 100 | 419194 | 0.231 | 0.335 | 0.439 | 0.748 | 2.291 | 7.3 | 57 | 0 |
| valkey-go/1-conn | 1000 | 1004730 | 0.981 | 1.092 | 1.407 | 2.711 | 4.903 | 5.5 | 304 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12981 | 0.072 | 0.087 | 0.141 | 0.397 | 1.648 | 54.6 | 28 | 0 |
| valkey-go/1-conn | 10 | 48732 | 0.203 | 0.271 | 0.450 | 0.680 | 2.043 | 40.4 | 30 | 0 |
| valkey-go/1-conn | 100 | 117654 | 0.819 | 0.957 | 1.412 | 2.055 | 2.837 | 18.7 | 58 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 12717 | 0.089 | 0.114 | 0.169 | 0.419 | 1.726 | 74.4 | 68 | 0 |
| valkey-go/docache | 10 | 142181 | 0.002 | 0.207 | 0.291 | 0.463 | 3.775 | 24.4 | 153 | 0 |
| valkey-go/docache | 100 | 472416 | 0.010 | 0.657 | 0.919 | 1.350 | 5.251 | 13.1 | 184 | 0 |
| valkey-go/docache | 1000 | 520112 | 0.017 | 6.183 | 7.963 | 10.463 | 19.759 | 12.0 | 417 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/docache | 1 | 10594 | 0.036 | 0.197 | 0.278 | 0.494 | 2.125 | 158.6 | 142 | 0 |
| valkey-go/docache | 10 | 344389 | 0.021 | 0.033 | 0.194 | 0.353 | 3.431 | 34.2 | 195 | 0 |
| valkey-go/docache | 100 | 563966 | 0.108 | 0.361 | 1.144 | 3.405 | 23.631 | 28.5 | 258 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 14991 | 0.063 | 0.075 | 0.106 | 0.313 | 1.594 | 46.3 | 35 | 0 |
| valkey-go/compat | 10 | 69461 | 0.139 | 0.196 | 0.265 | 0.476 | 1.657 | 44.1 | 39 | 0 |
| valkey-go/compat | 100 | 310748 | 0.319 | 0.439 | 0.578 | 0.956 | 2.491 | 16.3 | 65 | 0 |
| valkey-go/compat | 1000 | 885498 | 1.159 | 1.514 | 1.866 | 4.159 | 8.839 | 7.1 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/compat | 1 | 12528 | 0.075 | 0.090 | 0.144 | 0.402 | 1.679 | 55.9 | 35 | 0 |
| valkey-go/compat | 10 | 45320 | 0.212 | 0.298 | 0.497 | 0.827 | 1.682 | 68.0 | 41 | 0 |
| valkey-go/compat | 100 | 100702 | 1.014 | 1.312 | 1.886 | 2.793 | 5.403 | 34.2 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 15487 | 0.062 | 0.074 | 0.105 | 0.315 | 1.655 | 46.4 | 34 | 0 |
| valkey-go/hook | 10 | 60494 | 0.145 | 0.264 | 0.436 | 0.645 | 4.013 | 54.6 | 41 | 0 |
| valkey-go/hook | 100 | 321721 | 0.306 | 0.424 | 0.568 | 0.917 | 2.767 | 15.8 | 67 | 0 |
| valkey-go/hook | 1000 | 878405 | 1.166 | 1.526 | 2.020 | 3.119 | 6.407 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/hook | 1 | 12673 | 0.074 | 0.089 | 0.138 | 0.386 | 1.593 | 54.2 | 35 | 0 |
| valkey-go/hook | 10 | 46870 | 0.207 | 0.288 | 0.422 | 0.717 | 1.688 | 60.6 | 39 | 0 |
| valkey-go/hook | 100 | 102782 | 1.008 | 1.266 | 1.687 | 2.517 | 3.871 | 28.2 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 14996 | 0.064 | 0.076 | 0.110 | 0.327 | 1.621 | 50.4 | 37 | 0 |
| valkey-go/otel | 10 | 70092 | 0.137 | 0.195 | 0.268 | 0.527 | 1.803 | 48.0 | 41 | 0 |
| valkey-go/otel | 100 | 310029 | 0.315 | 0.443 | 0.668 | 1.450 | 2.699 | 18.3 | 65 | 0 |
| valkey-go/otel | 1000 | 869745 | 1.165 | 1.565 | 2.235 | 4.311 | 7.531 | 8.5 | 316 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel | 1 | 12449 | 0.075 | 0.091 | 0.134 | 0.397 | 1.715 | 56.2 | 35 | 0 |
| valkey-go/otel | 10 | 46058 | 0.210 | 0.292 | 0.451 | 0.765 | 1.599 | 66.4 | 40 | 0 |
| valkey-go/otel | 100 | 102858 | 0.999 | 1.270 | 1.727 | 2.599 | 4.155 | 31.1 | 67 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 10592 | 0.090 | 0.111 | 0.166 | 0.408 | 1.548 | 102.6 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 65525 | 0.143 | 0.209 | 0.354 | 0.672 | 1.828 | 60.9 | 41 | 0 |
| valkey-go/otel-sdk | 100 | 293992 | 0.309 | 0.475 | 1.129 | 2.281 | 3.741 | 27.2 | 68 | 0 |
| valkey-go/otel-sdk | 1000 | 791878 | 1.153 | 1.671 | 4.855 | 8.011 | 19.935 | 17.3 | 326 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/otel-sdk | 1 | 9500 | 0.100 | 0.122 | 0.196 | 0.446 | 1.586 | 108.0 | 37 | 0 |
| valkey-go/otel-sdk | 10 | 44430 | 0.213 | 0.305 | 0.562 | 0.936 | 1.728 | 76.8 | 42 | 0 |
| valkey-go/otel-sdk | 100 | 98472 | 1.029 | 1.365 | 2.075 | 3.507 | 5.011 | 43.0 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 11424 | 0.098 | 0.125 | 0.213 | 0.445 | 1.720 | 84.4 | 76 | 0 |
| valkey-go/aside | 10 | 135885 | 0.012 | 0.199 | 0.319 | 0.517 | 3.941 | 21.3 | 171 | 0 |
| valkey-go/aside | 100 | 193166 | 0.025 | 1.496 | 1.861 | 2.847 | 10.007 | 17.3 | 196 | 0 |
| valkey-go/aside | 1000 | 193084 | 0.026 | 14.663 | 15.903 | 22.687 | 25.231 | 17.5 | 393 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/aside | 1 | 1341 | 0.736 | 1.031 | 1.365 | 2.523 | 2.993 | 742.2 | 82 | 0 |
| valkey-go/aside | 10 | 84758 | 0.083 | 0.236 | 0.554 | 1.078 | 8.671 | 60.6 | 209 | 0 |
| valkey-go/aside | 100 | 51492 | 0.406 | 7.803 | 13.543 | 19.839 | 34.559 | 95.8 | 232 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 2392 | 0.402 | 0.480 | 0.714 | 1.195 | 1.957 | 517.5 | 30 | 0 |
| valkey-go/lock | 10 | 14312 | 0.670 | 0.840 | 1.428 | 2.203 | 3.129 | 187.1 | 35 | 0 |
| valkey-go/lock | 100 | 26823 | 3.709 | 4.179 | 5.335 | 6.499 | 7.695 | 101.2 | 62 | 3 |
| valkey-go/lock | 1000 | 14623 | 55.103 | 75.711 | 83.775 | 88.575 | 94.591 | 157.5 | 325 | 9261 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/lock | 1 | 235 | 4.171 | 4.691 | 6.011 | 8.043 | 8.367 | 5189.0 | 28 | 0 |
| valkey-go/lock | 10 | 1216 | 8.091 | 9.151 | 10.599 | 12.263 | 12.855 | 2094.2 | 33 | 0 |
| valkey-go/lock | 100 | 1726 | 57.343 | 59.903 | 68.287 | 73.087 | 75.519 | 1356.6 | 62 | 1 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 15143 | 0.063 | 0.076 | 0.107 | 0.302 | 1.660 | 45.7 | 36 | 0 |
| valkey-go | 10 | 71930 | 0.134 | 0.190 | 0.254 | 0.483 | 1.690 | 44.0 | 40 | 0 |
| valkey-go | 100 | 321148 | 0.308 | 0.425 | 0.563 | 0.952 | 1.895 | 16.0 | 67 | 0 |
| valkey-go | 1000 | 811234 | 1.271 | 1.659 | 2.081 | 3.313 | 5.883 | 6.8 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 13034 | 0.072 | 0.087 | 0.128 | 0.366 | 1.571 | 53.0 | 35 | 0 |
| valkey-go | 10 | 46511 | 0.209 | 0.291 | 0.431 | 0.691 | 1.350 | 61.6 | 39 | 0 |
| valkey-go | 100 | 84291 | 1.133 | 1.787 | 2.873 | 3.931 | 5.919 | 31.2 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 14264 | 0.067 | 0.076 | 0.112 | 0.358 | 1.570 | 52.6 | 28 | 0 |
| go-redis | 10 | 52585 | 0.189 | 0.245 | 0.337 | 0.591 | 1.813 | 59.6 | 33 | 0 |
| go-redis | 100 | 53179 | 1.806 | 2.026 | 3.589 | 4.055 | 5.847 | 66.8 | 70 | 0 |
| go-redis | 1000 | 52798 | 18.799 | 20.303 | 22.639 | 24.671 | 26.815 | 78.5 | 243 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11964 | 0.078 | 0.096 | 0.163 | 0.433 | 1.746 | 63.3 | 27 | 0 |
| go-redis | 10 | 38202 | 0.255 | 0.332 | 0.587 | 0.957 | 2.153 | 80.2 | 36 | 0 |
| go-redis | 100 | 38943 | 2.413 | 2.985 | 4.823 | 5.907 | 7.895 | 83.7 | 74 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 10101 | 0.092 | 0.121 | 0.190 | 0.433 | 1.629 | 101.6 | 30 | 0 |
| go-redis-autopipeline | 10 | 59820 | 0.155 | 0.209 | 0.362 | 0.794 | 1.751 | 21.1 | 35 | 0 |
| go-redis-autopipeline | 100 | 204946 | 0.477 | 0.580 | 0.794 | 1.472 | 2.067 | 8.0 | 59 | 0 |
| go-redis-autopipeline | 1000 | 448309 | 2.289 | 2.911 | 3.667 | 5.287 | 6.587 | 6.6 | 308 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 12002 | 0.077 | 0.096 | 0.158 | 0.433 | 1.537 | 63.2 | 29 | 0 |
| go-redis-autopipeline | 10 | 38318 | 0.255 | 0.330 | 0.579 | 0.963 | 1.905 | 79.4 | 37 | 0 |
| go-redis-autopipeline | 100 | 39413 | 2.413 | 2.779 | 4.811 | 5.627 | 9.167 | 83.7 | 75 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 7624 | 0.122 | 0.159 | 0.262 | 0.490 | 1.810 | 120.4 | 33 | 0 |
| glide | 10 | 34179 | 0.282 | 0.364 | 0.521 | 0.813 | 3.097 | 55.1 | 36 | 0 |
| glide | 100 | 35988 | 2.845 | 3.241 | 4.399 | 4.991 | 6.015 | 55.9 | 65 | 0 |
| glide | 1000 | 34571 | 28.623 | 32.159 | 39.167 | 42.879 | 43.807 | 58.3 | 233 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6441 | 0.143 | 0.193 | 0.348 | 0.645 | 2.657 | 126.5 | 31 | 0 |
| glide | 10 | 25171 | 0.375 | 0.556 | 0.967 | 1.571 | 2.405 | 94.4 | 35 | 0 |
| glide | 100 | 28755 | 3.469 | 4.451 | 5.503 | 6.447 | 7.943 | 85.6 | 64 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 15108 | 0.063 | 0.075 | 0.113 | 0.337 | 3.043 | 47.9 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 53329 | 0.187 | 0.237 | 0.309 | 0.541 | 1.892 | 50.8 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 53223 | 1.810 | 1.982 | 3.603 | 4.003 | 5.607 | 56.5 | 168 | 0 |
| valkey-go/no-autopipe | 1000 | 48349 | 20.223 | 22.495 | 25.327 | 41.215 | 52.447 | 52.3 | 1207 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-autopipe | 1 | 12932 | 0.073 | 0.087 | 0.142 | 0.407 | 1.630 | 54.1 | 30 | 0 |
| valkey-go/no-autopipe | 10 | 38732 | 0.256 | 0.323 | 0.448 | 0.708 | 1.897 | 64.5 | 54 | 0 |
| valkey-go/no-autopipe | 100 | 39950 | 2.401 | 2.649 | 4.779 | 5.287 | 7.075 | 68.8 | 260 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 12319 | 0.075 | 0.095 | 0.148 | 0.368 | 1.505 | 69.6 | 37 | 0 |
| valkey-go/always-pipe | 10 | 70645 | 0.136 | 0.195 | 0.269 | 0.500 | 1.908 | 44.0 | 41 | 0 |
| valkey-go/always-pipe | 100 | 313038 | 0.316 | 0.442 | 0.592 | 0.890 | 2.159 | 16.2 | 66 | 0 |
| valkey-go/always-pipe | 1000 | 848030 | 1.206 | 1.598 | 2.071 | 3.437 | 8.479 | 7.1 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/always-pipe | 1 | 9314 | 0.103 | 0.127 | 0.199 | 0.440 | 1.568 | 96.5 | 37 | 0 |
| valkey-go/always-pipe | 10 | 46771 | 0.208 | 0.288 | 0.421 | 0.695 | 1.905 | 60.8 | 39 | 0 |
| valkey-go/always-pipe | 100 | 99884 | 1.028 | 1.310 | 1.799 | 2.771 | 4.061 | 28.9 | 66 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 15458 | 0.062 | 0.073 | 0.107 | 0.272 | 1.722 | 46.4 | 35 | 0 |
| valkey-go/flush-20us | 10 | 71446 | 0.135 | 0.191 | 0.257 | 0.489 | 1.468 | 45.4 | 43 | 0 |
| valkey-go/flush-20us | 100 | 342454 | 0.290 | 0.393 | 0.526 | 0.816 | 2.213 | 13.8 | 67 | 0 |
| valkey-go/flush-20us | 1000 | 916116 | 1.118 | 1.475 | 1.852 | 2.979 | 6.239 | 6.5 | 313 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-20us | 1 | 12622 | 0.074 | 0.091 | 0.142 | 0.408 | 1.585 | 53.2 | 35 | 0 |
| valkey-go/flush-20us | 10 | 46626 | 0.208 | 0.289 | 0.427 | 0.697 | 1.953 | 62.7 | 38 | 0 |
| valkey-go/flush-20us | 100 | 101051 | 1.023 | 1.292 | 1.755 | 2.369 | 3.649 | 27.5 | 68 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 15491 | 0.062 | 0.072 | 0.102 | 0.344 | 1.544 | 46.2 | 37 | 0 |
| valkey-go/flush-100us | 10 | 60332 | 0.158 | 0.242 | 0.322 | 1.337 | 2.000 | 36.5 | 43 | 0 |
| valkey-go/flush-100us | 100 | 403880 | 0.238 | 0.340 | 0.481 | 0.750 | 2.347 | 9.1 | 67 | 0 |
| valkey-go/flush-100us | 1000 | 959425 | 1.062 | 1.428 | 1.773 | 2.693 | 5.523 | 5.5 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/flush-100us | 1 | 12801 | 0.073 | 0.089 | 0.130 | 0.393 | 1.460 | 53.9 | 35 | 0 |
| valkey-go/flush-100us | 10 | 44048 | 0.218 | 0.326 | 0.448 | 0.710 | 1.976 | 55.4 | 39 | 0 |
| valkey-go/flush-100us | 100 | 102980 | 0.984 | 1.260 | 1.813 | 2.415 | 3.591 | 23.1 | 65 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 15670 | 0.062 | 0.071 | 0.101 | 0.287 | 1.701 | 45.3 | 28 | 0 |
| valkey-go/1-conn | 10 | 81197 | 0.121 | 0.160 | 0.218 | 0.456 | 1.384 | 23.2 | 32 | 0 |
| valkey-go/1-conn | 100 | 414968 | 0.233 | 0.338 | 0.445 | 0.757 | 1.863 | 7.2 | 58 | 0 |
| valkey-go/1-conn | 1000 | 1019360 | 0.961 | 1.085 | 1.519 | 2.761 | 4.663 | 5.3 | 306 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/1-conn | 1 | 12437 | 0.074 | 0.096 | 0.151 | 0.405 | 1.684 | 57.1 | 27 | 0 |
| valkey-go/1-conn | 10 | 48443 | 0.204 | 0.273 | 0.456 | 0.694 | 1.777 | 41.0 | 31 | 0 |
| valkey-go/1-conn | 100 | 118260 | 0.818 | 0.938 | 1.417 | 2.183 | 2.939 | 18.7 | 58 | 0 |
mode=standalone fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 15307 | 0.062 | 0.073 | 0.101 | 0.296 | 1.504 | 46.6 | 37 | 0 |
| valkey-go/no-retry | 10 | 72108 | 0.134 | 0.190 | 0.260 | 0.469 | 1.389 | 43.8 | 41 | 0 |
| valkey-go/no-retry | 100 | 321470 | 0.307 | 0.425 | 0.570 | 0.882 | 2.271 | 15.9 | 67 | 0 |
| valkey-go/no-retry | 1000 | 888162 | 1.155 | 1.512 | 1.924 | 2.955 | 6.015 | 6.9 | 315 | 0 |
mode=standalone fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go/no-retry | 1 | 12971 | 0.073 | 0.086 | 0.132 | 0.383 | 1.558 | 52.6 | 37 | 0 |
| valkey-go/no-retry | 10 | 46800 | 0.208 | 0.287 | 0.423 | 0.675 | 1.709 | 60.9 | 41 | 0 |
| valkey-go/no-retry | 100 | 104237 | 0.986 | 1.252 | 1.726 | 2.483 | 4.077 | 28.8 | 65 | 0 |
