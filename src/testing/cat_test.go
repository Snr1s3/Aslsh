package testing

import (
	"aslsh/commands"
	"strings"
	"testing"
)

// =======================
// Testing cat.go
// =======================
func TestCatIsDir(t *testing.T) {
	got, _ := commands.Cat([]string{"cat", "."})
	if !strings.Contains(got, "is not a file") {
		t.Errorf("Cat() on directory = %q, want 'is not a file' error", got)
	}
}

func TestCatTooManyArgs(t *testing.T) {
	got, _ := commands.Cat([]string{"cat", "file1", "file2"})
	if !strings.Contains(got, "too many arguments") {
		t.Errorf("Cat() with too many args = %q, want 'too many arguments' error", got)
	}
}

func TestCatNeedsPath(t *testing.T) {
	got, _ := commands.Cat([]string{"cat"})
	if !strings.Contains(got, "needs a path") {
		t.Errorf("Cat() with no path = %q, want 'needs a path' error", got)
	}
}
