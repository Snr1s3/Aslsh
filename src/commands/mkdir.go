package commands

import (
	"log"
	"os"
)

func Mkdir(parts []string) string {
	if len(parts) == 2 {
		err := os.Mkdir(parts[1], 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		return "Dir created"
	}
	return "Error"
}
