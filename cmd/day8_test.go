package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay8(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day8_test.txt")
	ans := solveDay8(lines, false)
	if ans != 6440 {
		t.Errorf("ans should be 6440, got %d", ans)
	}
}
