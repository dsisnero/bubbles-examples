package main

// Golden harness for examples/file-picker, mirroring
// bubbletea-examples/spec/file_picker_parity_spec.cr: the spec points the
// picker at <examples>/temp/parity/file_picker (prepared with a_dir/inside.txt,
// b.txt and picked.txt), restricts the allowed types, uses memory input/output,
// no signal handler, 80x24, then Down, Up, Right, Left and 'q' with the spec's
// delays.

import (
	"os"
	"path/filepath"
	"time"

	"charm.land/bubbles/v2/filepicker"
	tea "charm.land/bubbletea/v2"

	"goldens/internal/capture"
)

// filePickerDir is the same directory the Crystal parity spec prepares, so both
// captures list identical entries (names and sizes appear in the output).
const filePickerDir = "/Volumes/extreme_ssd/repos/github.com/dsisnero/bubbletea.cr/bubbletea-examples/temp/parity/file_picker"

func prepareDir() error {
	if err := os.MkdirAll(filepath.Join(filePickerDir, "a_dir"), 0o755); err != nil {
		return err
	}
	for name, body := range map[string]string{
		filepath.Join("a_dir", "inside.txt"): "inside\n",
		"b.txt":                              "picked\n",
		"picked.txt":                         "picked\n",
	} {
		if err := os.WriteFile(filepath.Join(filePickerDir, name), []byte(body), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func goldenModel() tea.Model {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	fp.CurrentDirectory = filePickerDir
	return model{filepicker: fp}
}

func main() {
	if err := prepareDir(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	out := capture.Run(goldenModel(), capture.Options{Width: 80, Height: 24}, func(p *tea.Program) {
		capture.Sleep(120 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyDown})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyUp})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyRight})
		capture.Sleep(80 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: tea.KeyLeft})
		capture.Sleep(40 * time.Millisecond)
		p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	})
	os.Stdout.Write(out)
}
