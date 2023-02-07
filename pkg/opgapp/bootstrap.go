package opgapp

import (
	"os"
)

// Bootstrap creates the folder structure and (TODO) plist files
func Bootstrap(s *Settings) (err error) {

	// create the storage directory if it doesnt exist
	path := s.AccessKeys.Dir()
	if _, dirErr := os.Stat(path); os.IsNotExist(dirErr) {
		os.MkdirAll(path, os.ModePerm)
	}

	// create the current access key info if it doesnt exist
	currentFile := s.AccessKeys.CurrentFile()
	if _, fileErr := os.Stat(currentFile); os.IsNotExist(fileErr) {
		NewAccessKey(&s.AccessKeys).Save()
	}

	return
}
