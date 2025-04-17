package ui

import (
	"go-fyne/controller"
	"go-fyne/viewiface"

	"fyne.io/fyne/v2/widget"
)

func NewRequestButton(view viewiface.ViewInterface) *widget.Button {
	ctrl := controller.NewAppController(view)
	return widget.NewButton("Отправить запрос", func() {
		ctrl.FetchAndDisplayData()
	})
}
