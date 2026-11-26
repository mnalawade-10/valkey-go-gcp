#!/usr/bin/env python3
"""Summarize a full run (run_all.sh) into one Markdown file.

  summarize_official.py official-20261123 > official_summary.md

Reads each tool's results for the tag, takes the median across repeated runs, and uses benchstat's verdicts
(through latency/summarize_matrix.py) for the latency matrices: "~" means no real difference.
"""
import glob
import json
import re
import statistics
import sys
from collections import defaultdict
from pathlib import Path

ROOT = Path(__file__).resolve().parent
sys.path.insert(0, str(ROOT / "latency"))
from summarize_matrix import parse  # noqa: E402

TAG = sys.argv[1]
OUT = []
med = statistics.median
CLIENTS = ["valkey-go", "go-redis", "go-redis-autopipeline", "glide"]
NAMES = {"valkey-go": "valkey-go", "go-redis": "go-redis", "go-redis-autopipeline": "go-redis auto-pipelining",
         "glide": "GLIDE"}


def emit(*lines):
    OUT.extend(lines)


def rate(v):
    return "" if v is None else (f"{v / 1000:,.0f}k" if v >= 10000 else f"{v / 1000:.1f}k")


def msf(us):
    return "" if us is None else f"{us / 1000:.2f} ms"


def table(header, rows):
    emit("| " + " | ".join(header) + " |", "| :--- | " + " | ".join(":-" for _ in header[1:]) + " |")
    for r in rows:
        emit("| " + " | ".join(str(c) for c in r) + " |")
    emit("")


# ---- L0 memtier ----
def l0():
    runs = sorted(glob.glob(str(ROOT / f"memtier_l0/results-{TAG}-run*")))
    if not runs:
        return
    v = defaultdict(lambda: defaultdict(list))
    for r in runs:
        for f in glob.glob(r + "/*.json"):
            t = json.load(open(f))["ALL STATS"]["Totals"]
            p = t.get("Percentile Latencies", {})
            for k, key in (("ops", "Ops/sec"),):
                v[Path(f).stem][k].append(t[key])
            for k, key in (("p50", "p50.00"), ("p99", "p99.00"), ("p999", "p99.90")):
                if key in p:
                    v[Path(f).stem][k].append(p[key])
    emit("## L0: engine baseline (memtier)", "", f"Median of {len(runs)} runs of 60 s per setting.", "")
    rows = []
    for topo in ("standalone", "cluster"):
        for d in (100, 4000):
            for p in (1, 16):
                x = v.get(f"{topo}_d{d}_p{p}")
                if not x:
                    continue
                rows.append([{"standalone": "Single server", "cluster": "Cluster, 3 nodes"}[topo], f"{d:,} B", p,
                             f"{med(x['ops']):,.0f}", f"{med(x['p50']):.2f} ms", f"{med(x['p99']):.2f} ms",
                             f"{med(x['p999']):.2f} ms"])
    table(["Topology", "Value", "Pipeline", "Commands/s", "p50", "p99", "p99.9"], rows)


# ---- L1 README benchmark ----
L1 = re.compile(r"^BenchmarkL1/topo=(\w+)/op=(\w+)/keys=(\w+)/lib=([^/]+)/par=(\d+)/val=\d+-(\d+)\s+\d+\s+([\d.]+) ns/op"
                r"(?:\s+([\d.]+) cpu-ns/op)?")
L1LIBS = {"valkeygo-v1.0.80": "valkey-go", "goredis-v9.23.0": "go-redis",
          "goredis-v9.23.0-autopipeline": "go-redis-autopipeline", "glide-v2.6.1": "glide"}


def l1():
    f = ROOT / f"readme_repro/results_l1_{TAG}/l1.txt"
    if not f.exists():
        return
    ns, cpu, procs = defaultdict(list), defaultdict(list), 10
    for line in f.read_text().splitlines():
        m = L1.match(line)
        if not m:
            continue
        topo, op, keys, lib, par, procs, n, c = m.groups()
        key = (topo, op, keys, L1LIBS.get(lib, lib), int(par))
        ns[key].append(float(n))
        if c:
            cpu[key].append(float(c))
    procs = int(procs)
    runs = max((len(x) for x in ns.values()), default=0)
    emit("## L1: the README benchmark with today's clients", "",
         f"rueidis-benchmark, adapted; 64 B values; median of {runs} runs of 50,000 commands per case. "
         "Throughput in commands per second; CPU is the client process's CPU time per command.", "")
    for topo in ("single", "cluster"):
        for op in ("set", "get"):
            pars = sorted({k[4] for k in ns if k[0] == topo and k[1] == op and k[2] == "uniform"})
            if not pars:
                continue
            rows, crows = [], []
            for p in pars:
                vals = [ns.get((topo, op, "uniform", c, p)) for c in CLIENTS]
                cvals = [cpu.get((topo, op, "uniform", c, p)) for c in CLIENTS]
                rows.append([p * procs] + [rate(1e9 / med(x)) if x else "" for x in vals])
                crows.append([p * procs] + [f"{med(x) / 1000:.1f}" if x else "" for x in cvals])
            title = {"single": "Single server", "cluster": "Cluster, 3 nodes"}[topo]
            emit(f"**{title}, {op.upper()}, 1,024 keys: commands per second**", "")
            table(["Goroutines"] + [NAMES[c] for c in CLIENTS], rows)
            if any(any(r[1:]) for r in crows):
                emit(f"**{title}, {op.upper()}: client CPU per command (µs)**", "")
                table(["Goroutines"] + [NAMES[c] for c in CLIENTS], crows)
    rows = []
    for topo in ("single", "cluster"):
        for keys in ("one", "uniform", "zipf"):
            pars = sorted({k[4] for k in ns if k[0] == topo and k[1] == "set" and k[2] == keys})
            if not pars:
                continue
            p = pars[-1]
            rows.append([f"{topo}, {p * procs} goroutines", keys] +
                        [rate(1e9 / med(ns[(topo, 'set', keys, c, p)])) if (topo, 'set', keys, c, p) in ns else ""
                         for c in CLIENTS])
    emit("**Key shapes: SET at the highest concurrency, commands per second**", "")
    table(["Topology", "Keys"] + [NAMES[c] for c in CLIENTS], rows)


# ---- release sweep, limiter ----
def sweep():
    f = ROOT / f"readme_repro/results-{TAG}/sweep_summary.md"
    if f.exists():
        emit("## Release sweep", "", f.read_text().strip(), "")


def limiter():
    for f in sorted(glob.glob(str(ROOT / f"limiter_mem/results-{TAG}/*.txt"))):
        rows = [l for l in Path(f).read_text().splitlines() if l.startswith("|")]
        if rows:
            emit("## Rate-limiter memory", "", f"`{Path(f).name}`", "", *rows, "")


# ---- latency matrices ----
def mdir(name):
    d = ROOT / f"latency/results/{name}-{TAG}"
    return d if d.exists() else None


def short(name):
    m = re.search(r"load=([^/]+)/(inflight|rate)=(\d+)", name)
    load = {"mix": "80/20 mix", "batch10": "batch of 10 GETs", "mice": "small GETs", "elephants": "small GETs + 1 MB SETs"}
    return load.get(m.group(1), m.group(1)), m.group(2), int(m.group(3))


def clients_table(d, title, cols=CLIENTS, with_cpu=True):
    cl = parse(d / "clients-benchstat.txt")
    if "req/s" not in cl:
        return
    files = cl["req/s"]["cols"]
    idx = [files.index(c) for c in cols if c in files]
    hdr = ["Workload", "Level"] + [NAMES.get(files[i], files[i]) for i in idx]
    rows, crows = [], []
    for name, (medv, _) in cl["req/s"]["rows"].items():
        load, kind, n = short(name)
        p99 = cl.get("p99-us", {}).get("rows", {}).get(name, ([], []))[0]
        cpu = cl.get("cpu-us/req", {}).get("rows", {}).get(name, ([], []))[0]
        lvl = f"{n:,} in flight" if kind == "inflight" else f"{n:,}/s"
        rows.append([load, lvl] + [f"{rate(medv[i])} / {msf(p99[i])}" if i < len(p99) else rate(medv[i]) for i in idx])
        if cpu:
            crows.append([load, lvl] + [f"{cpu[i]:.1f}" for i in idx])
    emit(f"**{title}: requests per second / p99**", "")
    table(hdr, rows)
    if with_cpu and crows:
        emit(f"**{title}: client CPU per request (µs)**", "")
        table(hdr, crows)


def variants_tables(d, title):
    f = d / "variants-benchstat.txt"
    if not f.exists():
        return
    va = parse(f)
    cols = va["req/s"]["cols"][1:]
    for unit, what in (("req/s", "requests per second"), ("p99-us", "p99"), ("cpu-us/req", "client CPU per request")):
        if unit not in va:
            continue
        rows = []
        for name, (medv, deltas) in va[unit]["rows"].items():
            load, kind, n = short(name)
            rows.append([load, f"{n:,}"] + deltas + ["?"] * (len(cols) - len(deltas)))
        emit(f"**{title}: change in {what} against plain valkey-go** (`~` = no real difference)", "")
        table(["Workload", "In flight"] + [c.removeprefix("valkey-go_") for c in cols], rows)


def tls():
    d = mdir("matrix-tls-standalone-100B")
    if not d:
        return
    emit("## TLS off vs on", "")
    clients_table(d, "Without TLS (plain port of the same server)")
    rows = []
    pairs = CLIENTS + ["valkey-go_flush-100us"]
    data = {c: parse(d / f"tls-{c}-benchstat.txt") for c in pairs if (d / f"tls-{c}-benchstat.txt").exists()}
    base = data.get("valkey-go")
    if base:
        for name in base["req/s"]["rows"]:
            load, kind, n = short(name)
            cells = []
            for c in pairs:
                if c not in data:
                    cells.append("")
                    continue
                medv, deltas = data[c]["req/s"]["rows"].get(name, ([], []))
                p99 = data[c].get("p99-us", {}).get("rows", {}).get(name, ([], []))[0]
                cells.append(f"{rate(medv[1])} / {msf(p99[1])} ({deltas[0] if deltas else '?'})" if len(medv) > 1 else "")
            rows.append([load, f"{n:,}"] + cells)
        emit("**With TLS: requests per second / p99 (change in requests per second against the same client without TLS)**", "")
        table(["Workload", "In flight"] + [NAMES.get(c, "valkey-go, 100 µs write delay") for c in pairs], rows)


def cache():
    d = mdir("matrix-cache-standalone-100B")
    if not d:
        return
    hits = defaultdict(list)
    for line in (d / "valkey-go_docache.txt").read_text().splitlines():
        m = re.search(r"load=(\S+)/inflight=(\d+) .* ([\d.]+) cache-hit-pct", line)
        if m:
            hits[(m.group(1), int(m.group(2)))].append(float(m.group(3)))
    va = parse(d / "variants-benchstat.txt")
    rows = []
    for name, (medv, deltas) in va["req/s"]["rows"].items():
        m = re.search(r"load=(hot[\d.]+)/inflight=(\d+)", name)
        load, n = m.group(1), int(m.group(2))
        p99 = va["p99-us"]["rows"][name]
        cpu = va["cpu-us/req"]["rows"][name]
        h = hits.get((load, n))
        rows.append([load.replace("hot", "") + "% target", f"{med(h):.1f}%" if h else "", f"{n:,}",
                     rate(medv[0]), deltas[0] if deltas else "?", p99[1][0] if p99[1] else "?",
                     cpu[1][0] if cpu[1] else "?"])
    emit("## Client-side cache at fixed hit rates (DoCache against Do)", "")
    table(["Hits", "Measured", "In flight", "Do, requests/s", "DoCache: requests/s", "p99", "CPU per request"], rows)


def matrices():
    d = mdir("matrix-standalone-100B")
    if d:
        emit("## L2 and L3: single server, 100 B", "")
        clients_table(d, "Clients")
        variants_tables(d, "valkey-go features off vs on")
    d = mdir("rates-standalone-100B")
    if d:
        emit("## L2: fixed rates, and elephants and mice (single server, 100 B)", "")
        clients_table(d, "Clients at fixed rates", with_cpu=True)
    d = mdir("matrix-standalone-4000B")
    if d:
        emit("## L2: 4 KB values (single server)", "")
        clients_table(d, "Clients, 4 KB")
    d = mdir("matrix-cluster-100B")
    if d:
        emit("## L2: 3-node cluster, 100 B", "")
        clients_table(d, "Clients, cluster")
    cache()
    tls()


# ---- other tests ----
def stale():
    runs = sorted(glob.glob(str(ROOT / f"csc_stale/results-{TAG}/run*.txt")))
    if not runs:
        return
    v = defaultdict(lambda: defaultdict(list))
    order = []
    for f in runs:
        for line in Path(f).read_text().splitlines():
            c = [x.strip() for x in line.split("|")[1:-1]]
            if len(c) == 6 and c[1].isdigit():
                if c[0] not in order:
                    order.append(c[0])
                v[c[0]]["reads"].append(int(c[1]))
                v[c[0]]["stale"].append(float(c[4].rstrip("%")))
                v[c[0]]["lat"].append(c[2])
                v[c[0]]["for"].append(c[5])
    emit("## Stale-read meter", "", f"Median of {len(runs)} runs; latency and stale time from the run with the median stale share.", "")
    rows = []
    for name in order:
        x = v[name]
        i = sorted(range(len(x["stale"])), key=lambda j: x["stale"][j])[len(x["stale"]) // 2]
        rows.append([name, f"{med(x['reads']):,.0f}", x["lat"][i], f"{med(x['stale']):.2f}%", x["for"][i]])
    table(["Client", "Reads/s", "Read p50 / p99", "Stale reads", "Stale for: p50 / p99 / max"], rows)


def fleet():
    for mode in ("standalone", "cluster"):
        f = ROOT / f"fleet/results-{TAG}/fleet_{mode}.txt"
        if f.exists():
            lines = f.read_text().splitlines()
            emit(f"## Cold start and fleet footprint: {mode}", "", lines[0], "", *[l for l in lines if l.startswith("|")], "")


def failover():
    runs = sorted(glob.glob(str(ROOT / f"latency/results/failover-drill-{TAG}-run*")))
    if not runs:
        return
    res = defaultdict(list)
    for r in runs:
        for f in glob.glob(r + "/*.md"):
            client = Path(f).stem.replace("_", "/")
            secs, errors = [], None
            for line in Path(f).read_text().splitlines():
                c = [x.strip() for x in line.split("|")[1:-1]]
                if len(c) == 5 and c[0].isdigit():
                    secs.append((int(c[1]), int(c[2]), int(c[3])))
                elif len(c) == 11 and c[0] == client:
                    errors = int(c[10])
            if not secs:
                continue
            before = med([s[0] for s in secs[1:5]])
            after = [s[0] for s in secs[6:-1]]
            last = secs[-4:-1]  # the last full seconds
            recovered = bool(last) and all(s[1] + s[2] == 0 for s in last) and min(s[0] for s in last) >= 0.8 * before
            res[client].append((before, min(s[0] for s in secs[5:7]), med(after) if after else 0, errors,
                                sum(s[1] for s in secs), sum(s[2] for s in secs), recovered))
    emit("## Failover drill (local)", "", f"{len(runs)} runs per client; medians. Failover 5 s into a 20 s window; 10 requests in flight.", "")
    rows = []
    for client in ["valkey-go", "valkey-go/redirect", "go-redis", "glide"]:
        x = res.get(client)
        if not x:
            continue
        rows.append([client, f"{med([a[0] for a in x]):,.0f}", f"{med([a[1] for a in x]):,.0f}",
                     f"{med([a[2] for a in x]):,.0f}", f"{med([a[4] for a in x]):,.0f}", f"{med([a[5] for a in x]):,.0f}",
                     f"{sum(a[6] for a in x)} of {len(x)}"])
    table(["Client", "Requests/s before", "Slowest second around the failover", "Requests/s after",
           "READONLY errors", "Other errors", "Recovered: no errors and 80% of before, at the end"], rows)


emit(f"# Benchmark results: {TAG}", "")
for part in (l0, l1, matrices, sweep, limiter, stale, fleet, failover):
    try:
        part()
    except Exception as e:  # a missing or partial part shouldn't hide the others
        emit(f"*{part.__name__}: could not summarize ({e})*", "")
print("\n".join(OUT))
