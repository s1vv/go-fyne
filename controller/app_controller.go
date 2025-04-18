package controller

import (
	"go-fyne/handlers"
	"go-fyne/viewiface"

	"fyne.io/fyne/v2"
)

type AppController struct {
	View viewiface.ViewInterface
}

func NewAppController(view viewiface.ViewInterface) *AppController {
	return &AppController{View: view}
}

func (c *AppController) RunFetch(mode string, args ...func() string) {
	c.View.ShowProgress()
	c.View.SetStatus("Загрузка...")
	go func() {
		result, err := handlers.FetchData(mode, args...)
		fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Готово"})
		fyne.Do(func() {
			c.View.HideProgress()
			if err != nil {
				c.View.SetStatus("Ошибка: " + err.Error())
				return
			}
			c.View.SetStatus("Результат: " + result)
		})
	}()
}
