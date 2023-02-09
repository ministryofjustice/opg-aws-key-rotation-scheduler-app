package debugger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var _stdout io.Writer
var _stderr io.Writer

func SetupFileLogging(directory string) {

	stdoutfile := filepath.Join(directory, "stdout.log")
	stderrfile := filepath.Join(directory, "stderr.log")

	std, _ := os.OpenFile(stdoutfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	e, _ := os.OpenFile(stderrfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	_stdout = io.MultiWriter(os.Stdout, std)
	_stderr = io.MultiWriter(os.Stderr, e)

	//panic(stdoutfile)
}

func Log(message string, level int, values ...interface{}) func() {
	out := _stdout
	if LEVEL == ERR {
		out = _stderr
	}
	t := time.Now().UTC()
	return func() {
		str := fmt.Sprintf("[%s] (%s)\n", t, message)
		for _, v := range values {
			str += fmt.Sprintf("%v ", v)
		}
		str += "\n--\n"
		show := (level <= LEVEL)
		if show {
			out.Write([]byte(str))
		}

	}
}
