package main

// Golden harness for examples/list-fancy, mirroring
// bubbletea-examples/spec/list_fancy_h_a_parity_spec.cr: RNG seeded with 1,
// memory input/output, no signal handler, 80x24, "H" at 120ms, "a" at +40ms and
// q at +40ms.
//
// Run with GODEBUG=randseednop=0: math/rand.Seed is a no-op by default on
// current Go, and without it the shuffled item order changes every run.

import (
	"math/rand"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	// Mirrors RandomFancyItemGenerator.seed(1_i64) in the Crystal spec.
	rand.Seed(1)

	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'H', Text: "H"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'a', Text: "a"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
