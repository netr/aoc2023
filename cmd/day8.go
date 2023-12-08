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

	steps := 0
	node := "AAA"
	for {
		for _, c := range dirs {
			steps++
			if c == 'L' {
				node = paths[node][0]
				if node == "ZZZ" {
					return steps
				}
			} else if c == 'R' {
				node = paths[node][1]
				if node == "ZZZ" {
					return steps
				}
			}
		}
	}
}

func solveDay8_Part2(lines []string, js bool) int {
	var (
		nodes         []string
		zSteps        []int
		dirIdx, steps int
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

	// find the number of steps it takes to get to Z for each start node
	for _, node := range nodes {
		steps = 0
		dirIdx = 0
		for steps == 0 || !strings.HasSuffix(node, "Z") {
			steps += 1
			if dirs[dirIdx] == 'L' {
				node = paths[node][0]
			} else {
				node = paths[node][1]
			}
			dirIdx = (dirIdx + 1) % len(dirs)
		}
		zSteps = append(zSteps, steps)
	}

	// use first item as initial value for LCM
	ans := zSteps[0]
	zSteps = zSteps[1:]
	for _, t := range zSteps {
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
		// split the item into the name and the two nodes
		itemSplit = strings.SplitN(item, " = ", 2)
		name, dirs = itemSplit[0], itemSplit[1]
		// strip off the outer parens
		node = dirs[1 : len(dirs)-1]
		// split the node into the two nodes. space after comma is important
		dirNodes = strings.Split(node, ", ")
		itemMap[name] = dirNodes
	}
	return itemMap
}
