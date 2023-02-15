package gui

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/systray"
)

// how old a lock should be
const ()

// menu items - fyne specific
var (
	menuInformation *systray.MenuItem
	menuRotate      *systray.MenuItem
	menuQuit        *systray.MenuItem
	menu            *fyne.Menu
)

// preference based settings
var (
	tickDuration   time.Duration
	lockMaxAge     time.Duration
	dateTimeFormat string = ""
)

var (
	mu *sync.Mutex
)

var (
	Track tracker.Track
)

func StartApp(tr tracker.Track) {
	Track = tr
	mu = &sync.Mutex{}

	// get preferences
	// - times
	dateTimeFormat = pref.PREFERENCES.DateTimeFormat.Get()
	// - ticks
	tickDuration = pref.PREFERENCES.Tick.Get()
	// - locks
	lockMaxAge = pref.PREFERENCES.LockMaxAge.Get()

	systray.Run(SystraySetup, func() {
		defer debugger.Log("systray.Run", debugger.INFO, "exiting")()
	})
	// setup the app tray

}
