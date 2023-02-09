package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"

	"fyne.io/fyne/v2"
)

func SetupErrorMenu(err string) {
	defer debugger.Log("SetupErrorMenu()", debugger.INFO)()
	errorMsg := fyne.NewMenuItem("", func() {})
	errorMsg.Disabled = true
	errorMsg.Label = err
	_menu = fyne.NewMenu(_settings.Name, errorMsg)
}
