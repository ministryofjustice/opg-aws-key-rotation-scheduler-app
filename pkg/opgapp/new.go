package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/internal/project"
	"path/filepath"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	_settings        *Settings
	_labels          *Labels
	_errors          *ErrorMessages
	_os              *OsInfo
	_supported       *Supports
	_track           *AccessKeyTracker
	_rotateFrequency string
	_icons           ThemeIcons
)

var (
	_mu *sync.Mutex
)

var (
	_app             fyne.App
	_menuInformation *fyne.MenuItem
	_menuRotate      *fyne.MenuItem
	_menu            *fyne.Menu
)

func New(
	settingsFile string,
) {
	_mu = &sync.Mutex{}
	_app = app.NewWithID("opg-aws-key-rotation")

	_settings = LoadSettings(filepath.Join(project.ROOT_DIR, settingsFile))
	_rotateFrequency = _settings.RotationFrequency
	_os = _settings.Os()
	_labels = &_settings.Labels
	_errors = &_settings.Errors
	_supported = IsSupported(_settings, _app)
	_icons = _settings.Icons.Themed(_settings)

	Bootstrap(_settings)

	_track = _settings.AccessKeys.Current()

	// if everything is supported, use a standard menu setup and loop
	if _supported.Os && _supported.Desktop && _supported.AwsVault {
		SetupStandardMenu()

	} else if _supported.Os && _supported.Desktop {
		SetupErrorMenu(_errors.AwsVaultNotFoundError)
	}

	desk, _ := _app.(desktop.App)
	desk.SetSystemTrayMenu(_menu)
	_app.SetIcon(_icons.Default())
	UpdateMenu()
	_app.Run()

}
