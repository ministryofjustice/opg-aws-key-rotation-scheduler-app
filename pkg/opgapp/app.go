package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/vault"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	AppId   string = "com.opg-aws-key-rotation.app"
	AppName string = "OPG AWS Key Rotation"
)

var (
	App         fyne.App
	Preferences fyne.Preferences
	Window      fyne.Window
	Desktop     desktop.App
)

// this app specific items that are shared
var (
	Shell   shell.Shell     // System supported shell
	Os      osinfo.OsInfo   // Details about the Os
	Profile profile.Profile // Profile is what will be used in the aws-vault call
	Vault   vault.Vault     // Vault (aws-vault) setup
)

// configuration checks for the environment
var (
	Booting    bool = true
	IsDarkMode bool = false
	IsDesktop  bool = false
	DEBUG      bool = true
	PROFILING  bool = false
)

// debug preferences
const (
	debugPreferencesKey      string = "debug"
	debugFallback            bool   = false
	debugLevelPreferencesKey string = "debugger_level"
	debugLevelFallback       int    = debugger.DETAILED
	debugEnvVarName          string = "OPGAWSKEYROTATION_DEBUG"
)

// profiling preferences
const (
	profileEnvVarName string = "OPGAWSKEYROTATION_PROFILE"
)

func debugSetup() {
	// enable debug mode or not
	DEBUG = Preferences.BoolWithFallback(debugPreferencesKey, debugFallback)
	// also overwrite from os env vars
	if envVar := os.Getenv(debugEnvVarName); len(envVar) > 0 {
		DEBUG = true
	}
	// set the debugger level
	debugger.LEVEL = Preferences.IntWithFallback(debugLevelPreferencesKey, debugLevelFallback)
	if DEBUG {
		debugger.LEVEL = debugger.ALL
	}
}

func init() {
	// fyne app setup
	App = app.NewWithID(AppId)
	Preferences = App.Preferences()
	Window = App.NewWindow(AppName)
	Desktop, IsDesktop = App.(desktop.App)
	// our code setup
	Shell = shell.New()
	Os = osinfo.New()
	Profile = profile.New()
	Vault = vault.New()

	debugSetup()

	if envVar := os.Getenv(profileEnvVarName); len(envVar) > 0 {
		PROFILING = true
	}
}
