package gui

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"strconv"
	"time"

	"fyne.io/systray"
)

// UpdateMenu is called frequently via a go routine to check the status
// of access key, trigger gui updates and funcs to update the key
func UpdateMenu() {
	now := time.Now().UTC()
	at := Track.ExpiresAt()
	valid := Track.Valid()
	lock, lockErr := tracker.GetLock()

	debugger.Log("gui.UpdateMenu()",
		debugger.INFO,
		"now:\t\t"+now.String(),
		"rotateAt:\t"+at.String(),
		"valid:\t\t"+strconv.FormatBool(valid),
		"booting:\t"+strconv.FormatBool(cfg.IsBooting))()

	mu.Lock()

	// lock exists & is old
	if lockErr == nil && lock.Older(lockMaxAge) {
		MenuOldLock()
	}
	// lock exists
	if lockErr == nil {
		MenuLocked()
	} else if !valid && !cfg.IsBooting {
		MenuRotate()
	} else if !valid && cfg.IsBooting {
		MenuRotatingSoon()
	} else {
		menuInformation.SetTitle(
			fmt.Sprintf(labels.NextRotation, Track.ExpiresAt().Format(dateTimeFormat)))
	}

	mu.Unlock()
}

// MenuOldLock deals with a lock file that is aged, so presume
// that the rotate failed or was cancelled and therefore
// cleanup the file and carry on
func MenuOldLock() {
	defer debugger.Log("gui.MenuOldLock()", debugger.INFO, "Key is locked and too old, so removing...")()

	var err error = tracker.Unlock()
	if err != nil {
		panic(err)
	}
	menuRotate.Enable()
	systray.SetIcon(icons.Default(cfg.IsDarkMode))
}

// MenuLocked updates gui to show that there is a
// lock file present on the filesystem and there presume
// a rotate is in progress
func MenuLocked() {
	defer debugger.Log("gui.MenuLocked()", debugger.INFO, "Key is locked...")()
	menuRotate.Enable()
	menuInformation.SetTitle(labels.Locked)
	systray.SetIcon(icons.Locked(cfg.IsDarkMode))
	// menu.Refresh()
}

func MenuRotatingSoon() {
	defer debugger.Log("gui.MenuWillRotate()", debugger.INFO, "Key will be rotated, show warning")()
	systray.SetIcon(icons.RotatingSoon(cfg.IsDarkMode))
	menuInformation.SetTitle(labels.Rotating)
}

// MenuRotate handles the gui changes and func calls to change
// a key and show the status of that change
func MenuRotate() {
	debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotating key...")()
	var err error = tracker.SetLock(Track)
	if err != nil {
		panic(err)
	}
	systray.SetIcon(icons.Rotating(cfg.IsDarkMode))

	menuInformation.SetTitle(labels.Rotating)
	menuRotate.Disabled()

	command := cfg.Vault.Command(cfg.Profile, cfg.Os)
	sOut, sErr, err := cfg.Shell.Run(command, false, false)

	debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotate command finished", "stdOut:", sOut, "stdErr:", sErr)()
	if err == nil {
		err = tracker.Unlock()
		if err != nil {
			panic(err)
		}
		// new reacker
		Track, _ = tracker.SetCurrent(tracker.Clean())

		debugger.Log("gui.MenuRotate()", debugger.INFO, "Rotated successfully", "new tracker:", Track)()
		nextRotate := fmt.Sprintf(labels.NextRotation, Track.ExpiresAt().Format(dateTimeFormat))
		menuInformation.SetTitle(nextRotate)
		menuRotate.Enable()
		systray.SetIcon(icons.Default(cfg.IsDarkMode))
		cfg.Os.SystemMessage(cfg.Shell, cfg.AppBuiltName, []string{nextRotate}, labels.SystemSuccess)

	} else {
		debugger.Log("gui.MenuRotate()", debugger.ERR, "Rotate failed", "err:", err, "stdErr:", sErr.String())()
		cfg.Os.SystemMessage(cfg.Shell, cfg.AppBuiltName, []string{sErr.String(), err.Error()}, labels.SystemError)
		MenuLocked()
	}

}
