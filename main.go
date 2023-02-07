package main

import "opg-aws-key-rotation-scheduler-app/pkg/opgapp"

const (
	SettingsFile string = "/settings.json"
)

func main() {
	opgapp.New(SettingsFile)

}
