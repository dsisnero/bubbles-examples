package main

// Golden harness for examples/list-simple, mirroring
// bubbletea-examples/spec/list_simple_parity_spec.cr: memory input/output, no signal
// handler, 80x24, quit after 60ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(60 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
