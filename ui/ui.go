package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type AppUI struct {
	Layout   *fyne.Container
	status   *StatusBar
	progress *ProgressBar
}

func NewAppUI(fetchFunc func() (string, error)) *AppUI {
	status := NewStatusBar()
	progress := NewProgressBar()
	button := NewRequestButton(status, progress, fetchFunc)

	layout := container.NewVBox(
		button,
		progress.Widget,
		status.Widget,
	)

	return &AppUI{
		Layout:   layout,
		status:   status,
		progress: progress,
	}
}
