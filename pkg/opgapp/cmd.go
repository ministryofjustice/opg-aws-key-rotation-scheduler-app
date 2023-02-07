package opgapp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RotateCommand(
	s *Settings,
) error {
	osInfo := s.Os()
	osCfg := s.AwsVault.OsSpecific()
	// MAC ONLY
	// 	- aws-vault rotate identity --prompt=osascript
	zsh, _ := exec.LookPath(osInfo.Shell)
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

func LookupWithEnv(name string, s *Settings) (res string, err error) {
	output := new(strings.Builder)
	outerr := new(strings.Builder)

	sh, _ := exec.LookPath(s.Os().Shell)
	c := &exec.Cmd{
		Path: sh,
		Args: []string{
			"-s", "-c",
			s.Os().LoadProfile,
			fmt.Sprintf("%s %s", s.Os().Which, name),
		},
		Stdout: output,
		Stderr: outerr,
	}
	err = c.Run()
	res = strings.ReplaceAll(strings.ToLower(output.String()), "\n", "")
	return
}
