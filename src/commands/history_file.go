package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var path = "./.history"

func GetHistoryPath() string {
	return path
}
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
func CleanHistory() (string, error) {
	file, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to clean history: %w", err)
	}
	defer file.Close()
	return "history cleaned", nil
}

func History(parts []string) (string, error) {
	if len(parts) == 1 {
		return ReadHistory(path)
	}
	if len(parts) == 2 && parts[1] == "-c" {
		return CleanHistory()
	}
	return "aslsh: history: invalid arguments", nil
}
