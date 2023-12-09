package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay9(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day9_test.txt")
	ans := solveDay9(lines, false)
	if ans != 114 {
		t.Errorf("ans should be 114, got %d", ans)
	}

	ans = solveDay9(lines, true)
	if ans != 2 {
		t.Errorf("ans should be 2, got %d", ans)
	}
}

func Test_extrapolateHistory(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day9_test.txt")
	line := util.SplitToInts(lines[0], " ")
	ans := extrapolateHistory(line)
	if ans != 18 {
		t.Errorf("ans should be 18, got %d", ans)
	}

	line = util.SplitToInts(lines[1], " ")
	ans = extrapolateHistory(line)
	if ans != 28 {
		t.Errorf("ans should be 28, got %d", ans)
	}

	line = util.SplitToInts(lines[2], " ")
	ans = extrapolateHistory(line)
	if ans != 68 {
		t.Errorf("ans should be 68, got %d", ans)
	}
}

func Test_extrapolateHistory_Part2(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day9_test.txt")
	line := util.SplitToInts(lines[0], " ")
	ans := extrapolateHistoryPart2(line)
	if ans != -3 {
		t.Errorf("ans should be -3, got %d", ans)
	}

	line = util.SplitToInts(lines[1], " ")
	ans = extrapolateHistoryPart2(line)
	if ans != 0 {
		t.Errorf("ans should be 0, got %d", ans)
	}

	line = util.SplitToInts(lines[2], " ")
	ans = extrapolateHistoryPart2(line)
	if ans != 5 {
		t.Errorf("ans should be 5, got %d", ans)
	}
}
