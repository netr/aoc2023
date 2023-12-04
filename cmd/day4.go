package cmd

import (
	"fmt"
	"strings"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  `day 4`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day4.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/4")

		totalScore := 0
		for _, card := range lines {
			score := calculateScratchCard(card)
			totalScore += score
		}

		fmt.Println("Part1:", totalScore)

		cardsMap := cardsToMap(lines)
		totalScore = 0

		var cards []ScratchCard
		for _, card := range cardsMap {
			cards = append(cards, card)
		}

		for len(cards) > 0 {
			card := cards[0]
			cards = cards[1:]
			score := cardsMap[card.CardId].Score()
			for i := 1; i <= score; i++ {
				cards = append(cards, cardsMap[card.CardId+i])
			}
			totalScore += 1
		}

		fmt.Println("Part2:", totalScore)
	},
}

func removeDuplicateSpaces(s string) string {
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}

func calculateScratchCard(card string) int {
	game := strings.Split(card, ": ")[1]
	numbers := strings.Split(removeDuplicateSpaces(game), " | ")
	winning := strings.Split(numbers[0], " ")
	winMap := scratchNumsToMap(winning)
	ours := strings.Split(numbers[1], " ")

	score := 0
	for _, n := range ours {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if _, ok := winMap[util.MustAtoi(n)]; ok {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}

func calculateScratchCardAdd(card string) int {
	game := strings.Split(card, ": ")[1]
	numbers := strings.Split(removeDuplicateSpaces(game), " | ")
	winning := strings.Split(numbers[0], " ")
	winMap := scratchNumsToMap(winning)
	ours := strings.Split(numbers[1], " ")

	score := 0
	for _, n := range ours {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if _, ok := winMap[util.MustAtoi(n)]; ok {
			score += 1
		}
	}

	return score
}

type ScratchCard struct {
	CardId int
	Win    map[int]struct{}
	Our    map[int]struct{}
}

func (sc ScratchCard) Score() int {
	score := 0
	for n := range sc.Our {
		if _, ok := sc.Win[n]; ok {
			score += 1
		}
	}
	return score
}

func cardsToMap(cards []string) map[int]ScratchCard {
	cardMap := make(map[int]ScratchCard)
	for _, card := range cards {
		cardStr := strings.Split(card, ": ")[0]
		game := strings.Split(card, ": ")[1]
		cardGame := strings.Split(cardStr, " ")
		cardIdStr := cardGame[len(cardGame)-1]
		cardId := util.MustAtoi(cardIdStr)
		numbers := strings.Split(removeDuplicateSpaces(game), " | ")

		winning := strings.Split(numbers[0], " ")
		winMap := scratchNumsToMap(winning)

		ours := strings.Split(numbers[1], " ")
		ourMap := scratchNumsToMap(ours)

		cardMap[cardId] = ScratchCard{
			CardId: cardId,
			Win:    winMap,
			Our:    ourMap,
		}

	}

	return cardMap
}

func scratchNumsToMap(i []string) map[int]struct{} {
	m := make(map[int]struct{})
	for _, n := range i {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		m[util.MustAtoi(n)] = struct{}{}
	}
	return m
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
