#!/usr/bin/env python3
"""Summarize memtier JSON files (results/*.json) into a Markdown table."""
import json
import os
import sys

root = sys.argv[1]
rows = []
for name in sorted(os.listdir(root)):
    if not name.endswith(".json"):
        continue
    with open(os.path.join(root, name)) as f:
        d = json.load(f)
    totals = d["ALL STATS"]["Totals"]
    pct = totals.get("Percentile Latencies", {})
    topo, size, pipe = name[:-5].split("_")
    rows.append((topo, int(size[1:]), int(pipe[1:]), totals["Ops/sec"], totals.get("Latency", totals.get("Average Latency")),
                 pct.get("p50.00"), pct.get("p99.00"), pct.get("p99.90")))

print("| Topology | Value | Pipeline | Ops/s | Avg latency (ms) | p50 (ms) | p99 (ms) | p99.9 (ms) |")
print("| :--- | :- | :- | :- | :- | :- | :- | :- |")
for topo, size, pipe, ops, avg, p50, p99, p999 in sorted(rows, key=lambda r: (r[0] != "standalone", r[1], r[2])):
    print(f"| {topo} | {size} B | {pipe} | {ops:,.0f} | {avg:.3f} | {p50:.3f} | {p99:.3f} | {p999:.3f} |")
