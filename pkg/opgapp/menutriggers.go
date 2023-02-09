package opgapp

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"time"
)

// UpdateMenu is called frequently via a go routine to check the status
// of access key, trigger gui updates and funcs to update the key
func UpdateMenu() {
	now := time.Now().UTC()
	at := _track.RotateAt(_rotateFrequency)
	debugger.Log("UpdateMenu", debugger.INFO, at, now)()

	_mu.Lock()
	if _track.Locked() && _track.LockIsOld() {
		MenuKeyOldLock()
	}

	if _track.Locked() {
		MenuKeyLocked()
	} else if now.After(at) && !_booting {
		MenuRotate()
	} else {
		_menuInformation.Label = fmt.Sprintf(
			_labels.NextRotation,
			_track.RotateAt(_rotateFrequency).Format(_settings.DateTimeFormat),
		)
	}

	_mu.Unlock()
}

// MenuKeyOldLock deals with a lock file that is aged, so presume
// that the rotate failed or was cancelled and therefore
// cleanup the file and carry on
func MenuKeyOldLock() {
	debugger.Log("MenuKeyOldLock", debugger.INFO, "Key is locked and too old, so removing...")()
	_track.Unlock()
}

// MenuKeyLocked updates gui to show that there is a
// lock file present on the filesystem and there presume
// a rotate is in progress
func MenuKeyLocked() {
	debugger.Log("MenuKeyLocked", debugger.INFO, "Key is locked...")()
	_menuRotate.Disabled = false
	_menuInformation.Label = _labels.Locked
	_desk.SetSystemTrayIcon(_icons.Locked())
	_menu.Refresh()
}

// MenuRotate handles the gui changes and func calls to change
// a key and show the status of that change
func MenuRotate() {
	debugger.Log("MenuRotate", debugger.INFO, "Rotating key...")()
	_track.Lock()

	_desk.SetSystemTrayIcon(_icons.Rotating())
	_menuInformation.Label = _labels.Rotating
	_menuRotate.Disabled = true
	_menu.Refresh()

	err := RotateCommand(_settings)

	if err == nil {
		_track.Unlock()
		_track = _track.Rotate()

		_menuInformation.Label = fmt.Sprintf(
			_labels.NextRotation,
			_track.RotateAt(_rotateFrequency).Format(_settings.DateTimeFormat),
		)
		_menuRotate.Disabled = false
		_desk.SetSystemTrayIcon(_icons.Default())

	} else {
		MenuKeyLocked()
	}
	_menu.Refresh()

}
