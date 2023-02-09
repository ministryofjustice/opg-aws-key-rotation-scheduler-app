package opgapp

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"os"
)

// Bootstrap creates the folder structure
func Bootstrap(s *Settings) (err error) {
	var path string
	var currentFile string

	// create the storage directory if it doesnt exist
	path = s.AccessKeys.Dir()
	if _, dirErr := os.Stat(path); os.IsNotExist(dirErr) {
		os.MkdirAll(path, os.ModePerm)
	}

	// create the current access key info if it doesnt exist
	currentFile = s.AccessKeys.CurrentFile()
	if _, fileErr := os.Stat(currentFile); os.IsNotExist(fileErr) {
		NewAccessKey(&s.AccessKeys).Save()
	}
	defer debugger.Log("Bootstrap", debugger.INFO, "directory:", path, "\ncurrent:", currentFile)()
	return
}
