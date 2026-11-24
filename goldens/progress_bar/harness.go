package main

// Golden harness for examples/progress-bar, mirroring
// bubbletea-examples/spec/progress_bar_parity_spec.cr: memory input/output, no signal
// handler, 80x24, quit after 120ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{value: 50, state: tea.ProgressBarIndeterminate}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
