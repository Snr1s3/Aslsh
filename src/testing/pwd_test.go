package testing

import (
	"aslsh/commands"
	"os"
	"strings"
	"testing"
)

// =======================
// Testing pwd.go
// =======================
func TestPwd(t *testing.T) {
	got := commands.Pwd()
	if got == "" {
		t.Error("Pwd() returned an empty string")
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error: %v", err)
	}
	if got != wd {
		t.Errorf("Pwd() = %q, want %q", got, wd)
	}
}

func TestPwdNotEmpty(t *testing.T) {
	got := commands.Pwd()
	if len(strings.TrimSpace(got)) == 0 {
		t.Error("Pwd() returned only whitespace")
	}
}