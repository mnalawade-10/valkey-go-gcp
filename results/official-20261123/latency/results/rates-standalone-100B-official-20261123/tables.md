mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.137 | 0.200 | 0.358 | 0.993 | 1.640 | 1.117 | 1.297 | 119.8 | 241 | 0 |
| go-redis | 10000 | 9998 | 0.144 | 0.268 | 0.464 | 0.719 | 2.095 | 1.218 | 1.476 | 81.5 | 430 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.094 | 0.237 | 0.990 | 3.223 | 0.101 | 0.395 | 69.6 | 428 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.138 | 0.211 | 0.382 | 0.962 | 1.624 | 1.159 | 1.363 | 92.4 | 251 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.140 | 0.233 | 3.195 | 5.259 | 7.315 | 1.161 | 3.593 | 97.4 | 255 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.260 | 1.504 | 2.647 | 2.837 | 3.687 | 1.107 | 3.319 | 116.7 | 247 | 0 |
| go-redis-autopipeline | 10000 | 9998 | 0.394 | 1.569 | 2.597 | 2.909 | 3.685 | 1.120 | 2.737 | 33.5 | 450 | 0 |
| go-redis-autopipeline | 20000 | 19997 | 0.439 | 1.683 | 2.789 | 3.117 | 4.167 | 1.135 | 2.955 | 21.3 | 438 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 5000 | 0.565 | 1.517 | 2.555 | 2.761 | 2.995 | 1.122 | 2.705 | 55.8 | 251 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.262 | 1.546 | 3.353 | 4.711 | 6.011 | 1.117 | 3.679 | 57.2 | 259 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.207 | 0.280 | 0.522 | 1.058 | 1.599 | 0.889 | 1.186 | 162.2 | 243 | 0 |
| glide | 10000 | 10000 | 0.141 | 0.191 | 0.422 | 0.824 | 2.035 | 0.351 | 0.743 | 125.0 | 443 | 0 |
| glide | 20000 | 19999 | 0.157 | 0.257 | 0.540 | 1.397 | 2.655 | 0.245 | 0.721 | 86.0 | 444 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.188 | 0.272 | 0.466 | 1.112 | 1.608 | 0.945 | 1.247 | 125.2 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.188 | 0.287 | 3.397 | 4.979 | 5.863 | 0.936 | 3.659 | 132.3 | 281 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.152 | 0.258 | 0.424 | 0.716 | 1.846 | 1.107 | 1.371 | 145.5 | 239 | 0 |
| valkey-go | 10000 | 9999 | 0.102 | 0.199 | 0.349 | 0.820 | 1.760 | 1.151 | 1.350 | 97.0 | 269 | 0 |
| valkey-go | 20000 | 20000 | 0.086 | 0.125 | 0.244 | 0.616 | 1.923 | 0.346 | 0.603 | 76.5 | 375 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.143 | 0.212 | 0.377 | 1.041 | 2.637 | 1.154 | 1.359 | 87.0 | 244 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.143 | 0.221 | 1.219 | 2.935 | 3.485 | 1.147 | 1.628 | 90.2 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.268 | 1.501 | 2.647 | 3.017 | 4.011 | 1.108 | 3.397 | 113.2 | 245 | 0 |
| go-redis-autopipeline | 10000 | 9998 | 0.387 | 1.553 | 2.583 | 2.901 | 4.107 | 1.115 | 2.713 | 32.4 | 272 | 0 |
| go-redis-autopipeline | 20000 | 19999 | 0.417 | 1.650 | 2.687 | 3.007 | 3.395 | 1.131 | 2.837 | 21.2 | 362 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.222 | 1.517 | 2.529 | 2.781 | 3.845 | 1.117 | 2.705 | 55.0 | 259 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.274 | 1.526 | 3.073 | 4.311 | 5.019 | 1.116 | 3.485 | 58.5 | 264 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.210 | 0.281 | 0.513 | 1.075 | 1.623 | 0.895 | 1.187 | 164.6 | 243 | 0 |
| glide | 10000 | 10000 | 0.138 | 0.188 | 0.399 | 0.740 | 1.777 | 0.364 | 0.746 | 124.6 | 257 | 0 |
| glide | 20000 | 20000 | 0.138 | 0.236 | 0.563 | 1.680 | 2.909 | 0.249 | 0.758 | 86.5 | 273 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.187 | 0.265 | 0.455 | 1.105 | 1.810 | 0.934 | 1.220 | 125.0 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 4999 | 0.190 | 0.284 | 3.345 | 5.163 | 5.731 | 0.941 | 3.601 | 136.3 | 286 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.142 | 0.202 | 0.374 | 0.874 | 1.689 | 1.110 | 1.315 | 125.1 | 241 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.200 | 0.355 | 0.506 | 1.445 | 1.157 | 1.356 | 92.8 | 244 | 0 |
| valkey-go | 20000 | 20000 | 0.084 | 0.121 | 0.219 | 0.507 | 1.626 | 0.385 | 0.614 | 76.0 | 252 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.139 | 0.203 | 0.393 | 0.747 | 1.846 | 1.139 | 1.333 | 83.9 | 244 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.140 | 0.216 | 1.015 | 2.621 | 3.659 | 1.147 | 1.480 | 88.3 | 242 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.131 | 0.197 | 0.369 | 0.735 | 3.183 | 1.116 | 1.308 | 117.0 | 241 | 0 |
| go-redis | 10000 | 9999 | 0.147 | 0.280 | 0.489 | 0.926 | 2.151 | 1.210 | 1.481 | 81.7 | 259 | 0 |
| go-redis | 20000 | 20000 | 0.073 | 0.099 | 0.318 | 0.963 | 2.229 | 0.142 | 0.481 | 71.3 | 374 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.138 | 0.211 | 0.401 | 0.701 | 1.136 | 1.156 | 1.360 | 91.1 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.144 | 0.242 | 3.227 | 4.863 | 6.411 | 1.166 | 3.549 | 100.9 | 255 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.215 | 0.292 | 0.523 | 0.953 | 1.744 | 0.899 | 1.206 | 169.6 | 241 | 0 |
| glide | 10000 | 10000 | 0.135 | 0.184 | 0.404 | 0.846 | 1.943 | 0.358 | 0.736 | 122.6 | 306 | 0 |
| glide | 20000 | 20000 | 0.146 | 0.274 | 0.601 | 1.417 | 2.287 | 0.277 | 0.792 | 86.8 | 276 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.192 | 0.278 | 0.473 | 1.168 | 1.772 | 0.946 | 1.253 | 129.2 | 246 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.194 | 0.299 | 4.061 | 5.799 | 6.907 | 0.942 | 4.255 | 144.4 | 270 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.142 | 0.206 | 0.363 | 0.687 | 1.758 | 1.115 | 1.313 | 130.1 | 242 | 0 |
| valkey-go | 10000 | 10000 | 0.102 | 0.195 | 0.349 | 0.575 | 1.922 | 1.150 | 1.349 | 94.4 | 242 | 0 |
| valkey-go | 20000 | 20000 | 0.084 | 0.120 | 0.226 | 0.505 | 1.740 | 0.470 | 0.728 | 76.1 | 385 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.139 | 0.203 | 0.361 | 0.579 | 1.258 | 1.143 | 1.339 | 82.9 | 243 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.143 | 0.216 | 1.271 | 2.853 | 3.473 | 1.147 | 1.633 | 89.3 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.133 | 0.200 | 0.363 | 0.536 | 1.528 | 1.118 | 1.322 | 122.9 | 241 | 0 |
| go-redis | 10000 | 10000 | 0.147 | 0.281 | 0.474 | 0.809 | 1.835 | 1.216 | 1.483 | 83.3 | 412 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.096 | 0.256 | 0.860 | 2.729 | 0.109 | 0.412 | 70.7 | 409 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.142 | 0.219 | 0.401 | 0.642 | 1.584 | 1.162 | 1.374 | 92.3 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.142 | 0.234 | 2.957 | 4.583 | 5.671 | 1.162 | 3.329 | 99.2 | 253 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 0.605 | 1.524 | 2.643 | 2.905 | 3.887 | 1.118 | 3.319 | 146.4 | 245 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.395 | 1.561 | 2.621 | 2.901 | 3.911 | 1.115 | 2.753 | 33.0 | 270 | 0 |
| go-redis-autopipeline | 20000 | 19991 | 0.429 | 1.681 | 2.825 | 3.169 | 4.611 | 1.133 | 2.981 | 21.8 | 307 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 0.528 | 1.533 | 2.587 | 2.843 | 3.127 | 1.128 | 2.753 | 58.6 | 255 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.258 | 1.556 | 3.439 | 4.323 | 5.335 | 1.120 | 3.679 | 62.2 | 264 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.145 | 0.206 | 0.380 | 0.779 | 1.385 | 1.115 | 1.299 | 133.2 | 242 | 0 |
| valkey-go | 10000 | 10000 | 0.103 | 0.199 | 0.332 | 0.535 | 1.778 | 1.151 | 1.348 | 93.3 | 245 | 0 |
| valkey-go | 20000 | 20000 | 0.084 | 0.119 | 0.211 | 0.489 | 1.780 | 0.377 | 0.605 | 75.5 | 368 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.141 | 0.209 | 0.385 | 0.635 | 1.630 | 1.148 | 1.355 | 87.2 | 243 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.143 | 0.222 | 1.167 | 2.805 | 3.263 | 1.152 | 1.567 | 91.0 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.134 | 0.196 | 0.393 | 0.804 | 1.435 | 1.120 | 1.302 | 118.8 | 241 | 0 |
| go-redis | 10000 | 10000 | 0.143 | 0.264 | 0.460 | 0.695 | 1.856 | 1.213 | 1.460 | 80.8 | 263 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.097 | 0.329 | 0.921 | 2.807 | 0.165 | 0.500 | 70.3 | 403 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.142 | 0.218 | 0.414 | 0.773 | 1.526 | 1.166 | 1.384 | 92.9 | 253 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.143 | 0.236 | 3.013 | 4.723 | 5.359 | 1.164 | 3.423 | 98.5 | 253 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.216 | 1.514 | 2.641 | 2.957 | 5.079 | 1.115 | 3.233 | 125.8 | 247 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.383 | 1.546 | 2.587 | 2.875 | 3.811 | 1.117 | 2.719 | 33.0 | 280 | 0 |
| go-redis-autopipeline | 20000 | 19995 | 0.427 | 1.659 | 2.715 | 3.059 | 4.147 | 1.134 | 2.871 | 21.6 | 290 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 0.470 | 1.519 | 2.521 | 2.931 | 4.021 | 1.122 | 2.723 | 56.3 | 261 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.254 | 1.542 | 3.201 | 4.167 | 4.919 | 1.119 | 3.593 | 60.2 | 257 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.213 | 0.287 | 0.535 | 1.610 | 3.063 | 0.893 | 1.197 | 169.6 | 243 | 0 |
| glide | 10000 | 10000 | 0.135 | 0.185 | 0.411 | 0.992 | 1.790 | 0.383 | 0.787 | 123.3 | 261 | 0 |
| glide | 20000 | 20000 | 0.134 | 0.239 | 0.524 | 1.952 | 3.097 | 0.248 | 0.713 | 86.7 | 273 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.191 | 0.277 | 0.448 | 0.766 | 1.634 | 0.942 | 1.241 | 127.0 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.193 | 0.295 | 4.039 | 5.807 | 6.587 | 0.957 | 4.287 | 139.7 | 268 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.134 | 0.198 | 0.375 | 0.753 | 1.407 | 1.117 | 1.302 | 119.9 | 241 | 0 |
| go-redis | 10000 | 9999 | 0.145 | 0.272 | 0.469 | 0.747 | 1.519 | 1.213 | 1.476 | 80.5 | 442 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.095 | 0.278 | 0.826 | 1.678 | 0.122 | 0.438 | 68.9 | 436 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.141 | 0.219 | 0.401 | 1.000 | 1.349 | 1.158 | 1.373 | 91.6 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.140 | 0.230 | 2.921 | 4.531 | 5.647 | 1.156 | 3.353 | 97.4 | 255 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.276 | 1.507 | 2.645 | 2.831 | 3.661 | 1.109 | 3.347 | 114.1 | 247 | 0 |
| go-redis-autopipeline | 10000 | 9995 | 0.411 | 1.564 | 2.613 | 2.887 | 3.991 | 1.115 | 2.745 | 33.4 | 274 | 0 |
| go-redis-autopipeline | 20000 | 19994 | 0.445 | 1.690 | 2.807 | 3.149 | 4.187 | 1.135 | 2.953 | 21.8 | 306 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.225 | 1.536 | 2.563 | 2.833 | 3.327 | 1.121 | 2.741 | 57.9 | 255 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.275 | 1.549 | 3.393 | 4.391 | 5.083 | 1.119 | 3.663 | 60.2 | 264 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.216 | 0.292 | 0.514 | 0.718 | 1.796 | 0.890 | 1.187 | 170.6 | 241 | 0 |
| glide | 10000 | 10000 | 0.136 | 0.188 | 0.399 | 0.916 | 1.832 | 0.371 | 0.753 | 125.4 | 254 | 0 |
| glide | 20000 | 19999 | 0.143 | 0.238 | 0.588 | 1.616 | 2.569 | 0.243 | 0.756 | 86.1 | 397 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.192 | 0.269 | 0.489 | 1.419 | 3.191 | 0.945 | 1.281 | 131.0 | 250 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.198 | 0.302 | 3.989 | 5.847 | 6.927 | 0.968 | 4.195 | 140.9 | 261 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.157 | 0.244 | 0.420 | 1.392 | 2.491 | 1.123 | 1.365 | 165.6 | 239 | 0 |
| valkey-go | 10000 | 10000 | 0.104 | 0.196 | 0.347 | 0.546 | 1.825 | 1.153 | 1.352 | 97.8 | 246 | 0 |
| valkey-go | 20000 | 20000 | 0.086 | 0.125 | 0.222 | 0.424 | 1.591 | 0.389 | 0.603 | 76.7 | 376 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.142 | 0.211 | 0.374 | 1.026 | 1.443 | 1.143 | 1.346 | 87.5 | 244 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.141 | 0.213 | 1.266 | 2.749 | 4.099 | 1.146 | 1.626 | 88.2 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.280 | 1.497 | 2.631 | 2.889 | 3.747 | 1.103 | 3.347 | 111.0 | 243 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.392 | 1.557 | 2.579 | 2.887 | 3.051 | 1.113 | 2.723 | 32.0 | 442 | 0 |
| go-redis-autopipeline | 20000 | 19994 | 0.433 | 1.677 | 2.807 | 3.195 | 4.151 | 1.130 | 2.969 | 21.3 | 435 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 0.489 | 1.513 | 2.527 | 2.777 | 3.631 | 1.116 | 2.709 | 57.2 | 257 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.229 | 1.552 | 3.387 | 4.639 | 5.551 | 1.117 | 3.675 | 61.5 | 255 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.212 | 0.289 | 0.482 | 0.674 | 1.699 | 0.895 | 1.188 | 170.6 | 243 | 0 |
| glide | 10000 | 10000 | 0.137 | 0.187 | 0.395 | 0.848 | 1.971 | 0.396 | 0.775 | 123.9 | 447 | 0 |
| glide | 20000 | 20000 | 0.140 | 0.235 | 0.472 | 1.339 | 2.267 | 0.228 | 0.634 | 86.3 | 446 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.189 | 0.270 | 0.473 | 1.085 | 2.061 | 0.948 | 1.274 | 128.4 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.193 | 0.298 | 3.187 | 5.603 | 6.415 | 0.958 | 3.383 | 136.9 | 277 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.144 | 0.207 | 0.365 | 1.176 | 1.645 | 1.111 | 1.312 | 136.7 | 239 | 0 |
| valkey-go | 10000 | 9998 | 0.103 | 0.197 | 0.343 | 0.553 | 1.661 | 1.150 | 1.346 | 96.0 | 267 | 0 |
| valkey-go | 20000 | 20000 | 0.085 | 0.123 | 0.225 | 0.454 | 1.706 | 0.586 | 0.816 | 76.2 | 350 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.141 | 0.208 | 0.376 | 1.447 | 1.669 | 1.148 | 1.343 | 89.3 | 244 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.144 | 0.224 | 1.185 | 2.691 | 3.919 | 1.155 | 1.555 | 92.1 | 242 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.136 | 0.203 | 0.373 | 0.612 | 1.442 | 1.112 | 1.306 | 120.7 | 241 | 0 |
| go-redis | 10000 | 10000 | 0.148 | 0.279 | 0.488 | 0.717 | 1.760 | 1.218 | 1.483 | 80.9 | 388 | 0 |
| go-redis | 20000 | 20000 | 0.073 | 0.098 | 0.280 | 0.856 | 2.411 | 0.117 | 0.423 | 71.0 | 272 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.140 | 0.216 | 0.410 | 0.792 | 1.188 | 1.163 | 1.381 | 94.1 | 251 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.145 | 0.236 | 3.083 | 5.299 | 6.283 | 1.168 | 3.421 | 100.8 | 253 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.237 | 0.321 | 0.601 | 2.059 | 3.473 | 0.904 | 1.326 | 172.3 | 241 | 0 |
| glide | 10000 | 10000 | 0.141 | 0.201 | 0.445 | 1.388 | 2.157 | 0.344 | 0.729 | 124.0 | 266 | 0 |
| glide | 20000 | 19999 | 0.144 | 0.271 | 0.698 | 1.963 | 3.185 | 0.295 | 0.917 | 86.5 | 266 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.189 | 0.267 | 0.453 | 1.380 | 1.848 | 0.942 | 1.238 | 128.0 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.193 | 0.303 | 3.997 | 6.047 | 7.271 | 0.948 | 4.195 | 137.3 | 277 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.144 | 0.208 | 0.355 | 0.887 | 1.459 | 1.119 | 1.317 | 132.9 | 239 | 0 |
| valkey-go | 10000 | 10000 | 0.101 | 0.194 | 0.346 | 0.567 | 1.590 | 1.153 | 1.351 | 97.4 | 274 | 0 |
| valkey-go | 20000 | 20000 | 0.085 | 0.122 | 0.224 | 0.565 | 1.734 | 0.424 | 0.667 | 76.1 | 246 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.141 | 0.208 | 0.363 | 1.035 | 1.340 | 1.155 | 1.354 | 86.2 | 243 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.140 | 0.211 | 1.213 | 2.871 | 3.623 | 1.155 | 1.594 | 90.3 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.134 | 0.197 | 0.368 | 0.925 | 1.587 | 1.124 | 1.316 | 121.4 | 241 | 0 |
| go-redis | 10000 | 9998 | 0.142 | 0.267 | 0.475 | 1.311 | 1.855 | 1.213 | 1.472 | 82.1 | 426 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.095 | 0.261 | 0.893 | 1.767 | 0.115 | 0.402 | 70.0 | 422 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.140 | 0.219 | 0.404 | 0.676 | 1.816 | 1.168 | 1.385 | 92.6 | 251 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.143 | 0.236 | 2.973 | 5.139 | 6.623 | 1.162 | 3.347 | 99.7 | 253 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 0.532 | 1.533 | 2.645 | 2.971 | 3.917 | 1.119 | 3.135 | 152.1 | 245 | 0 |
| go-redis-autopipeline | 10000 | 9999 | 0.414 | 1.580 | 2.593 | 2.867 | 3.941 | 1.124 | 2.755 | 34.4 | 442 | 0 |
| go-redis-autopipeline | 20000 | 19993 | 0.448 | 1.694 | 2.735 | 3.059 | 3.999 | 1.136 | 2.903 | 22.3 | 434 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 5000 | 0.648 | 1.540 | 2.557 | 2.819 | 3.661 | 1.123 | 2.733 | 57.7 | 253 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.278 | 1.567 | 3.489 | 4.479 | 4.979 | 1.123 | 3.719 | 61.5 | 256 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.150 | 0.218 | 0.376 | 0.927 | 1.602 | 1.119 | 1.324 | 147.4 | 240 | 0 |
| valkey-go | 10000 | 9998 | 0.101 | 0.194 | 0.353 | 0.832 | 1.885 | 1.154 | 1.354 | 98.6 | 406 | 0 |
| valkey-go | 20000 | 20000 | 0.086 | 0.123 | 0.218 | 0.472 | 1.846 | 0.298 | 0.520 | 76.8 | 403 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.138 | 0.199 | 0.359 | 0.611 | 1.591 | 1.148 | 1.337 | 85.4 | 241 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.140 | 0.211 | 1.185 | 2.797 | 3.255 | 1.155 | 1.546 | 89.8 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.134 | 0.202 | 0.362 | 0.489 | 1.085 | 1.118 | 1.311 | 118.2 | 241 | 0 |
| go-redis | 10000 | 9998 | 0.146 | 0.274 | 0.499 | 1.204 | 1.687 | 1.214 | 1.481 | 80.4 | 391 | 0 |
| go-redis | 20000 | 20000 | 0.073 | 0.097 | 0.342 | 1.609 | 7.355 | 0.143 | 0.517 | 71.2 | 323 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.138 | 0.210 | 0.402 | 0.656 | 1.437 | 1.162 | 1.362 | 94.3 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.139 | 0.228 | 2.965 | 4.947 | 6.111 | 1.162 | 3.381 | 97.8 | 255 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.261 | 1.512 | 2.635 | 2.857 | 3.715 | 1.113 | 3.291 | 124.6 | 249 | 0 |
| go-redis-autopipeline | 10000 | 9999 | 0.396 | 1.565 | 2.597 | 2.891 | 4.599 | 1.124 | 2.743 | 34.1 | 256 | 0 |
| go-redis-autopipeline | 20000 | 19994 | 0.438 | 1.687 | 2.777 | 3.111 | 4.155 | 1.135 | 2.943 | 21.7 | 428 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 0.476 | 1.523 | 2.551 | 2.821 | 3.867 | 1.120 | 2.713 | 57.5 | 264 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.249 | 1.553 | 3.481 | 4.491 | 5.247 | 1.124 | 3.711 | 61.7 | 266 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.216 | 0.293 | 0.504 | 1.130 | 1.660 | 0.895 | 1.212 | 172.9 | 243 | 0 |
| glide | 10000 | 10000 | 0.134 | 0.184 | 0.402 | 0.936 | 1.896 | 0.389 | 0.786 | 122.6 | 439 | 0 |
| glide | 20000 | 19999 | 0.131 | 0.217 | 0.500 | 1.464 | 2.677 | 0.236 | 0.671 | 86.6 | 439 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 4999 | 0.190 | 0.269 | 0.478 | 1.220 | 1.790 | 0.947 | 1.265 | 127.6 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.192 | 0.296 | 3.601 | 5.899 | 7.071 | 0.941 | 3.809 | 137.0 | 276 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.134 | 0.201 | 0.370 | 0.827 | 1.400 | 1.124 | 1.325 | 122.7 | 241 | 0 |
| go-redis | 10000 | 9998 | 0.145 | 0.273 | 0.454 | 0.750 | 2.365 | 1.215 | 1.469 | 79.8 | 442 | 0 |
| go-redis | 20000 | 20000 | 0.071 | 0.096 | 0.353 | 0.878 | 2.975 | 0.299 | 0.643 | 69.5 | 439 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 4999 | 0.134 | 0.209 | 0.414 | 0.975 | 1.525 | 1.162 | 1.375 | 91.6 | 251 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.142 | 0.237 | 3.097 | 5.219 | 5.903 | 1.159 | 3.537 | 96.9 | 253 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 2000 | 1.274 | 1.505 | 2.647 | 2.901 | 3.853 | 1.109 | 3.341 | 113.5 | 241 | 0 |
| go-redis-autopipeline | 10000 | 9996 | 0.393 | 1.546 | 2.589 | 2.871 | 3.659 | 1.117 | 2.727 | 31.9 | 438 | 0 |
| go-redis-autopipeline | 20000 | 19997 | 0.429 | 1.670 | 2.797 | 3.241 | 4.227 | 1.133 | 2.947 | 21.6 | 435 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4998 | 1.259 | 1.519 | 2.565 | 2.795 | 3.803 | 1.121 | 2.741 | 55.4 | 259 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.278 | 1.540 | 3.195 | 4.375 | 4.655 | 1.123 | 3.555 | 59.7 | 264 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.219 | 0.297 | 0.522 | 1.251 | 1.633 | 0.905 | 1.245 | 174.6 | 243 | 0 |
| glide | 10000 | 10000 | 0.135 | 0.190 | 0.402 | 0.824 | 2.029 | 0.424 | 0.805 | 122.4 | 251 | 0 |
| glide | 20000 | 20000 | 0.133 | 0.241 | 0.564 | 1.525 | 2.257 | 0.265 | 0.761 | 85.9 | 448 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.193 | 0.279 | 0.455 | 1.015 | 1.890 | 0.947 | 1.254 | 129.6 | 248 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.189 | 0.293 | 3.633 | 6.067 | 7.195 | 0.939 | 3.883 | 134.4 | 282 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.140 | 0.202 | 0.367 | 0.813 | 1.391 | 1.113 | 1.310 | 127.5 | 240 | 0 |
| valkey-go | 10000 | 9999 | 0.104 | 0.194 | 0.330 | 0.591 | 1.846 | 1.156 | 1.349 | 93.3 | 245 | 0 |
| valkey-go | 20000 | 20000 | 0.083 | 0.120 | 0.223 | 0.442 | 1.459 | 0.494 | 0.722 | 75.6 | 444 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.141 | 0.210 | 0.381 | 0.830 | 1.338 | 1.145 | 1.343 | 85.2 | 241 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.139 | 0.209 | 1.061 | 2.601 | 3.321 | 1.145 | 1.474 | 87.9 | 246 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 2000 | 1999 | 1.281 | 1.506 | 2.633 | 2.859 | 3.759 | 1.112 | 3.355 | 116.0 | 247 | 0 |
| go-redis-autopipeline | 10000 | 9997 | 0.386 | 1.551 | 2.583 | 2.943 | 4.407 | 1.120 | 2.733 | 32.8 | 448 | 0 |
| go-redis-autopipeline | 20000 | 19992 | 0.430 | 1.668 | 2.785 | 3.129 | 4.387 | 1.135 | 2.937 | 21.4 | 443 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 0.781 | 1.530 | 2.545 | 2.847 | 3.045 | 1.127 | 2.735 | 57.1 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis-autopipeline | 5000 | 4999 | 1.110 | 1.546 | 3.167 | 4.159 | 4.851 | 1.124 | 3.515 | 61.6 | 266 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 2000 | 2000 | 0.213 | 0.294 | 0.523 | 1.110 | 1.721 | 0.906 | 1.230 | 170.7 | 241 | 0 |
| glide | 10000 | 10000 | 0.136 | 0.186 | 0.405 | 0.787 | 1.935 | 0.382 | 0.772 | 123.1 | 445 | 0 |
| glide | 20000 | 20000 | 0.155 | 0.265 | 0.568 | 1.477 | 2.401 | 0.260 | 0.762 | 86.7 | 445 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 4999 | 0.192 | 0.277 | 0.465 | 0.793 | 1.599 | 0.945 | 1.243 | 130.0 | 250 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| glide | 5000 | 5000 | 0.192 | 0.297 | 3.571 | 5.651 | 6.679 | 0.958 | 3.841 | 135.4 | 267 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2000 | 2000 | 0.141 | 0.203 | 0.380 | 0.777 | 1.447 | 1.113 | 1.310 | 131.6 | 240 | 0 |
| valkey-go | 10000 | 10000 | 0.106 | 0.200 | 0.354 | 0.617 | 1.822 | 1.153 | 1.353 | 92.3 | 249 | 0 |
| valkey-go | 20000 | 20000 | 0.085 | 0.124 | 0.235 | 0.493 | 1.654 | 0.416 | 0.654 | 76.0 | 265 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 4999 | 0.141 | 0.205 | 0.367 | 0.605 | 1.280 | 1.149 | 1.342 | 86.8 | 242 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 5000 | 5000 | 0.142 | 0.214 | 1.204 | 2.857 | 3.287 | 1.144 | 1.581 | 87.8 | 244 | 0 |
mode=standalone fixed rate, workload=GLIDE mix (80% GET, 20% SET) value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 2000 | 2000 | 0.135 | 0.198 | 0.361 | 0.965 | 1.644 | 1.119 | 1.312 | 122.5 | 241 | 0 |
| go-redis | 10000 | 10000 | 0.144 | 0.277 | 0.482 | 0.839 | 1.649 | 1.227 | 1.498 | 82.0 | 422 | 0 |
| go-redis | 20000 | 20000 | 0.072 | 0.095 | 0.257 | 0.757 | 1.942 | 0.146 | 0.434 | 69.7 | 416 | 0 |
mode=standalone fixed rate, workload=small GETs only value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.137 | 0.210 | 0.389 | 0.526 | 1.528 | 1.173 | 1.370 | 92.8 | 249 | 0 |
mode=standalone fixed rate, workload=small GETs, plus 50 SETs/s of 1024 KB on the same client value=100 B keys=100000 workers=512 warm-up=1s measured=5s per rate
Latency = from the call to the reply. Schedule lag = how late the load generator made the call (it must stay small for the rate to be real).

| Client | Target rate | Achieved | p50 (ms) | p90 (ms) | p99 (ms) | p99.9 (ms) | max (ms) | Schedule lag p99 (ms) | p99 from due time (ms) | Client CPU per request (µs) | Client RSS (MiB) | Errors |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| go-redis | 5000 | 5000 | 0.139 | 0.227 | 2.835 | 4.931 | 5.559 | 1.167 | 3.321 | 99.3 | 253 | 0 |
