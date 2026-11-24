package main

// Golden harness for examples/chat, mirroring
// bubbletea-examples/spec/chat_parity_spec.cr: no input, memory output, no signal
// handler, 30x8, then 'h', 'i', enter, quit after 50ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{
		NoInput: true,
		Width:   30,
		Height:  8,
	}, func(p *tea.Program) {
		capture.Sleep(150 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'h', Text: "h"})
		p.Send(tea.KeyPressMsg{Code: 'i', Text: "i"})
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		capture.Sleep(50 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
