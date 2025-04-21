package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewFileSelector(win fyne.Window, title string, ui *AppUI) *FileSelector {
	label := widget.NewLabel(title + ": не выбран")
	selector := &FileSelector{
		Label: label,
		Path:  "",
	}

	btn := widget.NewButton("Выбрать "+title, func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err == nil && uri != nil {
				selector.Path = uri.URI().Path()
				label.SetText(title + ": " + selector.Path)
			}
		}, win)
	})
	ui.FileSelectors[title] = selector

	selector.Container = container.NewVBox(btn, label)
	return selector
}
