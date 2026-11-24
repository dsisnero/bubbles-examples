package main

// Golden harness for examples/splash, mirroring
// bubbletea-examples/spec/splash_parity_spec.cr: the Crystal spec constructs the
// model with rate 0 and overrides Init to return nil, so no ticks arrive and the
// gradient is deterministic. Memory input/output, no signal handler, 40x10, 'q'
// at 60ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

type nilInit struct{ tea.Model }

func (n nilInit) Init() tea.Cmd { return nil }

func goldenModel() tea.Model {
	return nilInit{model{rate: 0}}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 40, Height: 10}, func(p *tea.Program) {
		capture.Sleep(60 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
