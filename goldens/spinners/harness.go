package main

// Golden harness for examples/spinners, mirroring
// bubbletea-examples/spec/spinners_parity_spec.cr: setup copied from the
// example's main(), memory input/output, no signal handler, 80x24, right arrow
// at 80ms then 'q' at 100ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	m := model{}
	m.resetSpinner()
	return m
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyRight})
		capture.Sleep(20 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
