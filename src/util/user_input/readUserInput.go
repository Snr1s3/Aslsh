package user_input


import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func ReadInputStr(cmd string) string {
	fmt.Print(cmd)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
