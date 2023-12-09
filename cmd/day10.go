package cmd

import (
	"fmt"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "Day 10",
	Long:  `Day 10`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day10.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/10")
		fmt.Println("Part1:", solveDay10(lines, false))
		fmt.Println("Part2:", solveDay10(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}

func solveDay10(lines []string, p2 bool) int {
	ans := 0
	return ans
}
