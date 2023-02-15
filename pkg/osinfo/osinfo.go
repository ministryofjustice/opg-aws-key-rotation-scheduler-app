package osinfo

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo/darwin"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
)

type OsInfo interface {
	Supported() bool
	Prompt() string
	DarkMode(sh shell.Shell) bool
	Errors(sh shell.Shell, appName string, errs []string)
}

func New() (os OsInfo) {
	os = &darwin.Darwin{}
	defer debugger.Log("osinfo.New()", debugger.VERBOSE, os)()
	return
}
