package main

// Golden harness for examples/realtime, mirroring
// bubbletea-examples/spec/realtime_parity_spec.cr: Init overridden to nil (no
// activity generator goroutine), memory input/output, no signal handler, 80x24,
// a spinner tick at 80ms, a responseMsg at +20ms and q at +20ms.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// parityModel mirrors the spec's RealtimeParityModel: Init is nil.
type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func main() {
	m := model{
		sub:     make(chan struct{}),
		spinner: spinner.New(),
	}

	out := capture.Run(parityModel{m}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(m.spinner.Tick())
		capture.Sleep(20 * time.Millisecond)
		p.Send(responseMsg{})
		capture.Sleep(20 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
