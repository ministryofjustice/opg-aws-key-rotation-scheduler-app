package main

import (
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/errors"
	"opg-aws-key-rotation-scheduler-app/pkg/gui"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
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
	if !cfg.Os.Supported() {
		errs = append(errs, errors.UnsupportedOs)
	}
	if !cfg.Shell.Supported() {
		errs = append(errs, errors.UnsupportedShell)
	}

	if len(errs) == 0 {
		p, pf, rs := cfg.Profile.Supported(cfg.Shell)
		if !p {
			errs = append(errs, errors.ProfileCLINotFound)
		}
		if !pf {
			errs = append(errs, errors.ProfileNotFound)
		}
		if !rs {
			errs = append(errs, errors.RegionNotSet)
		}
		if !cfg.Vault.Supported(cfg.Shell) {
			errs = append(errs, errors.VaultNotFound)
		}
	}
	return
}

func init() {
	// config the preferences data with info from cfg
	pref.PREFERENCES = pref.New(cfg.AppName, cfg.Preferences)

}

func main() {
	// turn on debug
	if pref.PREFERENCES.Debug.Get() {
		debugger.LEVEL = debugger.ALL
	}
	if pref.PREFERENCES.CpuProfiling.Get() {
		cpuF, _ := profiling()
		pprof.StartCPUProfile(cpuF)
		defer pprof.StopCPUProfile()
	}

	// key tracker - down outside of app to avoid cyclic imports
	Track = tracker.New()
	// check for support
	supportErrors = supported()
	if !cfg.IsDesktop {
		supportErrors = append(supportErrors, errors.IsNotDesktop)
	}

	if len(supportErrors) > 0 {
		window := gui.ErrorDialog(cfg.App, cfg.Window, supportErrors)
		window.ShowAndRun()
	} else {
		cfg.IsDarkMode = cfg.Os.DarkMode(cfg.Shell)
		gui.StartApp(Track)
	}

}
