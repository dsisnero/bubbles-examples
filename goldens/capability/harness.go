package main

// Golden harness for examples/capability, mirroring
// bubbletea-examples/spec/capability_parity_spec.cr: memory input/output, no
// signal handler, 80x24, quit after 40ms. The main() setup lines are folded
// into goldenModel().

import (
	"os"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	m := model{}
	m.input = textinput.New()
	m.input.Placeholder = "Enter capability name to request"
	m.input.Focus()
	return m
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
