package main

// Golden harness for examples/dynamic-textarea, mirroring
// bubbletea-examples/spec/dynamic_textarea_parity_spec.cr. The example only
// quits on ctrl+c.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(100 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'a', Text: "a"})
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		p.Send(tea.KeyPressMsg{Code: 'b', Text: "b"})
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
		p.Send(tea.KeyPressMsg{Code: 'c', Text: "c"})
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
