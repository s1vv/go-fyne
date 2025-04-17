package main

import (
	"go-fyne/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("MVC Example")
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	appUI := ui.NewAppUI()

	w.SetContent(appUI.Layout)
	w.ShowAndRun()
}
