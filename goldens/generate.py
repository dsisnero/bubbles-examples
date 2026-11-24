#!/usr/bin/env python3
"""Regenerate a Go golden from the pinned upstream example.

Run from anywhere; paths are resolved from this file:

    bubbletea-examples/goldens/generate.py   <- this script
    bubbletea-examples/goldens/<name>/harness.go
    bubbletea-examples/spec/golden/<name>.go.golden
    ../vendor/bubbletea/examples/<go-dir>/*.go

The harness's `model.go` (plus any sibling files the example needs) is
re-materialised from the pinned Go sources on every run, so a pin bump only
needs `python3 goldens/generate.py <names>` again. `func main()` is renamed to
`func exampleMain()` so the example's imports stay used.

Usage:
    python3 goldens/generate.py simple mouse views
"""

import subprocess
import sys
from pathlib import Path

GOLDENS_DIR = Path(__file__).resolve().parent
EXAMPLES_REPO = GOLDENS_DIR.parent
UPSTREAM_EXAMPLES = EXAMPLES_REPO.parent / "vendor" / "bubbletea" / "examples"
GOLDEN_DIR = EXAMPLES_REPO / "spec" / "golden"


def rename_main(src: str) -> str:
    needle = "func main("
    if needle not in src:
        raise SystemExit("no func main in example source")
    return src.replace(needle, "func exampleMain(", 1)


def go_dir_for(name: str) -> Path:
    """Map a Crystal example name to its Go directory."""
    aliases = {"list_fancy_h_a": "list-fancy"}
    candidates = [aliases.get(name, name), name, name.replace("_", "-")]
    for candidate in candidates:
        path = UPSTREAM_EXAMPLES / candidate
        if path.is_dir():
            return path
    raise SystemExit(f"no Go example for {name}")


def main() -> None:
    names = sys.argv[1:]
    if not names:
        raise SystemExit(__doc__)

    for name in names:
        package = GOLDENS_DIR / name
        harness = package / "harness.go"
        if not harness.is_file():
            raise SystemExit(f"{harness} is missing; write the harness first")

        example = go_dir_for(name)
        for source in sorted(example.glob("*.go")):
            if source.name.endswith("_test.go"):
                continue
            text = source.read_text()
            if "func main(" in text:
                text = rename_main(text)
            # main.go becomes model.go so it never shadows the harness's main.
            target = "model.go" if source.name == "main.go" else source.name
            (package / target).write_text(text)
            (package / "main.go").unlink(missing_ok=True)

        result = subprocess.run(
            ["go", "run", f"./{name}"],
            cwd=GOLDENS_DIR,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )
        if result.returncode != 0:
            raise SystemExit(
                f"[golden] {name}: harness failed\n{result.stderr.decode(errors='replace')}"
            )
        target = GOLDEN_DIR / f"{name}.go.golden"
        target.write_bytes(result.stdout)
        print(f"[golden] {name}: {len(result.stdout)} bytes -> {target}")


if __name__ == "__main__":
    main()
