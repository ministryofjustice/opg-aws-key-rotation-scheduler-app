package darwin

import (
	"fmt"
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
	stdout, _, err := sh.Run(args, false, true)
	if err == nil {
		mode := strings.ReplaceAll(strings.ToLower(stdout.String()), "\n", "")
		isDarkMode = (mode == dark)
	}
	defer debugger.Log("Darwin.DarkMode()", debugger.VERBOSE, isDarkMode)()
	return
}

func (os *Darwin) Errors(sh shell.Shell, app string, errors []string) {
	cmd := fmt.Sprintf(`%s -e 'tell app "%s" to display alert "Errors" message "- %s" '`, os.Prompt(), app, strings.Join(errors, "\n - "))
	sh.Run([]string{cmd}, false, false)
}
