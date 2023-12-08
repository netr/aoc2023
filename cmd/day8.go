package cmd

import (
	"fmt"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Day 8",
	Long:  `Day 8`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day8.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/8")
		fmt.Println("Part1:", solveDay8(lines, false))
		fmt.Println("Part2:", solveDay8(lines, true))
		// fmt.Println("Part2:", solveDay6(convertRacesToKerning(lines)))
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}

func solveDay8(lines []string, js bool) int {
	return 0
}
