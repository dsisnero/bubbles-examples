package main

// Golden harness for examples/progress-static, mirroring
// bubbletea-examples/spec/progress_static_parity_spec.cr: memory input/output,
// no signal handler, 80x24, Init overridden to nil, then five tickMsgs 20ms
// apart starting at 60ms.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// parityModel mirrors the spec's model, whose Init returns nil.
type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func main() {
	prog := progress.New(progress.WithScaled(true), progress.WithColors(pink, yellow))

	out := capture.Run(parityModel{model{progress: prog}}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(60 * time.Millisecond)
		for range 5 {
			p.Send(tickMsg(time.Now()))
			capture.Sleep(20 * time.Millisecond)
		}
	})
	os.Stdout.Write(out)
}
