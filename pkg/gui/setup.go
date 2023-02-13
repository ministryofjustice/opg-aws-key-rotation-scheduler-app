package gui

import (
	"opg-aws-key-rotation-scheduler-app/pkg/icons"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"opg-aws-key-rotation-scheduler-app/pkg/vault"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	tickDuration time.Duration = time.Minute
	lockMaxAge   time.Duration = 10 * time.Minute
)

var (
	menuInformation *fyne.MenuItem
	menuRotate      *fyne.MenuItem
	menu            *fyne.Menu
)

var (
	mu             *sync.Mutex
	dateTimeFormat string
	isDark         bool
	booting        bool = true
)

var (
	app         fyne.App
	desktopApp  desktop.App
	preferences fyne.Preferences
	window      fyne.Window
	zsh         shell.Shell
	os          osinfo.OsInfo
	prof        profile.Profile
	track       tracker.Track
	awsVault    vault.Vault
)

func StartApp(
	a fyne.App,
	dApp desktop.App,
	win fyne.Window,
	pref fyne.Preferences,
	sh shell.Shell,
	o osinfo.OsInfo,
	p profile.Profile,
	tr tracker.Track,
	v vault.Vault,
	isDarkMode bool,
	format string,
) {
	mu = &sync.Mutex{}

	app = a
	desktopApp = dApp
	preferences = pref
	window = win
	track = tr
	isDark = isDarkMode
	dateTimeFormat = format
	awsVault = v
	zsh = sh
	os = o
	prof = p

	SystraySetup()
	desktopApp.SetSystemTrayMenu(menu)
	desktopApp.SetSystemTrayIcon(icons.Default(isDark))

	UpdateMenu()
	// trigger the activation policy to remove the docker icon etc
	app.Lifecycle().SetOnStarted(func() {
		go func() {
			time.Sleep(200 * time.Millisecond)
			setActivationPolicy()
		}()
	})

	app.Run()
}
