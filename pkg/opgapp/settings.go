package opgapp

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	Name              string                 `json:"name"`
	DateTimeFormat    string                 `json:"date_time_format"`
	RotationFrequency string                 `json:"rotation_frequency"`
	AccessKeys        AccessKeyConfiguration `json:"access_keys"`
	AwsVault          AwsVaultConfig         `json:"aws_vault"`
	Labels            Labels                 `json:"labels"`
	Errors            ErrorMessages          `json:"errors"`
}

// ---
func LoadSettings(file string) (s *Settings) {
	s = &Settings{}
	content, _ := ioutil.ReadFile(file)
	json.Unmarshal([]byte(content), &s)
	return
}
