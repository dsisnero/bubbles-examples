package main

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func main() {
	// The Crystal spec re-derives the program environment from the golden's
	// EnvMsg line, so any stable explicit environment works here.
	out := capture.Run(newModel(), capture.Options{
		Width:  80,
		Height: 24,
		Env:    []string{"TERM=xterm-256color"},
		SetEnv: true,
	}, func(p *tea.Program) {
		capture.Sleep(60 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
