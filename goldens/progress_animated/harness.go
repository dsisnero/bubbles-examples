package main

// Golden harness for examples/progress-animated, mirroring
// bubbletea-examples/spec/progress_animated_parity_spec.cr: memory input/output,
// no signal handler, 80x24, one tickMsg at 80ms, then q 800ms later.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{
		progress: progress.New(progress.WithDefaultBlend()),
	}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tickMsg(time.Now()))
		capture.Sleep(800 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
