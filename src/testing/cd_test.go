package testing

import (
	"aslsh/commands"
	"os/user"
	"testing"
)

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