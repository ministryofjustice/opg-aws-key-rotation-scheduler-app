package opgapp

import "os/exec"

type AwsVaultOsSpecific struct {
	Command string `json:"command"`
	Profile string `json:"profile"`
	Prompt  string `json:"prompt"`
}

func (avos *AwsVaultOsSpecific) Installed() (installed bool) {
	installed = false

	if _, err := exec.LookPath(avos.Command); err == nil {
		installed = true
	}
	return

}
