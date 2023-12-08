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
	paths := lines[2:]
	pathMap := makePathMap(paths)

	// find all the start nodes
	for p := range pathMap {
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
				node = pathMap[node][0]
			} else {
				node = pathMap[node][1]
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

func makePathMap(paths []string) map[string][]string {
	// example path = "JXD = (NFK, KMD)"
	var (
		pathSplit  []string
		dirNodes   []string
		node, dirs string
	)
	pathMap := make(map[string][]string, len(paths))
	for _, p := range paths {
		// split the path into the name and the two nodes
		pathSplit = strings.SplitN(p, " = ", 2)
		node, dirs = pathSplit[0], pathSplit[1]
		// strip off the outer parens
		dirs = dirs[1 : len(dirs)-1]
		// split the node into the two nodes. space after comma is important
		dirNodes = strings.Split(dirs, ", ")
		pathMap[node] = dirNodes
	}
	return pathMap
}
