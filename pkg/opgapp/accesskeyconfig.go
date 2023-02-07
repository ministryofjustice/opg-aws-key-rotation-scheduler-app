package opgapp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type AccessKeyConfiguration struct {
	Directory       string `json:"directory"`
	CurrentFilePath string `json:"current_access_key"`
	LockFilePath    string `json:"lock_file"`
}

func (akc *AccessKeyConfiguration) Dir() (path string) {
	home, _ := os.UserHomeDir()
	path = strings.ReplaceAll(akc.Directory, "~/", home+"/")
	return
}

func (akc *AccessKeyConfiguration) CurrentFile() (f string) {
	f = filepath.Join(akc.Dir(), akc.CurrentFilePath)
	return
}

func (akc *AccessKeyConfiguration) LockFile() (f string) {
	f = filepath.Join(akc.Dir(), akc.LockFilePath)
	return
}

func (akc *AccessKeyConfiguration) Lock() (key *AccessKeyTracker) {
	lockfile := akc.LockFile()
	content, _ := ioutil.ReadFile(lockfile)
	json.Unmarshal([]byte(content), &key)
	return
}

func (akc *AccessKeyConfiguration) Current() (key *AccessKeyTracker) {
	lockfile := akc.CurrentFile()
	content, _ := ioutil.ReadFile(lockfile)
	json.Unmarshal([]byte(content), &key)
	return
}
