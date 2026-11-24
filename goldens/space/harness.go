package main

// Golden harness for examples/space, mirroring
// bubbletea-examples/spec/space_parity_spec.cr: 40x12, NoTTY color profile,
// 'q' at 120ms. The Crystal model is constructed with a fixed RNG seed
// (Random.new(1)), so the Go model seeds math/rand with 1 to keep the color
// grid reproducible.
//
// NOTE: this spec is marked `pending` upstream (high-volume styled redraw
// stream drifts from Go), so this golden is recorded but not asserted.

import (
	"math/rand"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	rand.Seed(1) //nolint:staticcheck // match Random.new(1) in the parity spec
	return model{}
}

func main() {
	profile := colorprofile.NoTTY
	out := capture.Run(goldenModel(), capture.Options{
		Width:   40,
		Height:  12,
		Profile: &profile,
	}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
