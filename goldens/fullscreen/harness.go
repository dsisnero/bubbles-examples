package main

// Golden harness for examples/fullscreen, mirroring
// bubbletea-examples/spec/fullscreen_parity_spec.cr: memory input/output, no
// signal handler, 80x24, 'q' at 80ms (before the first one second tick).

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model(5), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
