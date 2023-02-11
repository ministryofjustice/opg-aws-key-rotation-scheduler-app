package zsh

import (
	"strings"
	"testing"
)

func TestZshSelf(t *testing.T) {
	sh := Zsh{}
	path, err := sh.Self()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(path, "zsh") {
		t.Errorf("zsh missing: %v", path)
	}
}

func TestZshSearch(t *testing.T) {
	sh := Zsh{}
	// -- look for less, should always exist!
	path, stdout, stderr, err := sh.Search("less", false)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(path, "less") {
		t.Errorf("less missing: %v", path)
	}
	strerr := stderr.String()
	if len(strerr) > 0 {
		t.Errorf("unexpected error: %v", strerr)
	}

	str := stdout.String()
	if len(str) == 0 {
		t.Errorf("unexpected error: %v", str)
	}

	// -- something thats installed via brew etc and only in user profile
	path, stdout, stderr, err = sh.Search("jq", true)
	if err != nil {
		t.Errorf("unexpected an error: %v", err)
	}
	if !strings.Contains(path, "/homebrew") {
		t.Errorf("failed to find command: %v", path)
	}
	strerr = stderr.String()
	if len(strerr) > 0 {
		t.Errorf("unexpected error: %v", strerr)
	}

	str = stdout.String()
	if len(str) == 0 {
		t.Errorf("expected command path: %v", str)
	}
}
