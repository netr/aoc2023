/*
Copyright © 2023 Corey Jackson <programmatical@gmail.com>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Day 1",
	Long:  `Day 1`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day1.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/1")
		fmt.Println("Part1:", totalCalibration(lines, calibrate))
		fmt.Println("Part2:", totalCalibration(lines, calibrateWordNums))
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}

func calibrate(input string) int {
	ans := ""
	for i := 0; i < len(input); i++ {
		if util.IsNumber(input[i]) {
			ans += string(input[i])
		}
	}

	if len(ans) > 2 {
		ans = fmt.Sprintf("%s%s", string(ans[0]), string(ans[len(ans)-1:]))
	} else if len(ans) == 1 {
		ans = ans + ans
	}

	fin, err := strconv.Atoi(ans)
	if err != nil {
		panic(err)
	}

	return fin
}

func calibrateWordNums(input string) int {
	numMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	tmpInput := ""
	for i := 0; i < len(input); i++ {
		if util.IsNumber(input[i]) {
			tmpInput += string(input[i])
			continue
		}

		found := false
		for k, v := range numMap {
			if i+len(k) <= len(input) && input[i:i+len(k)] == k {
				tmpInput += strconv.Itoa(v)
				i += len(k) - 1
				found = true
				break
			}
		}

		if !found {
			tmpInput += string(input[i])
		}
	}

	return calibrate(tmpInput)
}

func totalCalibration(input []string, cb func(string) int) int {
	ttl := 0
	for _, line := range input {
		ttl += cb(line)
	}
	return ttl
}
