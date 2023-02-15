package appcfg

import (
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/vault"
	"strings"
)

const (
	AppId   string = "com.opg-aws-key-rotation.app"
	AppName string = "OPG AWS Key Rotation"
)

var (
	AppBuiltName string = strings.ReplaceAll(AppName, " ", "")
)

// this app specific items that are shared
var (
	Shell      shell.Shell     // System supported shell
	Os         osinfo.OsInfo   // Details about the Os
	Profile    profile.Profile // Profile is what will be used in the aws-vault call
	Vault      vault.Vault     // Vault (aws-vault) setup
	IsDarkMode bool
	IsBooting  bool
)

func init() {
	// our code setup
	Shell = shell.New()
	Os = osinfo.New()
	Profile = profile.New()
	Vault = vault.New()
}
