package gui

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"strconv"
	"time"
)

// UpdateMenu is called frequently via a go routine to check the status
// of access key, trigger gui updates and funcs to update the key
func UpdateMenu() {
	now := time.Now().UTC()
	at := track.ExpiresAt()
	valid := track.Valid()
	lock, lockErr := tracker.GetLock()

	debugger.Log("gui.UpdateMenu()", debugger.INFO, "now:\t\t"+now.String(), "rotateAt:\t"+at.String(), "valid:\t\t"+strconv.FormatBool(valid), "booting:\t"+strconv.FormatBool(booting))()

	mu.Lock()

	// lock exists & is old
	if lockErr == nil && lock.Older(lockMaxAge) {
		MenuOldLock()
	}

	// lock exists
	if lockErr == nil {
		MenuLocked()
	} else if !valid && !booting {
		MenuRotate()
	} else if !valid && booting {
		MenuRotatingSoon()
	} else {
		menuInformation.Label = fmt.Sprintf(
			labels.NextRotation,
			track.ExpiresAt().Format(dateTimeFormat),
		)
	}
	menu.Refresh()
	mu.Unlock()
}

// MenuOldLock deals with a lock file that is aged, so presume
// that the rotate failed or was cancelled and therefore
// cleanup the file and carry on
func MenuOldLock() {
	defer debugger.Log("gui.MenuOldLock()", debugger.INFO, "Key is locked and too old, so removing...")()
	tracker.Unlock()
}

// MenuLocked updates gui to show that there is a
// lock file present on the filesystem and there presume
// a rotate is in progress
func MenuLocked() {
	defer debugger.Log("gui.MenuLocked()", debugger.INFO, "Key is locked...")()
	menuRotate.Disabled = false
	menuInformation.Label = labels.Locked
	desktopApp.SetSystemTrayIcon(icons.Locked(isDark))
	menu.Refresh()
}

func MenuRotatingSoon() {
	defer debugger.Log("gui.MenuWillRotate()", debugger.INFO, "Key will be rotated, show warning")()
	desktopApp.SetSystemTrayIcon(icons.RotatingSoon(isDark))
	menuInformation.Label = labels.Rotating
	menu.Refresh()
}

// MenuRotate handles the gui changes and func calls to change
// a key and show the status of that change
func MenuRotate() {
	debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotating key...")()
	tracker.SetLock(track)

	desktopApp.SetSystemTrayIcon(icons.Rotating(isDark))
	menuInformation.Label = labels.Rotating
	menuRotate.Disabled = true
	menu.Refresh()

	command := awsVault.Command(prof, os)
	sOut, sErr, err := zsh.Run([]string{command}, false)

	debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotate command finished", "stdOut:", sOut, "stdErr:", sErr)()
	if err == nil {
		tracker.Unlock()
		// new reacker
		track, _ = tracker.SetCurrent(tracker.Clean())
		debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotated successfully", "new tracker:", track)()
		menuInformation.Label = fmt.Sprintf(
			labels.NextRotation,
			track.ExpiresAt().Format(dateTimeFormat),
		)
		menuRotate.Disabled = false
		desktopApp.SetSystemTrayIcon(icons.Default(isDark))

	} else {
		debugger.Log("gui.MenuRotate()", debugger.ERR, "Rotate failed", "err:", err, "stdErr:", sErr.String())()
		w := ErrorDialog(app, []string{sErr.String(), err.Error()})
		w.Show()
		MenuLocked()
	}
	menu.Refresh()

}