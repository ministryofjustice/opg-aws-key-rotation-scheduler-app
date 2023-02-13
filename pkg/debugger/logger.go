package debugger

import (
	"fmt"
	"io"
	"opg-aws-key-rotation-scheduler-app/pkg/storage"
	"os"
	"path/filepath"
	"time"
)

var _stdout io.Writer
var _stderr io.Writer

func init() {
	directory := storage.StorageDirectory()
	stdoutfile := filepath.Join(directory, "stdout.log")
	stderrfile := filepath.Join(directory, "stderr.log")

	std, _ := os.OpenFile(stdoutfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	e, _ := os.OpenFile(stderrfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	_stdout = io.MultiWriter(os.Stdout, std)
	_stderr = io.MultiWriter(os.Stderr, e)

	_stdout.Write([]byte("logger starting up\n"))
	_stderr.Write([]byte("logger starting up\n"))

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
			out.Write([]byte(str))
		}

	}
}
