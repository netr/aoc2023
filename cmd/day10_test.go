package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay10(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day10_test.txt")
	ans := solveDay10(lines, false)
	if ans != 114 {
		t.Errorf("ans should be 114, got %d", ans)
	}

	// ans = solveDay10(lines, true)
	// if ans != 2 {
	// 	t.Errorf("ans should be 2, got %d", ans)
	// }
}
