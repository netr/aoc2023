/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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

	// sort cards by score
	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards)-1; j++ {
			if cards[j].Score < cards[j+1].Score {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}

	// if scores are equal, sort by which hand has the higher first cards
	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards)-1; j++ {
			if cards[j].Score == cards[j+1].Score {
				q := 0
				for {
					h1 := convertCardToNum(string(cards[j].Hand[q]), js)
					h2 := convertCardToNum(string(cards[j+1].Hand[q]), js)
					if h1 == h2 {
						q++
						continue
					}
					if h1 < h2 {
						cards[j], cards[j+1] = cards[j+1], cards[j]
					}
					break
				}
			}
		}
	}

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

	// check for 5 of a kind
	goodCards := 0
	for k, v := range cards {
		if js {
			if k == "J" {
				continue
			}
		}
		if v == 5 {
			return int(FiveOfAKind)
		} else if v == 4 {
			if js && cards["J"] == 1 {
				return int(FiveOfAKind)
			}
			return int(FourOfAKind)
		} else if v == 3 {
			goodCards += 3
		} else if v == 2 {
			goodCards += 2
		}
	}

	if goodCards == 5 {
		return int(FullHouse)
	} else if goodCards == 4 {
		if js && cards["J"] == 1 {
			return int(FullHouse)
		}
		return int(TwoPair)
	} else if goodCards == 3 {
		if js {
			if cards["J"] == 1 {
				return int(FourOfAKind)
			} else if cards["J"] == 2 {
				return int(FiveOfAKind)
			}
		}
		return int(ThreeOfAKind)
	} else if goodCards == 2 {
		if js {
			if cards["J"] == 1 {
				return int(ThreeOfAKind)
			} else if cards["J"] == 2 {
				return int(FourOfAKind)
			} else if cards["J"] == 3 {
				return int(FiveOfAKind)
			}
		}
		return int(Pair)
	}

	if goodCards == 0 && js {
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

	return goodCards
}
