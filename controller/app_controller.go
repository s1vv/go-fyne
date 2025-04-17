package controller

import (
	"go-fyne/model"
	"go-fyne/viewiface"

	"fyne.io/fyne/v2"
)

type AppController struct {
	View viewiface.ViewInterface
}

func NewAppController(view viewiface.ViewInterface) *AppController {
	return &AppController{View: view}
}

func (c *AppController) FetchAndDisplayData() {
	c.View.ShowProgress()
	c.View.SetStatus("Загрузка...")

	go func() {
		data, err := model.FetchData()

		fyne.Do(func() {
			c.View.HideProgress()
			if err != nil {
				c.View.SetStatus("Ошибка: " + err.Error())
				return
			}
			c.View.SetStatus("Результат: " + data)
		})
	}()
}
