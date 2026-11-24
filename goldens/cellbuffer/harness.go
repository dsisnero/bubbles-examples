package main

// Golden harness for examples/cellbuffer, mirroring
// bubbletea-examples/spec/cellbuffer_parity_spec.cr: memory input/output, no
// signal handler, 80x24, and Init overridden to return nil (the Crystal spec
// subclasses the model for the same reason), quit on 'q' at 60ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/harmonica"

	"goldens/internal/capture"
)

// nilInit delegates to the wrapped model but never starts the animation ticker,
// like the parity spec's CellBufferParityModel.
type nilInit struct{ tea.Model }

func (n nilInit) Init() tea.Cmd { return nil }

func goldenModel() tea.Model {
	m := model{
		spring: harmonica.NewSpring(harmonica.FPS(fps), frequency, damping),
	}
	return nilInit{m}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(60 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
