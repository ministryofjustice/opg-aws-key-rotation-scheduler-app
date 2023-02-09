package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"runtime"
)

type AwsVaultConfig struct {
	Darwin AwsVaultOsSpecific `json:"darwin"`
}

func (avc *AwsVaultConfig) OsSpecific() (osSpecific AwsVaultOsSpecific) {

	switch runtime.GOOS {
	case "darwin":
		osSpecific = avc.Darwin
	}
	defer debugger.Log("AwsVaultConfig.OsSpecific()", debugger.VERBOSE, osSpecific)()
	return
}
