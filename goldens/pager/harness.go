package main

// Golden harness for examples/pager, mirroring
// bubbletea-examples/spec/pager_parity_spec.cr: the Crystal model loads
// src/artichoke.md at construction, so the Go model loads the pinned upstream
// copy of the same file. Memory input/output, no signal handler, 80x24, 'j' at
// 120ms then QuitMsg at 140ms.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	content, err := os.ReadFile("../../vendor/bubbletea/examples/pager/artichoke.md")
	if err != nil {
		panic(err)
	}
	return model{content: string(content)}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'j', Text: "j"})
		capture.Sleep(20 * time.Millisecond)
		p.Send(tea.QuitMsg{})
	})
	os.Stdout.Write(out)
}
