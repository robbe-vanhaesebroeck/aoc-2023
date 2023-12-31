package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type CardTuple struct {
	strength int
	count    int
}

type Hand struct {
	originalCards string
	cards         []CardTuple
	bid           int
}

const cardStrength = "23456789TJQKA"

func getStrengthEasy(cardFace byte) int {
	return strings.IndexByte(cardStrength, cardFace)
}

func getStrengthHard(cardFace byte) int {
	if cardFace == byte('J') {
		return -1
	}

	return getStrengthEasy(cardFace)
}

func getCompareCardTuple(getStrength func(face byte) int) func(c1 CardTuple, c2 CardTuple) int {
	return func(c1, c2 CardTuple) int {
		if c1.count != c2.count {
			return c2.count - c1.count
		}

		// If the count is the same, check the face of the card
		return c2.strength - c1.strength
	}
}

func parseHand(cards string, getStrength func(face byte) int) []CardTuple {
	counter := make(map[rune]int)
	for _, face := range cards {
		counter[face] += 1
	}

	tuples := make([]CardTuple, len(counter))

	cnt := 0
	for face, count := range counter {
		strength := getStrength(byte(face))
		t := CardTuple{strength, count}
		tuples[cnt] = t
		cnt++
	}

	slices.SortStableFunc(tuples, getCompareCardTuple(getStrength))

	return tuples
}

func parseInput(fileName string, getStrength func(face byte) int) []*Hand {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	hands := make([]*Hand, len(lines))

	for i, line := range lines {
		splitHand := strings.Fields(line)

		bid, err := strconv.Atoi(splitHand[1])
		errorHandling.Check(err)

		hands[i] = &Hand{cards: parseHand(splitHand[0], getStrength), bid: bid, originalCards: splitHand[0]}
	}

	return hands
}

func scoreHand(cards []CardTuple) int {
	// 5 of a kind
	if len(cards) == 1 {
		// 5 of a kind, early return
		return 0
	}

	// 4 of a kind
	if cards[0].count == 4 {
		// Add card length so 4 of a kind is always lower than 5 of a kind
		return 1
	}

	// Full house
	if cards[0].count == 3 && cards[1].count == 2 {
		return 2
	}

	// 3 of a kind
	if cards[0].count == 3 {
		return 3
	}

	// Double pair
	if cards[0].count == 2 && cards[1].count == 2 {
		return 4
	}

	// Pair
	if cards[0].count == 2 {
		return 5
	}

	// High card
	return 6

}

func getCompareHands(getStrength func(face byte) int) func(h1 *Hand, h2 *Hand) int {
	return func(h1 *Hand, h2 *Hand) int {
		s1, s2 := scoreHand(h1.cards), scoreHand(h2.cards)

		if s1 != s2 {
			return s2 - s1
		}

		for i := range h1.originalCards {
			c1 := h1.originalCards[i]
			c2 := h2.originalCards[i]

			strength1 := getStrength(c1)
			strength2 := getStrength(c2)

			if strength1 != strength2 {
				return strength1 - strength2
			}
		}
		return 0
	}
}

func transformJokers(hands []*Hand) {
	for _, hand := range hands {
		jokerIdx := slices.IndexFunc(hand.cards, func(h CardTuple) bool { return h.strength == getStrengthHard(byte('J')) })

		// If not found or if joker is the first (and only) card
		if jokerIdx < 0 || len(hand.cards) == 1 {
			continue
		}

		cnt := hand.cards[jokerIdx].count

		// Remove joker from the hand
		hand.cards = common.RemoveFromSlice(hand.cards, hand.cards[jokerIdx])

		// Add count to first element
		hand.cards[0].count += cnt
	}
}

func easy() {
	hands := parseInput("input.txt", getStrengthEasy)

	// Sort the hands reversed so the weakest comes first
	slices.SortStableFunc(hands, getCompareHands(getStrengthEasy))

	// Compute the total winnings
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("The total winnings are %d\n", winnings)
}

func hard() {
	hands := parseInput("input.txt", getStrengthHard)

	transformJokers(hands)

	// Sort the hands reversed so the weakest comes first
	slices.SortStableFunc(hands, getCompareHands(getStrengthHard))

	// Compute the total winnings
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("The total winnings are %d\n", winnings)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
