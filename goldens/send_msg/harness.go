package main

// Golden harness for examples/send-msg, mirroring
// bubbletea-examples/spec/send_msg_parity_spec.cr: memory input/output, no
// signal handler, 80x24, two resultMsgs at 140ms and +80ms, then q.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(newModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(140 * time.Millisecond)
		p.Send(resultMsg{duration: 120 * time.Millisecond, food: "an apple"})
		capture.Sleep(80 * time.Millisecond)
		p.Send(resultMsg{duration: 250 * time.Millisecond, food: "tacos"})
		capture.Sleep(1 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
