package main

// Golden harness for examples/package-manager, mirroring
// bubbletea-examples/spec/package_manager_parity_spec.cr: three fixed packages,
// Init overridden to nil, memory input/output, no signal handler, NoTTY color
// profile, 80x24, installedPkgMsg at 40ms and q at +40ms.
//
// Run with GODEBUG=randseednop=0: the example's downloadAndInstall uses the
// global math/rand for its simulated delay, and without a working seed the
// golden is not reproducible run to run.

import (
	"os"
	"time"

	"charm.land/bubbles/v2/progress"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"

	"goldens/internal/capture"
)

// parityModel mirrors the spec's PackageManagerParityModel: Init is nil, so no
// download/tick is started at startup.
type parityModel struct{ model }

func (m parityModel) Init() tea.Cmd { return nil }

func goldenModel() tea.Model {
	p := progress.New(
		progress.WithDefaultBlend(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return parityModel{model{
		packages: []string{"alpha-1.0.0", "beta-2.0.0", "gamma-3.0.0"},
		spinner:  s,
		progress: p,
	}}
}

func main() {
	profile := colorprofile.NoTTY
	out := capture.Run(goldenModel(), capture.Options{
		Width:   80,
		Height:  24,
		Profile: &profile,
	}, func(p *tea.Program) {
		capture.Sleep(40 * time.Millisecond)
		p.Send(installedPkgMsg("alpha-1.0.0"))
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
