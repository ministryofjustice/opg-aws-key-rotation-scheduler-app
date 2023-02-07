package opgapp

import (
	"fmt"
	"os"
	"os/exec"
)

func RotateCommand(
	s *Settings,
) error {

	osCfg := s.AwsVault.OsSpecific()
	// MAC ONLY
	// 	- aws-vault rotate identity --prompt=osascript
	zsh, _ := exec.LookPath(osCfg.Shell)
	c := &exec.Cmd{
		Path: zsh,
		Args: []string{
			"-s", "-c",

			fmt.Sprintf(
				"%s rotate %s --prompt=%s",
				osCfg.Command,
				osCfg.Profile,
				osCfg.Prompt,
			),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	return c.Run()

}
