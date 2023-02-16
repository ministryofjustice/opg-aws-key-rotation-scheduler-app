package pref

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"strings"
)

var envCache map[string]string = map[string]string{}

func env(appName string, name string, sh *shell.Shell) (v string, ok bool) {
	// generate env cache
	if _, set := envCache["_boot"]; !set {
		debugger.Log("pref.env()", debugger.INFO, "creating env cache")()
		_sh := *sh
		for k, v := range _sh.Env() {
			if strings.Contains(k, appName) {
				envCache[k] = v
			}
		}
		envCache["_boot"] = "set"
		debugger.Log("pref.env()", debugger.INFO, "env cache:", envCache)()
	}

	nm := appName + "_" + name
	v, ok = envCache[nm]
	defer debugger.Log("pref.env()", debugger.INFO, "name:\t"+name, "env:\t"+nm, "value:\t"+v)()
	return
}
