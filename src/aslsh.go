package main


import (
	"aslsh/util/user_input"
	"aslsh/commands/commands"
	"aslsh/util/hist_file"
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"log"
)

var cmd string = "$> "
func main() {
		initAslsh()
		exitB := false
		for !exitB {
			line := user_input.ReadInputStr(cmd)
			if line == "exit" {
				exitB = true
			} else {
				classifier(line)
			}
		}
}
func classifier(line string){
		re := regexp.MustCompile(`\s+`)
		line = re.ReplaceAllString(line, " ")
		parts := strings.Split(line, " ")
		switch parts[0]{
		case "echo":
			commands.Echo(parts)
			hist_file.SaveCommand(line)
		case "pwd":
			commands.Pwd(parts)
			hist_file.SaveCommand(line)
		case "clear":
			commands.Clear()
			hist_file.SaveCommand(line)
		case "history":
			hist_file.ReadHistory()
			hist_file.SaveCommand(line)
		case "help":
			commands.Help()
			hist_file.SaveCommand(line)
		case "alias":
			commands.SetAlias(parts)
			hist_file.SaveCommand(line)
		default:
			if len(line) != 0{
				fmt.Println("aslsh: "+parts[0]+": command not found")
			}
		}
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
}