package shell

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/shell/zsh"
	"strings"
)

type Shell interface {
	Supported() bool
	Self() (string, error)
	Profile() (profile string, err error)
	Search(commandName string, withProfile bool) (path string, stdout *strings.Builder, stderr *strings.Builder, err error)
	Run(args []string, withProfile bool) (stdout *strings.Builder, stderr *strings.Builder, err error)
}

// New the users shell for the app based off preferences
//   - Handle preferences needed, currently just returns zsh
func New() (s Shell) {
	s = &zsh.Zsh{}
	defer debugger.Log("shell.New()", debugger.VERBOSE, s)()
	return
}
