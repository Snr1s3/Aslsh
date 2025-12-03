package commands

import (
	"fmt"
	"os"
)

func Touch(name string) string {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Sprintf("Error creating file: %v", err)
	}
	defer file.Close()
	return ""
}