package commands

import (
	"fmt"
	"regexp"
	"golang.org/x/exp/slices"
)

var aliasMap map[string][]string = make(map[string][]string)
var commandList = []string{"cd", "exit", "pwd", "echo", "help", "history", "clear", "source","alias","cat", "ls"}
var validAliasStart = regexp.MustCompile(`^[^-=|><&;$` + "`" + `"'\\/:#@]`)

func Alias(alias []string) string {
    if len(alias) < 2 {
        return "Invalid alias command. Use 'alias set' or 'alias unset'.\nUsage: alias set [name] [command] [arguments...]\nUsage: alias unset [name]"
    }
    if alias[1] == "set" && len(alias) > 3 {
        return SetAlias(alias)
    } else if alias[1] == "unset" && len(alias) == 3 {
        return UnsetAlias(alias[2])
    } else if alias[1] == "show" && len(alias) == 2 {
        return ShowAlias()
    } else {
        return "Invalid alias command. Use 'alias set' or 'alias unset'.\nUsage: alias set [name] [command] [arguments...]\nUsage: alias unset [name]"
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
func SetAlias(alias []string) string {
    if !slices.Contains(commandList, alias[3]) {
        return fmt.Sprintf("Command not found: %s", alias[3])
    }
    if _, exists := aliasMap[alias[2]]; exists {
        return "Alias name starts with a forbidden character1."
    }
    if !validAliasStart.MatchString(alias[2]) || slices.Contains(commandList, alias[2]) {
        return "Alias name starts with a forbidden character2."
    }
    aliasMap[alias[2]] = alias[3:]
    return fmt.Sprintf("Alias '%s' set to: %v", alias[2], alias[3:])
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

func ShowAlias() string {
    if len(aliasMap) == 0 {
        return "No aliases set."
    }
    result := "Current aliases:\n"
    for name, cmd := range aliasMap {
        result += fmt.Sprintf("%s: %v\n", name, cmd)
    }
    return result
}

func UnsetAlias(aliasName string) string {
    delete(aliasMap, aliasName)
    return fmt.Sprintf("Alias: '%s' unset successfully.", aliasName)
}