mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10734 | 0.088 | 0.112 | 0.195 | 0.363 | 1.635 | 61.7 | 29 | 0 |
| go-redis | 10 | 87220 | 0.105 | 0.157 | 0.246 | 0.591 | 1.821 | 69.7 | 36 | 0 |
| go-redis | 100 | 149379 | 0.570 | 1.274 | 2.011 | 3.099 | 6.295 | 64.9 | 91 | 0 |
| go-redis | 1000 | 154609 | 1.697 | 17.551 | 21.327 | 25.231 | 27.423 | 64.3 | 337 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 5953 | 0.159 | 0.202 | 0.337 | 0.672 | 2.965 | 299.4 | 28 | 0 |
| go-redis | 10 | 37229 | 0.253 | 0.363 | 0.707 | 1.496 | 3.159 | 250.7 | 50 | 0 |
| go-redis | 100 | 46068 | 2.013 | 2.811 | 4.111 | 6.083 | 8.919 | 203.2 | 109 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 7040 | 0.135 | 0.169 | 0.273 | 0.513 | 1.687 | 133.0 | 33 | 0 |
| go-redis-autopipeline | 10 | 50513 | 0.143 | 0.361 | 0.527 | 0.951 | 2.593 | 84.6 | 45 | 0 |
| go-redis-autopipeline | 100 | 212145 | 0.408 | 0.801 | 1.442 | 2.267 | 4.707 | 35.4 | 76 | 0 |
| go-redis-autopipeline | 1000 | 462649 | 2.091 | 3.323 | 4.747 | 6.691 | 14.583 | 16.6 | 338 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6148 | 0.155 | 0.195 | 0.315 | 0.687 | 1.642 | 293.3 | 30 | 0 |
| go-redis-autopipeline | 10 | 36390 | 0.258 | 0.373 | 0.712 | 1.639 | 4.311 | 256.3 | 49 | 0 |
| go-redis-autopipeline | 100 | 46029 | 2.019 | 2.809 | 4.151 | 6.179 | 8.055 | 203.3 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6122 | 0.155 | 0.195 | 0.320 | 0.552 | 3.021 | 130.3 | 33 | 0 |
| glide | 10 | 29345 | 0.320 | 0.447 | 0.681 | 1.740 | 6.455 | 60.1 | 37 | 0 |
| glide | 100 | 39471 | 2.617 | 3.279 | 4.323 | 5.551 | 7.307 | 51.6 | 67 | 0 |
| glide | 1000 | 40260 | 24.527 | 28.527 | 33.375 | 38.815 | 41.759 | 52.7 | 221 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4724 | 0.192 | 0.273 | 0.432 | 1.325 | 3.145 | 198.7 | 35 | 0 |
| glide | 10 | 11806 | 0.827 | 1.065 | 1.524 | 2.505 | 5.471 | 140.3 | 37 | 0 |
| glide | 100 | 12441 | 8.303 | 11.199 | 13.711 | 16.127 | 18.159 | 135.4 | 71 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12370 | 0.078 | 0.097 | 0.170 | 0.310 | 1.658 | 50.4 | 33 | 0 |
| valkey-go | 10 | 87371 | 0.108 | 0.156 | 0.225 | 0.477 | 2.135 | 42.3 | 40 | 0 |
| valkey-go | 100 | 571016 | 0.158 | 0.258 | 0.438 | 1.045 | 4.103 | 11.5 | 64 | 0 |
| valkey-go | 1000 | 1860267 | 0.440 | 0.950 | 1.828 | 3.183 | 5.151 | 6.3 | 311 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 6651 | 0.141 | 0.185 | 0.293 | 1.220 | 2.263 | 250.3 | 34 | 0 |
| valkey-go | 10 | 42701 | 0.215 | 0.317 | 0.595 | 1.227 | 2.989 | 121.5 | 37 | 0 |
| valkey-go | 100 | 200385 | 0.446 | 0.708 | 1.736 | 2.727 | 5.015 | 48.1 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6520 | 0.143 | 0.189 | 0.321 | 0.555 | 1.492 | 140.7 | 33 | 0 |
| go-redis-autopipeline | 10 | 49429 | 0.146 | 0.366 | 0.541 | 1.003 | 2.059 | 86.4 | 45 | 0 |
| go-redis-autopipeline | 100 | 215161 | 0.402 | 0.793 | 1.421 | 2.207 | 4.219 | 35.0 | 73 | 0 |
| go-redis-autopipeline | 1000 | 458922 | 2.111 | 3.347 | 4.755 | 6.395 | 11.231 | 16.8 | 332 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 5810 | 0.162 | 0.209 | 0.350 | 0.666 | 1.766 | 306.8 | 30 | 0 |
| go-redis-autopipeline | 10 | 32954 | 0.272 | 0.432 | 0.928 | 2.343 | 6.399 | 267.3 | 51 | 0 |
| go-redis-autopipeline | 100 | 43794 | 2.083 | 2.991 | 4.463 | 6.331 | 10.823 | 220.6 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6218 | 0.149 | 0.191 | 0.328 | 1.021 | 3.025 | 130.5 | 33 | 0 |
| glide | 10 | 27665 | 0.339 | 0.495 | 0.724 | 1.508 | 2.365 | 63.1 | 37 | 0 |
| glide | 100 | 38110 | 2.683 | 3.475 | 4.395 | 5.431 | 7.047 | 53.1 | 66 | 0 |
| glide | 1000 | 39258 | 25.391 | 28.815 | 32.031 | 34.815 | 37.087 | 54.3 | 245 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4862 | 0.187 | 0.261 | 0.417 | 1.109 | 1.963 | 194.9 | 33 | 0 |
| glide | 10 | 11914 | 0.829 | 1.033 | 1.450 | 2.169 | 3.009 | 139.0 | 37 | 0 |
| glide | 100 | 12387 | 8.247 | 11.223 | 13.911 | 16.255 | 18.863 | 135.8 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12398 | 0.077 | 0.097 | 0.170 | 0.318 | 1.509 | 50.1 | 33 | 0 |
| valkey-go | 10 | 88540 | 0.107 | 0.153 | 0.218 | 0.458 | 2.079 | 42.1 | 39 | 0 |
| valkey-go | 100 | 570426 | 0.158 | 0.261 | 0.433 | 0.885 | 2.381 | 11.4 | 64 | 0 |
| valkey-go | 1000 | 1928327 | 0.450 | 0.819 | 1.620 | 2.905 | 5.967 | 6.4 | 314 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 6793 | 0.140 | 0.177 | 0.292 | 0.477 | 1.821 | 245.2 | 34 | 0 |
| valkey-go | 10 | 42950 | 0.214 | 0.315 | 0.567 | 1.209 | 2.377 | 120.6 | 40 | 0 |
| valkey-go | 100 | 203639 | 0.442 | 0.694 | 1.706 | 2.573 | 3.665 | 47.2 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10649 | 0.089 | 0.112 | 0.190 | 0.371 | 1.836 | 63.4 | 30 | 0 |
| go-redis | 10 | 86778 | 0.105 | 0.158 | 0.252 | 0.593 | 3.259 | 69.5 | 37 | 0 |
| go-redis | 100 | 152718 | 0.562 | 1.241 | 1.864 | 2.827 | 7.055 | 64.2 | 91 | 0 |
| go-redis | 1000 | 155950 | 3.569 | 16.495 | 18.991 | 20.351 | 23.487 | 64.3 | 342 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6109 | 0.156 | 0.193 | 0.321 | 0.597 | 1.825 | 294.8 | 28 | 0 |
| go-redis | 10 | 36821 | 0.256 | 0.369 | 0.696 | 1.455 | 3.123 | 250.4 | 49 | 0 |
| go-redis | 100 | 46831 | 1.988 | 2.735 | 4.067 | 6.303 | 9.983 | 201.8 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6322 | 0.149 | 0.186 | 0.313 | 0.515 | 2.065 | 128.7 | 33 | 0 |
| glide | 10 | 29108 | 0.322 | 0.474 | 0.681 | 1.014 | 2.019 | 62.1 | 37 | 0 |
| glide | 100 | 39859 | 2.597 | 3.267 | 4.103 | 5.563 | 7.171 | 50.8 | 67 | 0 |
| glide | 1000 | 40760 | 24.463 | 27.615 | 33.023 | 36.479 | 38.559 | 53.1 | 239 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4874 | 0.188 | 0.258 | 0.400 | 0.729 | 2.036 | 194.5 | 34 | 0 |
| glide | 10 | 11857 | 0.833 | 1.047 | 1.438 | 2.141 | 3.523 | 140.6 | 36 | 0 |
| glide | 100 | 12576 | 8.147 | 11.031 | 13.495 | 15.983 | 18.767 | 133.8 | 72 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12604 | 0.077 | 0.094 | 0.174 | 0.328 | 1.584 | 48.3 | 31 | 0 |
| valkey-go | 10 | 88974 | 0.107 | 0.153 | 0.218 | 0.374 | 1.946 | 42.1 | 38 | 0 |
| valkey-go | 100 | 577739 | 0.156 | 0.258 | 0.433 | 0.893 | 3.355 | 11.4 | 63 | 0 |
| valkey-go | 1000 | 2007830 | 0.440 | 0.760 | 1.540 | 2.789 | 5.235 | 6.4 | 313 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 6870 | 0.139 | 0.174 | 0.284 | 0.520 | 1.920 | 245.4 | 35 | 0 |
| valkey-go | 10 | 42576 | 0.215 | 0.318 | 0.594 | 1.346 | 2.955 | 121.5 | 37 | 0 |
| valkey-go | 100 | 204951 | 0.436 | 0.687 | 1.738 | 2.619 | 4.271 | 47.6 | 69 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10858 | 0.087 | 0.110 | 0.188 | 0.347 | 1.741 | 62.6 | 28 | 0 |
| go-redis | 10 | 87596 | 0.104 | 0.156 | 0.243 | 0.582 | 2.135 | 69.6 | 35 | 0 |
| go-redis | 100 | 151334 | 0.559 | 1.256 | 1.949 | 2.931 | 7.079 | 64.5 | 89 | 0 |
| go-redis | 1000 | 154439 | 1.373 | 18.175 | 20.287 | 22.559 | 24.383 | 64.2 | 343 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 5741 | 0.159 | 0.231 | 0.366 | 0.703 | 2.141 | 325.7 | 30 | 0 |
| go-redis | 10 | 37497 | 0.251 | 0.361 | 0.716 | 1.434 | 3.185 | 250.6 | 50 | 0 |
| go-redis | 100 | 46198 | 2.007 | 2.807 | 4.175 | 6.143 | 8.807 | 203.3 | 109 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6945 | 0.137 | 0.170 | 0.274 | 0.502 | 4.037 | 132.9 | 32 | 0 |
| go-redis-autopipeline | 10 | 50324 | 0.143 | 0.362 | 0.530 | 0.938 | 2.046 | 84.5 | 45 | 0 |
| go-redis-autopipeline | 100 | 223673 | 0.378 | 0.778 | 1.404 | 2.179 | 4.447 | 33.3 | 75 | 0 |
| go-redis-autopipeline | 1000 | 726747 | 1.331 | 2.008 | 3.065 | 5.199 | 10.031 | 12.6 | 338 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 5970 | 0.159 | 0.201 | 0.335 | 0.648 | 1.705 | 301.4 | 31 | 0 |
| go-redis-autopipeline | 10 | 36860 | 0.255 | 0.367 | 0.730 | 1.665 | 3.511 | 255.8 | 50 | 0 |
| go-redis-autopipeline | 100 | 46653 | 1.994 | 2.771 | 4.087 | 6.151 | 8.831 | 202.6 | 112 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12268 | 0.078 | 0.098 | 0.160 | 0.329 | 1.624 | 50.9 | 33 | 0 |
| valkey-go | 10 | 85402 | 0.109 | 0.159 | 0.258 | 0.826 | 2.729 | 42.8 | 37 | 0 |
| valkey-go | 100 | 570733 | 0.157 | 0.262 | 0.437 | 0.877 | 2.287 | 11.4 | 63 | 0 |
| valkey-go | 1000 | 1978089 | 0.431 | 0.788 | 1.662 | 2.959 | 5.331 | 6.3 | 316 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7045 | 0.135 | 0.170 | 0.278 | 0.482 | 1.776 | 238.5 | 34 | 0 |
| valkey-go | 10 | 43300 | 0.212 | 0.313 | 0.575 | 1.212 | 2.515 | 120.6 | 38 | 0 |
| valkey-go | 100 | 205537 | 0.433 | 0.688 | 1.758 | 2.587 | 4.023 | 47.8 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10824 | 0.087 | 0.110 | 0.188 | 0.341 | 1.636 | 62.3 | 30 | 0 |
| go-redis | 10 | 87538 | 0.105 | 0.156 | 0.242 | 0.572 | 2.069 | 69.6 | 36 | 0 |
| go-redis | 100 | 150179 | 0.562 | 1.284 | 2.024 | 2.963 | 6.323 | 65.1 | 91 | 0 |
| go-redis | 1000 | 154818 | 3.037 | 16.719 | 21.567 | 26.511 | 29.167 | 65.2 | 344 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 5952 | 0.159 | 0.202 | 0.337 | 0.635 | 2.231 | 298.7 | 28 | 0 |
| go-redis | 10 | 37183 | 0.252 | 0.364 | 0.747 | 1.520 | 3.465 | 250.6 | 49 | 0 |
| go-redis | 100 | 46642 | 1.991 | 2.771 | 4.131 | 6.115 | 8.879 | 203.0 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 7001 | 0.135 | 0.171 | 0.277 | 0.525 | 1.925 | 131.4 | 30 | 0 |
| go-redis-autopipeline | 10 | 50114 | 0.144 | 0.362 | 0.534 | 0.985 | 2.093 | 85.3 | 41 | 0 |
| go-redis-autopipeline | 100 | 217037 | 0.399 | 0.784 | 1.414 | 2.129 | 3.445 | 35.3 | 73 | 0 |
| go-redis-autopipeline | 1000 | 457943 | 2.117 | 3.353 | 4.743 | 6.475 | 11.111 | 16.8 | 327 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6032 | 0.158 | 0.199 | 0.324 | 0.664 | 1.633 | 298.7 | 32 | 0 |
| go-redis-autopipeline | 10 | 36702 | 0.256 | 0.369 | 0.710 | 1.537 | 2.849 | 257.0 | 51 | 0 |
| go-redis-autopipeline | 100 | 45887 | 2.014 | 2.833 | 4.219 | 6.475 | 9.183 | 207.4 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6209 | 0.152 | 0.192 | 0.324 | 0.562 | 1.581 | 131.4 | 33 | 0 |
| glide | 10 | 28807 | 0.328 | 0.466 | 0.681 | 1.010 | 2.253 | 62.9 | 37 | 0 |
| glide | 100 | 39248 | 2.627 | 3.321 | 4.287 | 5.575 | 7.595 | 51.7 | 66 | 0 |
| glide | 1000 | 39436 | 25.039 | 29.599 | 34.431 | 37.663 | 39.007 | 53.5 | 245 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4697 | 0.195 | 0.273 | 0.440 | 0.794 | 2.103 | 200.1 | 32 | 0 |
| glide | 10 | 11811 | 0.834 | 1.048 | 1.453 | 2.107 | 3.493 | 141.8 | 36 | 0 |
| glide | 100 | 12109 | 8.495 | 11.455 | 14.143 | 16.895 | 19.599 | 140.0 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10459 | 0.090 | 0.116 | 0.214 | 0.375 | 1.561 | 64.2 | 27 | 0 |
| go-redis | 10 | 85273 | 0.107 | 0.161 | 0.257 | 0.575 | 2.101 | 71.3 | 35 | 0 |
| go-redis | 100 | 115130 | 0.705 | 1.639 | 3.879 | 5.327 | 9.343 | 79.2 | 90 | 0 |
| go-redis | 1000 | 153601 | 2.535 | 17.311 | 19.951 | 21.983 | 24.207 | 65.3 | 344 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 5831 | 0.161 | 0.210 | 0.355 | 0.642 | 1.865 | 300.2 | 28 | 0 |
| go-redis | 10 | 36674 | 0.256 | 0.370 | 0.753 | 1.618 | 2.685 | 251.5 | 49 | 0 |
| go-redis | 100 | 43650 | 2.081 | 3.041 | 4.487 | 6.207 | 9.047 | 206.8 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6458 | 0.145 | 0.189 | 0.316 | 0.874 | 3.199 | 139.7 | 33 | 0 |
| go-redis-autopipeline | 10 | 50249 | 0.143 | 0.362 | 0.530 | 0.970 | 2.135 | 85.0 | 45 | 0 |
| go-redis-autopipeline | 100 | 218400 | 0.389 | 0.795 | 1.424 | 2.211 | 5.231 | 33.6 | 75 | 0 |
| go-redis-autopipeline | 1000 | 722033 | 1.335 | 2.025 | 3.211 | 4.883 | 9.087 | 12.5 | 337 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 5982 | 0.159 | 0.200 | 0.329 | 0.623 | 1.765 | 297.3 | 31 | 0 |
| go-redis-autopipeline | 10 | 37111 | 0.254 | 0.364 | 0.721 | 1.521 | 2.565 | 255.2 | 50 | 0 |
| go-redis-autopipeline | 100 | 46854 | 1.988 | 2.747 | 4.081 | 6.163 | 8.487 | 201.6 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6478 | 0.144 | 0.183 | 0.318 | 0.532 | 1.823 | 124.6 | 35 | 0 |
| glide | 10 | 28661 | 0.331 | 0.463 | 0.681 | 0.983 | 1.968 | 61.8 | 39 | 0 |
| glide | 100 | 39206 | 2.649 | 3.269 | 4.187 | 5.847 | 8.375 | 51.9 | 67 | 0 |
| glide | 1000 | 40418 | 24.559 | 28.127 | 31.455 | 32.991 | 35.455 | 51.9 | 226 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4993 | 0.182 | 0.254 | 0.402 | 0.696 | 3.155 | 191.3 | 32 | 0 |
| glide | 10 | 11865 | 0.834 | 1.040 | 1.412 | 1.917 | 2.901 | 139.3 | 34 | 0 |
| glide | 100 | 12601 | 8.167 | 10.983 | 13.463 | 16.015 | 19.151 | 133.8 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12211 | 0.078 | 0.101 | 0.171 | 0.331 | 1.584 | 50.9 | 33 | 0 |
| valkey-go | 10 | 91024 | 0.104 | 0.149 | 0.211 | 0.349 | 1.826 | 41.4 | 38 | 0 |
| valkey-go | 100 | 587218 | 0.154 | 0.251 | 0.418 | 0.824 | 2.635 | 11.3 | 62 | 0 |
| valkey-go | 1000 | 1999218 | 0.447 | 0.757 | 1.511 | 2.909 | 5.079 | 6.5 | 313 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7095 | 0.135 | 0.167 | 0.272 | 0.476 | 2.073 | 239.6 | 35 | 0 |
| valkey-go | 10 | 43822 | 0.209 | 0.309 | 0.563 | 1.229 | 2.771 | 119.4 | 38 | 0 |
| valkey-go | 100 | 202899 | 0.443 | 0.695 | 1.732 | 2.659 | 4.683 | 47.8 | 69 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6861 | 0.138 | 0.176 | 0.289 | 0.490 | 1.852 | 135.8 | 32 | 0 |
| go-redis-autopipeline | 10 | 49715 | 0.145 | 0.364 | 0.538 | 1.015 | 2.539 | 85.8 | 44 | 0 |
| go-redis-autopipeline | 100 | 222193 | 0.381 | 0.785 | 1.409 | 2.159 | 3.647 | 33.1 | 73 | 0 |
| go-redis-autopipeline | 1000 | 723249 | 1.338 | 2.020 | 3.085 | 4.855 | 8.727 | 12.5 | 340 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6072 | 0.156 | 0.199 | 0.329 | 0.670 | 3.061 | 295.1 | 30 | 0 |
| go-redis-autopipeline | 10 | 35118 | 0.267 | 0.386 | 0.789 | 1.711 | 3.393 | 257.4 | 49 | 0 |
| go-redis-autopipeline | 100 | 46322 | 2.011 | 2.775 | 4.135 | 6.095 | 8.735 | 202.1 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6515 | 0.142 | 0.183 | 0.303 | 0.541 | 2.625 | 125.5 | 33 | 0 |
| glide | 10 | 29209 | 0.323 | 0.468 | 0.668 | 1.079 | 2.075 | 61.6 | 37 | 0 |
| glide | 100 | 39761 | 2.593 | 3.273 | 4.299 | 5.771 | 7.475 | 51.3 | 68 | 0 |
| glide | 1000 | 40196 | 24.703 | 27.887 | 31.423 | 34.239 | 36.031 | 53.6 | 219 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4876 | 0.188 | 0.257 | 0.408 | 0.706 | 1.879 | 195.5 | 32 | 0 |
| glide | 10 | 11976 | 0.823 | 1.045 | 1.420 | 2.041 | 2.795 | 139.3 | 34 | 0 |
| glide | 100 | 12513 | 8.287 | 11.143 | 13.487 | 15.951 | 18.863 | 135.2 | 68 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12117 | 0.077 | 0.103 | 0.190 | 0.343 | 1.641 | 50.6 | 33 | 0 |
| valkey-go | 10 | 90876 | 0.104 | 0.150 | 0.212 | 0.358 | 2.119 | 41.4 | 38 | 0 |
| valkey-go | 100 | 573649 | 0.158 | 0.257 | 0.429 | 0.871 | 2.801 | 11.5 | 62 | 0 |
| valkey-go | 1000 | 2040427 | 0.440 | 0.728 | 1.477 | 2.845 | 5.467 | 6.5 | 312 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7016 | 0.135 | 0.176 | 0.283 | 0.541 | 1.831 | 234.7 | 36 | 0 |
| valkey-go | 10 | 43722 | 0.210 | 0.309 | 0.565 | 1.200 | 2.373 | 118.5 | 42 | 0 |
| valkey-go | 100 | 205002 | 0.438 | 0.687 | 1.723 | 2.613 | 4.013 | 47.8 | 71 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11037 | 0.086 | 0.107 | 0.178 | 0.348 | 1.738 | 60.7 | 28 | 0 |
| go-redis | 10 | 87087 | 0.105 | 0.157 | 0.245 | 0.590 | 2.000 | 69.9 | 34 | 0 |
| go-redis | 100 | 145514 | 0.551 | 1.363 | 2.399 | 3.301 | 6.267 | 65.7 | 91 | 0 |
| go-redis | 1000 | 146753 | 1.917 | 18.271 | 31.151 | 33.343 | 36.095 | 65.6 | 342 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6158 | 0.154 | 0.194 | 0.327 | 0.662 | 1.892 | 290.8 | 30 | 0 |
| go-redis | 10 | 37263 | 0.252 | 0.364 | 0.722 | 1.528 | 3.773 | 249.9 | 49 | 0 |
| go-redis | 100 | 46749 | 1.991 | 2.761 | 4.053 | 6.383 | 9.575 | 203.6 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6130 | 0.154 | 0.196 | 0.319 | 0.543 | 1.908 | 132.0 | 33 | 0 |
| glide | 10 | 29306 | 0.324 | 0.450 | 0.660 | 1.054 | 3.097 | 61.3 | 39 | 0 |
| glide | 100 | 39599 | 2.585 | 3.327 | 4.379 | 5.799 | 7.859 | 52.2 | 69 | 0 |
| glide | 1000 | 39488 | 25.263 | 28.671 | 32.639 | 36.319 | 38.655 | 53.5 | 235 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4741 | 0.197 | 0.266 | 0.404 | 0.754 | 1.757 | 200.0 | 34 | 0 |
| glide | 10 | 11733 | 0.842 | 1.055 | 1.436 | 2.000 | 2.609 | 140.7 | 37 | 0 |
| glide | 100 | 12570 | 8.263 | 11.007 | 13.343 | 15.535 | 18.591 | 134.2 | 72 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12364 | 0.077 | 0.097 | 0.166 | 0.314 | 1.688 | 51.1 | 33 | 0 |
| valkey-go | 10 | 89788 | 0.106 | 0.151 | 0.214 | 0.371 | 1.996 | 42.0 | 36 | 0 |
| valkey-go | 100 | 602010 | 0.151 | 0.243 | 0.391 | 0.813 | 2.621 | 11.3 | 62 | 0 |
| valkey-go | 1000 | 1946445 | 0.460 | 0.802 | 1.504 | 2.847 | 6.103 | 6.6 | 313 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7147 | 0.133 | 0.169 | 0.277 | 0.482 | 1.920 | 233.8 | 34 | 0 |
| valkey-go | 10 | 44054 | 0.208 | 0.306 | 0.566 | 1.185 | 2.477 | 118.3 | 39 | 0 |
| valkey-go | 100 | 202755 | 0.443 | 0.700 | 1.728 | 2.621 | 4.255 | 47.1 | 69 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10875 | 0.087 | 0.111 | 0.191 | 0.356 | 2.939 | 60.5 | 28 | 0 |
| go-redis | 10 | 86513 | 0.106 | 0.159 | 0.250 | 0.565 | 2.275 | 69.9 | 34 | 0 |
| go-redis | 100 | 150734 | 0.573 | 1.236 | 1.964 | 2.871 | 5.851 | 64.5 | 88 | 0 |
| go-redis | 1000 | 155803 | 1.500 | 17.999 | 20.751 | 23.599 | 25.631 | 63.9 | 344 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6072 | 0.157 | 0.196 | 0.329 | 0.645 | 2.010 | 296.1 | 28 | 0 |
| go-redis | 10 | 37541 | 0.251 | 0.359 | 0.706 | 1.432 | 3.551 | 249.5 | 49 | 0 |
| go-redis | 100 | 46028 | 2.018 | 2.795 | 4.119 | 6.203 | 8.751 | 203.8 | 111 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 7033 | 0.135 | 0.169 | 0.269 | 0.513 | 1.913 | 132.0 | 30 | 0 |
| go-redis-autopipeline | 10 | 50082 | 0.144 | 0.363 | 0.533 | 0.992 | 2.859 | 84.8 | 43 | 0 |
| go-redis-autopipeline | 100 | 202734 | 0.439 | 0.813 | 1.436 | 2.233 | 3.731 | 37.5 | 68 | 0 |
| go-redis-autopipeline | 1000 | 389552 | 2.425 | 4.175 | 6.159 | 8.115 | 12.927 | 19.3 | 330 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6032 | 0.158 | 0.198 | 0.323 | 0.613 | 1.771 | 298.7 | 32 | 0 |
| go-redis-autopipeline | 10 | 32111 | 0.277 | 0.473 | 0.868 | 1.924 | 4.595 | 286.6 | 51 | 0 |
| go-redis-autopipeline | 100 | 39447 | 2.399 | 3.375 | 5.047 | 6.531 | 10.471 | 214.8 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12167 | 0.078 | 0.100 | 0.172 | 0.349 | 1.580 | 50.9 | 35 | 0 |
| valkey-go | 10 | 88953 | 0.107 | 0.153 | 0.215 | 0.351 | 2.014 | 42.0 | 38 | 0 |
| valkey-go | 100 | 581724 | 0.155 | 0.254 | 0.431 | 0.827 | 2.387 | 11.5 | 64 | 0 |
| valkey-go | 1000 | 2007375 | 0.448 | 0.755 | 1.567 | 2.829 | 5.051 | 6.5 | 315 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7049 | 0.136 | 0.169 | 0.275 | 0.445 | 1.702 | 239.1 | 34 | 0 |
| valkey-go | 10 | 43179 | 0.212 | 0.314 | 0.571 | 1.195 | 2.197 | 120.8 | 38 | 0 |
| valkey-go | 100 | 202923 | 0.439 | 0.702 | 1.754 | 2.685 | 4.255 | 48.1 | 68 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10782 | 0.088 | 0.112 | 0.185 | 0.356 | 1.655 | 62.2 | 28 | 0 |
| go-redis | 10 | 87873 | 0.104 | 0.155 | 0.244 | 0.601 | 3.221 | 69.2 | 34 | 0 |
| go-redis | 100 | 152529 | 0.560 | 1.252 | 1.877 | 2.813 | 5.643 | 64.2 | 89 | 0 |
| go-redis | 1000 | 155816 | 2.549 | 17.071 | 19.423 | 21.615 | 24.239 | 64.1 | 343 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6058 | 0.156 | 0.200 | 0.345 | 0.612 | 1.781 | 291.0 | 28 | 0 |
| go-redis | 10 | 36697 | 0.255 | 0.371 | 0.774 | 1.628 | 4.335 | 249.3 | 49 | 0 |
| go-redis | 100 | 47266 | 1.974 | 2.725 | 3.963 | 5.807 | 8.111 | 200.3 | 108 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6950 | 0.137 | 0.170 | 0.269 | 0.471 | 3.249 | 133.6 | 32 | 0 |
| go-redis-autopipeline | 10 | 50671 | 0.142 | 0.360 | 0.528 | 0.949 | 2.371 | 84.0 | 45 | 0 |
| go-redis-autopipeline | 100 | 207260 | 0.428 | 0.799 | 1.415 | 2.253 | 4.151 | 36.9 | 71 | 0 |
| go-redis-autopipeline | 1000 | 402105 | 2.359 | 3.995 | 5.835 | 7.739 | 11.471 | 18.7 | 341 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6084 | 0.156 | 0.197 | 0.326 | 0.624 | 1.691 | 293.8 | 30 | 0 |
| go-redis-autopipeline | 10 | 37254 | 0.252 | 0.362 | 0.720 | 1.621 | 3.285 | 253.6 | 51 | 0 |
| go-redis-autopipeline | 100 | 46699 | 1.988 | 2.765 | 4.103 | 6.179 | 8.223 | 202.2 | 110 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6196 | 0.156 | 0.188 | 0.301 | 0.536 | 1.793 | 133.4 | 33 | 0 |
| glide | 10 | 28311 | 0.332 | 0.482 | 0.718 | 1.072 | 2.283 | 63.6 | 37 | 0 |
| glide | 100 | 40682 | 2.507 | 3.307 | 4.259 | 5.247 | 7.715 | 50.3 | 66 | 0 |
| glide | 1000 | 41779 | 23.791 | 27.695 | 31.695 | 32.831 | 35.007 | 51.1 | 247 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4858 | 0.187 | 0.268 | 0.422 | 0.740 | 2.029 | 193.7 | 34 | 0 |
| glide | 10 | 12069 | 0.816 | 1.023 | 1.411 | 2.009 | 2.561 | 136.9 | 36 | 0 |
| glide | 100 | 12569 | 8.223 | 11.127 | 13.639 | 16.127 | 18.479 | 134.7 | 72 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 11211 | 0.086 | 0.104 | 0.179 | 0.344 | 1.487 | 60.3 | 28 | 0 |
| go-redis | 10 | 88672 | 0.104 | 0.153 | 0.234 | 0.587 | 1.902 | 69.9 | 37 | 0 |
| go-redis | 100 | 152001 | 0.565 | 1.241 | 1.896 | 2.825 | 6.451 | 64.6 | 92 | 0 |
| go-redis | 1000 | 156283 | 2.151 | 17.359 | 19.695 | 22.239 | 24.031 | 64.3 | 341 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6052 | 0.157 | 0.198 | 0.330 | 0.654 | 1.954 | 295.1 | 29 | 0 |
| go-redis | 10 | 37585 | 0.250 | 0.358 | 0.734 | 1.613 | 3.241 | 248.2 | 49 | 0 |
| go-redis | 100 | 46878 | 1.988 | 2.761 | 4.035 | 6.095 | 8.335 | 202.7 | 108 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 7081 | 0.134 | 0.168 | 0.271 | 0.495 | 1.777 | 130.8 | 30 | 0 |
| go-redis-autopipeline | 10 | 50633 | 0.142 | 0.361 | 0.524 | 0.953 | 2.727 | 84.7 | 44 | 0 |
| go-redis-autopipeline | 100 | 217104 | 0.399 | 0.785 | 1.404 | 2.147 | 3.589 | 34.5 | 70 | 0 |
| go-redis-autopipeline | 1000 | 463171 | 2.095 | 3.309 | 4.659 | 6.395 | 11.743 | 16.5 | 336 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6099 | 0.156 | 0.196 | 0.324 | 0.660 | 1.718 | 296.4 | 30 | 0 |
| go-redis-autopipeline | 10 | 37312 | 0.252 | 0.363 | 0.708 | 1.542 | 3.041 | 253.2 | 52 | 0 |
| go-redis-autopipeline | 100 | 46795 | 1.988 | 2.759 | 4.069 | 6.239 | 8.879 | 202.7 | 113 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6220 | 0.152 | 0.191 | 0.311 | 0.517 | 2.497 | 132.9 | 33 | 0 |
| glide | 10 | 28188 | 0.336 | 0.477 | 0.691 | 0.972 | 2.089 | 62.3 | 36 | 0 |
| glide | 100 | 40522 | 2.571 | 3.181 | 3.939 | 5.527 | 6.827 | 50.6 | 67 | 0 |
| glide | 1000 | 39819 | 24.847 | 28.623 | 32.799 | 35.711 | 37.887 | 52.7 | 241 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4849 | 0.191 | 0.257 | 0.394 | 0.704 | 1.789 | 196.2 | 32 | 0 |
| glide | 10 | 11851 | 0.833 | 1.046 | 1.438 | 1.995 | 2.603 | 139.0 | 35 | 0 |
| glide | 100 | 12393 | 8.255 | 11.223 | 13.935 | 16.687 | 20.591 | 136.9 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12477 | 0.077 | 0.096 | 0.165 | 0.322 | 1.468 | 49.5 | 32 | 0 |
| valkey-go | 10 | 89702 | 0.106 | 0.151 | 0.210 | 0.350 | 2.079 | 41.9 | 37 | 0 |
| valkey-go | 100 | 592113 | 0.153 | 0.248 | 0.405 | 0.823 | 2.577 | 11.2 | 63 | 0 |
| valkey-go | 1000 | 2058155 | 0.427 | 0.728 | 1.551 | 2.941 | 6.815 | 6.3 | 313 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7098 | 0.134 | 0.170 | 0.279 | 0.529 | 2.457 | 235.4 | 36 | 0 |
| valkey-go | 10 | 43748 | 0.210 | 0.310 | 0.564 | 1.188 | 2.755 | 118.6 | 39 | 0 |
| valkey-go | 100 | 203854 | 0.438 | 0.703 | 1.706 | 2.637 | 4.167 | 47.2 | 71 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 7063 | 0.135 | 0.168 | 0.267 | 0.479 | 1.648 | 133.3 | 31 | 0 |
| go-redis-autopipeline | 10 | 50510 | 0.142 | 0.361 | 0.527 | 1.002 | 2.305 | 84.7 | 43 | 0 |
| go-redis-autopipeline | 100 | 225160 | 0.375 | 0.776 | 1.398 | 2.093 | 3.585 | 33.1 | 70 | 0 |
| go-redis-autopipeline | 1000 | 724689 | 1.330 | 2.015 | 3.175 | 5.347 | 9.527 | 12.6 | 342 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 1 | 6162 | 0.155 | 0.192 | 0.314 | 0.628 | 1.710 | 293.2 | 32 | 0 |
| go-redis-autopipeline | 10 | 37238 | 0.252 | 0.362 | 0.719 | 1.593 | 3.197 | 254.5 | 50 | 0 |
| go-redis-autopipeline | 100 | 43289 | 2.079 | 3.105 | 4.983 | 7.263 | 14.631 | 205.1 | 109 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 6396 | 0.148 | 0.183 | 0.292 | 0.503 | 2.303 | 129.3 | 33 | 0 |
| glide | 10 | 29639 | 0.319 | 0.454 | 0.661 | 1.022 | 2.141 | 60.4 | 39 | 0 |
| glide | 100 | 39684 | 2.589 | 3.289 | 4.395 | 5.575 | 7.071 | 51.7 | 67 | 0 |
| glide | 1000 | 40143 | 24.863 | 27.887 | 32.415 | 36.127 | 37.279 | 51.9 | 249 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 1 | 4894 | 0.186 | 0.262 | 0.408 | 0.684 | 1.809 | 193.1 | 32 | 0 |
| glide | 10 | 11944 | 0.825 | 1.045 | 1.437 | 2.047 | 2.951 | 138.6 | 35 | 0 |
| glide | 100 | 12632 | 8.187 | 10.991 | 13.391 | 15.535 | 18.543 | 135.2 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 12758 | 0.077 | 0.092 | 0.155 | 0.311 | 1.674 | 48.0 | 35 | 0 |
| valkey-go | 10 | 89087 | 0.107 | 0.153 | 0.216 | 0.369 | 1.939 | 42.0 | 38 | 0 |
| valkey-go | 100 | 591461 | 0.153 | 0.248 | 0.414 | 0.852 | 2.291 | 11.3 | 64 | 0 |
| valkey-go | 1000 | 2033891 | 0.436 | 0.754 | 1.504 | 2.913 | 9.575 | 6.4 | 315 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1 | 7146 | 0.134 | 0.165 | 0.272 | 0.451 | 1.658 | 235.3 | 35 | 0 |
| valkey-go | 10 | 44082 | 0.208 | 0.306 | 0.569 | 1.228 | 2.219 | 117.9 | 40 | 0 |
| valkey-go | 100 | 203524 | 0.437 | 0.698 | 1.787 | 2.697 | 3.929 | 47.5 | 70 | 0 |
mode=cluster fixed concurrency, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 10502 | 0.091 | 0.116 | 0.188 | 0.352 | 2.601 | 62.0 | 30 | 0 |
| go-redis | 10 | 74051 | 0.117 | 0.206 | 0.349 | 0.662 | 2.297 | 79.5 | 37 | 0 |
| go-redis | 100 | 152207 | 0.562 | 1.256 | 1.840 | 2.745 | 6.487 | 64.7 | 91 | 0 |
| go-redis | 1000 | 156768 | 2.619 | 17.007 | 19.231 | 21.775 | 24.207 | 64.0 | 345 | 0 |
mode=cluster fixed concurrency, workload=batches of 10 GETs, each sent as one pipeline value=100 B keys=100000 warm-up=1s measured=5s per level
Latency = from the call to the reply, per request (a batch is one request). Client CPU per request includes the load generator's own small, equal overhead.

| Client | In flight | Requests/s | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 1 | 6106 | 0.156 | 0.194 | 0.324 | 0.649 | 3.083 | 294.6 | 30 | 0 |
| go-redis | 10 | 37308 | 0.252 | 0.363 | 0.726 | 1.527 | 3.599 | 250.4 | 49 | 0 |
| go-redis | 100 | 46978 | 1.986 | 2.745 | 4.037 | 6.079 | 9.263 | 202.1 | 107 | 0 |
