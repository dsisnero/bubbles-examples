package main

// Golden harness for examples/set-terminal-color, mirroring
// bubbletea-examples/spec/set_terminal_color_parity_spec.cr (whose model is a
// transcription of this example): memory input/output, no signal handler,
// 80x24, then set foreground, background and cursor colors with the spec's
// delays, and quit.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	ti := textinput.New()
	ti.Placeholder = "#ff00ff"
	ti.CharLimit = 156
	ti.SetWidth(20)
	ti.SetVirtualCursor(false)
	return model{
		ti: ti,
	}
}

func sendChars(p *tea.Program, s string) {
	for _, r := range s {
		p.Send(tea.KeyPressMsg{Code: r, Text: string(r)})
	}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		// Foreground -> violet
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		capture.Sleep(40 * time.Millisecond)
		sendChars(p, "#6b50ff")
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})

		// Background -> green
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		capture.Sleep(40 * time.Millisecond)
		sendChars(p, "#00ff00")
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})

		// Cursor -> gray
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		capture.Sleep(40 * time.Millisecond)
		sendChars(p, "#808080")
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
