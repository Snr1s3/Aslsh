package commands

import (
	"os/user"
)

func Whoami() string {
	u, err := user.Current()
	if err != nil {
		return "unknown user"
	}
	return u.Username
}
