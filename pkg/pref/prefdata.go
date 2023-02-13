package pref

import (
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"time"

	"fyne.io/fyne/v2"
)

type PrefData[T string | bool | int | time.Duration] struct {
	Key       string
	Fallback  string
	ParseFunc func(res string) (T, error)

	value   interface{}
	pref    *fyne.Preferences
	appName *string
}

// Get is a small wrapper around the application prefences provided
// by fyne an allowing for ennvironment variable overwrites
func (pd *PrefData[T]) Get() (value T) {
	if pd.value != nil {
		return pd.value.(T)
	}
	pf := pd.ParseFunc
	p := *pd.pref
	got := p.StringWithFallback(pd.Key, pd.Fallback)
	// overwite from env variables
	if envVal, ok := env(*pd.appName, pd.Key); ok {
		got = envVal
	}
	value, _ = pf(got)
	defer debugger.Log("PrefData.Get()", debugger.INFO, "pref:"+pd.Key, "value", value)()
	return
}

func newPD[T string | bool | int | time.Duration](
	appName *string,
	pref *fyne.Preferences,
	key string,
	fallback string,
	parse func(res string) (T, error),
) PrefData[T] {
	return PrefData[T]{Key: key, Fallback: fallback, ParseFunc: parse, appName: appName, pref: pref}
}
