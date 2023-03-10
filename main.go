package main

import (
	_ "embed"
	"opg-aws-key-rotation-scheduler-app/pkg/cfg"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/errors"
	"opg-aws-key-rotation-scheduler-app/pkg/gui"
	"opg-aws-key-rotation-scheduler-app/pkg/labels"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/storage"
	"opg-aws-key-rotation-scheduler-app/pkg/tracker"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strconv"
)

//go:embed preferences.json
var preferences string

var (
	Track         tracker.Track
	supportErrors []string = []string{}
)

func profiling() (cpuFile *os.File) {
	cpuFileName := "cpu.prof"
	pDir := storage.ProfileDirectory()
	cpu := filepath.Clean(filepath.Join(pDir, cpuFileName))
	debugger.Log("main.profiling()", debugger.INFO, "cpu profile file:\t"+cpu)()
	cpuFile, _ = os.Create(cpu)
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
	debugger.Create()
	// config the preferences data with info from cfg
	pref.PREFERENCES = pref.New(cfg.AppBuiltName, preferences, cfg.Shell)
}

func main() {
	var err error
	// turn on debug
	debug := pref.PREFERENCES.Debug.Get()
	debugger.Log("main", debugger.INFO, "debug:\t"+strconv.FormatBool(debug))()
	if debug {
		debugger.LEVEL = debugger.ALL
	}

	prof := pref.PREFERENCES.CpuProfiling.Get()
	debugger.Log("main", debugger.INFO, "cpu profiling:\t"+strconv.FormatBool(prof))()
	if prof {
		cpuF := profiling()
		err = pprof.StartCPUProfile(cpuF)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	// key tracker - down outside of app to avoid cyclic imports
	Track = tracker.New()
	// check for support
	supportErrors = supported()

	if len(supportErrors) == 0 {
		cfg.IsDarkMode = cfg.Os.DarkMode(cfg.Shell)
		gui.StartApp(Track)
	} else {
		cfg.Os.SystemMessage(cfg.Shell, cfg.AppBuiltName, supportErrors, labels.SystemError)
	}

}
