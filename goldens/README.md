# Go golden harnesses

`spec/golden/*.go.golden` are byte captures of the upstream Go examples in
`../../vendor/bubbletea/examples` at the pinned revision. They are the oracle
for `spec/*_parity_spec.cr`: the Crystal example must produce exactly these
bytes.

**Never regenerate a golden from the Crystal port.** A golden written from this
repository's own output would make the parity spec tautological.

## Layout

Each `<example>/` package holds:

- `harness.go` — hand-written: builds the example's model, replays the same
  script as `spec/<example>_parity_spec.cr`, and writes the capture to stdout.
- `model.go` (plus any sibling files such as `words.go` or `packages.go`) —
  generated from the pinned Go example by `generate.py`. `func main()` is
  renamed to `func exampleMain()` so the example's imports stay used.

`internal/capture` is the shared runner: memory input/output, no signal
handler, fixed window size, optional colour profile and environment, then a
scripted message sequence.

`generate.py` re-materialises the model files and rewrites the golden:

```bash
cd bubbletea-examples/goldens
python3 generate.py simple mouse views
```

Run it after bumping `vendor/bubbletea`, then review the golden diff. The
harness module needs Go and resolves `charm.land/bubbletea/v2` through
`replace ... => ../../vendor/bubbletea`.

## Harness notes

Most harnesses call `capture.Run(model, opts, script)` with the expression the
Go example's own `main()` builds. A few need more:

- `autocomplete` — the upstream example fetches its repo list over HTTP, so the
  harness wraps the model and replaces only `Init` with a fixed list, matching
  the Crystal spec.
- `cellbuffer`, `splash`, `stopwatch`, `timer` — wrap the model so `Init`
  returns nil, mirroring the Crystal specs' subclasses.
- `table` — the spec builds its own model (5 rows, height 7, focused), so the
  harness reproduces that model in Go rather than the upstream example's.
- `focus_blur` — builds the program inline because it needs a genuinely empty
  environment (`capture.Options.SetEnv`).
- `package_manager` and `list_fancy_h_a` — depend on the global RNG; Go's
  `rand.Seed` is a no-op unless `GODEBUG=randseednop=0`, and their harness
  headers say so.

## What is not asserted

`spec/textarea_parity_spec.cr` and `spec/paginator_parity_spec.cr` capture a
variable number of renderer frames (the former is not reproducible in Go
either), and `file_picker`, `focus_blur` and `set_terminal_color` hit known
differences in the `bubbles` and `ultraviolet` ports. Those specs are `pending`
with the reason recorded inline.
