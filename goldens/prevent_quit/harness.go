package main

// Golden harness for examples/prevent-quit, mirroring
// bubbletea-examples/spec/prevent_quit_parity_spec.cr: memory input/output, no
// signal handler, 80x24, the example's filter installed, h at 120ms, i at
// +40ms, esc at +40ms and y at +60ms. Nothing sends a QuitMsg directly; the
// model quits itself.

import (
	"bytes"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

func main() {
	var in bytes.Buffer
	var out bytes.Buffer

	p := tea.NewProgram(initialModel(),
		tea.WithInput(&in),
		tea.WithOutput(&out),
		tea.WithoutSignalHandler(),
		tea.WithWindowSize(80, 24),
		tea.WithFilter(filter),
	)

	go func() {
		time.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'h', Text: "h"})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'i', Text: "i"})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEsc})
		time.Sleep(60 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'y', Text: "y"})
	}()

	if _, err := p.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(out.Bytes())
}
