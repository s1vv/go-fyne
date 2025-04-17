package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type AppUI struct {
	Layout      *fyne.Container
	StatusBar   *StatusBar
	ProgressBar *ProgressBar
}

func NewAppUI() *AppUI {
	status := NewStatusBar()
	progress := NewProgressBar()

	ui := &AppUI{
		Layout:      container.NewVBox(),
		StatusBar:   status,
		ProgressBar: progress,
	}

	ui.Layout.Add(NewRequestButton(ui))
	ui.Layout.Add(progress.Widget)
	ui.Layout.Add(status.Widget)

	return ui
}

// Реализация интерфейса
func (a *AppUI) SetStatus(text string) {
	a.StatusBar.SetText(text)
}

func (a *AppUI) ShowProgress() {
	a.ProgressBar.Show()
}

func (a *AppUI) HideProgress() {
	a.ProgressBar.Hide()
}
