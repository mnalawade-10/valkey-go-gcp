# Benchmark results: official-20261123

## L0: engine baseline (memtier)

Median of 3 runs of 60 s per setting.

| Topology | Value | Pipeline | Commands/s | p50 | p99 | p99.9 |
| :--- | :- | :- | :- | :- | :- | :- |
| Single server | 100 B | 1 | 52,224 | 1.83 ms | 3.74 ms | 4.74 ms |
| Single server | 100 B | 16 | 522,293 | 2.91 ms | 5.82 ms | 6.46 ms |
| Single server | 4,000 B | 1 | 48,074 | 2.01 ms | 4.01 ms | 4.42 ms |
| Single server | 4,000 B | 16 | 239,304 | 5.66 ms | 14.14 ms | 17.41 ms |
| Cluster, 3 nodes | 100 B | 1 | 150,548 | 0.85 ms | 1.78 ms | 2.65 ms |
| Cluster, 3 nodes | 100 B | 16 | 927,469 | 4.19 ms | 6.37 ms | 7.65 ms |
| Cluster, 3 nodes | 4,000 B | 1 | 142,551 | 0.94 ms | 2.19 ms | 2.80 ms |
| Cluster, 3 nodes | 4,000 B | 16 | 460,373 | 8.64 ms | 13.76 ms | 16.25 ms |

## L1: the README benchmark with today's clients

rueidis-benchmark, adapted; 64 B values; median of 10 runs of 50,000 commands per case. Throughput in commands per second; CPU is the client process's CPU time per command.

**Single server, SET, 1,024 keys: commands per second**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 10 | 74k | 54k | 59k | 36k |
| 80 | 300k | 55k | 175k | 38k |
| 640 | 853k | 55k | 370k | 37k |

**Single server, SET: client CPU per command (µs)**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 10 | 42.1 | 58.2 | 21.3 | 49.7 |
| 80 | 16.2 | 61.3 | 9.3 | 49.5 |
| 640 | 6.3 | 68.7 | 7.8 | 50.3 |

**Single server, GET, 1,024 keys: commands per second**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 10 | 76k | 55k | 62k | 36k |
| 80 | 301k | 56k | 192k | 35k |
| 640 | 920k | 56k | 385k | 34k |

**Single server, GET: client CPU per command (µs)**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 10 | 43.0 | 59.2 | 21.0 | 51.8 |
| 80 | 17.1 | 61.5 | 8.9 | 54.2 |
| 640 | 6.9 | 69.4 | 7.8 | 55.9 |

**Cluster, 3 nodes, SET, 1,024 keys: commands per second**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 40 | 266k | 148k | 111k | 34k |
| 160 | 851k | 157k | 241k | 43k |
| 960 | 1,896k | 158k | 461k | 43k |

**Cluster, 3 nodes, SET: client CPU per command (µs)**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 40 | 18.8 | 52.7 | 40.3 | 52.5 |
| 160 | 8.1 | 50.6 | 21.6 | 44.7 |
| 960 | 5.1 | 51.7 | 12.7 | 45.8 |

**Cluster, 3 nodes, GET, 1,024 keys: commands per second**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 40 | 268k | 146k | 110k | 31k |
| 160 | 876k | 161k | 246k | 39k |
| 960 | 1,745k | 159k | 381k | 39k |

**Cluster, 3 nodes, GET: client CPU per command (µs)**

| Goroutines | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- |
| 40 | 19.0 | 53.4 | 41.5 | 57.8 |
| 160 | 7.9 | 50.4 | 21.4 | 50.1 |
| 960 | 5.5 | 51.6 | 14.7 | 51.5 |

**Key shapes: SET at the highest concurrency, commands per second**

| Topology | Keys | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| single, 640 goroutines | one | 856k | 55k | 370k | 38k |
| single, 640 goroutines | uniform | 853k | 55k | 370k | 37k |
| single, 640 goroutines | zipf | 845k | 56k | 388k | 37k |
| cluster, 960 goroutines | uniform | 1,896k | 158k | 461k | 43k |
| cluster, 960 goroutines | zipf | 1,813k | 159k | 560k | 42k |

## L2 and L3: single server, 100 B

**Clients: requests per second / p99**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 15k / 0.10 ms | 14k / 0.11 ms | 9.9k / 0.19 ms | 7.6k / 0.27 ms |
| 80/20 mix | 10 in flight | 72k / 0.25 ms | 52k / 0.34 ms | 60k / 0.36 ms | 34k / 0.54 ms |
| 80/20 mix | 100 in flight | 321k / 0.57 ms | 53k / 3.62 ms | 205k / 0.81 ms | 36k / 4.38 ms |
| 80/20 mix | 1,000 in flight | 811k / 2.11 ms | 53k / 22.47 ms | 458k / 3.64 ms | 35k / 36.99 ms |
| batch of 10 GETs | 1 in flight | 13k / 0.13 ms | 12k / 0.16 ms | 12k / 0.17 ms | 6.6k / 0.35 ms |
| batch of 10 GETs | 10 in flight | 47k / 0.42 ms | 38k / 0.58 ms | 38k / 0.58 ms | 26k / 0.95 ms |
| batch of 10 GETs | 100 in flight | 95k / 1.83 ms | 39k / 4.81 ms | 39k / 4.81 ms | 29k / 5.40 ms |

**Clients: client CPU per request (µs)**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 46.2 | 53.5 | 103.7 | 121.1 |
| 80/20 mix | 10 in flight | 43.8 | 60.3 | 21.1 | 55.5 |
| 80/20 mix | 100 in flight | 15.9 | 66.7 | 8.1 | 56.2 |
| 80/20 mix | 1,000 in flight | 6.8 | 78.3 | 6.6 | 58.3 |
| batch of 10 GETs | 1 in flight | 53.6 | 63.0 | 63.6 | 123.1 |
| batch of 10 GETs | 10 in flight | 60.8 | 80.2 | 80.3 | 93.2 |
| batch of 10 GETs | 100 in flight | 28.7 | 82.6 | 83.5 | 85.1 |

**valkey-go features off vs on: change in requests per second against plain valkey-go** (`~` = no real difference)

| Workload | In flight | no-autopipe | always-pipe | flush-20us | flush-100us | 1-conn | no-retry | docache | compat | hook | otel | otel-sdk | aside | lock |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 | ~ | -19.24% | ~ | +1.70% | ~ | ~ | -18.16% | -1.85% | ~ | -2.85% | -29.59% | -23.72% | -84.33% |
| 80/20 mix | 10 | -26.56% | ~ | ~ | -15.75% | +12.60% | ~ | +94.88% | -3.66% | ~ | -2.18% | -7.67% | +89.38% | -80.18% |
| 80/20 mix | 100 | -83.41% | ~ | +6.96% | +25.86% | +30.39% | ~ | +43.89% | -4.30% | ~ | -2.89% | -8.59% | -39.65% | -91.53% |
| 80/20 mix | 1,000 | -94.05% | +7.99% | +11.93% | +18.08% | +25.47% | +9.79% | -35.54% | +8.66% | +9.73% | +7.14% | -3.10% | -76.27% | -98.17% |
| batch of 10 GETs | 1 | ~ | -26.70% | ~ | ~ | ~ | ~ | -17.41% | -3.58% | ~ | -3.45% | -26.57% | -89.68% | -98.16% |
| batch of 10 GETs | 10 | -16.75% | ~ | ~ | -7.09% | +3.38% | ~ | +622.16% | -3.26% | ~ | -2.05% | -4.21% | +79.79% | -97.41% |
| batch of 10 GETs | 100 | -58.21% | +7.07% | +8.46% | +9.47% | +23.67% | +8.16% | +484.43% | +6.59% | +8.59% | +7.82% | +5.69% | -45.48% | -98.19% |

**valkey-go features off vs on: change in p99 against plain valkey-go** (`~` = no real difference)

| Workload | In flight | no-autopipe | always-pipe | flush-20us | flush-100us | 1-conn | no-retry | docache | compat | hook | otel | otel-sdk | aside | lock |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 | ~ | +36.67% | ~ | ~ | ~ | ~ | +72.86% | ~ | ~ | +5.24% | +56.67% | +86.19% | +590.48% |
| 80/20 mix | 10 | +27.31% | ~ | ~ | +24.17% | -15.13% | ~ | +16.50% | +5.70% | ~ | +5.30% | +35.95% | +25.15% | +454.81% |
| 80/20 mix | 100 | +531.40% | ~ | -9.39% | -15.79% | -23.16% | ~ | +68.07% | +2.54% | ~ | +15.79% | +96.23% | +226.05% | +856.32% |
| 80/20 mix | 1,000 | +1096.16% | -9.75% | -11.65% | -14.61% | -32.42% | ~ | +276.71% | -9.37% | -8.28% | +7.02% | +127.94% | +656.50% | +3851.38% |
| batch of 10 GETs | 1 | ~ | +43.82% | ~ | ~ | ~ | ~ | +111.24% | ~ | ~ | ~ | +46.07% | +929.21% | +4381.65% |
| batch of 10 GETs | 10 | +2.38% | ~ | ~ | +9.03% | +8.08% | ~ | -53.68% | +18.41% | ~ | +8.67% | +29.22% | +31.47% | +2412.83% |
| batch of 10 GETs | 100 | +162.83% | -7.53% | -5.87% | ~ | -23.58% | ~ | -36.60% | +5.65% | -7.89% | ~ | +10.18% | +636.84% | +3418.72% |

**valkey-go features off vs on: change in client CPU per request against plain valkey-go** (`~` = no real difference)

| Workload | In flight | no-autopipe | always-pipe | flush-20us | flush-100us | 1-conn | no-retry | docache | compat | hook | otel | otel-sdk | aside | lock |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 | ~ | +51.78% | ~ | ~ | ~ | ~ | +62.47% | ~ | ~ | +8.30% | +119.89% | +80.69% | +1018.94% |
| 80/20 mix | 10 | +15.86% | ~ | +3.42% | -17.48% | -47.17% | ~ | -44.15% | +0.95% | ~ | +9.29% | +38.44% | -51.61% | +323.72% |
| 80/20 mix | 100 | +247.15% | ~ | -13.57% | -44.01% | -54.58% | ~ | -18.71% | +2.98% | ~ | +14.45% | +69.31% | +8.21% | +524.89% |
| 80/20 mix | 1,000 | +663.69% | ~ | -6.37% | -19.77% | -22.18% | ~ | +73.65% | +4.10% | ~ | +25.99% | +150.44% | +156.44% | +2075.70% |
| batch of 10 GETs | 1 | ~ | +75.53% | ~ | ~ | ~ | ~ | +192.86% | +6.88% | ~ | +4.81% | +102.50% | +1283.29% | +9499.65% |
| batch of 10 GETs | 10 | +5.88% | ~ | ~ | -8.46% | -32.55% | ~ | -42.74% | +12.21% | ~ | +8.13% | +25.52% | ~ | +3327.80% |
| batch of 10 GETs | 100 | +133.00% | ~ | -5.53% | -19.35% | -34.90% | ~ | ~ | +22.02% | ~ | +11.38% | +47.30% | +236.43% | +4554.67% |

## L2: fixed rates, and elephants and mice (single server, 100 B)

**Clients at fixed rates: requests per second / p99**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 2,000/s | 2.0k / 0.38 ms | 2.0k / 0.37 ms | 2.0k / 2.64 ms | 2.0k / 0.52 ms |
| 80/20 mix | 10,000/s | 10k / 0.35 ms | 10.0k / 0.47 ms | 10.0k / 2.59 ms | 10k / 0.40 ms |
| 80/20 mix | 20,000/s | 20k / 0.22 ms | 20k / 0.28 ms | 20k / 2.79 ms | 20k / 0.56 ms |
| small GETs | 5,000/s | 5.0k / 0.38 ms | 5.0k / 0.40 ms | 5.0k / 2.55 ms | 5.0k / 0.47 ms |
| small GETs + 1 MB SETs | 5,000/s | 5.0k / 1.19 ms | 5.0k / 2.99 ms | 5.0k / 3.37 ms | 5.0k / 3.62 ms |

**Clients at fixed rates: client CPU per request (µs)**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 2,000/s | 133.1 | 120.3 | 116.3 | 170.6 |
| 80/20 mix | 10,000/s | 95.2 | 81.2 | 33.0 | 123.6 |
| 80/20 mix | 20,000/s | 76.1 | 70.1 | 21.6 | 86.5 |
| small GETs | 5,000/s | 86.5 | 92.5 | 57.1 | 128.2 |
| small GETs + 1 MB SETs | 5,000/s | 89.6 | 98.8 | 60.9 | 136.9 |

## L2: 4 KB values (single server)

**Clients, 4 KB: requests per second / p99**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 14k / 0.13 ms | 13k / 0.15 ms | 8.9k / 0.24 ms | 7.3k / 0.29 ms |
| 80/20 mix | 10 in flight | 62k / 0.34 ms | 45k / 0.54 ms | 46k / 0.59 ms | 28k / 0.66 ms |
| 80/20 mix | 100 in flight | 220k / 1.04 ms | 48k / 3.95 ms | 130k / 1.54 ms | 32k / 5.01 ms |
| 80/20 mix | 1,000 in flight | 256k / 9.60 ms | 48k / 24.37 ms | 194k / 7.53 ms | 30k / 40.29 ms |
| 80/20 mix | 2,000/s | 2.0k / 0.39 ms | 2.0k / 0.39 ms | 2.0k / 2.64 ms | 2.0k / 0.53 ms |
| 80/20 mix | 10,000/s | 10k / 0.38 ms | 10.0k / 0.50 ms | 10.0k / 2.66 ms | 10k / 0.44 ms |

**Clients, 4 KB: client CPU per request (µs)**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 51.5 | 59.4 | 116.5 | 127.6 |
| 80/20 mix | 10 in flight | 49.8 | 72.8 | 27.9 | 68.2 |
| 80/20 mix | 100 in flight | 20.5 | 71.3 | 12.8 | 63.3 |
| 80/20 mix | 1,000 in flight | 18.8 | 83.2 | 9.9 | 66.8 |
| 80/20 mix | 2,000/s | 141.3 | 129.4 | 127.0 | 176.2 |
| 80/20 mix | 10,000/s | 99.2 | 86.4 | 35.8 | 126.6 |

## L2: 3-node cluster, 100 B

**Clients, cluster: requests per second / p99**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 12k / 0.17 ms | 11k / 0.19 ms | 7.0k / 0.27 ms | 6.2k / 0.32 ms |
| 80/20 mix | 10 in flight | 89k / 0.22 ms | 87k / 0.25 ms | 50k / 0.53 ms | 29k / 0.68 ms |
| 80/20 mix | 100 in flight | 580k / 0.43 ms | 151k / 1.96 ms | 217k / 1.41 ms | 40k / 4.29 ms |
| 80/20 mix | 1,000 in flight | 2,003k / 1.55 ms | 155k / 20.12 ms | 463k / 4.70 ms | 40k / 32.53 ms |
| batch of 10 GETs | 1 in flight | 7.0k / 0.28 ms | 6.1k / 0.33 ms | 6.1k / 0.33 ms | 4.9k / 0.41 ms |
| batch of 10 GETs | 10 in flight | 44k / 0.57 ms | 37k / 0.72 ms | 37k / 0.72 ms | 12k / 1.44 ms |
| batch of 10 GETs | 100 in flight | 204k / 1.73 ms | 47k / 4.09 ms | 46k / 4.14 ms | 13k / 13.57 ms |

**Clients, cluster: client CPU per request (µs)**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 50.5 | 62.1 | 133.2 | 130.4 |
| 80/20 mix | 10 in flight | 42.0 | 69.8 | 84.8 | 62.0 |
| 80/20 mix | 100 in flight | 11.4 | 64.6 | 34.8 | 51.6 |
| 80/20 mix | 1,000 in flight | 6.4 | 64.3 | 16.6 | 52.9 |
| batch of 10 GETs | 1 in flight | 238.8 | 295.6 | 296.9 | 195.2 |
| batch of 10 GETs | 10 in flight | 120.0 | 250.4 | 256.0 | 139.3 |
| batch of 10 GETs | 100 in flight | 47.7 | 203.1 | 203.0 | 135.2 |

## Client-side cache at fixed hit rates (DoCache against Do)

| Hits | Measured | In flight | Do, requests/s | DoCache: requests/s | p99 | CPU per request |
| :--- | :- | :- | :- | :- | :- | :- |
| 50% target | 50.2% | 1 | 15k | +17.21% | +64.25% | +11.02% |
| 50% target | 50.0% | 10 | 74k | +18.13% | +54.33% | -21.77% |
| 50% target | 51.2% | 100 | 348k | -45.29% | +217.83% | +30.73% |
| 90% target | 89.9% | 1 | 16k | +392.77% | +33.97% | -73.26% |
| 90% target | 90.0% | 10 | 74k | +452.24% | +18.73% | -80.97% |
| 90% target | 90.2% | 100 | 346k | +164.58% | +162.26% | -52.17% |
| 99% target | 99.0% | 1 | 16k | +2931.48% | -11.54% | -95.53% |
| 99% target | 99.0% | 10 | 74k | +2807.19% | -50.40% | -94.55% |
| 99% target | 99.0% | 100 | 347k | +820.92% | -21.04% | -77.07% |

## TLS off vs on

**Without TLS (plain port of the same server): requests per second / p99**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 16k / 0.11 ms | 14k / 0.12 ms | 9.7k / 0.21 ms | 7.6k / 0.27 ms |
| 80/20 mix | 10 in flight | 73k / 0.26 ms | 52k / 0.34 ms | 60k / 0.36 ms | 33k / 0.55 ms |
| 80/20 mix | 100 in flight | 328k / 0.54 ms | 53k / 3.60 ms | 203k / 0.77 ms | 36k / 4.41 ms |
| 80/20 mix | 1,000 in flight | 905k / 1.85 ms | 53k / 23.18 ms | 447k / 3.63 ms | 35k / 38.09 ms |
| batch of 10 GETs | 1 in flight | 13k / 0.13 ms | 12k / 0.17 ms | 12k / 0.17 ms | 6.6k / 0.34 ms |
| batch of 10 GETs | 10 in flight | 48k / 0.40 ms | 38k / 0.55 ms | 38k / 0.56 ms | 26k / 0.97 ms |
| batch of 10 GETs | 100 in flight | 105k / 1.66 ms | 40k / 4.79 ms | 40k / 4.78 ms | 29k / 5.48 ms |

**Without TLS (plain port of the same server): client CPU per request (µs)**

| Workload | Level | valkey-go | go-redis | go-redis auto-pipelining | GLIDE |
| :--- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 in flight | 46.8 | 54.0 | 106.7 | 120.3 |
| 80/20 mix | 10 in flight | 43.7 | 61.3 | 21.1 | 56.6 |
| 80/20 mix | 100 in flight | 15.7 | 68.1 | 8.2 | 56.7 |
| 80/20 mix | 1,000 in flight | 6.9 | 79.5 | 6.8 | 59.0 |
| batch of 10 GETs | 1 in flight | 53.0 | 64.2 | 64.8 | 123.4 |
| batch of 10 GETs | 10 in flight | 60.2 | 79.8 | 80.3 | 93.2 |
| batch of 10 GETs | 100 in flight | 28.7 | 83.6 | 83.9 | 84.8 |

**With TLS: requests per second / p99 (change in requests per second against the same client without TLS)**

| Workload | In flight | valkey-go | go-redis | go-redis auto-pipelining | GLIDE | valkey-go, 100 µs write delay |
| :--- | :- | :- | :- | :- | :- | :- |
| 80/20 mix | 1 | 13k / 0.13 ms (-19.14%) | 12k / 0.14 ms (-14.81%) | 8.0k / 0.26 ms (-17.69%) | 6.6k / 0.31 ms (-13.00%) | 13k / 0.13 ms (-18.43%) |
| 80/20 mix | 10 | 36k / 0.58 ms (-50.93%) | 37k / 0.45 ms (-28.62%) | 52k / 0.40 ms (-12.66%) | 30k / 0.63 ms (-9.43%) | 49k / 0.38 ms (-18.87%) |
| 80/20 mix | 100 | 37k / 8.02 ms (-88.72%) | 38k / 5.06 ms (-29.11%) | 190k / 0.83 ms (-6.23%) | 32k / 4.74 ms (-8.98%) | 210k / 1.12 ms (-49.09%) |
| 80/20 mix | 1,000 | 37k / 54.73 ms (-95.91%) | 37k / 32.58 ms (-29.40%) | 419k / 3.79 ms (-6.24%) | 32k / 40.73 ms (-6.73%) | 219k / 10.00 ms (-77.53%) |
| batch of 10 GETs | 1 | 10k / 0.17 ms (-20.43%) | 9.9k / 0.20 ms (-16.72%) | 9.8k / 0.20 ms (-16.55%) | 5.8k / 0.39 ms (-13.04%) | 11k / 0.17 ms (-19.63%) |
| batch of 10 GETs | 10 | 28k / 0.72 ms (-41.20%) | 29k / 0.63 ms (-24.28%) | 29k / 0.63 ms (-24.45%) | 23k / 0.96 ms (-9.15%) | 35k / 0.54 ms (-22.63%) |
| batch of 10 GETs | 100 | 29k / 9.83 ms (-72.17%) | 30k / 6.37 ms (-25.19%) | 29k / 6.43 ms (-26.13%) | 26k / 5.76 ms (-11.80%) | 45k / 5.76 ms (-58.18%) |

## Release sweep

| valkey-go version | 10 gor., 64 B | 10 gor., 256 B | 10 gor., 1024 B | 80 gor., 64 B | 80 gor., 256 B | 80 gor., 1024 B | 640 gor., 64 B | 640 gor., 256 B | 640 gor., 1024 B | Geomean |
| :- | :- | :- | :- | :- | :- | :- | :- | :- | :- | :- |
| v1.0.40 | -12% | -10% | -10% | -18% | -11% | -27% | -20% | +2% | +2% | -12% |
| v1.0.50 | -11% | -11% | -10% | -19% | -14% | -25% | -19% | -0% | -1% | -13% |
| v1.0.60 | -12% | -12% | -9% | -19% | -16% | -27% | -18% | +1% | +2% | -13% |
| v1.0.70 | -12% | -13% | -11% | -19% | -11% | -31% | -19% | -5% | +5% | -13% |
| v1.0.80 | -13% | -12% | -11% | -20% | -17% | -28% | -21% | +1% | +1% | -14% |

## Rate-limiter memory

`valkey_valkey_9.0-alpine_n1000000.txt`

| Limiter | Keys per identifier | Bytes per identifier (used_memory) | vs valkeylimiter fixed window | MEMORY USAGE, one identifier | Allow p50 | Allow p99 | Allow/s (64 goroutines) |
| :--- | :- | :- | :- | :- | :- | :- | :- |
| valkeylimiter fixed window | 2.00 | 208 | +0% | 120 B | 89µs | 171µs | 104071 |
| valkeylimiter GCRA | 1.00 | 132 | -36% | 88 B | 106µs | 198µs | 66582 |
| valkeylimiter fixed window (prefix rl) | 2.00 | 184 | -11% | 96 B | 90µs | 164µs | 100160 |
| valkeylimiter GCRA (prefix rl) | 1.00 | 116 | -44% | 72 B | 107µs | 206µs | 69078 |
| valkeycompatrate (GCRA) | 1.00 | 108 | -48% | 64 B | 110µs | 201µs | 70507 |
| go-redis redis_rate v10 (GCRA) | 1.00 | 107 | -48% | 64 B | 97µs | 169µs | 32527 |

## Stale-read meter

Median of 3 runs; latency and stale time from the run with the median stale share.

| Client | Reads/s | Read p50 / p99 | Stale reads | Stale for: p50 / p99 / max |
| :--- | :- | :- | :- | :- |
| valkey-go (no cache) | 63,675 | 118.9µs / 235.9µs | 0.00% | none |
| valkey-go DoCache | 7,216,177 | 700ns / 1.6µs | 0.02% | 12µs / 279µs / 1ms |
| go-redis CSC | 4,509,394 | 1.4µs / 3.4µs | 2.29% | 1.434ms / 4.963ms / 11ms |
| GLIDE cache (time-based) | 94,285 | 81.8µs / 223.5µs | 99.50% | 5.185812s / 9.858378s / 9.998s |
| GLIDE cache (server-assisted) | 92,214 | 80µs / 263.7µs | 0.09% | 96µs / 412µs / 1ms |

## Cold start and fleet footprint: standalone

mode=standalone nodes=1 instances=50 load=32 goroutines/instance for 3s on 64 keys, app vCPUs (GOMAXPROCS)=4, median of 5 runs

| Library | Connections per instance, idle | Connections per instance under load (busiest node) | Instances until 32,000 connections on a node | Server memory per instance, under load | First reply p50 / p99 / max | All instances up | Failed GETs under load | GETs/s, all instances |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 1.8 | 4.0 | 8000 | 63 KB | 16.3ms / 18.6ms / 18.6ms | 20ms | 0 | 210017 |
| valkey-go (ForceSingleClient) | 1.8 | 4.0 | 8000 | 64 KB | 19ms / 21.2ms / 21.2ms | 22ms | 0 | 213825 |
| go-redis | 1.0 | 32.0 | 1000 | 525 KB | 4.6ms / 5.8ms / 5.8ms | 6ms | 0 | 45947 |
| go-redis-autopipeline | 1.0 | 1.0 | 32000 | 16 KB | 7.4ms / 9.3ms / 9.3ms | 10ms | 0 | 739418 |
| glide | 1.0 | 1.0 | 32000 | 16 KB | 4.8ms / 5.8ms / 5.8ms | 11ms | 221 (0.010%) | 141091 |

## Cold start and fleet footprint: cluster

mode=cluster nodes=3 instances=50 load=32 goroutines/instance for 3s on 64 keys, app vCPUs (GOMAXPROCS)=4, median of 5 runs

| Library | Connections per instance, idle | Connections per instance under load (busiest node) | Instances until 32,000 connections on a node | Server memory per instance, under load | First reply p50 / p99 / max | All instances up | Failed GETs under load | GETs/s, all instances |
| :--- | :- | :- | :- | :- | :- | :- | :- | :- |
| valkey-go | 2.3 | 1.0 | 32000 | 35 KB | 29.6ms / 38.2ms / 38.2ms | 40ms | 0 | 277084 |
| go-redis | 2.0 | 20.0 | 1600 | 914 KB | 27.4ms / 33ms / 33ms | 34ms | 0 | 97558 |
| go-redis-autopipeline | 2.0 | 2.4 | 13445 | 93 KB | 10.7ms / 11.9ms / 11.9ms | 12ms | 0 | 214221 |
| glide | 6.0 | 2.0 | 16000 | 86 KB | 17.3ms / 25.6ms / 25.6ms | 28ms | 690 (0.036%) | 127917 |

## Failover drill (local)

3 runs per client; medians. Failover 5 s into a 20 s window; 10 requests in flight.

| Client | Requests/s before | Slowest second around the failover | Requests/s after | READONLY errors | Other errors | Recovered: no errors and 80% of before, at the end |
| :--- | :- | :- | :- | :- | :- | :- |
| valkey-go | 65,120 | 6,425 | 57,846 | 201,507 | 0 | 0 of 3 |
| valkey-go/redirect | 65,304 | 52,183 | 64,465 | 0 | 0 | 3 of 3 |
| go-redis | 47,336 | 395 | 391 | 1,351 | 0 | 0 of 3 |
| glide | 34,216 | 3,627 | 26,921 | 0 | 93,494 | 0 of 3 |

