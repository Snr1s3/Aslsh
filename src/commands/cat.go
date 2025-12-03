package commands

import (
	"fmt"
	"os"
)

func Cat(parts []string) (string, error) {
	if len(parts) == 1 {
		return "aslsh: cat: needs a path", nil
	} else if len(parts) > 2 {
		return "aslsh: cat: too many arguments", nil
	}
	path := parts[1]
	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "cat: " + path + ": No such file or directory", nil
	}
	if f.IsDir() {
		return "cat: " + path + ": Is not a file", nil
	}
	if os.IsPermission(err) {
		return "cat: " + path + ": Permission denied", nil
	}
	data, erro := os.ReadFile(path)
	if erro != nil {
		return "", fmt.Errorf("Error: %w", erro)
	}
	return string(data), nil
}