package debugger

import (
	"fmt"
	"io"
	"io/fs"
	"opg-aws-key-rotation-scheduler-app/pkg/storage"
	"os"
	"path/filepath"
	"time"
)

var _stdout io.Writer
var _stderr io.Writer
var _fileMode fs.FileMode = 0755
var err error

func init() {

	directory := storage.StorageDirectory()
	stdoutfile := filepath.Clean(filepath.Join(directory, "stdout.log"))
	stderrfile := filepath.Clean(filepath.Join(directory, "stderr.log"))

	std, _ := os.OpenFile(stdoutfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, _fileMode)
	e, _ := os.OpenFile(stderrfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, _fileMode)

	_stdout = io.MultiWriter(os.Stdout, std)
	_stderr = io.MultiWriter(os.Stderr, e)

	_, err = _stdout.Write([]byte("logger starting up\n"))
	if err != nil {
		panic(err)
	}
	_, err = _stderr.Write([]byte("logger starting up\n"))
	if err != nil {
		panic(err)
	}

}

func Log(message string, level int, values ...interface{}) func() {
	out := _stdout
	if LEVEL == ERR {
		out = _stderr
	}
	t := time.Now().UTC()
	return func() {
		str := fmt.Sprintf("[%s](level:%s) %s\n", t, levelToString[level], message)
		for _, v := range values {
			str += fmt.Sprintf("%v\n", v)
		}
		str += "---------\n"
		show := (level <= LEVEL)
		if show {
			_, err = out.Write([]byte(str))
			if err != nil {
				panic(err)
			}
		}

	}
}
