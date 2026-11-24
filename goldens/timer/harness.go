package main

// Golden harness for examples/timer, mirroring
// bubbletea-examples/spec/timer_parity_spec.cr: the Crystal spec subclasses
// TimerModel, uses a one-hour tick interval (so no tick fires during the
// capture) and overrides init to return nil. Memory input/output, no signal
// handler, 80x24, then 's', 's', 'r', 'q' with 40ms between sends.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// goldenModel mirrors TimerParityModel's `init` override (nil instead of
// m.timer.Init()).
type goldenModel struct {
	model
}

func (goldenModel) Init() tea.Cmd { return nil }

func newGoldenModel() goldenModel {
	m := model{
		timer: timer.New(timeout, timer.WithInterval(time.Hour)),
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
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
		help: help.New(),
	}
	m.keymap.start.SetEnabled(false)
	return goldenModel{m}
}

func main() {
	out := capture.Run(newGoldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(100 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 's', Text: "s"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 's', Text: "s"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'r', Text: "r"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
