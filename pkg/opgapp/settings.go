package opgapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Settings struct {
	Name string `json:"name"`

	AccessKeys struct {
		Directory string `json:"directory"`
		Current   string `json:"current_access_key"`
		LockFile  string `json:"lock_file"`
	} `json:"access_keys"`

	AwsVault struct {
		Command       string `json:"command"`
		Profile       string `json:"profile"`
		Shell         string `json:"shell"`
		Prompt        string `json:"prompt"`
		NotFoundError string `json:"not_found_error"`
	} `json:"aws_vault"`

	Menu struct {
		Rotation struct {
			Frequency  string `json:"frequency"`
			DateFormat string `json:"date_format"`
		} `json:"rotation"`
		Labels struct {
			NextRotation string `json:"next_rotation"`
			Rotating     string `json:"rotating"`
			Locked       string `json:"locked"`
		} `json:"labels"`
	} `json:"menu"`
}

func (s *Settings) Directory() (path string) {
	home, _ := os.UserHomeDir()
	path = strings.ReplaceAll(s.AccessKeys.Directory, "~/", home+"/")
	return
}

func (s *Settings) CurrentAccessKeyFilepath() (f string) {
	f = fmt.Sprintf("%s/%s", s.Directory(), s.AccessKeys.Current)
	return
}

func (s *Settings) LockFilepath() (f string) {
	f = fmt.Sprintf("%s/%s", s.Directory(), s.AccessKeys.LockFile)
	return
}

func LoadSettings(file string) (s Settings) {
	s = Settings{}
	content, _ := ioutil.ReadFile(file)
	json.Unmarshal([]byte(content), &s)
	return
}
