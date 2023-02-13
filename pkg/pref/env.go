package pref

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"os"
	"strings"
)

func env(appName string, name string) (v string, ok bool) {
	ok = false
	app := strings.ReplaceAll(appName, " ", "")
	name = fmt.Sprintf("%s_%s", app, name)
	if envVar := os.Getenv(name); len(envVar) > 0 {
		v = envVar
		ok = true
	}
	defer debugger.Log("pref.env()", debugger.INFO, "key:"+name, "value", v)()
	return
}
