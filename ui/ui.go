package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AppUI struct {
	Layout      *fyne.Container
	StatusBar   *StatusBar
	ProgressBar *ProgressBar
	file1Path   string
	file2Path   string
}

func NewAppUI(win fyne.Window) *AppUI {
	status := NewStatusBar()
	progress := NewProgressBar()

	file1Label := widget.NewLabel("Файл 1: не выбран")
	file2Label := widget.NewLabel("Файл 2: не выбран")

	ui := &AppUI{
		Layout:      container.NewVBox(),
		StatusBar:   status,
		ProgressBar: progress,
	}

	btnSelectFile1 := widget.NewButton("Выбрать файл 1", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err == nil && uri != nil {
				ui.file1Path = uri.URI().Path()
				file1Label.SetText("Файл 1: " + ui.file1Path)
			}
		}, win)
	})

	btnSelectFile2 := widget.NewButton("Выбрать файл 2", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err == nil && uri != nil {
				ui.file2Path = uri.URI().Path()
				file2Label.SetText("Файл 2: " + ui.file2Path)
			}
		}, win)
	})

	ui.Layout.Add(widget.NewLabel("Действия:"))
	ui.Layout.Add(NewButton(ui,
		"Объединить файлы",
		"python",
		func() string { return ui.file1Path },
		func() string { return ui.file2Path },
	)) // пути будут внутри обработчика
	ui.Layout.Add(NewButton(ui, "Запрос get json", "network"))

	ui.Layout.Add(widget.NewLabel("Файлы:"))
	ui.Layout.Add(btnSelectFile1)
	ui.Layout.Add(file1Label)
	ui.Layout.Add(btnSelectFile2)
	ui.Layout.Add(file2Label)

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
