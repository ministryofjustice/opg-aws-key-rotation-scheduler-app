package awsvault

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestAwsVault(t *testing.T) {

	v := &AwsVault{}
	fmt.Println(v.Supported(shell.New()))
}
