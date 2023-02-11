package errors

const (
	VaultNotFound      string = "Vault not found. Please check your $PATH."
	ProfileCLINotFound string = "'%s' command not found. Please check $PATH is configured correctly."
	ProfileNotFound    string = "Profile '%s' not found. Please check you `config` file is setup correctly."
	RegionNotSet       string = "Profile '%s' requires a region to be configured. Please check your `config` file"
	IsNotDesktop       string = "This application only supports desktop operating systems"
	UnsupportedOs      string = "Operating system is not supported"
	UnsupportedShell   string = "Zsh does not seem to be supported on this machine."
)
