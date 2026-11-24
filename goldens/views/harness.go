package main

// Golden harness for examples/views, mirroring
// bubbletea-examples/spec/views_parity_spec.cr: memory input/output, no signal
// handler, 80x24, down at 80ms, enter at +20ms, six frameMsgs, then q.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{0, false, 10, 0, 0, false, false}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(20 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		for range 6 {
			p.Send(frameMsg{})
		}
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
