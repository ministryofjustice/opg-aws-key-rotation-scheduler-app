package opgapp

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"os/exec"
	"strings"
)

func RotateCommand(
	s *Settings,
) (err error) {
	var (
		shellCmd       string
		osInfo         *OsInfo
		osCfg          AwsVaultOsSpecific
		output, outerr *strings.Builder
	)

	osInfo = _os
	osCfg = s.AwsVault.OsSpecific()
	output = new(strings.Builder)
	outerr = new(strings.Builder)

	// MAC ONLY
	// 	- aws-vault rotate identity --prompt=osascript
	zsh, _ := exec.LookPath(osInfo.Shell)
	shellCmd = fmt.Sprintf(
		"%s && %s rotate %s --prompt=%s",
		osInfo.LoadProfile,
		osCfg.Command,
		osCfg.Profile,
		osCfg.Prompt,
	)
	c := &exec.Cmd{
		Path:   zsh,
		Args:   []string{"-s", "-c", shellCmd},
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
	defer debugger.Log("RotateCommand()", debugger.INFO, osInfo, "\ncmd:", shellCmd, "\nstdout:", output, "\nerrors:", outerr, "\n", err)()

	return

}

// LookupWithEnv intended to be used in similar way as `exec.Lookup` but
// this loads in the users profile (`.zprofile`) first to ensure
// generated paths from brew etc are loaded
//   - uses `which` on macos
func LookupWithEnv(name string, s *Settings) (res string, err error) {
	var (
		shellCmd       string
		osInfo         *OsInfo
		output, outerr *strings.Builder
	)

	output = new(strings.Builder)
	outerr = new(strings.Builder)
	osInfo = _os
	shellCmd = fmt.Sprintf("%s %s", osInfo.Which, name)

	sh, _ := exec.LookPath(osInfo.Shell)
	c := &exec.Cmd{
		Path:   sh,
		Args:   []string{"-s", "-c", osInfo.LoadProfile, shellCmd},
		Stdout: output,
		Stderr: outerr,
	}
	err = c.Run()
	res = strings.ReplaceAll(strings.ToLower(output.String()), "\n", "")
	defer debugger.Log("LookupWithEnv()", debugger.INFO, "looking for:", name, "\ncmd", shellCmd, "\n", osInfo, "\nres:", res, "\n", err)()
	return
}
