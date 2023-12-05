package cmd

import (
	"testing"
)

func Test_SolveDay5(t *testing.T) {
	expected := 15
	lines := []string{
		"1 2 3 4 5 6 7 8 9 10 11 12 13 14 15",
	}

	if solveDay5(lines) != expected {
		t.Errorf("Expected %d, got %d", expected, solveDay5(lines))
	}
}
