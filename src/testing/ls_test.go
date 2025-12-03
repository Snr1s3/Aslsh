package testing

import (
	"aslsh/commands"
	"strings"
	"testing"
)
// =======================
// Testing ls.go
// =======================
func TestLs(t *testing.T) {
	got, err := commands.Ls()
	if err != nil {
		t.Errorf("Ls() error: %v", err)
	}
	if len(got) == 0 {
		t.Error("Ls() returned empty string")
	}
}

func TestLsNonEmpty(t *testing.T) {
	got, err := commands.Ls()
	if err != nil {
		t.Errorf("Ls() error: %v", err)
	}
	if len(strings.TrimSpace(got)) == 0 {
		t.Error("Ls() returned only whitespace")
	}
}