package commands
import (
	"fmt"
	"golang.org/x/exp/slices"
    "regexp"
)

var aliasMap map[string][]string = make(map[string][]string)
var commandList = []string{"cd", "exit", "pwd", "echo", "help", "history", "clear"}
var validAliasStart = regexp.MustCompile(`^[^-=|><&;$` + "`" + `"'\\/:#@]`)
func SetAlias(alias []string){
    if len(alias) < 3 {
        fmt.Println("Usage: alias [name] [command] [arguments...]")
        return
    }
    if !slices.Contains(commandList, alias[2]) {
        fmt.Println("Command not found:", alias[2])
        return
    }
    if !validAliasStart.MatchString(alias[1]) {
        fmt.Println("Alias name starts with a forbidden character.")
        return
    }
    aliasMap[alias[1]] = alias[2:]
    fmt.Printf("Alias '%s' set to: %v\n", alias[1], alias[2:])
}

func UnsetAlias(){
	rem
	fmt.Println("Alias unset successfully.")
}