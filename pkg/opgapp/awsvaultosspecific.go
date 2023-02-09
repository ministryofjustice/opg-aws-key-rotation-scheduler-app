package opgapp

import "opg-aws-key-rotation-scheduler-app/pkg/debugger"

type AwsVaultOsSpecific struct {
	Command string `json:"command"`
	Profile string `json:"profile"`
	Prompt  string `json:"prompt"`
}

func (avos *AwsVaultOsSpecific) Installed(settings *Settings) (installed bool) {
	installed = false
	if _, err := LookupWithEnv(avos.Command, settings); err == nil {
		installed = true
	}
	defer debugger.Log("AwsVaultConfig.Installed()", debugger.VERBOSE, "installed:", installed, "\ncommand:", avos.Command)()
	return

}
