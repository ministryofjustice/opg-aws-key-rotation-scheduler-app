package opgapp

import (
	"fmt"
	"os/exec"
	"strings"
)

func RotateCommand(
	s *Settings,
) (err error) {
	osInfo := _os
	osCfg := s.AwsVault.OsSpecific()

	output := new(strings.Builder)
	outerr := new(strings.Builder)

	// MAC ONLY
	// 	- aws-vault rotate identity --prompt=osascript
	zsh, _ := exec.LookPath(osInfo.Shell)
	c := &exec.Cmd{
		Path: zsh,
		Args: []string{
			"-s", "-c",
			fmt.Sprintf(
				"%s && %s rotate %s --prompt=%s",
				osInfo.LoadProfile,
				osCfg.Command,
				osCfg.Profile,
				osCfg.Prompt,
			),
		},
		Stdout: output,
		Stderr: outerr,
	}
	err = c.Run()

	// res := strings.ReplaceAll(strings.ToLower(output.String()), "\n", "")
	// _app.SendNotification(fyne.NewNotification("vault", res))
	// if err != nil {
	// 	_app.SendNotification(fyne.NewNotification("vault err", err.Error()))
	// }
	// _app.SendNotification(fyne.NewNotification("vault err", outerr.String()))

	return

}

// LookupWithEnv intended to be used in similar way as `exec.Lookup` but
// this loads in the users profile (`.zprofile`) first to ensure
// generated paths from brew etc are loaded
//   - uses `which` on macos
func LookupWithEnv(name string, s *Settings) (res string, err error) {
	output := new(strings.Builder)
	outerr := new(strings.Builder)

	sh, _ := exec.LookPath(s.Os().Shell)
	c := &exec.Cmd{
		Path: sh,
		Args: []string{
			"-s", "-c",
			s.Os().LoadProfile,
			fmt.Sprintf("%s %s", s.Os().Which, name),
		},
		Stdout: output,
		Stderr: outerr,
	}
	err = c.Run()
	res = strings.ReplaceAll(strings.ToLower(output.String()), "\n", "")
	return
}
