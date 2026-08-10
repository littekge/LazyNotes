// Package ui builds the application ui
package ui

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

var app *tview.Application

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
			node := tview.NewTreeNode(file.Name()).
				SetSelectable(true)
			target.AddChild(node)
			if file.IsDir() {
				addNodes(node, filepath.Join(path, file.Name()))
			}
		}
	}

	// path is hardcoded for now, dynamic configuration will be added in a later version
	addNodes(root, "./test_notes_dir")
	tv := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	tv.SetBorder(true)
	tv.SetFocusFunc(func() { tv.SetBorderStyle(LazyNotesFocusStyle()) })
	tv.SetBlurFunc(func() { tv.SetBorderStyle(LazyNotesBlurStyle()) })
	return tv
}

func BuildApp() {
	setDrawVars()
	app = tview.NewApplication()
	mainContainer := tview.NewFlex().
		AddItem(buildTreeView(), 0, 1, true).
		AddItem(buildNoteView(), 0, 2, false)
	if err := app.SetRoot(mainContainer, true).Run(); err != nil {
		panic(err)
	}
}
