package main

import (
	"embed"
	_ "embed"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
)

const (
	SettingsFile string = "/settings.json"
)

//go:embed settings.json
var settings []byte

//go:embed icons/*
var icons embed.FS

func main() {
	opgapp.New(settings, icons)

}
