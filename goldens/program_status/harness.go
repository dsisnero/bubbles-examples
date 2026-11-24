package main

// Golden harness for examples/program-status, mirroring
// bubbletea-examples/spec/program_status_parity_spec.cr.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyRight})
		capture.Sleep(440 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyLeft})
		capture.Sleep(200 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
