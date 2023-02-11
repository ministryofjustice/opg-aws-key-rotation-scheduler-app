package gui

import (
	"image/color"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func errorMenu(app fyne.App) (menu *fyne.MainMenu) {
	file := fyne.NewMenu(
		"File",
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)
	menu = fyne.NewMainMenu(file)

	return
}

func ErrorDialog(
	app fyne.App,
	messages []string,
) fyne.Window {
	debugger.Log("gui.ErrorDialog()", debugger.ERR, "Error messages", strings.Join(messages, "\n"))()
	window := app.NewWindow(opgapp.AppName + " - Error!")
	window.SetMainMenu(errorMenu(app))

	box := container.NewVBox()
	for _, message := range messages {
		text := canvas.NewText(message, color.White)
		text.Alignment = fyne.TextAlignCenter
		box.Add(text)
	}

	window.SetContent(box)
	return window
}