package main

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/internal/project"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	SettingsFile string = "/settings.json"
)

func updateInformationMenu(
	info *fyne.MenuItem,
	key *opgapp.AccessKeyTracker,
	s *opgapp.Settings,
) {
	now := time.Now().UTC()
	at := key.RotateAt(s.Menu.Rotation.Frequency)

	// - if lock file exists, change menu to show
	// - if the key is old, start the rotation process
	// - otherwise show the details of the key
	if key.Locked(s.LockFilepath()) {
		info.Label = s.Menu.Labels.Locked

	} else if now.After(at) {
		info.Label = s.Menu.Labels.Rotating
		key.Lock(s.LockFilepath())
		opgapp.RotateCommand(s)
		// 	key = key.Rotate(s.CurrentAccessKeyFilepath())
		// 	at = key.RotateAt(s.Menu.Rotation.Frequency)

	} else {
		info.Label = fmt.Sprintf(s.Menu.Labels.NextRotation, at.Format(s.Menu.Rotation.DateFormat))
	}

}

func main() {
	var accessKey opgapp.AccessKeyTracker
	var settings opgapp.Settings

	var rotate, information, noVault *fyne.MenuItem
	var menu *fyne.Menu
	var hasVault bool = false

	settings = opgapp.LoadSettings(project.ROOT_DIR + SettingsFile)
	// look for aws vault on the host machine
	if _, err := exec.LookPath(settings.AwsVault.Command); err == nil {
		hasVault = true
	}

	// create base files and structure for the app
	opgapp.Bootstrap(settings)
	// fetch last key data
	accessKey, _ = opgapp.CurrentAccessKey(settings)

	// create the app menus
	rotate = fyne.NewMenuItem("Rotate", func() {})
	information = fyne.NewMenuItem("", func() {})
	information.Disabled = true

	noVault = fyne.NewMenuItem(settings.AwsVault.NotFoundError, func() {})
	noVault.Disabled = true

	// generate the app
	a := app.New()

	if desk, ok := a.(desktop.App); ok {
		// if aws vault is installed, run the app as standard
		if hasVault {
			menu = fyne.NewMenu(settings.Name, rotate, information)
			desk.SetSystemTrayMenu(menu)
			updateInformationMenu(information, &accessKey, &settings)

			go func() {
				updateInformationMenu(information, &accessKey, &settings)
			}()

		} else {
			// otherwise show an error in the app
			menu = fyne.NewMenu(settings.Name, noVault)
			desk.SetSystemTrayMenu(menu)
		}
	}

	a.Run()
}
