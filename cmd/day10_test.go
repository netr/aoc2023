package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay10(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day10_test.txt")
	ans := solveDay10(lines, false)
	if ans != 8 {
		t.Errorf("ans should be 8, got %d", ans)
	}
}
