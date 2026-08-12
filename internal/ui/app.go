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

// sets the tview variables that determine how the UI is drawn.
func setDrawVars() {
	// set unfocused borders
	tview.Borders.Horizontal = tview.BoxDrawingsLightHorizontal
	tview.Borders.Vertical = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeft = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRight = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeft = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRight = tview.BoxDrawingsLightArcUpAndLeft
	// set focused borders
	// focused borders are intentionally the same as unfocused borders, elements
	// use SetFocusFunc change the color when they are focused.
	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightArcDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightArcDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightArcUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightArcUpAndLeft
}

// Builds the note viewing pane.
func buildNoteView() *tview.TextView {
	tv := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetText("Notes appear here...")
	tv.SetBorder(true)
	tv.SetFocusFunc(func() { tv.SetBorderStyle(LazyNotesFocusStyle()) })
	tv.SetBlurFunc(func() { tv.SetBorderStyle(LazyNotesBlurStyle()) })
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

	tv.SetBorder(true)
	tv.SetFocusFunc(func() { tv.SetBorderStyle(LazyNotesFocusStyle()) })
	tv.SetBlurFunc(func() { tv.SetBorderStyle(LazyNotesBlurStyle()) })
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
