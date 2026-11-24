package main

// Golden harness for examples/tabs, mirroring
// bubbletea-examples/spec/tabs_parity_spec.cr: setup copied from the example's
// main() (dark styles), memory input/output, no signal handler, 80x24, right
// arrow at 80ms, left arrow at 120ms, 'q' at 160ms. The spec reads
// golden/tabs_arrows.go.golden.

import (
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	tabs := []string{"Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"}
	tabContent := []string{"Lip Gloss Tab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}
	return model{Tabs: tabs, TabContent: tabContent, styles: newStyles(true)}
}

func main() {
	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyRight})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyLeft})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
