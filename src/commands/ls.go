package commands

import (
	"os"
	"strings"
)

func Ls() (string, error) {
	var dirs []string
	var files []string

	entries, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, "d "+entry.Name())
		} else {
			files = append(files, "f "+entry.Name())
		}
	}

	result := strings.Join(append(dirs, files...), "\n")
	return result, nil
}
