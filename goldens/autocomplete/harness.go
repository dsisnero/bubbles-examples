package main

// Harness for examples/autocomplete.
//
// The upstream example fetches its repo list over HTTP, which is not
// reproducible. The Crystal parity spec replaces that with a fixed list, so
// this harness wraps the pinned Go model and only swaps Init to deliver the
// same list (the rest of the model, keymap and view are the upstream ones).

import (
	"os"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

type fixedRepos struct{ inner model }

func (m fixedRepos) Init() tea.Cmd {
	return tea.Batch(
		func() tea.Msg {
			return gotReposSuccessMsg{
				{Name: "bubbletea"},
				{Name: "bubbles"},
				{Name: "lipgloss"},
			}
		},
		textinput.Blink,
	)
}

func (m fixedRepos) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	inner, cmd := m.inner.Update(msg)
	if next, ok := inner.(model); ok {
		return fixedRepos{inner: next}, cmd
	}
	return m, cmd
}

func (m fixedRepos) View() tea.View { return m.inner.View() }

func main() {
	out := capture.Run(fixedRepos{inner: initialModel()}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'b', Text: "b"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'u', Text: "u"})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
	})
	os.Stdout.Write(out)
}
