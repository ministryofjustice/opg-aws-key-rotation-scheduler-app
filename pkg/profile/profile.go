package profile

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/profile/awsprofile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
)

type Profile interface {
	Name() string
	Supported(sh shell.Shell) (installed bool, identityProfileFound bool, regionSet bool)
}

func New() (p Profile) {
	p = &awsprofile.AwsProfile{}
	defer debugger.Log("profile.New()", debugger.VERBOSE, p)()
	return
}
