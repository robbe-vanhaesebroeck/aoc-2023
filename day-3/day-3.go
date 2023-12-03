package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var symbolsRegex = regexp.MustCompile(`[^A-Za-z0-9\.]`)
var gearRegex = regexp.MustCompile(`\*`)
var digitsRegex = regexp.MustCompile(`\d+`)

func parseInput(fileName string) []string {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	return strings.Split(string(data), "\n")
}

func checkSpecialChar(lines []string, digitsRange []int, lineNumber int) bool {
	if lineNumber < 0 || lineNumber >= len(lines) {
		return false
	}

	line := lines[lineNumber]

	specialCharRanges := symbolsRegex.FindAllStringIndex(line, -1)

	for _, specialCharRange := range specialCharRanges {
		if specialCharRange[0] >= digitsRange[0]-1 && specialCharRange[0] <= digitsRange[1] {
			return true
		}
	}

	return false
}

func findGearDigits(lines []string, specialCharIdx int, lineNumber int) [2]int {
	var gearDigits [2]int
	for i := lineNumber - 1; i <= lineNumber+1; i++ {
		if i < 0 || i >= len(lines) {
			break
		}

		line := lines[i]
		digitsRanges := digitsRegex.FindAllStringIndex(line, -1)

		for _, digitsRange := range digitsRanges {
			if specialCharIdx >= digitsRange[0]-1 && specialCharIdx <= digitsRange[1] {
				val, err := strconv.Atoi(line[digitsRange[0]:digitsRange[1]])

				errorHandling.Check(err)

				if gearDigits[0] == 0 {
					gearDigits[0] = val
				} else if gearDigits[1] == 0 {
					gearDigits[1] = val
					return gearDigits
				}
			}
		}
	}

	return gearDigits
}

func easy() {
	lines := parseInput("input.txt")

	sum := 0
	for i, line := range lines {
		digitsRanges := digitsRegex.FindAllStringIndex(line, -1)

		for _, digitsRange := range digitsRanges {
			prevLineMatch := checkSpecialChar(lines, digitsRange, i-1)
			currentLineMatch := checkSpecialChar(lines, digitsRange, i)
			nextLineMatch := checkSpecialChar(lines, digitsRange, i+1)

			if prevLineMatch || currentLineMatch || nextLineMatch {
				val, err := strconv.Atoi(line[digitsRange[0]:digitsRange[1]])
				errorHandling.Check(err)

				sum += val
			}
		}
	}

	fmt.Printf("The sum is %d\n", sum)
}

func hard() {
	lines := parseInput("input.txt")

	sum := 0

	for i, line := range lines {
		gearIndices := gearRegex.FindAllStringIndex(line, -1)

		for _, gearIdx := range gearIndices {
			gears := findGearDigits(lines, gearIdx[0], i)

			sum += gears[0] * gears[1]
		}

	}

	fmt.Printf("The sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
