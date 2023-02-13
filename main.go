package main

import (
	"opg-aws-key-rotation-scheduler-app/pkg/errors"
	"opg-aws-key-rotation-scheduler-app/pkg/gui"
	. "opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"opg-aws-key-rotation-scheduler-app/pkg/storage"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"os"
	"path/filepath"
	"runtime/pprof"
)

var (
	Track         tracker.Track
	supportErrors []string = []string{}
)

func profiling() (cpuFile *os.File, memoryFile *os.File) {
	pDir := storage.ProfileDirectory()
	cpu := filepath.Join(pDir, "cpu.prof")
	memory := filepath.Join(pDir, "memory.prof")
	cpuFile, _ = os.Create(cpu)
	memoryFile, _ = os.Create(memory)
	return

}

// supported checks that all requirements for this app are met
// and returned a slice of error messages for anything that is not
// Checks:
//   - OS support (darwin)
//   - Shell support (zsh)
//   - Profile (aws profile installed, identity configured, region set on identity)
//   - Vault (aws vault found within $shell)
func supported() (errs []string) {
	if !Os.Supported() {
		errs = append(errs, errors.UnsupportedOs)
	}
	if !Shell.Supported() {
		errs = append(errs, errors.UnsupportedShell)
	}

	if len(errs) == 0 {
		p, pf, rs := Profile.Supported(Shell)
		if !p {
			errs = append(errs, errors.ProfileCLINotFound)
		}
		if !pf {
			errs = append(errs, errors.ProfileNotFound)
		}
		if !rs {
			errs = append(errs, errors.RegionNotSet)
		}
		if !Vault.Supported(Shell) {
			errs = append(errs, errors.VaultNotFound)
		}
	}
	return
}

func main() {
	if PROFILING {
		cpuF, _ := profiling()
		pprof.StartCPUProfile(cpuF)
		defer pprof.StopCPUProfile()
	}
	// key tracker - down outside of app to avoid cyclic imports
	Track = tracker.New()
	// check for support
	supportErrors = supported()
	if !IsDesktop {
		supportErrors = append(supportErrors, errors.IsNotDesktop)
	}

	if len(supportErrors) > 0 {
		window := gui.ErrorDialog(App, Window, supportErrors)
		window.ShowAndRun()
	} else {
		IsDarkMode = Os.DarkMode(Shell)
		gui.StartApp(Track)
	}

}
