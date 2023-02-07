package opgapp

type Labels struct {
	Init         string `json:"init"`
	NextRotation string `json:"next_rotation"`
	Rotating     string `json:"rotating"`
	Locked       string `json:"locked"`
}

type ErrorMessages struct {
	AwsVaultNotFoundError string `json:"aws_vault_not_found"`
	NotDesktop            string `json:"not_desktop"`
	UnsupportedOs         string `json:"unsupported_os"`
}
