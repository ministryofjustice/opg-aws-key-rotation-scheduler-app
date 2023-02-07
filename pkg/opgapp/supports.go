package opgapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// track what is supported
type Supports struct {
	Os       bool `default:"false"`
	AwsVault bool `default:"false"`
	Desktop  bool `default:"false"`
}

func IsSupported(
	settings *Settings,
	app fyne.App,
) (s Supports) {
	s = Supports{Os: false, AwsVault: false, Desktop: false}

	spec := settings.AwsVault.OsSpecific()
	// os is supported if a profile is returned from aws vault config
	// then see if vault is installed
	if len(spec.Command) > 0 {
		s.Os = true
		s.AwsVault = spec.Installed()
	}
	// check desktop
	_, desktop := app.(desktop.App)
	s.Desktop = desktop

	return
}
