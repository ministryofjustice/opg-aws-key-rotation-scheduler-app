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
	path string = ".opg/aws-key-rotation"
)

func init() {
	// directory storage
	UserHome, _ := os.UserHomeDir()
	storageDirectory = filepath.Join(UserHome, path)
}

func ProfileDirectory() string {
	perfProfileDirectory, _ = os.Getwd()
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
