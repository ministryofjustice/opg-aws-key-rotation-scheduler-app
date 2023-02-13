package gui

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	"time"

	"fyne.io/fyne/v2"
)

func SystraySetup() {
	menuRotate = fyne.NewMenuItem(labels.Rotate, func() {
		mu.Lock()
		MenuRotate()
		mu.Unlock()
	})

	menuInformation = fyne.NewMenuItem(
		fmt.Sprintf(labels.NextRotation, Track.ExpiresAt().Format(dateTimeFormat)),
		func() {},
	)
	menuInformation.Disabled = true

	split := fyne.NewMenuItemSeparator()

	menu = fyne.NewMenu(cfg.AppName, menuRotate, split, menuInformation)
	debugger.Log("gui.SystraySetup()", debugger.INFO, "generated menu")()

	go func() {
		for range time.Tick(tickDuration) {
			cfg.IsBooting = false
			debugger.Log("tick", debugger.INFO)()
			UpdateMenu()
		}
	}()
}
