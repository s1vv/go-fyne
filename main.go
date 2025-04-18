package main

import (
	"fmt"
	"go-fyne/ui"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	os.Setenv("LANGUAGE", "en_US.UTF-8")
	os.Setenv("LANG", "en_US.UTF-8")

	a := app.NewWithID("com.s1v.gotool")
	w := a.NewWindow("MVC Example")
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)
	fmt.Println(fyne.CurrentApp().Driver().Device().Locale())

	appUI := ui.NewAppUI(w)

	w.SetContent(appUI.Layout)
	w.ShowAndRun()
}
