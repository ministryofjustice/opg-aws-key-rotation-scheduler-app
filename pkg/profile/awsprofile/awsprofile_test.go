package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestAwsProfileSupported(t *testing.T) {

	p := &AwsProfile{}

	installed, profile, region := p.Supported(shell.New())

	fmt.Println(installed, profile, region)

}
