package main

// Golden harness for examples/table, mirroring both captures in
// bubbletea-examples/spec/table_parity_spec.cr. The spec builds its own model:
// the same columns/styles as the example's main(), but only the first five rows
// and no explicit width, so the harness reproduces that model.
//
//   go run ./goldens/table plain   -> golden/table.go.golden
//   go run ./goldens/table arrows  -> golden/table_arrows_esc.go.golden

import (
	"os"
	"time"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

func goldenModel() tea.Model {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "Sao Paulo", "Brazil", "22,429,800"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	t.SetStyles(table.DefaultStyles())

	return model{t}
}

func main() {
	script := "plain"
	if len(os.Args) > 1 {
		script = os.Args[1]
	}

	mode := os.Getenv("GOLDEN_SCRIPT")
	if mode != "" {
		script = mode
	}

	var run func(p *tea.Program)
	switch script {
	case "plain":
		run = func(p *tea.Program) {
			capture.Sleep(80 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: 'j', Text: "j"})
			p.Send(tea.KeyPressMsg{Code: 'j', Text: "j"})
			p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
		}
	case "arrows":
		run = func(p *tea.Program) {
			capture.Sleep(80 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
			capture.Sleep(40 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: tea.KeyEsc})
			capture.Sleep(40 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
			capture.Sleep(40 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: tea.KeyEsc})
			capture.Sleep(40 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: tea.KeyUp})
			capture.Sleep(40 * time.Millisecond)
			p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
		}
	default:
		panic("unknown script: " + script)
	}

	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, run)
	os.Stdout.Write(out)
}
