package commands

import(
	"os"
	"log"
	"bufio"
	"fmt"
)

func ReadHistory(){
	path := "./.history"
	file, err := os.Open(path)
	if err != nil {
			log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%d: %s\n", lineNumber, line)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file: %s", err)
	}
}