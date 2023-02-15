package pref

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"time"
)

type PrefData[T string | bool | int | time.Duration] struct {
	Key       string
	Fallback  string
	ParseFunc func(res string) (T, error)

	value   interface{}
	sh      *shell.Shell
	pref    *map[string]string
	appName *string
}

// Get is a small wrapper around the application prefences provided
// by fyne an allowing for ennvironment variable overwrites
func (pd *PrefData[T]) Get() (value T) {
	// if we have looked this valuf up already, return it directly
	if pd.value != nil {
		defer debugger.Log("PrefData.Get()", debugger.INFO, "[cached]", "pref:"+pd.Key, "value", value)()
		return pd.value.(T)
	}
	var got string = pd.Fallback
	var p map[string]string = *pd.pref
	// get from preferences
	if val, ok := p[pd.Key]; ok {
		got = val
	}
	// overwite from env variables
	if envVal, ok := env(*pd.appName, pd.Key, pd.sh); ok {
		got = envVal
	}
	pf := pd.ParseFunc
	value, _ = pf(got)
	defer debugger.Log("PrefData.Get()", debugger.INFO, "pref:"+pd.Key, "value", value)()
	return
}

func newPD[T string | bool | int | time.Duration](
	appName *string,
	pref *map[string]string,
	sh *shell.Shell,
	key string,
	fallback string,
	parse func(res string) (T, error),
) PrefData[T] {
	return PrefData[T]{Key: key, Fallback: fallback, ParseFunc: parse, appName: appName, pref: pref, sh: sh}
}
