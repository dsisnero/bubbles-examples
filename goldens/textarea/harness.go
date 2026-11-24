package main

// Golden harness for examples/textarea, mirroring
// bubbletea-examples/spec/textarea_parity_spec.cr: the Crystal spec subclasses
// TextareaModel and turns the virtual cursor ON, uses no input, memory output,
// a TERM=xterm-256color environment, no signal handler, 60x12, then types four
// lines (enter between them) and quits 40ms later.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	m := initialModel()
	m.textarea.SetVirtualCursor(true) // TextareaParityModel#initialize
	return m
}

func main() {
	lines := []string{
		"Once upon a time,",
		"there were three bears.",
		"Papa Bear made porridge.",
		"Baby Bear took a walk.",
	}

	out := capture.Run(goldenModel(), capture.Options{
		NoInput: true,
		Width:   60,
		Height:  12,
		Env:     []string{"TERM=xterm-256color"},
	}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		for i, line := range lines {
			for _, r := range line {
				p.Send(tea.KeyPressMsg{Code: r, Text: string(r)})
			}
			if i != len(lines)-1 {
				p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
			}
		}
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
