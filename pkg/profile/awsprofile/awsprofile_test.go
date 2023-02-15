package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestAwsProfileSupported(t *testing.T) {
	s := shell.New()
	pref.PREFERENCES = pref.New("test-app", "{}", s)

	p := &AwsProfile{}

	installed, profile, region := p.Supported(s)

	fmt.Println(installed, profile, region)

}
