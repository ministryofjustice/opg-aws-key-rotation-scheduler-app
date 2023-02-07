package opgapp

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

	return

}
