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
		fmt.Println("Part2:", solveDay8_Part2(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}

func solveDay8(lines []string, js bool) int {
	dirs := lines[0]
	items := lines[2:]
	paths := makePathMap(items)

	times := 0
	curVal := "AAA"
	for {
		for _, c := range dirs {
			times++
			if c == 'L' {
				curVal = paths[curVal][0]
				if curVal == "ZZZ" {
					break
				}
			} else if c == 'R' {
				curVal = paths[curVal][1]
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

func solveDay8_Part2(lines []string, js bool) int {
	var (
		nodes         []string
		zTimes        []int
		dirIdx, times int
	)
	dirs := lines[0]
	items := lines[2:]
	paths := makePathMap(items)

	// find all the start nodes
	for p := range paths {
		if strings.HasSuffix(p, "A") {
			nodes = append(nodes, p)
		}
	}

	// find the number of times it takes to get to Z for each start node
	for _, node := range nodes {
		times = 0
		dirIdx = 0
		for times == 0 || !strings.HasSuffix(node, "Z") {
			times += 1
			if dirs[dirIdx] == 'L' {
				node = paths[node][0]
			} else {
				node = paths[node][1]
			}
			dirIdx = (dirIdx + 1) % len(dirs)
		}
		zTimes = append(zTimes, times)
	}

	// use first item as initial value for LCM
	ans := zTimes[0]
	zTimes = zTimes[1:]
	for _, t := range zTimes {
		ans = util.LCM(ans, t)
	}
	return ans
}

func makePathMap(items []string) map[string][]string {
	var (
		itemSplit        []string
		dirNodes         []string
		name, dirs, node string
	)

	itemMap := make(map[string][]string, len(items))
	for _, item := range items {
		itemSplit = strings.Split(item, " = ")
		name = itemSplit[0]
		dirs = itemSplit[1]
		node = strings.TrimLeft(dirs, "(")
		node = strings.TrimRight(node, ")")
		dirNodes = strings.Split(node, ", ")
		itemMap[name] = dirNodes
	}
	return itemMap
}
