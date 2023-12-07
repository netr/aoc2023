package cmd

import (
	"testing"

	"github.com/netr/aoc/util"
)

func Test_SolveDay7(t *testing.T) {
	lines := util.ReadFileIntoSlice("../data/day7_test.txt")
	ans := solveDay7(lines, false)
	if ans != 6440 {
		t.Errorf("ans should be 6440, got %d", ans)
	}

	ans = solveDay7(lines, true)
	if ans != 5905 {
		t.Errorf("ans should be 5905, got %d", ans)
	}
}

func Test_identifyHand(t *testing.T) {
	val := identifyHand("32T3K", false)
	if val != 2 {
		t.Errorf("val should be 2, got %d", val)
	}
	val = identifyHand("T55J5", false)
	if val != 3 {
		t.Errorf("val should be 3, got %d", val)
	}
	val = identifyHand("KK677", false)
	if val != 4 {
		t.Errorf("val should be 4, got %d", val)
	}
	val = identifyHand("KTJJT", false)
	if val != 4 {
		t.Errorf("val should be 4, got %d", val)
	}
	val = identifyHand("QQQJA", false)
	if val != 3 {
		t.Errorf("val should be 3, got %d", val)
	}
}

func Test_identifyHand2(t *testing.T) {
	val := identifyHand("32T3K", true)
	if val != int(Pair) {
		t.Errorf("val should be %d, got %d", int(Pair), val)
	}
	val = identifyHand("T55J5", true)
	if val != int(FourOfAKind) {
		t.Errorf("val should be %d, got %d", int(FourOfAKind), val)
	}
	val = identifyHand("KK677", true)
	if val != int(TwoPair) {
		t.Errorf("val should be %d, got %d", int(TwoPair), val)
	}
	val = identifyHand("KTJJT", true)
	if val != int(FourOfAKind) {
		t.Errorf("val should be %d, got %d", int(FourOfAKind), val)
	}
	val = identifyHand("QQQJA", true)
	if val != int(FourOfAKind) {
		t.Errorf("val should be %d, got %d", int(FourOfAKind), val)
	}
}
