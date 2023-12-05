package cmd

import (
	"fmt"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Day 5",
	Long:  `Day 5`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day5.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/5")
		fmt.Println("Part1:", solveSchematic(lines))
		fmt.Println("Part2:", solveSchematicGearRatio(lines))
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}

func solveDay5(lines []string) int {
	return 0
}
