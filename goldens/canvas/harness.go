package main

// Golden harness for examples/canvas, mirroring
// bubbletea-examples/spec/canvas_parity_spec.cr: memory input/output, no
// signal handler, 80x24, key 'x' at 80ms then 'q' at 120ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'x', Text: "x"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
