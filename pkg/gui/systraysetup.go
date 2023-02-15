package gui

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	"time"

	"fyne.io/systray"
)

func SystraySetup() {

	systray.SetIcon(icons.Default(cfg.IsDarkMode))
	menuRotate = systray.AddMenuItem(labels.Rotate, labels.Rotate)
	menuRotate.Enable()
	go func() {
		<-menuRotate.ClickedCh
		mu.Lock()
		MenuRotate()
		mu.Unlock()
	}()

	systray.AddSeparator()

	infoLabel := fmt.Sprintf(labels.NextRotation, Track.ExpiresAt().Format(dateTimeFormat))
	menuInformation = systray.AddMenuItem(infoLabel, infoLabel)
	menuInformation.Disable()

	menuQuit = systray.AddMenuItem("Quit", "Quit")
	menuQuit.Enable()
	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()

	UpdateMenu()

	go func() {
		for range time.Tick(tickDuration) {
			cfg.IsBooting = false
			debugger.Log("tick", debugger.INFO)()
			UpdateMenu()
		}
	}()
}
