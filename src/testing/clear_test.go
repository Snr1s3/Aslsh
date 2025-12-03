package testing

import (
	"aslsh/commands"
	"strings"
	"testing"
)

// =======================
// Testing clear.go
// =======================
func TestClear(t *testing.T) {
	got := commands.Clear()
	want := "\033[H\033[2J"
	if got != want {
		t.Errorf("Clear() = %q, want %q", got, want)
	}
}

func TestClearEscapeSequence(t *testing.T) {
	got := commands.Clear()
	if !strings.Contains(got, "\033") {
		t.Errorf("Clear() does not contain escape sequence: %q", got)
	}
}