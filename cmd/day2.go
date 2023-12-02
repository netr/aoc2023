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
		fmt.Println("Part1:", playCubeGames(lines))
		fmt.Println("Part2:", playCubeGamesForMinimum(lines))
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

func tryCubeGame(game string) (int, bool) {
	id, played := splitCubeGame(game)

	for _, curGame := range played {
		cubes := strings.Split(curGame, ", ")
		for _, curCube := range cubes {
			color, count := parseCube(curCube)
			if !isCubePossible(color, count) {
				return 0, false
			}
		}
	}

	return id, true
}

// splitCubeGame splits a game string into a game id and a slice of games played
func splitCubeGame(game string) (int, []string) {
	gameSplit := strings.Split(game, ": ")
	gameId := strings.Replace(gameSplit[0], "Game ", "", -1)
	gameIdInt, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}
	gamesPlayed := strings.Split(gameSplit[1], "; ")
	return gameIdInt, gamesPlayed
}

func getCubeGameMinimums(game string) int {
	_, played := splitCubeGame(game)

	cubeMap := make(map[string]int, 3)
	cubeMap["red"] = 0
	cubeMap["green"] = 0
	cubeMap["blue"] = 0

	for _, curGame := range played {
		cubes := strings.Split(curGame, ", ")
		for _, curCube := range cubes {
			color, count := parseCube(curCube)
			if cubeMap[color] < count {
				cubeMap[color] = count
			}
		}
	}

	power := 1
	for _, v := range cubeMap {
		power *= v
	}
	return power
}

func parseCube(cube string) (string, int) {
	s := strings.Split(cube, " ")
	count, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	color := s[1]
	return color, count
}

func playCubeGames(games []string) int {
	total := 0
	for _, game := range games {
		if id, ok := tryCubeGame(game); ok {
			total += id
		}
	}
	return total
}

func playCubeGamesForMinimum(games []string) int {
	total := 0
	for _, game := range games {
		power := getCubeGameMinimums(game)
		total += power
	}
	return total
}
