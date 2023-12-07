/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Day 7",
	Long:  `Day 7`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day7.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/7")
		fmt.Println("Part1:", solveDay7(lines, false))
		fmt.Println("Part2:", solveDay7(lines, true))
		// fmt.Println("Part2:", solveDay6(convertRacesToKerning(lines)))
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}

func solveDay7(lines []string, js bool) int {
	cards := make([]Cards, len(lines))
	for i, line := range lines {
		ls := strings.Split(line, " ")
		hand := ls[0]
		// bid := ls[1]
		cards[i] = Cards{
			Hand:  hand,
			Bid:   util.MustAtoi(ls[1]),
			Score: identifyHand(hand, js),
		}
	}

	// if scores are equal, sort by which hand has the higher first cards
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Score == cards[j].Score {
			q := 0
			for {
				h1 := convertCardToNum(string(cards[i].Hand[q]), js)
				h2 := convertCardToNum(string(cards[j].Hand[q]), js)
				if h1 == h2 {
					q++
					continue
				}
				return h1 > h2
			}
		}
		return cards[i].Score > cards[j].Score
	})

	ans := 0
	for i, c := range cards {
		ans += c.Bid * (len(cards) - i)
	}
	return ans
}

func convertCardToNum(c string, js bool) int {
	if c == "T" {
		return 10
	} else if c == "J" {
		if js {
			return 0
		}
		return 11
	} else if c == "Q" {
		return 12
	} else if c == "K" {
		return 13
	} else if c == "A" {
		return 14
	}
	return util.MustAtoi(c)
}

type Cards struct {
	Hand  string
	Bid   int
	Score int
}

type HandType int

const (
	HighCard HandType = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func identifyHand(hand string, js bool) int {
	cards := make(map[string]int)
	for _, c := range hand {
		cards[string(c)]++
	}

	gc := []int{}
	for k, v := range cards {
		if js && k == "J" {
			continue
		}
		switch v {
		case 2:
			gc = append(gc, int(Pair))
		case 3:
			gc = append(gc, int(ThreeOfAKind))
		case 4:
			if js && cards["J"] == 1 {
				return int(FiveOfAKind)
			}
			return int(FourOfAKind)
		case 5:
			return int(FiveOfAKind)
		}
	}

	// two hands found, don't need extra checks, can all return
	if len(gc) == 2 {
		if (gc[0] == int(Pair) && gc[1] == int(ThreeOfAKind)) ||
			(gc[0] == int(ThreeOfAKind) && gc[1] == int(Pair)) {
			return int(FullHouse)
		} else if gc[0] == int(Pair) && gc[1] == int(Pair) {
			if js && cards["J"] == 1 {
				return int(FullHouse)
			}
			return int(TwoPair)
		}
	}

	// one hand found, if no jokers, return the hand, else process
	if len(gc) == 1 {
		if !js {
			return gc[0]
		}

		// check jokers, if jokers are present, return the next highest hand
		switch gc[0] {
		case int(Pair):
			switch cards["J"] {
			case 1:
				return int(ThreeOfAKind)
			case 2:
				return int(FourOfAKind)
			case 3:
				return int(FiveOfAKind)
			default:
				return int(Pair)
			}
		case int(ThreeOfAKind):
			switch cards["J"] {
			case 1:
				return int(FourOfAKind)
			case 2:
				return int(FiveOfAKind)
			default:
				return int(ThreeOfAKind)
			}
		}
	}

	// no hands found, check for joker hands
	if js {
		if cards["J"] == 1 {
			return int(Pair)
		} else if cards["J"] == 2 {
			return int(ThreeOfAKind)
		} else if cards["J"] == 3 {
			return int(FourOfAKind)
		} else if cards["J"] >= 4 {
			return int(FiveOfAKind)
		}
	}

	return 0
}
