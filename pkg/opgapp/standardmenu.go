package opgapp

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

func SetupStandardMenu() {

	_menuRotate = fyne.NewMenuItem(_labels.Rotate, func() {
		_mu.Lock()
		MenuRotate()
		_mu.Unlock()
	})

	_menuInformation = fyne.NewMenuItem(
		fmt.Sprintf(_labels.NextRotation, _track.RotateAt(_rotateFrequency).Format(_settings.DateTimeFormat)),
		func() {},
	)
	_menuInformation.Disabled = true

	split := fyne.NewMenuItemSeparator()

	_menu = fyne.NewMenu(_settings.Name, _menuRotate, split, _menuInformation)

	go func() {
		dur := time.Minute
		for range time.Tick(dur) {
			_booting = false
			fmt.Println("tick")
			UpdateMenu()
		}
	}()

}
