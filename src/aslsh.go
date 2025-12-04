package main

import (
	"aslsh/commands"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/chzyer/readline"
)
var path = "/home/asegura/.ASLSH_history"
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
        readline.PcItem("cat"),
        readline.PcItem("touch"),
        readline.PcItem("mv"),
		)
	rl, err := readline.NewEx(&readline.Config{
        Prompt:       cmd,
		HistoryFile:  path,
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
        line, erro := rl.Readline()
        if erro != nil || err != nil {
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
func classifier(parts []string) bool {
    parts = commands.GetAlias(parts)
    var output string
    var err error

    switch parts[0] {
    case "echo":
        output = commands.Echo(parts)
    case "pwd":
        output = commands.Pwd()
    case "clear":
        output = commands.Clear()
    case "history":
        output, err = commands.ReadHistory(path)
    case "help":
        output = commands.Help()
    case "alias":
        output = commands.Alias(parts)
    case "exit":
        return true
    case "source":
        initAslsh()
        return false
    case "cd":
        output = commands.Cd(parts)
    case "ls":
        output, err = commands.Ls()
    case "cat":
        output, err = commands.Cat(parts)
    case "touch":
        output = commands.Touch(parts[1])
    case "mv":
        output = commands.Mv(parts)
    default:
        if len(parts[0]) != 0 {
            output = "aslsh: " + parts[0] + ": command not found"
        }
    }

    if err != nil {
        fmt.Println("Error:", err)
    } else if len(output) > 0 {
        fmt.Println(output)
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