package commands

import (
	"fmt"
	"os"
)

func Touch(name string){
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
}