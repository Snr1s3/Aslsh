package commands

import (
	"fmt"
	"log"
	"os"
)

func Cat(parts []string){
	var path string
		if len(parts) == 1{
			return
		} else if len(parts) > 2{
			fmt.Println("aslsh: cd: too many arguments")
			return
		} else{
			path = parts[1]
		}
	f, err := os.Stat(path);
	if os.IsNotExist(err){
		fmt.Println("cat: "+path+": No such a file or directory")
	}
	if f.IsDir(){
		fmt.Println("cat: "+path+": Is not a files")
		return
	}
	if os.IsPermission(err) {
		fmt.Println("cat: "+path+": Permission denied")
		return
	}
	data, erro :=os.ReadFile(path)
		if erro != nil {
			fmt.Print("Error: ")
			log.Fatal(erro)
			return
		}
		str := string(data)
		fmt.Println(str)
}