package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var resultsRegex = regexp.MustCompile(`(?P<number>\d+) (?P<color>red|blue|green)`)

func parseInput(fileName string) [](map[string][]int) {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	games := strings.Split(string(data), "\n")

	gameResults := make([](map[string][]int), len(games))

	for gameIdx, game := range games {
		allResults := strings.Split(game, ":")[1]
		results := strings.Split(allResults, ";")

		resultMap := make(map[string][]int)

		for _, r := range results {
			matches := resultsRegex.FindAllStringSubmatch(r, -1)

			for _, match := range matches {
				num, err := strconv.Atoi(match[1])
				errorHandling.Check(err)

				color := match[2]

				if val, ok := resultMap[color]; ok {
					resultMap[color] = append(val, num)
				} else {
					intSlice := []int{num}
					resultMap[color] = intSlice
				}
			}
		}

		gameResults[gameIdx] = resultMap
	}

	return gameResults
}

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

func isLegalGame(game map[string][]int) bool {
	redResults := slices.Max(game["red"])
	greenResults := slices.Max(game["green"])
	blueResults := slices.Max(game["blue"])

	// Return true if we have a legal max for all colours
	return redResults <= maxRedCubes && greenResults <= maxGreenCubes && blueResults <= maxBlueCubes
}

func findMinimalPower(game map[string][]int) int {
	redResults := slices.Max(game["red"])
	greenResults := slices.Max(game["green"])
	blueResults := slices.Max(game["blue"])

	return redResults * greenResults * blueResults
}

func easy() {
	games := parseInput("input.txt")

	sum := 0
	for idx, game := range games {
		if isLegalGame(game) {
			sum += idx + 1
		}
	}

	fmt.Printf("The final sum is %d\n", sum)
}

func hard() {
	games := parseInput("input.txt")

	sum := 0
	for _, game := range games {
		sum += findMinimalPower(game)
	}

	fmt.Printf("The final sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
