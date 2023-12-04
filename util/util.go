package util

import (
	"bufio"
	"os"
	"strconv"
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

func IsNumber[T byte | rune](char T) bool {
	return char >= '0' && char <= '9'
}

func AllDirs() [][]int {
	dirs := [][]int{}
	dirs = append(dirs, Dirs()...)
	dirs = append(dirs, Diags()...)
	return dirs
}

func Dirs() [][]int {
	return [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
}

func Diags() [][]int {
	return [][]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}
}

func RuneSliceToMap(symbols ...rune) map[rune]struct{} {
	symbolsMap := map[rune]struct{}{}
	for _, symbol := range symbols {
		symbolsMap[symbol] = struct{}{}
	}
	return symbolsMap
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
