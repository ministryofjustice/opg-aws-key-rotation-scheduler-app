package cfg

import (
	"opg-aws-key-rotation-scheduler-app/pkg/cfg/appcfg"
	"opg-aws-key-rotation-scheduler-app/pkg/cfg/guicfg"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/vault"

	"fyne.io/fyne/v2"
)

// -- all gui related settings merged int
const (
	AppName = guicfg.AppName
	AppID   = guicfg.AppId
)

var (
	App fyne.App
)

var (
	Shell   shell.Shell     // System supported shell
	Os      osinfo.OsInfo   // Details about the Os
	Profile profile.Profile // Profile is what will be used in the aws-vault call
	Vault   vault.Vault     // Vault (aws-vault) setup

	IsDarkMode bool
	IsBooting  bool
)

func init() {
	App = guicfg.App

	// -- app code required vars
	Shell = appcfg.Shell     // System supported shell
	Os = appcfg.Os           // Details about the Os
	Profile = appcfg.Profile // Profile is what will be used in the aws-vault call
	Vault = appcfg.Vault     // Vault (aws-vault) setup

	IsDarkMode = appcfg.IsDarkMode
	IsBooting = appcfg.IsBooting

}
