package commands
import (
	"fmt"
)
func Echo(parts []string){
		for i := 1; i< len(parts); i++ {
			fmt.Print(parts[i])
			if i < len(parts)-1{
				fmt.Print(" ")
			}
		}
		fmt.Println()
}