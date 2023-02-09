package opgapp

import (
	"encoding/json"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"runtime"
)

type Settings struct {
	Name              string                 `json:"name"`
	DateTimeFormat    string                 `json:"date_time_format"`
	RotationFrequency string                 `json:"rotation_frequency"`
	AccessKeys        AccessKeyConfiguration `json:"access_keys"`
	AwsVault          AwsVaultConfig         `json:"aws_vault"`
	Labels            Labels                 `json:"labels"`
	Errors            ErrorMessages          `json:"errors"`
	Icons             MenuIcons              `json:"icons"`
	OsData            struct {
		Darwin OsInfo `json:"darwin"`
	} `json:"os"`
}

type OsInfo struct {
	Shell       string `json:"shell"`
	Theme       string `json:"theme"`
	LoadProfile string `json:"load_profile"`
	Which       string `json:"which"`
}

// Os returns the specific OS info for this os
func (s *Settings) Os() (info *OsInfo) {
	switch runtime.GOOS {
	case "darwin":
		info = &s.OsData.Darwin
	}
	defer debugger.Log("OS info", debugger.VERBOSE, info)()
	return
}

func (s *Settings) JsonBytes() (content string) {
	b, _ := json.Marshal(s)
	return string(b)
}

// ---
func LoadSettings(content []byte) (s *Settings) {

	s = &Settings{}
	json.Unmarshal(content, &s)
	return
}
