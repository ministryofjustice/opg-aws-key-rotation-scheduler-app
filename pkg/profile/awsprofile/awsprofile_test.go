package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"

	"fyne.io/fyne/v2/app"
)

func TestAwsProfileSupported(t *testing.T) {
	app.NewWithID("test-app")
	pref.PREFERENCES = pref.New("test-app", "{}")

	p := &AwsProfile{}

	installed, profile, region := p.Supported(shell.New())

	fmt.Println(installed, profile, region)

}
