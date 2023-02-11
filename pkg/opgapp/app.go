package opgapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	App         fyne.App
	Preferences fyne.Preferences
)

const (
	AppId   string = "com.opg-aws-key-rotation.app"
	AppName string = "OPG AWS Key Rotation"
)

func init() {
	App = app.NewWithID(AppId)
	Preferences = App.Preferences()
}
