package main

// Golden harness for examples/result, mirroring
// bubbletea-examples/spec/result_parity_spec.cr: memory input/output, no signal
// handler, 80x24, down at 100ms and enter at +40ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(100 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
	})
	os.Stdout.Write(out)
}
