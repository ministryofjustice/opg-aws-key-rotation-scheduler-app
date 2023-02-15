package zsh

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	which    string = "which"
	load     string = "source"
	profiles string = ".zprofile,.profile"
)

const (
	self     string = "zsh"
	darwin   string = "darwin"
	notFound string = "not found"
)

const (
	ERR_NO_PROFILE string = "no user profiles found"
)

// -- main shell
type Zsh struct {
	self string
}

// Self finds the system path for zsh bin
func (sh *Zsh) Self() (path string, err error) {
	if len(sh.self) == 0 {
		path, err = exec.LookPath(self)
		sh.self = path
		defer debugger.Log("Zsh.Self()", debugger.INFO, "path:\t"+path, "err:", err)()
	} else {
		path = sh.self
		defer debugger.Log("Zsh.Self()", debugger.INFO, "[cached]", "path:\t"+path, "err:", err)()
	}
	return
}

// profile is internal and looks for possible user profile files
// as they might be in various locations
func (sh *Zsh) profilePath() (path string, err error) {
	home, _ := os.UserHomeDir()
	files := strings.Split(profiles, ",")
	for _, profile := range files {
		profilePath := filepath.Join(home, profile)
		if _, fErr := os.Stat(profilePath); fErr == nil {
			return profilePath, nil
		}
	}
	err = fmt.Errorf(ERR_NO_PROFILE)
	defer debugger.Log("Zsh.profile()", debugger.VERBOSE, "profile:", path, "error:", err)()
	return
}

func (sh *Zsh) Supported() (supported bool) {
	supported = runtime.GOOS == darwin
	defer debugger.Log("Zsh.Supported()", debugger.INFO, supported)()
	return

}

// Profile tries to return command to load a profile
func (sh *Zsh) Profile() (profile string, err error) {
	var profileFile string
	profile = ""
	if profileFile, err = sh.profilePath(); err == nil {
		profile = fmt.Sprintf("%s %s; ", load, profileFile)
	} else {
		defer debugger.Log("Zsh.Profile()", debugger.VERBOSE, "withProfile requested, but no profiles found")()
	}
	return
}

// Search will look for the `commandName` passed using `which` by running a zsh shell and adding commands
//   - if `withProfile` is true then `source` is used to load in a profile (if found)
func (sh *Zsh) Search(commandName string, withProfile bool) (path string, stdout *strings.Builder, stderr *strings.Builder, err error) {
	var pre string = ""

	// work out the profile loading
	if withProfile {
		pre, _ = sh.Profile()
		defer debugger.Log("Zsh.Search()", debugger.VERBOSE, "withProfile requested", "profile:", pre)()
	}

	cmd := fmt.Sprintf("%s%s %s", pre, which, commandName)
	stdout, stderr, err = sh.Run([]string{cmd}, false, true)
	path = stdout.String()

	// if there is an error, and the result contains "not found", or if
	// the result says not found return a more custom set of info
	if err != nil && strings.Contains(path, notFound) || strings.Contains(path, notFound) {
		err = fmt.Errorf(stdout.String())
		stdout = new(strings.Builder)
		path = ""
	} else {
		path = strings.TrimSuffix(path, "\n")

	}
	defer debugger.Log("Zsh.Search()", debugger.VERBOSE, "commandName:", commandName, "path:", path, "err:", err)()
	return
}

// Run creates and executes a cmd using zsh shell as a wrapper
func (sh *Zsh) Run(args []string, withProfile bool, withCache bool) (stdout *strings.Builder, stderr *strings.Builder, err error) {
	var pre string
	var pErr error

	stdout = new(strings.Builder)
	stderr = new(strings.Builder)
	shell, _ := sh.Self()

	cmdArgs := []string{"-s", "-c"}
	if withProfile {
		if pre, pErr = sh.Profile(); pErr == nil {
			cmdArgs = append(cmdArgs, pre)
		}
		defer debugger.Log("Zsh.Run()", debugger.VERBOSE, "withProfile requested", "profile:", pre)()
	}
	cmdArgs = append(cmdArgs, args...)

	if cached, found := fromCache(shell, cmdArgs); withCache && found {
		stdout = cached.Stdout
		stderr = cached.Stderr
		err = cached.Err
		defer debugger.Log("Zsh.Run()", debugger.INFO, "[cached]", "shell:", shell, "args:", cmdArgs, "err:", err)()
	} else {
		c := &exec.Cmd{
			Path:   shell,
			Args:   cmdArgs,
			Stdout: stdout,
			Stderr: stderr,
		}
		err = c.Run()
		defer debugger.Log("Zsh.Run()", debugger.INFO, "shell:", shell, "args:", cmdArgs, "err:", err)()
		defer toCache(shell, cmdArgs, stdout, stderr, err)
	}

	return
}
