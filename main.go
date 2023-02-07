package main

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/internal/project"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/k0kubun/pp"
)

const (
	SettingsFile string = "/settings.json"
)

// Shard vars between top level funcstions
var (
	settings         *opgapp.Settings
	accessKeyTracker *opgapp.AccessKeyTracker
	supports         opgapp.Supports
)

func main() {
	mu := &sync.Mutex{}
	var rotate, information, errorMsg *fyne.MenuItem
	var menu *fyne.Menu
	var a fyne.App = app.New()

	settings = opgapp.LoadSettings(project.ROOT_DIR + SettingsFile)
	// supported checks
	supports = opgapp.IsSupported(settings, a)

	// create base files and structure for the app
	opgapp.Bootstrap(settings)
	// fetch last key data
	accessKeyTracker, _ = opgapp.CurrentAccessKey(settings)

	// create the app menus
	//	- rotate
	rotate = fyne.NewMenuItem(settings.Labels.Rotate, func() {})
	//	- information
	at := accessKeyTracker.RotateAt(settings.RotationFrequency)
	information = fyne.NewMenuItem(fmt.Sprintf(settings.Labels.NextRotation, at.Format(settings.DateTimeFormat)), func() {})
	information.Disabled = true
	// - error message about app state (missing requirements etc)
	errorMsg = fyne.NewMenuItem("", func() {})
	errorMsg.Disabled = true

	desk, _ := a.(desktop.App)
	// happy path, all supported
	if supports.Os && supports.Desktop && supports.AwsVault {
		menu = fyne.NewMenu(settings.Name, rotate, information)
		desk.SetSystemTrayMenu(menu)

		go func() {
			dur := time.Minute
			for range time.Tick(dur) {
				pp.Println("tick")
				accessKeyTracker = opgapp.UpdateMenu(information, menu, accessKeyTracker, settings, mu)
			}
		}()
	} else if supports.Os && supports.Desktop {
		// AWS Vault is not installed
		errorMsg.Label = settings.Errors.AwsVaultNotFoundError
		menu = fyne.NewMenu(settings.Name, errorMsg)
		desk.SetSystemTrayMenu(menu)
	}

	accessKeyTracker = opgapp.UpdateMenu(information, menu, accessKeyTracker, settings, mu)
	a.Run()

}
