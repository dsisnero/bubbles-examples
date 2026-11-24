package main

// Golden harness for examples/mouse, mirroring
// bubbletea-examples/spec/mouse_parity_spec.cr:
// memory input/output, no signal handler, 80x24, quit after 60ms.

import (
	"bytes"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

func main() {
	var in bytes.Buffer
	var out bytes.Buffer

	p := tea.NewProgram(model{},
		tea.WithInput(&in),
		tea.WithOutput(&out),
		tea.WithoutSignalHandler(),
		tea.WithWindowSize(80, 24),
	)

	go func() {
		time.Sleep(60 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	}()

	if _, err := p.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(out.Bytes())
}
