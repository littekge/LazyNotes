// Package mdparse reads from a markdown file (.md) and formats the contents of the file in accordance to the tview text format.
package mdparse

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

func ParseText(path string) (string, error) {
	// Ensure that 'path' is the path to an existing markdown file
	if filepath.Ext(path) != ".md" {
		return "", errors.New("parser: invalid file extension" + filepath.Ext(path))
	}
	if _, err := os.Stat(path); err != nil {
		return "", errors.New("parser: file" + path + "does not exist")
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes), nil
}

func MParseText(source io.Writer) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}
	return "", nil
}
