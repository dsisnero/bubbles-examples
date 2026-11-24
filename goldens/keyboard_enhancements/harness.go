package main

// Golden harness for examples/keyboard-enhancements, mirroring
// bubbletea-examples/spec/keyboard_enhancements_parity_spec.cr: memory input/output, no signal
// handler, 80x24, quit after 80ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
