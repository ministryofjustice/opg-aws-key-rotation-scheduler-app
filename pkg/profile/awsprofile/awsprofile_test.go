package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestAwsProfileSupported(t *testing.T) {
	pref.PREFERENCES = pref.New("test-app", "{}")

	p := &AwsProfile{}

	installed, profile, region := p.Supported(shell.New())

	fmt.Println(installed, profile, region)

}
