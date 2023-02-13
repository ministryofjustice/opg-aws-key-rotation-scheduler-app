package gui

import (
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	. "opg-aws-key-rotation-scheduler-app/pkg/opgapp"
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

// items handled by preferences
var (
	// ticks - how often we check status of track
	tickDurationPreferencesKey string = "check_frequency"
	tickDurationFallback       string = "1m"
	tickDuration               time.Duration

	lockMaxAgePreferencesKey string = "lock_max_age"
	lockMaxAgeFallback       string = "10m"
	lockMaxAge               time.Duration

	// date time formats
	dateTimePreferencesKey string = "date_time_format"
	dateTimeFormatFallback string = "02-Jan-2006 15:04"
	dateTimeFormat         string = ""
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
	dateTimeFormat = Preferences.StringWithFallback(dateTimePreferencesKey, dateTimeFormatFallback)
	// - ticks
	tickDurStr := Preferences.StringWithFallback(tickDurationPreferencesKey, tickDurationFallback)
	tickDuration, _ = time.ParseDuration(tickDurStr)
	// - locks
	lockStr := Preferences.StringWithFallback(lockMaxAgePreferencesKey, lockMaxAgeFallback)
	lockMaxAge, _ = time.ParseDuration(lockStr)

	// setup the app tray
	SystraySetup()
	Desktop.SetSystemTrayMenu(menu)
	Desktop.SetSystemTrayIcon(icons.Default(IsDarkMode))

	UpdateMenu()
	// trigger the activation policy to remove the docker icon etc
	App.Lifecycle().SetOnStarted(func() {
		go func() {
			time.Sleep(200 * time.Millisecond)
			setActivationPolicy()
		}()
	})

	App.Run()
}
