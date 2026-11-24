package main

// Golden harness for examples/focus-blur, mirroring
// bubbletea-examples/spec/focus_blur_parity_spec.cr: memory input/output, no
// signal handler, an EMPTY environment (Tea.with_environment(Environ.new([]))),
// 80x24, then blur, 't', 't', focus, 't', 'q' with 40ms between sends.
//
// The shared capture helper only forwards a non-empty Env, and an empty
// environment is not the same as "use os.Environ()", so the program is built
// here directly.

import (
	"bytes"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

func main() {
	var in bytes.Buffer
	var out bytes.Buffer

	p := tea.NewProgram(model{focused: true, reporting: true},
		tea.WithInput(&in),
		tea.WithOutput(&out),
		tea.WithoutSignalHandler(),
		tea.WithEnvironment([]string{}),
		tea.WithWindowSize(80, 24),
	)

	go func() {
		time.Sleep(80 * time.Millisecond)
		p.Send(tea.BlurMsg{})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 't', Text: "t"})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 't', Text: "t"})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.FocusMsg{})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 't', Text: "t"})
		time.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	}()

	if _, err := p.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(out.Bytes())
}
