package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay8(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day8_test.txt")
	ans := solveDay8(lines, false)
	if ans != 6 {
		t.Errorf("ans should be 6, got %d", ans)
	}
}

func Test_SolveDay8_Part2(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day8_test2.txt")
	ans := solveDay8_Part2(lines, false)
	if ans != 6 {
		t.Errorf("ans should be 6, got %d", ans)
	}
}
