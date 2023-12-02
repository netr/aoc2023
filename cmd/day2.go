/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Day 2",
	Long:  `Day 2`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day2.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/2")
		fmt.Println("Part1:", playGames(lines))
		// fmt.Println("Part2:", totalCalibration(lines, calibrateWordNums))
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}

func isCubePossible(color string, count int) bool {
	switch color {
	case "blue":
		return count <= 14
	case "green":
		return count <= 13
	case "red":
		return count <= 12
	default:
		return false
	}
}

func playGame(game string) (int, bool) {
	gameSplit := strings.Split(game, ": ")
	gameId := strings.Replace(gameSplit[0], "Game ", "", -1)

	gamesPlayed := strings.Split(gameSplit[1], "; ")
	for _, curGame := range gamesPlayed {
		cubes := strings.Split(curGame, ", ")
		for _, curCube := range cubes {
			cubeSplit := strings.Split(curCube, " ")
			cubeCount, err := strconv.Atoi(cubeSplit[0])
			if err != nil {
				panic(err)
			}

			if !isCubePossible(cubeSplit[1], cubeCount) {
				return 0, false
			}
		}
	}

	gameIdInt, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	return gameIdInt, true
}

func playGames(games []string) int {
	total := 0
	for _, game := range games {
		if id, ok := playGame(game); ok {
			total += id
		}
	}
	return total
}
