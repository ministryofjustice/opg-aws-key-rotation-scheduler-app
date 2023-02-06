package opgapp

import (
	"fmt"
	"os"
	"os/exec"
)

func RotateCommand(
	s *Settings,
) {

	// -- MAC ONLY
	zsh, _ := exec.LookPath(s.AwsVault.Shell) //"zsh"

	c := &exec.Cmd{
		Path: zsh,
		Args: []string{
			"-s", "-c",
			// "aws-vault rotate identity --prompt=osascript"
			fmt.Sprintf(
				"%s rotate %s --prompt=%s",
				s.AwsVault.Command,
				s.AwsVault.Profile,
				s.AwsVault.Prompt,
			),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	c.Run()

}
