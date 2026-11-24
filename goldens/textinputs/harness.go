package main

// Golden harness for examples/textinputs, mirroring
// bubbletea-examples/spec/textinputs_parity_spec.cr: memory input/output, no
// signal handler, 80x24, "hi" typed into the first (focused) input, then tab to
// the next inputs while typing, a ctrl+r cursor-mode cycle, then tab to the
// submit button and enter to quit.
//
// Init is nil so the example's textinput.Blink tick never fires: a blink tick
// changes the virtual cursor's cell style, and its phase against the renderer's
// frame ticks makes the capture depend on scheduling. Every phase below is
// separated by a sleep longer than a frame for the same reason — a burst of
// keys otherwise coalesces into a variable number of frames.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func main() {
	out := capture.Run(parityModel{initialModel()}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(100 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'h', Text: "h"})
		p.Send(tea.KeyPressMsg{Code: 'i', Text: "i"})
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})
		capture.Sleep(80 * time.Millisecond)

		p.Send(tea.KeyPressMsg{Code: 'a', Text: "a"})
		p.Send(tea.KeyPressMsg{Code: '@', Text: "@"})
		p.Send(tea.KeyPressMsg{Code: 'b', Text: "b"})
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})
		capture.Sleep(80 * time.Millisecond)

		p.Send(tea.KeyPressMsg{Code: '1', Text: "1"})
		p.Send(tea.KeyPressMsg{Code: '2', Text: "2"})
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'r', Mod: tea.ModCtrl})
		capture.Sleep(80 * time.Millisecond)

		p.Send(tea.KeyPressMsg{Code: tea.KeyTab})
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyEnter})
	})
	os.Stdout.Write(out)
}
