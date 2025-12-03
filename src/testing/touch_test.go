package testing

import (
	"aslsh/commands"
	"os"
	"strings"
	"testing"
)

// =======================
// Testing touch.go and cat.go
// =======================
func TestTouchAndCat(t *testing.T) {
	filename := "testfile.txt"
	defer os.Remove(filename)
	msg := commands.Touch(filename)
	if msg != "" {
		t.Errorf("Touch() error: %s", msg)
	}
	content := "hello aslsh"
	os.WriteFile(filename, []byte(content), 0644)
	got, err := commands.Cat([]string{"cat", filename})
	if err != nil {
		t.Errorf("Cat() error: %v", err)
	}
	if got != content {
		t.Errorf("Cat() = %q, want %q", got, content)
	}
	got, err = commands.Cat([]string{"cat", "nofile.txt"})
	if got != "cat: nofile.txt: No such file or directory" {
		t.Errorf("Cat() non-existent = %q, want error message", got)
	}
}

// =======================
// Testing touch.go
// =======================
func TestTouchError(t *testing.T) {
	msg := commands.Touch("/nonexistentdir/testfile.txt")
	if !strings.Contains(msg, "Error creating file") {
		t.Errorf("Touch() error message not returned for bad path: %s", msg)
	}
}


