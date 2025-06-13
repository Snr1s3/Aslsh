package main


import (
	"aslsh/commands/commands"
	"github.com/chzyer/readline"
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"log"
)

var cmd string = "$> "
var exitB bool = false
func main() {
	initAslsh()
	completer := readline.NewPrefixCompleter(
        readline.PcItem("echo"),
        readline.PcItem("pwd"),
        readline.PcItem("clear"),
        readline.PcItem("history"),
        readline.PcItem("help"),
        readline.PcItem("alias"),
        readline.PcItem("source"),
        readline.PcItem("cd"),
        readline.PcItem("ls"),
		)
	rl, err := readline.NewEx(&readline.Config{
        Prompt:       cmd,
		HistoryFile: "./.history",
		HistoryLimit:  1000,
        AutoComplete: completer,
    })
    if err != nil {
        log.Fatalf("readline error: %v", err)
    }
    defer rl.Close()
    for !exitB {
				rl, err = readline.NewEx(&readline.Config{
						Prompt:       cmd,
				HistoryFile: "./.history",
				HistoryLimit:  1000,
						AutoComplete: completer,
				})
        line, err := rl.Readline()
        if err != nil {
            break
        }
        if line == "exit" {
            exitB = true
        } else {
			re := regexp.MustCompile(`\s+`)
			line = re.ReplaceAllString(line, " ")
			parts := strings.Split(line, " ")
            exitB = classifier(parts)
        }
    }
}
func classifier(parts []string) bool{
		parts = commands.GetAlias(parts)
		switch parts[0]{
		case "echo":
			commands.Echo(parts)
		case "pwd":
			commands.Pwd()
		case "clear":
			commands.Clear()
		case "history":
			commands.ReadHistory()
		case "help":
			commands.Help()
		case "alias":
			commands.Alias(parts)
		case "exit":
			return true
		case "source":
			initAslsh()
		case "cd":
			commands.Cd(parts[1])
		case "ls":
			commands.Ls()
		default:
			if len(parts[0]) != 0{
				fmt.Println("aslsh: "+parts[0]+": command not found")
			}
		}
		return false
}

func initAslsh(){
		path := "./.aslshrc"
		file, err := os.Open(path)
		if err != nil {
				log.Fatalf("failed to open file: %s", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
				line := scanner.Text()
				parseRc(line)

		}
		if err := scanner.Err(); err != nil {
				log.Fatalf("error reading file: %s", err)
		}
}
func parseRc(line string){
		parts := strings.Split(line, "=")
		if parts[0] == "ps"{
			cmd = parts[1]
		}
		if parts[0] == "alias"{
			parts := strings.Split(parts[1], " ")
			commands.AliasRc(parts)
		}
}