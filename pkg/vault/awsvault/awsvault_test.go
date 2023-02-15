package awsvault

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"

	"fyne.io/fyne/v2/app"
)

func TestAwsVault(t *testing.T) {
	app.NewWithID("test-app")
	pref.PREFERENCES = pref.New("test-app", "{}")

	v := &AwsVault{}
	fmt.Println(v.Supported(shell.New()))
}
