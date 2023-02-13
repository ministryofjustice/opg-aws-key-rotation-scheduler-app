package main

import (
	"opg-aws-key-rotation-scheduler-app/pkg/errors"
	"opg-aws-key-rotation-scheduler-app/pkg/gui"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"opg-aws-key-rotation-scheduler-app/pkg/vault"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	App         fyne.App
	Window      fyne.Window
	Desktop     desktop.App
	Preferences fyne.Preferences

	Shell   shell.Shell     // System supported shell
	Os      osinfo.OsInfo   // Details about the Os
	Profile profile.Profile // Profile is what will be used in the aws-vault call
	Vault   vault.Vault
	Track   tracker.Track

	DarkMode  bool
	IsDesktop bool

	supportErrors []string = []string{}

	dateTimePreferencesKey string = "date_time_format"
	dateTimeFormat         string = ""
	dateTimeFormatFallback string = "02-Jan-2006 15:04"
)

// supported checks that all requirements for this app are met
// and returned a slice of error messages for anything that is not
// Checks:
//   - OS support (darwin)
//   - Shell support (zsh)
//   - Profile (aws profile installed, identity configured, region set on identity)
//   - Vault (aws vault found within $shell)
func supported() (errs []string) {
	if !Os.Supported() {
		errs = append(errs, errors.UnsupportedOs)
	}
	if !Shell.Supported() {
		errs = append(errs, errors.UnsupportedShell)
	}

	if len(errs) == 0 {
		p, pf, rs := Profile.Supported(Shell)
		if !p {
			errs = append(errs, errors.ProfileCLINotFound)
		}
		if !pf {
			errs = append(errs, errors.ProfileNotFound)
		}
		if !rs {
			errs = append(errs, errors.RegionNotSet)
		}
		if !Vault.Supported(Shell) {
			errs = append(errs, errors.VaultNotFound)
		}
	}
	return
}

func main() {
	// main app
	App = opgapp.App
	Preferences = opgapp.Preferences
	Window = opgapp.Window
	// os / config items
	Shell = shell.New()
	Os = osinfo.New()
	Profile = profile.New()
	Vault = vault.New()
	// key tracker
	Track = tracker.New()
	// get the datetime format
	dateTimeFormat = Preferences.StringWithFallback(dateTimePreferencesKey, dateTimeFormatFallback)
	// check for support
	supportErrors = supported()
	Desktop, IsDesktop = App.(desktop.App)
	if !IsDesktop {
		supportErrors = append(supportErrors, errors.IsNotDesktop)
	}

	if len(supportErrors) > 0 {
		window := gui.ErrorDialog(App, Window, supportErrors)
		window.ShowAndRun()
	} else {
		DarkMode = Os.DarkMode(Shell)
		gui.StartApp(
			App,
			Desktop,
			Window,
			Preferences,
			Shell,
			Os,
			Profile,
			Track,
			Vault,
			DarkMode,
			dateTimeFormat)
	}

}
