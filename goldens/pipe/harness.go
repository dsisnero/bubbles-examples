package main

// Golden harness for examples/pipe, mirroring
// bubbletea-examples/spec/pipe_parity_spec.cr: model built from the literal
// "artichoke" (the spec does not read stdin), memory input/output, no signal
// handler, 80x24, Enter at 120ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	return newModel("artichoke")
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
	})
	os.Stdout.Write(out)
}
