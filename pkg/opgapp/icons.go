package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"

	"fyne.io/fyne/v2"
)

// MenuItems are loaded from the settings file and
// contain icons for the various states of the app
type MenuIcons struct {
	Default  MenuIcon `json:"default"`
	Rotating MenuIcon `json:"rotating"`
	Locked   MenuIcon `json:"locked"`
}

// MenuItem is parsed from settings and contains
// a white and black file path to aid with OS
// theme switching
type MenuIcon struct {
	Black string `json:"black"`
	White string `json:"white"`
}

// Themed fetches the OS theme / mode and uses that to then
// provide a specific set of icons (ThemedIcons) to match
// the settings.
// Defaults to icons suitable for light themes (black images)
func (m *MenuIcons) Themed(s *Settings) (icons ThemeIcons) {
	icons = ThemeIcons{DefaultIcon: m.Default.Black, RotatingIcon: m.Rotating.Black, LockedIcon: m.Locked.Black}
	if OsTheme() == "dark" {
		icons = ThemeIcons{DefaultIcon: m.Default.White, RotatingIcon: m.Rotating.White, LockedIcon: m.Locked.White}
	}
	defer debugger.Log("Themed icons", debugger.VERBOSE, icons)()
	return
}

// ThemeIcons is a created by the menuicons to hold dark/light specific versions
type ThemeIcons struct {
	DefaultIcon  string
	RotatingIcon string
	LockedIcon   string
}

// Default generates a Resource based on the default file path
func (ti ThemeIcons) Default() (r fyne.Resource) {
	path := ti.DefaultIcon
	content, _ := _iconFS.ReadFile(path)
	r = fyne.NewStaticResource("default-icon", content)
	return
}

func (ti ThemeIcons) Rotating() (r fyne.Resource) {
	path := ti.RotatingIcon
	content, _ := _iconFS.ReadFile(path)
	r = fyne.NewStaticResource("rotating-icon", content)
	return
}

func (ti ThemeIcons) Locked() (r fyne.Resource) {
	path := ti.LockedIcon
	content, _ := _iconFS.ReadFile(path)
	r = fyne.NewStaticResource("locked-icon", content)
	return
}
