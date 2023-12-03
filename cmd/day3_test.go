package cmd

import (
	"testing"
)

func Test_SolveSchematic(t *testing.T) {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	if solveSchematic(schematic) != 4361 {
		t.Errorf("should be 4361")
	}
}

func Test_SolveSchematicGearRatio(t *testing.T) {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	if solveSchematicGearRatio(schematic) != 467835 {
		t.Errorf("should be 467835")
	}
}

func Test_GetPartNumberFrom(t *testing.T) {
	number := getPartNumberFrom("..35..633.", 3)

	if number != 35 {
		t.Errorf("number should be 35, got %d", number)
	}
}
