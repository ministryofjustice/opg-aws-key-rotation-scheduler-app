package icons

import (
	"embed"
	_ "embed"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
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

func Default(isDarkMode bool) (i []byte) {
	i = LightModeDefault
	if isDarkMode {
		i = DarkModeDefault
	}
	defer debugger.Log("icons.Default()", debugger.INFO, isDarkMode)()
	return
}

func Locked(isDarkMode bool) (i []byte) {
	i = LightModeLocked
	if isDarkMode {
		i = DarkModeLocked
	}
	defer debugger.Log("icons.Locked()", debugger.INFO, isDarkMode)()
	return
}

func Rotating(isDarkMode bool) (i []byte) {
	i = LightModeRotating
	if isDarkMode {
		i = DarkModeRotating
	}
	defer debugger.Log("icons.Rotating()", debugger.INFO, isDarkMode)()
	return
}

func RotatingSoon(isDarkMode bool) (i []byte) {
	i = LightModeRotatingSoon
	if isDarkMode {
		i = DarkModeRotatingSoon
	}
	defer debugger.Log("icons.RotatingSoon()", debugger.INFO, isDarkMode)()
	return
}
