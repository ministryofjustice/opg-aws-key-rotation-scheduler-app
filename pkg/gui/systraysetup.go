package gui

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	. "opg-aws-key-rotation-scheduler-app/pkg/opgapp"
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

	menu = fyne.NewMenu(AppName, menuRotate, split, menuInformation)
	debugger.Log("gui.SystraySetup()", debugger.INFO, "generated menu")()

	go func() {
		for range time.Tick(tickDuration) {
			Booting = false
			debugger.Log("tick", debugger.INFO)()
			UpdateMenu()
		}
	}()
}
