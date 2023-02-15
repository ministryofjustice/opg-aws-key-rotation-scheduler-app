package osinfo

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo/darwin"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
)

type OsInfo interface {
	Supported() bool
	PromptCommand() string
	DarkMode(sh shell.Shell) bool
	SystemMessage(sh shell.Shell, appName string, msgs []string, msgType string)
}

func New() (os OsInfo) {
	os = &darwin.Darwin{}
	defer debugger.Log("osinfo.New()", debugger.VERBOSE, os)()
	return
}
