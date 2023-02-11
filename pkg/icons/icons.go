package icons

import (
	"embed"
	_ "embed"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"

	"fyne.io/fyne/v2"
)

var (
	//go:embed files/*
	fs embed.FS
	//go:embed files/main.png
	AppIcon []byte
	//go:embed files/menu.darkmode.png
	DarkModeDefault []byte
	//go:embed files/locked.png
	DarkModeLocked []byte
	//go:embed files/rotating.png
	DarkModeRotating []byte
	//go:embed files/rotating.png
	DarkModeRotatingSoon []byte

	//go:embed files/menu.lightmode.png
	LightModeDefault []byte
	//go:embed files/locked.png
	LightModeLocked []byte
	//go:embed files/rotating.png
	LightModeRotating []byte
	//go:embed files/rotating.png
	LightModeRotatingSoon []byte
)

func Default(isDarkMode bool) (r fyne.Resource) {
	r = fyne.NewStaticResource("default", LightModeDefault)
	if isDarkMode {
		r = fyne.NewStaticResource("default", DarkModeDefault)
	}
	defer debugger.Log("icons.Default()", debugger.INFO, isDarkMode)()
	return
}

func Locked(isDarkMode bool) (r fyne.Resource) {
	r = fyne.NewStaticResource("locked", LightModeLocked)
	if isDarkMode {
		r = fyne.NewStaticResource("locked", DarkModeLocked)
	}
	defer debugger.Log("icons.Locked()", debugger.INFO, isDarkMode)()
	return
}

func Rotating(isDarkMode bool) (r fyne.Resource) {
	r = fyne.NewStaticResource("rotating", LightModeRotating)
	if isDarkMode {
		r = fyne.NewStaticResource("rotating", DarkModeRotating)
	}
	defer debugger.Log("icons.Rotating()", debugger.INFO, isDarkMode)()
	return
}

func RotatingSoon(isDarkMode bool) (r fyne.Resource) {
	r = fyne.NewStaticResource("rotating-soon", LightModeRotatingSoon)
	if isDarkMode {
		r = fyne.NewStaticResource("rotating-soon", DarkModeRotatingSoon)
	}
	defer debugger.Log("icons.RotatingSoon()", debugger.INFO, isDarkMode)()
	return
}
