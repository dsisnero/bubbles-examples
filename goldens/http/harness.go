package main

// Golden harness for examples/http, mirroring
// bubbletea-examples/spec/http_parity_spec.cr: the spec model overrides Init to
// return nil (so no real HTTP request is made), memory input/output, no signal
// handler, 80x24, and a synthetic statusMsg(200) at 40ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// parityModel mirrors the spec's HttpParityModel: Init does not issue the real
// network request.
type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func main() {
	out := capture.Run(parityModel{model{}}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(statusMsg(200))
	})
	os.Stdout.Write(out)
}
