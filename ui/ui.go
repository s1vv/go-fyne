package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FileSelector struct {
	Container *fyne.Container
	Path      string
	Label     *widget.Label
}

type AppUI struct {
	Layout       *fyne.Container
	StatusBar    *StatusBar
	ProgressBar  *ProgressBar
	FileSelector map[string]FileSelector
}

func NewAppUI(win fyne.Window) *AppUI {
	status := NewStatusBar()
	progress := NewProgressBar()
	var btnSelectFile1 FileSelector
	var btnSelectFile2 FileSelector
	fileSelecrors := map[string]FileSelector{
		"selector1": *btnSelectFile1,
		"selector2": *btnSelectFile2,
	}

	ui := &AppUI{
		Layout:       container.NewVBox(),
		StatusBar:    status,
		ProgressBar:  progress,
		FileSelector: fileSelecrors,
	}

	btnSelectFile1 = NewFileSelector(win, "Файл 1")
	btnSelectFile1 = NewFileSelector(win, "Файл 1")

	ui.Layout.Add(widget.NewLabel("Действия:"))
	ui.Layout.Add(NewButton(
		ui,
		"Объединить файлы",
		"python",
		func() string { return ui.FileSelector["selector1"] },
		func() string { return ui.file2Path },
	)) // пути будут внутри обработчика
	ui.Layout.Add(NewButton(ui, "Запрос get json", "network"))

	ui.Layout.Add(widget.NewLabel("Файлы:"))

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
