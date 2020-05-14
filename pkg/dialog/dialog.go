// package diaglog encapsulates standard user messaging for all standard cli behavior.
// It is distinct from https://github.com/puppetlabs/horsehead/blob/master/logging/logger.go
// which should be used for structured debug logging, intended for developers or for semi-technical
// users running in debug-mode. This package is for polished messages that are leveled
// but unstructured. All messages are hidden in json output mode, under the assumption
// that users will want to pipe json output to a file or another process
package dialog

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/puppetlabs/relay/pkg/config"
)

type Dialog interface {
	WithStdout(io.Writer) Dialog
	WithStderr(io.Writer) Dialog

	Progress(string)
	Progressf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Error(string)
	Errorf(string, ...interface{})

	WriteString(string) error

	// Table returns a table for formatting for output.
	Table() Table
}

type TextDialog struct {
	p      *Progress
	stdout io.Writer
	stderr io.Writer
}

func (d *TextDialog) WithStdout(w io.Writer) Dialog {
	return &TextDialog{stdout: w, stderr: d.stderr, p: d.p}
}

func (d *TextDialog) WithStderr(w io.Writer) Dialog {
	return &TextDialog{stdout: d.stdout, stderr: w, p: d.p}
}

func withNewLine(str string) string {
	if len(str) == 0 {
		return ""
	}

	if str[len(str)-1] != '\n' {
		return str + "\n"
	}

	return str
}

func (d *TextDialog) completeProgress() {
	if d.p != nil {
		d.p.Complete()
		d.p = nil
	}
}

func (d *TextDialog) Info(message string) {
	d.completeProgress()

	fmt.Fprintf(d.stdout, withNewLine(message))
}

func (d *TextDialog) Infof(message string, args ...interface{}) {
	d.completeProgress()

	fmt.Fprintf(d.stdout, withNewLine(message), args...)
}

func (d *TextDialog) Error(msg string) {
	d.completeProgress()

	fmt.Fprintf(d.stderr, "%s%s", color.RedString("Error:"), withNewLine(msg))
}

func (d *TextDialog) Errorf(msg string, args ...interface{}) {
	d.completeProgress()

	str := fmt.Sprintf(msg, args...)
	fmt.Fprintf(d.stderr, "%s%s", color.RedString("Error:"), withNewLine(str))
}

func (d *TextDialog) Progress(msg string) {
	d.completeProgress()

	d.p = NewProgress(d.stdout, msg)
	d.p.Start()
}

func (d *TextDialog) Progressf(msg string, args ...interface{}) {
	d.completeProgress()

	d.p = NewProgress(d.stdout, fmt.Sprintf(msg, args...))
	d.p.Start()
}

func (d *TextDialog) WriteString(c string) error {
	_, err := io.WriteString(d.stdout, c)
	return err
}

func (d *TextDialog) Table() Table {
	return &textTable{w: d.stdout}
}

type JSONDialog struct {
	stdout, stderr io.Writer
}

func (d *JSONDialog) WithStdout(w io.Writer) Dialog {
	return &JSONDialog{stdout: w, stderr: d.stderr}
}

func (d *JSONDialog) WithStderr(w io.Writer) Dialog {
	return &JSONDialog{stdout: d.stdout, stderr: w}
}

func (d *JSONDialog) Progress(string) {
	// noop
}

func (d *JSONDialog) Progressf(string, ...interface{}) {
	// noop
}

func (d *JSONDialog) Info(string) {
	// noop
}

func (d *JSONDialog) Infof(string, ...interface{}) {
	// noop
}

func (d *JSONDialog) Error(msg string) {
	fmt.Fprintf(d.stderr, "%s%s", color.RedString("Error:"), msg)
}

func (d *JSONDialog) Errorf(msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	fmt.Fprintf(d.stderr, "%s%s", color.RedString("Error:"), str)
}

func (d *JSONDialog) WriteString(string) error {
	// noop
	return nil
}

func (d *JSONDialog) Table() Table {
	return &jsonTable{w: d.stdout}
}

func FromConfig(cfg *config.Config) Dialog {
	switch cfg.Out {
	case config.OutputTypeJSON:
		return &JSONDialog{stdout: os.Stdout, stderr: os.Stderr}
	default:
		return &TextDialog{stdout: os.Stdout, stderr: os.Stderr}
	}
}
