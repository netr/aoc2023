package cmd

import (
	"testing"
)

func Test_SolveScrathCard(t *testing.T) {
	cards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	totalScore := 0
	for _, card := range cards {
		score := calculateScratchCard(card)
		totalScore += score
	}

	if totalScore != 13 {
		t.Errorf("Expected 13, got %d", totalScore)
	}
}

func Test_SolveScrathCardRecursive(t *testing.T) {
	cardStrings := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	cardsMap := cardsToMap(cardStrings)
	totalScore := 0

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

	if totalScore != 30 {
		t.Errorf("Expected 30, got %d", totalScore)
	}
}
