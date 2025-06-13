package commands

import (
		"fmt"
		"os"
)

func Cd(path string){
		err := os.Chdir(path)
		if err != nil {
				fmt.Println("bash: cd: "+path+": No such file or directory")
				return
		}
}