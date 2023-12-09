package cmd

import (
	"fmt"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Day 9",
	Long:  `Day 9`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day9.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/9")
		fmt.Println("Part1:", solveDay9(lines, false))
		fmt.Println("Part2:", solveDay9(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}

func solveDay9(lines []string, p2 bool) int {
	ans := 0
	for _, line := range lines {
		nums := util.SplitToInts(line, " ")
		if !p2 {
			ans += extrapolateHistory(nums)
		} else {
			ans += extrapolateHistoryPart2(nums)
		}
	}
	return ans
}

func drawHistory(history [][]int) [][]int {
	for {
		history = append(history, []int{})
		idx := len(history) - 1
		zeros := 0
		for i := 0; i < len(history[idx-1])-1; i++ {
			diff := history[idx-1][i+1] - history[idx-1][i]
			history[idx] = append(history[idx], diff)
			if history[idx][len(history[idx])-1] == 0 {
				zeros++
			}
		}
		if zeros == len(history[idx]) {
			break
		}
	}
	return history
}

func extrapolateHistory(line []int) int {
	history := [][]int{}
	history = append(history, line)
	history = drawHistory(history)

	// iterate over poop from the bottom up
	history[len(history)-1] = append(history[len(history)-1], 0)
	for i := len(history) - 2; i >= 0; i-- {
		before := history[i+1][len(history[i+1])-1]
		curr := history[i][len(history[i])-1]
		history[i] = append(history[i], before+curr)
	}

	return history[0][len(history[0])-1]
}

func extrapolateHistoryPart2(line []int) int {
	history := [][]int{}
	history = append(history, line)
	history = drawHistory(history)

	// iterate over poop from the bottom up
	history[len(history)-1] = append([]int{0}, history[len(history)-1]...)
	for i := len(history) - 2; i >= 0; i-- {
		before := history[i+1][0]
		curr := history[i][0]
		history[i] = append([]int{curr - before}, history[i]...)
	}

	return history[0][0]
}
