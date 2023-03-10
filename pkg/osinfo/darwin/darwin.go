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

var sysMessage string = `%s -e 'tell app "%s" to display alert "%s" message "%s" '`

type Darwin struct{}

func (os *Darwin) PromptCommand() string {
	defer debugger.Log("Darwin.PromptCommand()", debugger.VERBOSE, prompt)()
	return prompt
}

func (os *Darwin) Supported() (supported bool) {
	supported = runtime.GOOS == darwin
	defer debugger.Log("Darwin.Supported()", debugger.INFO, supported)()
	return

}

func (os *Darwin) DarkMode(sh shell.Shell) (isDarkMode bool) {
	isDarkMode = false
	stdout, _, err := sh.Run(darkModeCommand, false, true)
	if err == nil {
		mode := strings.ReplaceAll(strings.ToLower(stdout.String()), "\n", "")
		isDarkMode = (mode == dark)
	}
	defer debugger.Log("Darwin.DarkMode()", debugger.VERBOSE, isDarkMode)()
	return
}

func (os *Darwin) SystemMessage(sh shell.Shell, appName string, msgs []string, msgType string) {

	cmd := fmt.Sprintf(sysMessage, os.PromptCommand(), appName, msgType, strings.Join(msgs, "\n"))
	_, _, err := sh.Run(cmd, false, false)
	if err != nil {
		panic(err)
	}

}
