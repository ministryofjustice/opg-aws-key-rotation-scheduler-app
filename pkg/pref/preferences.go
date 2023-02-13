package pref

import (
	"time"
)

type AppPreferences struct {
	// log info
	Debug        PrefData[bool]
	CpuProfiling PrefData[bool]
	// formats
	DateTimeFormat PrefData[string]
	// key tracking
	Tick            PrefData[time.Duration]
	TrackerLifetime PrefData[string]
	LockMaxAge      PrefData[time.Duration]
	// aws cli
	ProfileTool     PrefData[string]
	ProfileIdentity PrefData[string]
	// aws vault
	VaultTool PrefData[string]
}
