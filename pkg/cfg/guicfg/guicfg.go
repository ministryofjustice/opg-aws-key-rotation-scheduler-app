package guicfg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	AppId   string = "com.opg-aws-key-rotation.app"
	AppName string = "OPG AWS Key Rotation"
)

var (
	App       fyne.App
	Window    fyne.Window
	Desktop   desktop.App
	IsDesktop bool
)

func init() {
	// fyne app setup
	App = app.NewWithID(AppId)
	Window = App.NewWindow(AppName)
	Desktop, IsDesktop = App.(desktop.App)
}
