package main

// Golden harness for examples/eyes, mirroring
// bubbletea-examples/spec/eyes_parity_spec.cr: memory input/output, no signal
// handler, 80x24, quit after 120ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
