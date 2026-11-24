package main

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	out := capture.Run(initialModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(240 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
