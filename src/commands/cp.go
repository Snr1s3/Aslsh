package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Cp(parts []string) string {
	var oldPath, newPath string
	if len(parts) < 3 {
		return "aslsh: cp: needs the old path and the new path"
	} else if len(parts) > 3 {
		return "aslsh: cp: too many arguments"
	} else {
		oldPath = parts[1]
		newPath = parts[2]
	}
	if exists, errMsg := PathExists(oldPath); !exists {
		return "aslsh: cp: " + errMsg
	}
	if exists, _ := PathExists(newPath); exists {
		return fmt.Sprintf("aslsh: cp: destination '%s' already exists", newPath)
	}
	// Copy file
	src, err := os.Open(oldPath)
	if err != nil {
		return fmt.Sprintf("aslsh: cp: failed to open '%s': %v", oldPath, err)
	}
	defer src.Close()

	dstPath := newPath
	if fi, err := os.Stat(newPath); err == nil && fi.IsDir() {
		// If destination is a directory, append the filename
		dstPath = filepath.Join(newPath, filepath.Base(oldPath))
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Sprintf("aslsh: cp: failed to create '%s': %v", newPath, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Sprintf("aslsh: cp: failed to copy '%s' to '%s': %v", oldPath, newPath, err)
	}
	return ""
}
