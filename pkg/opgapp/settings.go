package opgapp

import (
	"encoding/json"
	"io/ioutil"
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
	Shell string `json:"shell"`
	Theme string `json:"theme"`
}

// Os returns the specific OS info for this os
func (s *Settings) Os() (info OsInfo) {
	switch runtime.GOOS {
	case "darwin":
		info = s.OsData.Darwin
	}
	return
}

// ---
func LoadSettings(file string) (s *Settings) {
	s = &Settings{}
	content, _ := ioutil.ReadFile(file)
	json.Unmarshal([]byte(content), &s)
	return
}
