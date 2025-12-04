package commands

import (
	"fmt"
	"os"
)

func Mv(parts []string) string {
	var oldPath, newPath string
	if len(parts) < 3 {
		return "aslsh: mv: needs the old path and the new path"
	} else if len(parts) > 3 {
		return "aslsh: mv: too many arguments"
	} else {
		oldPath = parts[1]
		newPath = parts[2]
	}
	if exists, errMsg := PathExists(oldPath); !exists {
		return "aslsh: mv: " + errMsg
	}
	if exists, _ := PathExists(newPath); exists {
		return fmt.Sprintf("aslsh: mv: destination '%s' already exists", newPath)
	}
	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Sprintf("aslsh: mv: failed to move '%s' to '%s': %v", oldPath, newPath, err)
	}
	return ""
}

func PathExists(path string) (bool, string) {
	_, err := os.Stat(path)
	if err == nil {
		return true, ""
	}
	if os.IsNotExist(err) {
		return false, fmt.Sprintf("Error: path '%s' does not exist.", path)
	}
	return false, fmt.Sprintf("Error checking path '%s': %v", path, err)
}