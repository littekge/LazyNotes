// Package ui builds the application ui
package ui

import (
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

// Builds a basic box (unused function)
func buildBox() *tview.Box {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Note")
	box.SetFocusFunc(func() { box.SetBorderStyle(LazyNotesFocusStyle()) })
	box.SetBlurFunc(func() { box.SetBorderStyle(LazyNotesBlurStyle()) })
	return box
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

func BuildApp() {
	setDrawVars()
	// box := buildNoteView()
	app = tview.NewApplication()
	mainContainer := tview.NewFlex().
		AddItem(buildNoteView(), 0, 1, false)
	if err := app.SetRoot(mainContainer, true).Run(); err != nil {
		panic(err)
	}
}
