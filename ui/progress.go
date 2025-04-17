package ui

import "fyne.io/fyne/v2/widget"

type ProgressBar struct {
	Widget *widget.ProgressBarInfinite
}

func NewProgressBar() *ProgressBar {
	p := widget.NewProgressBarInfinite()
	p.Hide()
	return &ProgressBar{Widget: p}
}

func (p *ProgressBar) Show() {
	p.Widget.Show()
}

func (p *ProgressBar) Hide() {
	p.Widget.Hide()
}
