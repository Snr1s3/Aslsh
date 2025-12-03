package testing

import (
	"aslsh/commands"
	"testing"
)

// =======================
// Testing echo.go
// =======================
func TestEcho(t *testing.T) {
	got := commands.Echo([]string{"echo", "hello", "world"})
	want := "hello world"
	if got != want {
		t.Errorf("Echo() = %q, want %q", got, want)
	}
}

func TestEchoEmpty(t *testing.T) {
	got := commands.Echo([]string{"echo"})
	if got != "" {
		t.Errorf("Echo() with no args = %q, want empty string", got)
	}
}