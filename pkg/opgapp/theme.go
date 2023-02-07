package opgapp

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// OsTheme runs an exec command to try and work out the
// OS theme being used at the moment.
// Defaults to light
//   - Only supports MacOs
//   - shell and command to run comes from settings.json
func OsTheme() (mode string) {
	mode = "light"
	switch runtime.GOOS {
	case "darwin":
		osInfo := _os
		sh, _ := exec.LookPath(osInfo.Shell)
		output := new(strings.Builder)
		c := &exec.Cmd{
			Path: sh,
			Args: []string{
				"-s", "-c", osInfo.Theme,
			},
			Stdout: output,
			Stderr: os.Stderr,
		}
		err := c.Run()
		if err == nil {
			// lower case, trim new lines
			mode = strings.ReplaceAll(strings.ToLower(output.String()), "\n", "")
		}
	}
	return
}
