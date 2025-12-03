package commands

import (
	"fmt"
	"log"
	"os"
)

func Ls(){
		var dirs []string
		var files []string

		entries, err := os.ReadDir(".")
		if err != nil {
				log.Fatal(err)
		}

		for _, entry := range entries {
				if entry.IsDir() {
						dirs = append(dirs, entry.Name())
				} else {
						files = append(files, entry.Name())
				}
		}


		for _, dir := range dirs{
				fmt.Println("d "+dir)
		}

		for _,file := range files{
				fmt.Println("f "+file)
		}

}
