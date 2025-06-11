package commands
import (
	"fmt"
)
var cd =      "cd	  Change the current directory\n"
var exit =    "exit	  Exit the shell\n"
var pwd =     "pwd	  Print the current working directory\n"
var echo =    "echo	  Print arguments to the terminal\n"
var export =  "export	  Set environment variables\n"
var unset =   "unset	  Unset environment variables\n"
var alias =   "alias     Define command shortcuts\n"
var help =    "help      Show help for built-in commands\n"
var history = "history	  Show command history\n"
var clear =   "clear	  Clear the terminal screen"
func Help(){
	fmt.Println(cd+exit+pwd+echo+export+unset+alias+help+history+clear)
}