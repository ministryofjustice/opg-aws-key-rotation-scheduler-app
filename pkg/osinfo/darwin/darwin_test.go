package darwin

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"testing"
)

func TestDarkMode(t *testing.T) {
	os := &Darwin{}
	dark := os.DarkMode(shell.New())

	fmt.Printf("dark: %v\n", dark)

}
