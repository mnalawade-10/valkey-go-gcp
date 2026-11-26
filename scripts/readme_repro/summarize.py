#!/usr/bin/env python3
"""Summarize the README rerun: medians per client, and the ratios the report needs.

Usage: summarize.py results/   (reads single_set.txt, single_get.txt, cluster_set.txt)
Prints Markdown tables. Throughput = 1e9 / median ns/op (parallel benchmarks report wall time per op).
"""
import os
import re
import statistics
import sys
from collections import defaultdict

LINE = re.compile(r"^Benchmark(\w+)/lib=([^/]+)/par=(\d+)/val=(\d+)-\d+\s+\d+\s+([\d.]+) ns/op.*?([\d.]+) B/op\s+([\d.]+) allocs/op")

# The README chart (rueidis-benchmark fd05043, MacBook Pro M1 Pro, Dec 2022): SET ns/op.
README = {
    (1, 64): (3201, 5489), (1, 256): (3256, 5687), (1, 1024): (3243, 5299),
    (8, 64): (803.4, 5499), (8, 256): (834.4, 5723), (8, 1024): (923.8, 5289),
    (64, 64): (531.9, 7607), (64, 256): (624.7, 7570), (64, 1024): (693.9, 8030),
}


def load(path):
    data = defaultdict(lambda: {"ns": [], "b": [], "allocs": []})
    if not os.path.exists(path):
        return data
    for line in open(path, encoding="utf-8"):
        m = LINE.match(line)
        if not m:
            continue
        bench, lib, par, val, ns, b, allocs = m.groups()
        d = data[(bench, lib, int(par), int(val))]
        d["ns"].append(float(ns))
        d["b"].append(float(b))
        d["allocs"].append(float(allocs))
    return data


def med(d, key, lib, field="ns"):
    v = d.get(key[:1] + (lib,) + key[1:])
    if not v or not v[field]:
        return None
    return statistics.median(v[field])


def spread(d, key, lib):
    v = d.get(key[:1] + (lib,) + key[1:])
    if not v or len(v["ns"]) < 2:
        return None
    m = statistics.median(v["ns"])
    return 100 * (max(v["ns"]) - min(v["ns"])) / 2 / m


def kops(ns):
    return "" if ns is None else f"{1e9 / ns / 1000:,.0f}k"


def ratio(a, b):
    """How many times faster a is than b, from ns/op (b / a)."""
    if a is None or b is None:
        return ""
    return f"{b / a:.1f}×"


def change(new, old):
    if new is None or old is None:
        return ""
    return f"{(old / new - 1) * 100:+.0f}%"  # throughput change, new vs old


def main():
    root = sys.argv[1]
    s = load(os.path.join(root, "single_set.txt"))
    g = load(os.path.join(root, "single_get.txt"))
    c = load(os.path.join(root, "cluster_set.txt"))

    RO, GO, VG, GN, GA = "rueidis-v0.0.90", "goredis-v9.0.0-rc.2", "valkeygo-v1.0.80", "goredis-v9.23.0", "goredis-v9.23.0-autopipeline"

    def table(title, d, bench, pars, readme=False):
        print(f"\n### {title}\n")
        hdr = "| Goroutines | Value | rueidis v0.0.90 | valkey-go v1.0.80 | Change | go-redis rc.2 | go-redis v9.23.0 | Change | go-redis AutoPipeline | valkey-go ÷ go-redis | valkey-go ÷ AutoPipeline |"
        if readme:
            hdr += " README ratio (M1, 2022) |"
        print(hdr)
        print("|" + " :- |" * (hdr.count("|") - 1))
        for p in pars:
            for v in (64, 256, 1024):
                k = (bench, p, v)
                ro, go, vg, gn, ga = (med(d, k, lib) for lib in (RO, GO, VG, GN, GA))
                row = f"| {p * 10} | {v} B | {kops(ro)} | {kops(vg)} | {change(vg, ro)} | {kops(go)} | {kops(gn)} | {change(gn, go)} | {kops(ga)} | {ratio(vg, gn)} | {ratio(vg, ga)} |"
                if readme:
                    r = README.get((p, v))
                    row += f" {r[1] / r[0]:.1f}× |" if r else " |"
                print(row)

    table("SET, single server (the README chart), ops/s", s, "SingleClientSet", (1, 8, 64), readme=True)
    table("GET, single server, ops/s", g, "SingleClientGet", (1, 8, 64))
    table("SET, 3-node cluster, ops/s", c, "ClusterClientSet", (4, 16, 96))

    print("\n### GET with client-side caching (single server), ops/s\n")
    print("| Goroutines | Value | rueidis v0.0.90 + CSC | valkey-go v1.0.80 + CSC | go-redis v9.23.0 + CSC (experimental) |")
    print("| :- | :- | :- | :- | :- |")
    for p in (1, 8, 64):
        for v in (64, 256, 1024):
            k = ("SingleClientGet", p, v)
            print(f"| {p * 10} | {v} B | {kops(med(g, k, RO + '-csc'))} | {kops(med(g, k, VG + '-csc'))} | {kops(med(g, k, GN + '-csc'))} |")

    print("\n### Memory per command (median B/op and allocs/op), SET single server, 640 goroutines, 64 B\n")
    print("| Client | B/op | allocs/op |")
    print("| :- | :- | :- |")
    k = ("SingleClientSet", 64, 64)
    for lib in (RO, VG, GO, GN, GA):
        print(f"| {lib} | {med(s, k, lib, 'b'):.0f} | {med(s, k, lib, 'allocs'):.0f} |" if med(s, k, lib, "b") is not None else f"| {lib} | | |")

    print("\n### Run-to-run spread (half the min-max range, % of median), worst case per client\n")
    worst = defaultdict(float)
    for d in (s, g, c):
        for (bench, lib, p, v), vals in d.items():
            if len(vals["ns"]) >= 2:
                m = statistics.median(vals["ns"])
                worst[lib] = max(worst[lib], 100 * (max(vals["ns"]) - min(vals["ns"])) / 2 / m)
    for lib, w in sorted(worst.items()):
        print(f"- {lib}: ±{w:.0f}%")


if __name__ == "__main__":
    main()
