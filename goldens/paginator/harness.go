package main

// Golden harness for examples/paginator, mirroring
// bubbletea-examples/spec/paginator_parity_spec.cr: memory input/output, no signal
// handler, 80x24, quit after 80ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(newModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
