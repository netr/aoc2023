package cmd

import (
	"container/list"
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
		// fmt.Println("Part2:", solveDay10(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}

func solveDay10(lines []string, p2 bool) int {
	grid := buildGrid(lines)
	sr, sc := getS(grid)

	visited := map[[2]int]bool{}
	visited[[2]int{sr, sc}] = true
	queue := list.New()
	queue.PushBack([2]int{sr, sc})

	is_valid_up := func(ch rune) bool {
		return ch == 'S' || ch == '|' || ch == 'J' || ch == 'L'
	}
	is_valid_down := func(ch rune) bool {
		return ch == 'S' || ch == '|' || ch == '7' || ch == 'F'
	}
	is_valid_left := func(ch rune) bool {
		return ch == 'S' || ch == '-' || ch == 'J' || ch == '7'
	}
	is_valid_right := func(ch rune) bool {
		return ch == 'S' || ch == '-' || ch == 'L' || ch == 'F'
	}

	for queue.Len() > 0 {
		qi, ok := queue.Remove(queue.Front()).([2]int)
		if !ok {
			panic("not ok")
		}
		row, col := qi[0], qi[1]
		cur := grid[row][col]

		if row > 0 && is_valid_up(cur) && is_valid_down(grid[row-1][col]) && !visited[[2]int{row - 1, col}] {
			visited[[2]int{row - 1, col}] = true
			queue.PushBack([2]int{row - 1, col})
		}

		if row < len(grid)-1 && is_valid_down(cur) && is_valid_up(grid[row+1][col]) && !visited[[2]int{row + 1, col}] {
			visited[[2]int{row + 1, col}] = true
			queue.PushBack([2]int{row + 1, col})
		}

		if col > 0 && is_valid_left(cur) && is_valid_right(grid[row][col-1]) && !visited[[2]int{row, col - 1}] {
			visited[[2]int{row, col - 1}] = true
			queue.PushBack([2]int{row, col - 1})
		}

		if col < len(grid[row])-1 && is_valid_right(cur) && is_valid_left(grid[row][col+1]) && !visited[[2]int{row, col + 1}] {
			visited[[2]int{row, col + 1}] = true
			queue.PushBack([2]int{row, col + 1})
		}
	}

	return len(visited) / 2
}

func getS(grid [][]rune) (int, int) {
	for r, row := range grid {
		for c := range row {
			if grid[r][c] == 'S' {
				return r, c
			}
		}
	}
	return -1, -1
}

func buildGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
