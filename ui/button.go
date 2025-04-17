package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewRequestButton(status *StatusBar, progress *ProgressBar, fetchFunc func() (string, error)) *widget.Button {
	return widget.NewButton("Отправить запрос", func() {
		progress.Show()
		status.SetText("Загрузка...")

		go func() {
			result, err := fetchFunc()
			fyne.Do(func() {
				progress.Hide()
				if err != nil {
					status.SetText("Ошибка: " + err.Error())
				} else {
					status.SetText("Успех: " + result)
				}
			})
		}()
	})
}
