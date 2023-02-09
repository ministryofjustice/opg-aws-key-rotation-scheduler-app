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

While rotating your keys a lock file is created (`~/.opg/aws-key-rotation/locked.v0`) to act as a semaphore and block access / execution relating to key changes. 

While in this state the `Rotate now` menu item is also disabled.

### Locked

If the application crashes, is quit out or you select `cancel` option while entering key chain / MFA details the application will enter a locked state. The system tracy icon is visually different to show this (coloured red). 

After an hour the lock file will be automatically removed and the key rotation process restarted, but if you wish to trigger this beforehand you can remove the lock file manually (`rm ~/.opg/aws-key-rotation/locked.v0`).

### Default

Every minute (this will be reduced after testing) the application checks the age of its tracking data (stored in a file at `~/.opg/aws-key-rotation/current.v0`) against the configured rotating interval.

If it is older, then a rotation is triggered.


