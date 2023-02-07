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

// mutex lock

func updateInformationMenu(
	info *fyne.MenuItem,
	menu *fyne.Menu,
	key *opgapp.AccessKeyTracker,
	s *opgapp.Settings,
	mu *sync.Mutex,
) *opgapp.AccessKeyTracker {

	now := time.Now().UTC()
	at := key.RotateAt(s.RotationFrequency)
	lockfile := s.AccessKeys.LockFile()
	current := s.AccessKeys.CurrentFile()

	pp.Printf("[%s] next rotation at [%s]\n", now, at)

	mu.Lock()

	if key.Locked(lockfile) && time.Since(key.LockedAt(lockfile)) > time.Hour {
		pp.Println("Key is locked and too old, so removing...")
		key.Unlock(lockfile)
	} else if key.Locked(lockfile) {
		pp.Println("Key is locked...")
		info.Label = s.Labels.Locked
		// add icon change - LOCKED
		menu.Refresh()
	} else if now.After(at) {
		pp.Println("Rotating key...")
		key.Lock(lockfile)
		// add icon change - UPDATING
		info.Label = s.Labels.Rotating
		menu.Refresh()

		opgapp.RotateCommand(s)
		key.Unlock(lockfile)
		key = key.Rotate(current)

		// add icon change - NORMAL
		at = key.RotateAt(s.RotationFrequency)
		info.Label = fmt.Sprintf(s.Labels.NextRotation, at.Format(s.DateTimeFormat))
	} else {
		at = key.RotateAt(s.RotationFrequency)
		info.Label = fmt.Sprintf(s.Labels.NextRotation, at.Format(s.DateTimeFormat))

	}
	menu.Refresh()

	mu.Unlock()
	return key
}

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
	rotate = fyne.NewMenuItem("Rotate", func() {})
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
			for range time.Tick(time.Minute) {
				pp.Println("tick")
				accessKeyTracker = updateInformationMenu(information, menu, accessKeyTracker, settings, mu)
			}
		}()
	} else if supports.Os && supports.Desktop {
		// AWS Vault is not installed
		errorMsg.Label = settings.Errors.AwsVaultNotFoundError
		menu = fyne.NewMenu(settings.Name, errorMsg)
		desk.SetSystemTrayMenu(menu)
	}

	a.Run()
}
