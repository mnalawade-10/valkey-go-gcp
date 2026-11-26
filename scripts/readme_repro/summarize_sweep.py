#!/usr/bin/env python3
"""Summarize sweep_set.txt: each valkey-go version against rueidis v0.0.90 measured in the same process.

Prints, per version and case, median ops/s for both and the change (valkey-go vs rueidis v0.0.90),
plus a geometric-mean change per version.
"""
import math
import re
import statistics
import sys
from collections import defaultdict

LINE = re.compile(r"^Benchmark\w+/lib=([^/]+)/par=(\d+)/val=(\d+)-\d+\s+\d+\s+([\d.]+) ns/op")
path = sys.argv[1]
sections = []  # (version, {(lib, par, val): [ns...]})
for line in open(path, encoding="utf-8"):
    if line.startswith("== valkey-go "):
        sections.append((line.split()[-1], defaultdict(list)))
        continue
    m = LINE.match(line)
    if m and sections:
        lib, par, val, ns = m.groups()
        sections[-1][1][(lib, int(par), int(val))].append(float(ns))

cases = [(p, v) for p in (1, 8, 64) for v in (64, 256, 1024)]
print("| valkey-go version | " + " | ".join(f"{p * 10} gor., {v} B" for p, v in cases) + " | Geomean |")
print("|" + " :- |" * (len(cases) + 2))
for version, data in sections:
    cells, logs = [], []
    for p, v in cases:
        new = data.get((f"valkeygo-{version}", p, v))
        old = data.get(("rueidis-v0.0.90", p, v))
        if not new or not old:
            cells.append("")
            continue
        r = statistics.median(old) / statistics.median(new)  # throughput ratio new/old
        logs.append(math.log(r))
        cells.append(f"{(r - 1) * 100:+.0f}%")
    gm = f"{(math.exp(sum(logs) / len(logs)) - 1) * 100:+.0f}%" if logs else ""
    print(f"| {version} | " + " | ".join(cells) + f" | {gm} |")
