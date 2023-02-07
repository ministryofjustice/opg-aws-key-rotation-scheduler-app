package opgapp

import (
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
	return
}
