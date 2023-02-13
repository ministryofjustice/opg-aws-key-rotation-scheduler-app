# OPG AWS Key Rotation Scheduler

This code base uses `golang` and `fyne` to create `macOS` applications (both `arm64` and `amd64`) that when installed will trigger the rotation of your identity access keys on a set interval.

While trailing the app between arch versions the interval is set as `two hours`, with a check running every minute. This will be lengthed when more testing has been done.

## Requirements

- `aws-vault` installed and configured

## Assumptions

- Default shell is `zsh`
- Default shell profile `.zprofile` exists and contains any custom pathing details (for example $PATH changes for homebrew installed versions of `aws-vault`)
- AWS primary profile is named `identity`

## Supports

- MacOS 10.11+
- Dark & Light mode

## Application states

The app itself has 3 states - default, rotating and locked - each of which show a differing system tray icon to highlight it.

### Rotating

You will be prompted for you key chain password and then for MFA codes for your AWS account. When these prompts occur you should see a visual change in the systray icon (generally orange) and the message state which `iam` role you are providing credintials for.

While rotating your keys a lock file is created (`~/.opg/aws-key-rotation/lock.v1`) to act as a semaphore and block access / execution relating to key changes. 

While in this state the `Rotate now` menu item is also disabled.

### Locked

If the application crashes, is quit out or you select `cancel` option while entering key chain / MFA details the application will enter a locked state. The system tracy icon is visually different to show this (coloured red). 

After an hour the lock file will be automatically removed and the key rotation process restarted, but if you wish to trigger this beforehand you can remove the lock file manually (`rm ~/.opg/aws-key-rotation/lock.v1`).

### Default

Every minute (this will be reduced after testing) the application checks the age of its tracking data (stored in a file at `~/.opg/aws-key-rotation/current.v1`) against the configured rotating interval.

If it is older, then a rotation is triggered.


# Setup for local development / build

Clone the repository to you local drive and from the root directory run:
```sh
make
```

If successful, an application will be created at:
```sh
./builds/darwin_${ARCH}/
```
Where `${ARCH}` is either `amd64` or `arm64`.

## Details

If you want to adjust this app or create a build locally then you will need to ensure some items are confgured and installed first:

- `macOS` >= 10.11
- `go` >= 1.19
- env variable `${GOBIN}` set to a folder path that exists
- `${GOBIN}` path is present within your `${PATH}`
- `fyne` >= 2.3.0 installed within `${GOBIN}` 
  - install manually by running  `go install fyne.io/fyne/v2/cmd/fyne@v2.3.0`

The `Makefile` will attempt to resolve these for you within the `requirements` target.

To create an app for your macOs version, from the root of the directory run `make` and check the content of the `./builds/` folder.

If you want to create application for all supported architectures, use `make all`


## Overrides

You can overwrite certain features of the application using environment variables rather than rebuilding the whole application:

Enable debug:
```sh
export OPGAWSKEYROTATION_DEBUG="1"
```

Enable CPU profiling:
```sh
export OPGAWSKEYROTATION_PROFILE="1"
```

Change rotation frequency:
```sh
export OPGAWSKEYROTATION_TRACKER_LIFETIME=2m
```


## CPU Profiling

When enbabled (via env var `OPGAWSKEYROTATION_PROFILE`) the application will genrate a cpu profile file that can then be check with `pprof`. This file is stored in the same folder as executable.

To run the app with profiling (based on running from this repositories directory):
```sh
export OPGAWSKEYROTATION_PROFILE="1"; go run main.go
```

To review the profile:
```sh
go tool pprof ${EXECUTABLE_NAME} cpu.prof
```
