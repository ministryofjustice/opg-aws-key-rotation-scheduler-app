package opgapp

import (
	"fyne.io/fyne/v2"
)

func SetupErrorMenu(err string) {
	errorMsg := fyne.NewMenuItem("", func() {})
	errorMsg.Disabled = true
	errorMsg.Label = err
	_menu = fyne.NewMenu(_settings.Name, errorMsg)
}
