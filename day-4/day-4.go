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

func determinePoints(card ScratchCard) int {
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

	if numMatches > 0 {
		return 1 << (numMatches - 1)
	}

	return 0
}

func easy() {
	scratchCards := parseInput("input.txt")

	totalPoints := 0

	for _, card := range scratchCards {
		totalPoints += determinePoints(card)
	}

	fmt.Printf("The points total is %d\n", totalPoints)
}

func main() {
	fmt.Println("Part one")
	easy()
}
