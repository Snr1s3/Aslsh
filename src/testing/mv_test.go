package testing

import (
    "aslsh/commands"
    "os"
    "testing"
)

func TestMvSuccess(t *testing.T) {
    src := "mv_test_src.txt"
    dst := "mv_test_dst.txt"
    os.WriteFile(src, []byte("test"), 0644)
    defer os.Remove(dst)

    msg := commands.Mv([]string{"mv", src, dst})
    if msg != "" {
        t.Errorf("Expected success, got error: %s", msg)
    }
    // Source should not exist, destination should exist
    if _, err := os.Stat(src); !os.IsNotExist(err) {
        t.Errorf("Source file still exists after move")
    }
    if _, err := os.Stat(dst); err != nil {
        t.Errorf("Destination file does not exist after move")
    }
    os.Remove(dst)
}

func TestMvMissingArgs(t *testing.T) {
    msg := commands.Mv([]string{"mv", "a"})
    want := "aslsh: mv: needs the old path and the new path"
    if msg != want {
        t.Errorf("Expected missing args error: %q, got: %q", want, msg)
    }
}

func TestMvTooManyArgs(t *testing.T) {
    msg := commands.Mv([]string{"mv", "a", "b", "c"})
    want := "aslsh: mv: too many arguments"
    if msg != want {
        t.Errorf("Expected too many args error: %q, got: %q", want, msg)
    }
}

func TestMvSourceNotExist(t *testing.T) {
    msg := commands.Mv([]string{"mv", "no_such_file.txt", "dst.txt"})
    want := "aslsh: mv: Error: path 'no_such_file.txt' does not exist."
    if msg != want {
        t.Errorf("Expected source not exist error: %q, got: %q", want, msg)
    }
}

func TestMvDestExists(t *testing.T) {
    src := "mv_test_src2.txt"
    dst := "mv_test_dst2.txt"
    os.WriteFile(src, []byte("test"), 0644)
    os.WriteFile(dst, []byte("test2"), 0644)
    defer os.Remove(src)
    defer os.Remove(dst)

    msg := commands.Mv([]string{"mv", src, dst})
    want := "aslsh: mv: destination 'mv_test_dst2.txt' already exists"
    if msg != want {
        t.Errorf("Expected destination exists error: %q, got: %q", want, msg)
    }
    os.Remove(src)
    os.Remove(dst)
}