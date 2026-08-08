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

func BuildApp() {
	setDrawVars()

	app = tview.NewApplication()
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("test").
		SetBorderStyle(LazyNotesBlurStyle())
	box.SetFocusFunc(func() { box.SetBorderStyle(LazyNotesFocusStyle()) })
	box.SetBlurFunc(func() { box.SetBorderStyle(LazyNotesBlurStyle()) })
	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
