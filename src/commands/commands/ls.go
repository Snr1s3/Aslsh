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
		
		for _, entry := range entries {
				if entry.IsDir() {
						fmt.Println("DIR:", entry.Name())
				} else {
						fmt.Println("FILE:", entry.Name())
				}
		}


		for dir := range dirs{
				fmt.Println(dir)
		}

		for file := range files{
				fmt.Println(file)
		}

}
