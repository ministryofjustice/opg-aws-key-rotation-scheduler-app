package vault

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/osinfo"
	"opg-aws-key-rotation-scheduler-app/pkg/profile"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"opg-aws-key-rotation-scheduler-app/pkg/vault/awsvault"
)

type Vault interface {
	Supported(sh shell.Shell) bool
	Command(profile profile.Profile, os osinfo.OsInfo) string
}

func New() (v Vault) {
	v = &awsvault.AwsVault{}
	defer debugger.Log("vault.New()", debugger.VERBOSE, v)()
	return
}
