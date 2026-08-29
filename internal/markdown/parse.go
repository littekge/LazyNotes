// Package markdown reads from a markdown file (.md) and formats the contents of the file in accordance to the tview text format.
package markdown

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark/v2/parser"
)

// Valid ensures that 'path' is the path to an existing markdown file
func Valid(path string) (bool, error) {
	if filepath.Ext(path) != ".md" {
		return false, errors.New("parser: invalid file extension" + filepath.Ext(path))
	}
	if _, err := os.Stat(path); err != nil {
		return false, errors.New("parser: file" + path + "does not exist")
	}
	return true, nil
}

// RenderText uses the glamour library to render a markdown file into ANSI text compatible with tview
func RenderText(path string) (string, error) {
	if _, err := Valid(path); err != nil {
		return "", err
	}
	rawBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	p := parser.New()
	doc := p.Parse(rawBytes)

	r := New()
	var buf bytes.Buffer
	if err := r.Render(&buf, rawBytes, doc); err != nil {
		panic(err)
	}
	return buf.String(), nil
}
