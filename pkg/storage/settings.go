package storage

import (
	"io/fs"
	"os"
	"path/filepath"
)

var (
	UserHome              string
	storageDirectory      string
	perfProfileDirectory  string
	StoragePermissionMode fs.FileMode = 0755
)

const (
	path     string = ".opg/aws-key-rotation"
	perfPath string = "pref"
)

func init() {
	// directory storage
	UserHome, _ := os.UserHomeDir()
	storageDirectory = filepath.Clean(filepath.Join(UserHome, path))
}

func ProfileDirectory() string {
	perfProfileDirectory = filepath.Clean(filepath.Join(StorageDirectory(), perfPath))
	if _, dirErr := os.Stat(perfProfileDirectory); os.IsNotExist(dirErr) {
		err := os.MkdirAll(perfProfileDirectory, StoragePermissionMode)
		if err != nil {
			panic(err)
		}
	}
	return perfProfileDirectory
}

func StorageDirectory() string {
	if _, dirErr := os.Stat(storageDirectory); os.IsNotExist(dirErr) {
		err := os.MkdirAll(storageDirectory, StoragePermissionMode)
		if err != nil {
			panic(err)
		}
	}
	return storageDirectory
}
