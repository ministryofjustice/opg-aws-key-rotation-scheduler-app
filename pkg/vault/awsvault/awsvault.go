package awsvault

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"strconv"
)

const (
	rotate string = "%s rotate %s --prompt=%s"
)

var (
	path string = ""
)

type AwsVault struct{}

func (v *AwsVault) Supported(sh shell.Shell) (supported bool) {
	command := pref.PREFERENCES.VaultTool.Get()
	p, _, _, err := sh.Search(command, true)
	supported = (err == nil)
	if supported {
		path = p
	}
	defer debugger.Log("AwsVault.Supported()", debugger.INFO, "supported:\t"+strconv.FormatBool(supported), "path:\t\t"+path)()
	return
}

func (v *AwsVault) Command(p profile.Profile, os osinfo.OsInfo) (str string) {
	str = fmt.Sprintf(rotate, path, p.Name(), os.Prompt())
	defer debugger.Log("AwsVault.Command()", debugger.INFO, str)()
	return
}
