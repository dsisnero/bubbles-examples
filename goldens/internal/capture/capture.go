// Package capture runs a bubbletea example model the way the Crystal parity
// specs do: memory input/output, no signal handler, a fixed window size and a
// scripted message sequence, returning the raw output bytes.
package capture

import (
	"bytes"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
)

// Options mirror the Tea program options used by the Crystal specs.
type Options struct {
	Input   string // initial bytes available on stdin
	NoInput bool   // Tea.with_input(nil)
	Width   int
	Height  int
	Profile *colorprofile.Profile // optional with_color_profile
	Env     []string              // with_environment; empty slice means empty env
	// SetEnv forwards Env even when it is empty, which the focus_blur spec
	// needs: an empty environment changes Go's terminal capability set.
	SetEnv bool
}

// Run starts the program, runs script against it, and returns the output bytes.
func Run(m tea.Model, o Options, script func(*tea.Program)) []byte {
	var in bytes.Buffer
	in.WriteString(o.Input)

	var out bytes.Buffer

	opts := []tea.ProgramOption{
		tea.WithOutput(&out),
		tea.WithoutSignalHandler(),
	}
	if o.NoInput {
		opts = append(opts, tea.WithInput(nil))
	} else {
		opts = append(opts, tea.WithInput(&in))
	}
	if o.Width > 0 {
		opts = append(opts, tea.WithWindowSize(o.Width, o.Height))
	}
	if o.Profile != nil {
		opts = append(opts, tea.WithColorProfile(*o.Profile))
	}
	if o.SetEnv || len(o.Env) > 0 {
		opts = append(opts, tea.WithEnvironment(o.Env))
	}

	p := tea.NewProgram(m, opts...)

	if script != nil {
		go script(p)
	}

	if _, err := p.Run(); err != nil {
		panic(err)
	}

	return out.Bytes()
}

// Sleep pauses the script for the given duration.
func Sleep(d time.Duration) { time.Sleep(d) }
