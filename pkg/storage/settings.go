package storage

import (
	"io/fs"
	"os"
	"path/filepath"
)

var (
	UserHome              string
	storageDirectory      string
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

func StorageDirectory() string {
	if _, dirErr := os.Stat(storageDirectory); os.IsNotExist(dirErr) {
		os.MkdirAll(storageDirectory, StoragePermissionMode)
	}
	return storageDirectory
}
