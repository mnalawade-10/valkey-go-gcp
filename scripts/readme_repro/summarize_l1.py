#!/usr/bin/env python3
"""Summarize BenchmarkL1 output: median commands/s per client, by topology, operation, key shape and goroutines.

Usage: summarize_l1.py results_l1/l1.txt
"""
import re
import statistics
import sys
from collections import defaultdict

LINE = re.compile(r"^BenchmarkL1/topo=(\w+)/op=(\w+)/keys=(\w+)/lib=([^/]+)/par=(\d+)/val=(\d+)-(\d+)\s+\d+\s+([\d.]+) ns/op")
data = defaultdict(list)
procs = None
for line in open(sys.argv[1], encoding="utf-8"):
    m = LINE.match(line)
    if not m:
        continue
    topo, op, keys, lib, par, val, procs, ns = m.groups()
    data[(topo, op, keys, lib, int(par))].append(float(ns))

VG, GR, GA, GL = "valkeygo-v1.0.80", "goredis-v9.23.0", "goredis-v9.23.0-autopipeline", "glide-v2.6.1"


def ops(key):
    v = data.get(key)
    return 1e9 / statistics.median(v) if v else None


def k(x):
    return "" if x is None else f"{x / 1000:,.0f}k"


def ratio(a, b):
    return "" if a is None or b is None else f"{a / b:.1f}×"


gor = int(procs or 10)
for topo in ("single", "cluster"):
    for op in ("set", "get"):
        rows = sorted({(kk, p) for (t, o, kk, lib, p) in data if t == topo and o == op},
                      key=lambda r: (["one", "uniform", "zipf"].index(r[0]), r[1]))
        if not rows:
            continue
        print(f"\n#### {topo}, {op.upper()} (64 B values), commands/s\n")
        print("| Keys | Goroutines | valkey-go | go-redis (default pool) | go-redis AutoPipeline | GLIDE | valkey-go ÷ go-redis | valkey-go ÷ AutoPipeline | valkey-go ÷ GLIDE |")
        print("| :--- | :- | :- | :- | :- | :- | :- | :- | :- |")
        for kk, p in rows:
            vg, gr, ga, gl = (ops((topo, op, kk, lib, p)) for lib in (VG, GR, GA, GL))
            print(f"| {kk} | {p * gor} | {k(vg)} | {k(gr)} | {k(ga)} | {k(gl)} | {ratio(vg, gr)} | {ratio(vg, ga)} | {ratio(vg, gl)} |")
