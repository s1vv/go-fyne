package ui

import "fyne.io/fyne/v2/widget"

type StatusBar struct {
	Widget *widget.Label
}

func NewStatusBar() *StatusBar {
	label := widget.NewLabel("Ожидание запроса...")
	return &StatusBar{Widget: label}
}

func (s *StatusBar) SetText(text string) {
	s.Widget.SetText(text)
}
