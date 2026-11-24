package main

// Golden harness for examples/progress-download, mirroring
// bubbletea-examples/spec/progress_download_parity_spec.cr: no real download
// (pw nil, Init nil), memory input/output, no signal handler, 80x24,
// progressMsg(0.42) at 40ms and q at +40ms.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// parityModel mirrors the spec's ProgressDownloadParityModel: Init is nil.
type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func main() {
	out := capture.Run(parityModel{model{
		progress: progress.New(progress.WithDefaultBlend()),
	}}, capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(progressMsg(0.42))
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
