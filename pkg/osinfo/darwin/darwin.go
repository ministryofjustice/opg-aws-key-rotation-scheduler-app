package darwin

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"runtime"
	"strings"
)

const (
	prompt          string = "osascript"
	darkModeCommand string = "defaults read -g AppleInterfaceStyle"
	darwin          string = "darwin"
	dark            string = "dark"
)

type Darwin struct{}

func (os *Darwin) Prompt() string {

	defer debugger.Log("Darwin.Prompt()", debugger.VERBOSE, prompt)()
	return prompt
}

func (os *Darwin) Supported() (supported bool) {
	supported = runtime.GOOS == darwin
	defer debugger.Log("Darwin.Supported()", debugger.INFO, supported)()
	return

}

func (os *Darwin) DarkMode(sh shell.Shell) (isDarkMode bool) {
	isDarkMode = false
	args := []string{darkModeCommand}
	stdout, _, err := sh.Run(args, false)
	if err == nil {
		mode := strings.ReplaceAll(strings.ToLower(stdout.String()), "\n", "")
		isDarkMode = (mode == dark)
	}
	defer debugger.Log("Darwin.DarkMode()", debugger.VERBOSE, isDarkMode)()
	return
}
