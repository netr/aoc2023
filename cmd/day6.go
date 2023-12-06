/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day6Cmd represents the day5 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Day 6",
	Long:  `Day 6`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day6.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/6")

		fmt.Println("Part1:", solveDay6(lines))
		fmt.Println("Part2:", solveDay6(convertRacesToKerning(lines)))
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}

func parseRaces(lines []string) ([]int, []int) {
	time := strings.Split(lines[0], ":")
	dist := strings.Split(lines[1], ":")
	timeStr := strings.TrimSpace(time[1])
	distStr := strings.TrimSpace(dist[1])

	var timeInts []int
	var distInts []int

	for _, t := range strings.Split(timeStr, " ") {
		if t != "" {
			timeInts = append(timeInts, util.MustAtoi(t))
		}
	}

	for _, d := range strings.Split(distStr, " ") {
		if d != "" {
			distInts = append(distInts, util.MustAtoi(d))
		}
	}

	return timeInts, distInts
}

func convertRacesToKerning(lines []string) []string {
	times, dists := parseRaces(lines)
	timeStr, distStr := "", ""
	for _, t := range times {
		timeStr += fmt.Sprintf("%d", t)
	}
	for _, d := range dists {
		distStr += fmt.Sprintf("%d", d)
	}

	races := []string{
		fmt.Sprintf("Time: %s", timeStr),
		fmt.Sprintf("Distance: %s", distStr),
	}
	return races
}

func solveDay6(races []string) int {
	times, dists := parseRaces(races)

	waysToWin := []int{}
	for raceIdx, raceTime := range times {
		ways := 0
		for i := 1; i <= raceTime; i++ {
			ttl := (raceTime - i) * i
			if dists[raceIdx] < ttl {
				ways++
			}
		}
		waysToWin = append(waysToWin, ways)
		fmt.Println("ways", ways)
	}

	ans := 1
	for i := 0; i < len(waysToWin); i++ {
		ans *= waysToWin[i]
	}

	fmt.Println("ans", ans)
	return ans
}

func solveDay6Kerning(races []string) int {
	times, dists := parseRaces(races)
	raceTime := times[0]
	raceDist := dists[0]

	ways := 0
	for i := 1; i <= raceTime; i++ {
		ttl := (raceTime - i) * i
		if raceDist < ttl {
			ways++
		}
	}

	return ways
}
