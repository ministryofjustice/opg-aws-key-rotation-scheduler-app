package opgapp

import (
	"io/ioutil"
	"opg-aws-key-rotation-scheduler-app/internal/project"
	"path/filepath"

	"fyne.io/fyne/v2"
	"github.com/k0kubun/pp"
)

// MenuItems are loaded from the settings file and
// contain icons for the various states of the app
type MenuIcons struct {
	Default  MenuIcon `json:"default"`
	Locked   MenuIcon `json:"locked"`
	Rotating MenuIcon `json:"rotating"`
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
	icons = ThemeIcons{DefaultIcon: m.Default.Black, LockedIcon: m.Locked.Black, RotatingIcon: m.Rotating.Black}
	mode := OsTheme(s)
	pp.Println(mode)
	if OsTheme(s) == "dark" {
		icons = ThemeIcons{DefaultIcon: m.Default.White, LockedIcon: m.Locked.White, RotatingIcon: m.Rotating.White}
	}
	return
}

// ThemeIcons is a created by the menuicons to hold dark/light specific versions
type ThemeIcons struct {
	DefaultIcon  string
	LockedIcon   string
	RotatingIcon string
}

// Default generates a Resource based on the default file path
func (ti *ThemeIcons) Default() (r fyne.Resource) {
	path := filepath.Join(project.ROOT_DIR, ti.DefaultIcon)
	content, _ := ioutil.ReadFile(path)
	r = fyne.NewStaticResource("default-icon", content)
	return
}

// Locked generates a Resource based on the locked file path
func (ti *ThemeIcons) Locked() (r fyne.Resource) {
	path := filepath.Join(project.ROOT_DIR, ti.LockedIcon)
	content, _ := ioutil.ReadFile(path)
	r = fyne.NewStaticResource("locked-icon", content)
	return
}

// Rotating generates a Resource based on the rotating file path
func (ti *ThemeIcons) Rotating() (r fyne.Resource) {
	path := filepath.Join(project.ROOT_DIR, ti.RotatingIcon)
	content, _ := ioutil.ReadFile(path)
	r = fyne.NewStaticResource("rotating-icon", content)
	return
}
