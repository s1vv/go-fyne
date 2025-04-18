package ui

import (
	"go-fyne/controller"
	"go-fyne/viewiface"

	"fyne.io/fyne/v2/widget"
)

func NewButton(view viewiface.ViewInterface, nameBtn string, cmd string, args ...func() string) *widget.Button {
	ctrl := controller.NewAppController(view)
	return widget.NewButton(nameBtn, func() {
		ctrl.RunFetch(cmd, args...)
	})
}
