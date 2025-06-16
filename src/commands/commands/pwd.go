package commands

import (
	"fmt"
	"log"
	"os"
)
func Pwd(){
		dir, err:= os.Getwd()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(dir)
}