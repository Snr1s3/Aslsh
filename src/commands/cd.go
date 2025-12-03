package commands

import (
	"os"
	"os/user"
)

func Cd(parts []string) string {
	var path string
	if len(parts) == 1 {
		user, err := user.Current()
		if err != nil {
			return "Error getting current user: " + err.Error()
		}
		path = "/home/" + user.Username
	} else if len(parts) > 2 {
		return "aslsh: cd: too many arguments"
	} else {
		path = parts[1]
	}
	err := os.Chdir(path)
	if err != nil {
		return "aslsh: cd: " + path + ": No such file or directory"
	}
	return ""
}