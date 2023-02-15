package awsvault

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestAwsVault(t *testing.T) {
	pref.PREFERENCES = pref.New("test-app", "{}")

	v := &AwsVault{}
	fmt.Println(v.Supported(shell.New()))
}
