package cmd

import (
	"testing"
)

func Test_SolveDay6(t *testing.T) {
	races := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	// first race lasts 7 ms, best distance is 9 mm
	// Your toy boat has a starting speed of zero millimeters per millisecond. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.
	ans := solveDay6(races)
	if ans != 288 {
		t.Errorf("ans should be 288, got %d", ans)
	}

	ans2 := solveDay6Kerning(convertRacesToKerning(races))
	if ans2 != 71503 {
		t.Errorf("ans2 should be 71503, got %d", ans2)
	}

	ans3 := solveDay6Quadratic(convertRacesToKerning(races))
	if ans3 != 71503 {
		t.Errorf("ans3 should be 71503, got %d", ans3)
	}

}

func Test_ParseRaces(t *testing.T) {
	races := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	times, dists := parseRaces(races)
	if len(times) != 3 {
		t.Errorf("should be 3 times, got %d", len(times))
	}
	if len(dists) != 3 {
		t.Errorf("should be 3 dists, got %d", len(dists))
	}
}
