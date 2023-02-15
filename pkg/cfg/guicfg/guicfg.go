package guicfg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	AppId   string = "com.opg-aws-key-rotation.app"
	AppName string = "OPG AWS Key Rotation"
)

var (
	App fyne.App
)

func init() {
	// fyne app setup
	App = app.NewWithID(AppId)
}
