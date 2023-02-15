package pref

import (
	"encoding/json"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"time"
)

var (
	_preferences map[string]string
	_appName     string

	PREFERENCES *AppPreferences = &AppPreferences{}
	// all the fallbacks
	Fbs = map[string]string{
		"debug":              "false",
		"cpu_profiling":      "false",
		"date_time_format":   "02-Jan-2006 15:04",
		"tick":               "1m",
		"rotation_frequency": "168h",
		"lock_max_age":       "5m",
		"profile_cli_tool":   "aws",
		"profile_name":       "identity",
		"vault_tool":         "aws-vault",
	}
)

func New(appName string, appPreferencesFileContent string, sh shell.Shell) (ap *AppPreferences) {
	_appName = appName
	err := json.Unmarshal([]byte(appPreferencesFileContent), &_preferences)
	if err != nil {
		panic(err)
	}
	ap = &AppPreferences{
		Debug:           newPD[bool](&_appName, &_preferences, &sh, "debug", Fbs["debug"], strToBool),
		CpuProfiling:    newPD[bool](&_appName, &_preferences, &sh, "cpu_profiling", Fbs["cpu_profiling"], strToBool),
		DateTimeFormat:  newPD[string](&_appName, &_preferences, &sh, "date_time_format", Fbs["date_time_format"], strToStr),
		Tick:            newPD[time.Duration](&_appName, &_preferences, &sh, "tick", Fbs["tick"], strToDuration),
		TrackerLifetime: newPD[string](&_appName, &_preferences, &sh, "rotation_frequency", Fbs["rotation_frequency"], strToStr),
		LockMaxAge:      newPD[time.Duration](&_appName, &_preferences, &sh, "lock_max_age", Fbs["lock_max_age"], strToDuration),
		ProfileTool:     newPD[string](&_appName, &_preferences, &sh, "profile_cli_tool", Fbs["profile_cli_tool"], strToStr),
		ProfileIdentity: newPD[string](&_appName, &_preferences, &sh, "profile_name", Fbs["profile_name"], strToStr),
		VaultTool:       newPD[string](&_appName, &_preferences, &sh, "vault_tool", Fbs["vault_tool"], strToStr),
	}
	return
}
