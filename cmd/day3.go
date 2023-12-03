/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Day 3",
	Long:  `Day 3`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day3.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/3")
		fmt.Println("Part1:", solveSchematic(lines))
		fmt.Println("Part2:", solveSchematicGearRatio(lines))
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}

func schematicToMatrix(schematic []string) [][]rune {
	mat := make([][]rune, len(schematic))
	for i, line := range schematic {
		mat[i] = []rune(line)
	}
	return mat
}

func solveSchematic(schematics []string) int {
	mat := schematicToMatrix(schematics)
	symbolsMap := util.RuneSliceToMap('#', '$', '%', '&', '*', '+', '-', '/', '=', '@')
	dirs := util.AllDirs()

	count := 0
	for i, schematic := range schematics {
		for j, char := range schematic {

			partsSeen := map[int]struct{}{}
			if _, ok := symbolsMap[char]; ok {

				for _, dir := range dirs {
					row, col := i+dir[0], j+dir[1]
					if util.IsNumber(mat[row][col]) {
						partNumber := getPartNumberFrom(schematics[row], col)
						if _, ok := partsSeen[partNumber]; ok {
							continue
						}
						partsSeen[partNumber] = struct{}{}
						count += partNumber
					}
				}
			}
		}
	}

	return count
}

func solveSchematicGearRatio(schematics []string) int {
	mat := schematicToMatrix(schematics)
	symbolsMap := util.RuneSliceToMap('*')
	dirs := util.AllDirs()

	count := 0
	for i, schematic := range schematics {
		for j, char := range schematic {

			partsSeen := map[int]struct{}{}
			if _, ok := symbolsMap[char]; ok {
				for _, dir := range dirs {
					row, col := i+dir[0], j+dir[1]
					if util.IsNumber(mat[row][col]) {
						partNumber := getPartNumberFrom(schematics[row], col)
						if _, ok := partsSeen[partNumber]; ok {
							continue
						}
						partsSeen[partNumber] = struct{}{}
					}
				}
			}

			if len(partsSeen) == 2 {
				ttl := 1
				for partNumber := range partsSeen {
					ttl *= partNumber
				}
				count += ttl
			}
		}
	}

	return count
}

func getPartNumberFrom(schematic string, i int) int {
	var chars []byte

	if util.IsNumber(schematic[i]) {
		chars = []byte{schematic[i]}
	} else {
		return 0
	}

	k := i - 1
	j := i + 1
	for {
		if k >= 0 {
			if util.IsNumber(schematic[k]) {
				chars = append([]byte{schematic[k]}, chars...)
				k--
			} else {
				k = -1
			}
		}

		if j < len(schematic) {
			if util.IsNumber(schematic[j]) {
				chars = append(chars, schematic[j])
				j++
			} else {
				j = len(schematic)
			}
		}

		if k == -1 && j == len(schematic) {
			break
		}
	}

	ans, err := strconv.Atoi(string(chars))
	if err != nil {
		panic(err)
	}

	return ans
}
