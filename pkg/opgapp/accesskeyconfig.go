package opgapp

import (
	"fmt"
	"os"
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
	f = fmt.Sprintf("%s/%s", akc.Dir(), akc.CurrentFilePath)
	return
}

func (akc *AccessKeyConfiguration) LockFile() (f string) {
	f = fmt.Sprintf("%s/%s", akc.Dir(), akc.LockFilePath)
	return
}
