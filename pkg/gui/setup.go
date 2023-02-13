package gui

import (
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"sync"
	"time"

	"fyne.io/fyne/v2"
)

// how old a lock should be
const ()

// menu items - fyne specific
var (
	menuInformation *fyne.MenuItem
	menuRotate      *fyne.MenuItem
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

	// setup the app tray
	SystraySetup()
	cfg.Desktop.SetSystemTrayMenu(menu)
	cfg.Desktop.SetSystemTrayIcon(icons.Default(cfg.IsDarkMode))

	UpdateMenu()
	// trigger the activation policy to remove the docker icon etc
	cfg.App.Lifecycle().SetOnStarted(func() {
		go func() {
			time.Sleep(200 * time.Millisecond)
			setActivationPolicy()
		}()
	})

	cfg.App.Run()
}
