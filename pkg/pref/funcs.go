package pref

import (
	"strconv"
	"time"
)

func strToStr(str string) (string, error) {
	return str, nil
}
func strToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
func strToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
func strToDuration(str string) (time.Duration, error) {
	return time.ParseDuration(str)
}
