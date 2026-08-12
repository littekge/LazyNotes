// Package ui builds the application ui
package ui

import (
	"os"
	"path/filepath"

	"github.com/littekge/LazyNotes/internal/mdparse"
	"github.com/rivo/tview"
)

var (
	app      *tview.Application
	treeView *tview.TreeView
	noteView *tview.TextView
)

// Builds the note viewing pane.
func buildNoteView() *tview.TextView {
	tv := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetText("Notes appear here...")
	Decorate(tv.Box, "Detail")
	return tv
}

func buildTreeView() *tview.TreeView {
	root := tview.NewTreeNode("Notes")

	// recursive helper function to build the node tree
	var addNodes func(*tview.TreeNode, string)
	addNodes = func(target *tview.TreeNode, path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fileString := filepath.Join(path, file.Name())
			node := tview.NewTreeNode(file.Name()).
				SetSelectable(true).
				SetReference(fileString)
			target.AddChild(node)
			if file.IsDir() {
				addNodes(node, fileString)
			}
		}
	}

	// path is hardcoded for now, dynamic configuration will be added in a later version
	addNodes(root, "./test_notes_dir")
	tv := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// displays text of file on change
	tv.SetChangedFunc(func(node *tview.TreeNode) {
		filePath, ok := node.GetReference().(string)
		if !ok {
			noteView.SetText("Notes appear here...")
			return
		}
		mdtext, err := mdparse.ParseText(filePath)
		if err != nil {
			noteView.SetText(err.Error())
		} else {
			noteView.SetText(mdtext)
		}
	})
	Decorate(tv.Box, "Notes")
	return tv
}

func BuildApp() {
	setDrawVars()
	app = tview.NewApplication()
	treeView = buildTreeView()
	noteView = buildNoteView()
	mainContainer := tview.NewFlex().
		AddItem(treeView, 0, 1, true).
		AddItem(noteView, 0, 2, false)
	if err := app.SetRoot(mainContainer, true).Run(); err != nil {
		panic(err)
	}
}
