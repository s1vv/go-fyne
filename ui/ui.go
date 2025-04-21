package ui

import (
	"fmt"

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
	Layout        *fyne.Container
	StatusBar     *StatusBar
	ProgressBar   *ProgressBar
	FileSelectors map[string]*FileSelector
}

func (ui *AppUI) GetPath(key string) string {
	fs, ok := ui.FileSelectors[key]
	for p := range ui.FileSelectors {
		fmt.Println(ui.FileSelectors[p])
	}
	if !ok {
		return ""
	}
	return fs.Path
}

func NewAppUI(win fyne.Window) *AppUI {
	status := NewStatusBar()
	progress := NewProgressBar()

	ui := &AppUI{
		Layout:        container.NewVBox(),
		StatusBar:     status,
		ProgressBar:   progress,
		FileSelectors: make(map[string]*FileSelector),
	}

	btnSelectFile1 := NewFileSelector(win, "file 1", ui)
	btnSelectFile2 := NewFileSelector(win, "file 2", ui)

	ui.Layout.Add(widget.NewLabel("Действия:"))
	ui.Layout.Add(NewButton(
		ui,
		"Объединить файлы",
		"python",
		func() string { return ui.GetPath("file 1") },
		func() string { return ui.GetPath("file 2") },
	)) // пути будут внутри обработчика
	ui.Layout.Add(NewButton(ui, "Запрос get json", "network"))

	ui.Layout.Add(widget.NewLabel("Файлы:"))
	ui.Layout.Add(btnSelectFile1.Container)
	ui.Layout.Add(btnSelectFile2.Container)
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
