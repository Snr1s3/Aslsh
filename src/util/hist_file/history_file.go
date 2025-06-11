package hist_file

import(
	"os"
	"log"
	"io"
	"bufio"
	"fmt"
)

func SaveCommand(command string){
		path := "./.history"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.WriteString(file, command+"\n")
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
}

func ReadHistory(){
		path := "./.history"
		file, err := os.Open(path)
		if err != nil {
				log.Fatalf("failed to open file: %s", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
				log.Fatalf("error reading file: %s", err)
		}
}