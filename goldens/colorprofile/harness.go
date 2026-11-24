package main

// Golden harness for examples/colorprofile, mirroring
// bubbletea-examples/spec/colorprofile_parity_spec.cr: memory input/output, no
// signal handler, 80x24, TrueColor profile forced, quit after 40ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lucasb-eyer/go-colorful"

	"goldens/internal/capture"
)

func main() {
	myFancyColor, _ = colorful.Hex("#6b50ff")
	profile := colorprofile.TrueColor

	out := capture.Run(model{}, capture.Options{
		Width:   80,
		Height:  24,
		Profile: &profile,
	}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
