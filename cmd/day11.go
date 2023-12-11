/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day11 called")
		solveday11()
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}

func solveday11() {
	space := []string{}
	times := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		// Find empty rows
		if !strings.Contains(line, "#") {
			// Add expansion to rows
			for i := 0; i < times; i++ {
				space = append(space, strings.Repeat(".", len(line)))
			}
		}
		space = append(space, line)
	}

	// Find empty columns
	cols := []int{}
	for c := 0; c < len(space[0]); c++ {
		count := 0
		for r := 0; r < len(space); r++ {
			if space[r][c] == '#' {
				break
			}
			count++
		}
		if count == len(space) {
			cols = append(cols, c+(len(cols)*times))
		}
	}

	// Add expansion to columns
	repeat := strings.Repeat(".", times)
	for _, c := range cols {
		for r := 0; r < len(space); r++ {
			space[r] = space[r][:c+1] + repeat + space[r][c+1:]
		}
	}

	galaxy := 1
	galaxies := []struct{ r, c int }{}
	for r, row := range space {
		if strings.Contains(row, "#") {
			for c, col := range row {
				if col == '#' {
					space[r] = strings.Replace(space[r], "#", fmt.Sprintf("%d", galaxy), 1)
					galaxies = append(galaxies, struct{ r, c int }{r, c})
					galaxy++
				}
			}
		}
	}

	count := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			gi0 := galaxies[i].r
			gj0 := galaxies[j].r
			gi1 := galaxies[i].c
			gj1 := galaxies[j].c
			count += abs(gi0-gj0) + abs(gi1-gj1)
		}
	}
	fmt.Println(count)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
