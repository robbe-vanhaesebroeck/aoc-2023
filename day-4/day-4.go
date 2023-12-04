package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	winning []int
	have    []int
}

var scratchCardRegex = regexp.MustCompile(`Card +\d+: +(?P<winning>(?:\d+ +)*)\| +(?P<have>(?:\d+ *)*)`)

func convertStrToIntList(numbersStr string) []int {
	splitStr := strings.Fields(numbersStr)

	intList := make([]int, len(splitStr))

	for i, str := range splitStr {
		val, err := strconv.Atoi(str)
		errorHandling.Check(err)

		intList[i] = val
	}

	return intList
}

func parseInput(fileName string) []ScratchCard {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	cards := strings.Split(string(data), "\n")

	scratchCards := make([]ScratchCard, len(cards))

	for i, card := range cards {
		// Parse the numbers
		match := scratchCardRegex.FindStringSubmatch(card)

		winningIndex := scratchCardRegex.SubexpIndex("winning")
		haveIndex := scratchCardRegex.SubexpIndex("have")

		winningNumbersStr := match[winningIndex]
		haveNumbersStr := match[haveIndex]

		winningNumbers := convertStrToIntList(winningNumbersStr)
		haveNumbers := convertStrToIntList(haveNumbersStr)

		scratchCards[i] = ScratchCard{winning: winningNumbers, have: haveNumbers}
	}

	return scratchCards
}

func determineNumMatches(card ScratchCard) int {
	numMatches := 0
	for _, winning := range card.winning {
		for _, have := range card.have {
			if winning == have {
				numMatches += 1
				// Count each winning number just once
				break
			}
		}
	}

	return numMatches
}

func convertMatchesToPoints(numMatches int) int {
	if numMatches > 0 {
		return 1 << (numMatches - 1)
	}

	return 0
}

func determinePoints(card ScratchCard) int {
	numMatches := determineNumMatches(card)

	return convertMatchesToPoints(numMatches)
}

func easy() {
	scratchCards := parseInput("input.txt")

	totalPoints := 0

	for _, card := range scratchCards {
		totalPoints += determinePoints(card)
	}

	fmt.Printf("The points total is %d\n", totalPoints)
}

func hard() {
	scratchCards := parseInput("input.txt")

	numCards := make([]int, len(scratchCards))
	matches := make([]int, len(scratchCards))

	for i, card := range scratchCards {
		numCards[i] += 1
		numMatches := determineNumMatches(card)
		matches[i] = numMatches

		for j := 1; j <= numMatches; j++ {
			numCards[i+j] += numCards[i]
		}
	}

	totalCards := 0

	for _, c := range numCards {
		totalCards += c
	}

	fmt.Printf("The total number of scratch cards is %d\n", totalCards)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
