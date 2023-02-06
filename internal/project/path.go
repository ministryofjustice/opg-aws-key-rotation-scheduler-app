package project

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	ROOT_DIR   = filepath.Join(filepath.Dir(b), "../..")
)
