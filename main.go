package main

import (
	"go-fyne/service"
	"go-fyne/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Request Example")

	appUI := ui.NewAppUI(service.FetchData)

	w.SetContent(appUI.Layout)
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	w.ShowAndRun()
}
