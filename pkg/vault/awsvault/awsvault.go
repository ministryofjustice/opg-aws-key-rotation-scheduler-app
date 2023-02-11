package awsvault

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
)

const (
	command string = "aws-vault"
	rotate  string = "aws-vault rotate %s --prompt=%s"
)

type AwsVault struct{}

func (v *AwsVault) Supported(sh shell.Shell) (supported bool) {
	_, _, _, err := sh.Search(command, true)
	supported = (err == nil)
	defer debugger.Log("AwsVault.Supported()", debugger.INFO, supported)()
	return
}

func (v *AwsVault) Command(p profile.Profile, os osinfo.OsInfo) (str string) {
	str = fmt.Sprintf(rotate, p.Name(), os.Prompt())
	defer debugger.Log("AwsVault.Command()", debugger.INFO, str)()
	return
}
