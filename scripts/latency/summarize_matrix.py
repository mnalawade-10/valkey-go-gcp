#!/usr/bin/env python3
"""Turn run_matrix.sh's benchstat output into Markdown tables.

  summarize_matrix.py results/matrix-standalone-100B-<tag>

Writes summary.md next to the inputs:
- clients: median requests/s and p99 for each client, per workload and requests in flight;
- variants: each valkey-go variant against plain valkey-go, as a % change in requests/s, p99 and CPU
  per request. benchstat's test decides: "~" means no real difference at its default alpha (0.05).
"""
import re
import sys
from pathlib import Path

SUFFIX = {"": 1, "k": 1e3, "M": 1e6, "G": 1e9}
NUM = r"(\d+(?:\.\d+)?)([kMG]?)"


def parse(path):
    """benchstat text -> {unit: {"cols": [file names], "rows": {name: (medians, deltas)}}}."""
    out, cols, unit = {}, [], None
    for line in Path(path).read_text().splitlines():
        if "│" in line:
            cells = [c.strip() for c in line.split("│")[1:-1]]
            if any(c.endswith(".txt") for c in cells):
                cols = [c.removesuffix(".txt") for c in cells]
            elif cells and cells[0]:
                unit = cells[0].split()[0]
                out.setdefault(unit, {"cols": cols, "rows": {}})
            continue
        if unit is None or "±" not in line:
            continue
        name, rest = line.split(None, 1)
        medians = [float(v) * SUFFIX[s] for v, s in re.findall(NUM + r"\s*±", rest)]
        deltas = re.findall(r"(~|[+-]\d+(?:\.\d+)?%)\s*\(p=", rest)
        out[unit]["rows"][name] = (medians, deltas)
    return out


def short(name):
    m = re.search(r"load=([^/]+)/inflight=(\d+)", name)
    load = {"mix": "80/20 mix", "batch10": "batch of 10 GETs"}.get(m.group(1), m.group(1))
    return load, int(m.group(2))


def fmt_rate(v):
    return f"{v / 1000:,.0f}k" if v >= 10000 else f"{v / 1000:.1f}k"


def main():
    d = Path(sys.argv[1])
    lines = [f"# Off/on matrix and client comparison: {d.name}", ""]
    env = (d / "env.txt").read_text().strip().splitlines() if (d / "env.txt").exists() else []
    if env:
        lines += ["```text", *env, "```", ""]

    cl = parse(d / "clients-benchstat.txt")
    if "req/s" in cl:
        cols = cl["req/s"]["cols"]
        lines += ["## Clients: median requests per second, and p99 (ms)", "",
                  "| Workload | In flight | " + " | ".join(cols) + " |",
                  "| :--- | :- | " + " | ".join(":-" for _ in cols) + " |"]
        for name, (med, _) in cl["req/s"]["rows"].items():
            load, n = short(name)
            p99 = cl.get("p99-us", {}).get("rows", {}).get(name, ([], []))[0]
            cells = [f"{fmt_rate(r)} / {p / 1000:.2f}" if i < len(p99) else fmt_rate(r)
                     for i, (r, p) in enumerate(zip(med, p99 or [0] * len(med)))]
            lines.append(f"| {load} | {n:,} | " + " | ".join(cells) + " |")
        lines.append("")

    vf = d / "variants-benchstat.txt"
    if vf.exists():
        va = parse(vf)
        cols = va["req/s"]["cols"][1:]
        rows = list(va["req/s"]["rows"])
        lines += ["## valkey-go variants against plain valkey-go", "",
                  "Change in requests/s, then p99, then client CPU per request. `~` = no real difference (benchstat).", ""]
        for unit, title in (("req/s", "Requests per second"), ("p99-us", "p99 latency"), ("cpu-us/req", "Client CPU per request")):
            if unit not in va:
                continue
            lines += [f"**{title}**", "",
                      "| Workload | In flight | valkey-go | " + " | ".join(c.removeprefix("valkey-go_") for c in cols) + " |",
                      "| :--- | :- | :- | " + " | ".join(":-" for _ in cols) + " |"]
            for name in rows:
                med, deltas = va[unit]["rows"].get(name, ([], []))
                load, n = short(name)
                base = (fmt_rate(med[0]) if unit == "req/s" else
                        f"{med[0] / 1000:.2f} ms" if unit == "p99-us" else f"{med[0]:.1f} µs") if med else "?"
                lines.append(f"| {load} | {n:,} | {base} | " + " | ".join(deltas + ["?"] * (len(cols) - len(deltas))) + " |")
            lines.append("")
    (d / "summary.md").write_text("\n".join(lines))
    print("\n".join(lines))


if __name__ == "__main__":
    main()
