package main

// Golden harness for examples/stopwatch, mirroring
// bubbletea-examples/spec/stopwatch_parity_spec.cr: setup copied from the
// example's main(), Init overridden to return nil (as the Crystal spec's
// StopwatchParityModel does), memory input/output, no signal handler, 80x24,
// 'q' at 40ms.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/stopwatch"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

type nilInit struct{ tea.Model }

func (n nilInit) Init() tea.Cmd { return nil }

func goldenModel() tea.Model {
	m := model{
		stopwatch: stopwatch.New(stopwatch.WithInterval(time.Millisecond)),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "q"),
				key.WithHelp("q", "quit"),
			),
		},
		help: help.New(),
	}

	m.keymap.start.SetEnabled(false)
	return nilInit{m}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
