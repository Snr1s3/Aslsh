package testing

import (
	"aslsh/commands"
	"os"
	"os/user"
	"strings"
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

// =======================
// Testing alias.go
// =======================
func TestAliasSetShowUnset(t *testing.T) {
	setMsg := commands.Alias([]string{"alias", "set", "ll", "ls"})
	if !strings.Contains(setMsg, "Alias 'll' set to") {
		t.Errorf("Alias set failed: %s", setMsg)
	}
	showMsg := commands.Alias([]string{"alias", "show"})
	if !strings.Contains(showMsg, "ll: [ls]") {
		t.Errorf("Alias show failed: %s", showMsg)
	}
	unsetMsg := commands.Alias([]string{"alias", "unset", "ll"})
	if !strings.Contains(unsetMsg, "unset successfully") {
		t.Errorf("Alias unset failed: %s", unsetMsg)
	}
}

// =======================
// Testing help.go
// =======================
func TestHelp(t *testing.T) {
	expected := `cd       Change the current directory
exit     Exit the shell
pwd      Print the current working directory
echo     Print arguments to the terminal
export   Set environment variables
unset    Unset environment variables
alias    Define command shortcuts
help     Show help for built-in commands
history  Show command history
source   Reload .aslshrc
clear    Clear the terminal screen`
	if strings.TrimSpace(expected) != strings.TrimSpace(commands.Help()) {
		t.Errorf("The message should be:\n%s\nyou got\n%s", expected, commands.Help())
	}
}

// =======================
// Testing cd.go
// =======================
func TestCdToHome(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	home := "/home/" + usr.Username

	commands.Cd([]string{"cd"})
	got := commands.Pwd()
	if got != home {
		t.Errorf("Cd to home failed: got %s, want %s", got, home)
	}
}

func TestCdToCurrentAndBack(t *testing.T) {
	original := commands.Pwd()
	commands.Cd([]string{"cd", ".."})
	parent := commands.Pwd()
	if parent == original {
		t.Error("Cd .. did not change directory")
	}
	commands.Cd([]string{"cd", original})
	back := commands.Pwd()
	if back != original {
		t.Errorf("Failed to cd back to original: got %s, want %s", back, original)
	}
}

func TestCdNonExistent(t *testing.T) {
	original := commands.Pwd()
	path := "/this/does/not/exist"
	cd := commands.Cd([]string{"cd", path})
	after := commands.Pwd()
	if after != original {
		t.Errorf("Cd to non-existent dir changed directory: got %s, want %s", after, original)
	}
	if cd != "aslsh: cd: "+path+": No such file or directory" {
		t.Errorf("Error should be: \"aslsh: cd: "+path+": No such file or directory\" instead of \"%s\"", cd)
	}
}

func TestCdTooManyArgs(t *testing.T) {
	original := commands.Pwd()
	cd := commands.Cd([]string{"cd", "foo", "bar"})
	after := commands.Pwd()
	if after != original {
		t.Errorf("Cd with too many args changed directory: got %s, want %s", after, original)
	}
	if cd != "aslsh: cd: too many arguments" {
		t.Errorf("Error should be: aslsh: cd: too many arguments instead of %s", cd)
	}
}

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