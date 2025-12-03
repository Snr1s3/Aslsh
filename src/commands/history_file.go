package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadHistory(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, fmt.Sprintf("%d: %s", lineNumber, line))
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return strings.Join(lines, "\n"), nil
}