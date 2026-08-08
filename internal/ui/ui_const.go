package ui

import "github.com/gdamore/tcell/v2"

func LazyNotesBlurStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorWhite)
}

func LazyNotesFocusStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.ColorYellow)
}
