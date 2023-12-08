package cmd

import (
	"fmt"
	"strings"

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
		fmt.Println("Part2:", solveDay8_2(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}

func solveDay8(lines []string, js bool) int {
	dirs := lines[0]
	items := lines[2:]

	newItems := make(map[string][]string, len(items))
	newItemsIdx := make(map[string]int, len(items))
	for idx, item := range items {
		i := strings.Split(item, " = ")
		x := strings.TrimLeft(i[1], "(")
		x = strings.TrimRight(x, ")")
		is := strings.Split(x, ", ")
		newItems[i[0]] = is
		newItemsIdx[i[0]] = idx
	}

	times := 0
	curVal := "AAA"
	for {
		for _, c := range dirs {
			times++
			if c == 'L' {
				curVal = newItems[curVal][0]
				if curVal == "ZZZ" {
					break
				}
			} else if c == 'R' {
				curVal = newItems[curVal][1]
				if curVal == "ZZZ" {
					break
				}
			}
		}
		if curVal == "ZZZ" {
			break
		}
	}

	return times
}

func solveDay8_2(lines []string, js bool) int {
	dirs := lines[0]
	items := lines[2:]

	newItems := make(map[string][]string, len(items))
	newItemsIdx := make(map[string]int, len(items))
	for idx, item := range items {
		i := strings.Split(item, " = ")
		x := strings.TrimLeft(i[1], "(")
		x = strings.TrimRight(x, ")")
		is := strings.Split(x, ", ")
		newItems[i[0]] = is
		newItemsIdx[i[0]] = idx
	}

	fmt.Println(dirs)

	// times := 0
	var starts []string
	for k := range newItems {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	var startTimes [][]int

	for _, start := range starts {
		times := 0
		startDirs := dirs
		firstMatch := ""
		timesRes := []int{}

		for {
			for times == 0 || !strings.HasSuffix(start, "Z") {
				times += 1
				if startDirs[0] == 'L' {
					start = newItems[start][0]
				} else {
					start = newItems[start][1]
				}
				startDirs = startDirs[1:] + startDirs[0:1]
			}
			timesRes = append(timesRes, times)

			if firstMatch == "" {
				firstMatch = start
				times = 0
			} else if start == firstMatch {
				break
			}
		}

		startTimes = append(startTimes, timesRes)
	}

	var nums []int
	for _, t := range startTimes {
		nums = append(nums, t[0])
	}
	lcm := nums[0]
	nums = nums[1:]
	for _, num := range nums {
		lcm = lcm * num / gcd(lcm, num)
	}
	return lcm
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
