package commands

import (
	"fmt"
	"os"
	"os/user"
)

func Cd(parts []string){
		var path string
		if len(parts) == 1{
			user, err := user.Current()
			if err != nil {
				fmt.Println("Error getting current user:", err)
				return
			}
			path = "/home/"+user.Username
		} else if len(parts) > 2{
			fmt.Println("aslsh: cd: too many arguments")
			return
		} else{
			path = parts[1]
		}
		err := os.Chdir(path)
		if err != nil {
				fmt.Println("aslsh: cd: "+path+": No such file or directory")
				return
		}
		
}