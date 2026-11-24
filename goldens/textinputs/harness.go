package main

// Golden harness for examples/textinputs, mirroring
// bubbletea-examples/spec/textinputs_parity_spec.cr: memory input/output, no
// signal handler, 80x24, "hi" typed into the first (focused) input, tab to the
// next inputs while typing, a ctrl+r cursor-mode cycle, then tab to the submit
// button and enter to quit.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(100 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'h', Text: "h"})
		p.Send(tea.KeyPressMsg{Code: 'i', Text: "i"})
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})

		p.Send(tea.KeyPressMsg{Code: 'a', Text: "a"})
		p.Send(tea.KeyPressMsg{Code: '@', Text: "@"})
		p.Send(tea.KeyPressMsg{Code: 'b', Text: "b"})
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})

		p.Send(tea.KeyPressMsg{Code: '1', Text: "1"})
		p.Send(tea.KeyPressMsg{Code: '2', Text: "2"})
		p.Send(tea.KeyPressMsg{Code: 'r', Mod: tea.ModCtrl})

		capture.Sleep(60 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
	})
	os.Stdout.Write(out)
}
