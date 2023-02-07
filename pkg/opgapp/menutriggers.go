package opgapp

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

// UpdateMenu is called frequently via a go routine to check the status
// of access key, trigger gui updates and funcs to update the key
func UpdateMenu() {
	now := time.Now().UTC()
	at := _track.RotateAt(_rotateFrequency)
	pp.Printf("[%s] next rotation at [%s]\n", now, at)

	_mu.Lock()
	if _track.Locked() && _track.LockIsOld() {
		MenuKeyOldLock()
	}

	if _track.Locked() {
		MenuKeyLocked()
	} else if now.After(at) {
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
	pp.Println("Key is locked and too old, so removing...")
	_track.Unlock()
}

// MenuKeyLocked updates gui to show that there is a
// lock file present on the filesystem and there presume
// a rotate is in progress
func MenuKeyLocked() {
	pp.Println("Key is locked...")
	_menuInformation.Label = _labels.Locked
	_menu.Refresh()
}

// MenuRotate handles the gui changes and func calls to change
// a key and show the status of that change
func MenuRotate() {
	pp.Println("Rotating key...")
	_track.Lock()

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

	} else {
		MenuKeyLocked()
	}
	_menu.Refresh()

}
