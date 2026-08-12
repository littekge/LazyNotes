package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	blurBorderColor  = tcell.ColorGray
	focusBorderColor = tcell.ColorYellow
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

func LazyNotesBlurStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(blurBorderColor)
}

func LazyNotesFocusStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(focusBorderColor)
}

func Decorate(elem *tview.Box, title string) {
	elem.SetBorder(true)
	elem.SetTitle(" " + title + " ")
	elem.SetBorderStyle(LazyNotesBlurStyle())
	elem.SetFocusFunc(func() { elem.SetBorderStyle(LazyNotesFocusStyle()) })
	elem.SetBlurFunc(func() { elem.SetBorderStyle(LazyNotesBlurStyle()) })
}
