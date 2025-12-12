package testing

import (
	"aslsh/commands"
	"os"
	"strings"
	"testing"
)

// =======================
// Testing history_file.go
// =======================
func TestReadHistory(t *testing.T) {
	filename := "testhistory.txt"
	defer os.Remove(filename)
	lines := []string{"first", "second"}
	os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	got, err := commands.ReadHistory(filename)
	if err != nil {
		t.Errorf("ReadHistory() error: %v", err)
	}
	want := "1: first\n2: second"
	if got != want {
		t.Errorf("ReadHistory() = %q, want %q", got, want)
	}
}

func TestReadHistoryError(t *testing.T) {
	_, err := commands.ReadHistory("/nonexistentdir/history.txt")
	if err == nil {
		t.Error("ReadHistory() should return error for non-existent file")
	}
}

func TestCleanHistory(t *testing.T) {
	filename := commands.GetHistoryPath()
	os.WriteFile(filename, []byte("line1\nline2"), 0644)

	_, err := commands.CleanHistory()
	if err != nil {
		t.Fatalf("CleanHistory() error: %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("File not empty after CleanHistory(), got: %q", string(data))
	}
}
