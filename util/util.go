package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileIntoSlice(path string) []string {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines
}

func IsNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
