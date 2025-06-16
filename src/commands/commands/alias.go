package commands

import (
	"fmt"
	"regexp"

	"golang.org/x/exp/slices"
)

var aliasMap map[string][]string = make(map[string][]string)
var commandList = []string{"cd", "exit", "pwd", "echo", "help", "history", "clear", "source","alias","cat"}
var validAliasStart = regexp.MustCompile(`^[^-=|><&;$` + "`" + `"'\\/:#@]`)
func Alias(alias []string) {
    if len(alias) < 2 {
        fmt.Println("Invalid alias command. Use 'alias set' or 'alias unset'.")
        fmt.Println("Usage: alias set [name] [command] [arguments...]")
        fmt.Println("Usage: alias unset [name]")
        return
    }
    if alias[1] == "set" && len(alias) > 3 {
        SetAlias(alias)
    } else if alias[1] == "unset" && len(alias) == 3 {
        UnsetAlias(alias[2])
    } else if alias[1] == "show" && len(alias) == 2 {
        ShowAlias()
    }else {
        fmt.Println("Invalid alias command. Use 'alias set' or 'alias unset'.")
        fmt.Println("Usage: alias set [name] [command] [arguments...]")
        fmt.Println("Usage: alias unset [name]")
    }
}

func GetAlias(parts []string) []string {
    if len(parts) < 1 {
        return parts
    }
    if alias, exists := aliasMap[parts[0]]; exists {
        return append(alias, parts[1:]...)
    }
    return parts
}
func SetAlias(alias []string){
    
    if !slices.Contains(commandList, alias[3]) {
        fmt.Println("Command not found:", alias[3])
        return
    }
		
		if _, exists := aliasMap[alias[2]]; exists {
				fmt.Println("Alias name starts with a forbidden character1.")
				return
		}
    if !validAliasStart.MatchString(alias[2]) || slices.Contains(commandList, alias[2]){
        fmt.Println("Alias name starts with a forbidden character2.")
        return
    }
    aliasMap[alias[2]] = alias[3:]
    fmt.Printf("Alias '%s' set to: %v\n", alias[2], alias[3:])
}

func AliasRc(alias []string){
    if !slices.Contains(commandList, alias[1]) {
        fmt.Println("Command not found:", alias[1])
				fmt.Println(alias[1])
        return
    }
		if _, exists := aliasMap[alias[0]]; exists {
				return
		}
    if !validAliasStart.MatchString(alias[0]) || slices.Contains(commandList, alias[0]){
        fmt.Println("Alias name starts with a forbidden character.")
				fmt.Println(alias)
        return
    }
    aliasMap[alias[0]] = alias[1:]
}

func ShowAlias() {
    if len(aliasMap) == 0 {
        fmt.Println("No aliases set.")
        return
    }
    fmt.Println("Current aliases:")
    for name, cmd := range aliasMap {
        fmt.Printf("%s: %v\n", name, cmd)
    }
}
func UnsetAlias(aliasName string) {
    delete(aliasMap, aliasName)
	fmt.Printf("Alias: '%s' unset successfully.\n",aliasName)
}