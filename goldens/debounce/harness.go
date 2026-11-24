package main

// Golden harness for examples/debounce, mirroring
// bubbletea-examples/spec/debounce_parity_spec.cr: memory input/output, no
// signal handler, 80x24, key 'a' at 80ms, then the model quits itself once the
// one second debounce tick with the matching tag arrives.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(model{}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'a', Text: "a"})
	})
	os.Stdout.Write(out)
}
