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

func solveDay8_2(lines []string, js bool) int {
	dirs := lines[0]
	items := lines[2:]
	paths := makePathMap(items)

	var nodes []string
	for p := range paths {
		if strings.HasSuffix(p, "A") {
			nodes = append(nodes, p)
		}
	}

	var nodeTimes []int
	dirIdx := 0
	times := 0
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
		nodeTimes = append(nodeTimes, times)
	}

	lcm := nodeTimes[0]
	nodeTimes = nodeTimes[1:]
	for _, t := range nodeTimes {
		lcm = util.LCM(lcm, t)
	}
	return lcm
}

func makePathMap(items []string) map[string][]string {
	newItems := make(map[string][]string, len(items))
	for _, item := range items {
		i := strings.Split(item, " = ")
		x := strings.TrimLeft(i[1], "(")
		x = strings.TrimRight(x, ")")
		is := strings.Split(x, ", ")
		newItems[i[0]] = is
	}
	return newItems
}
