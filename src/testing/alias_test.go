package testing

import (
	"aslsh/commands"
	"strings"
	"testing"
)

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

func TestAliasInvalidCommand(t *testing.T) {
	msg := commands.Alias([]string{"alias", "set", "badalias", "notacommand"})
	if !strings.Contains(msg, "Command not found") {
		t.Errorf("Alias() invalid command = %q, want 'Command not found' error", msg)
	}
}

func TestAliasInvalidUsage(t *testing.T) {
	msg := commands.Alias([]string{"alias"})
	if !strings.Contains(msg, "Invalid alias command") {
		t.Errorf("Alias() invalid usage = %q, want 'Invalid alias command' error", msg)
	}
}

func TestAliasShowEmpty(t *testing.T) {
	commands.Alias([]string{"alias", "unset", "ll"})
	msg := commands.Alias([]string{"alias", "show"})
	if !strings.Contains(msg, "No aliases set") {
		t.Errorf("Alias() show with no aliases = %q, want 'No aliases set'", msg)
	}
}

