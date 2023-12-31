package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(fileName string) *[]*[]int {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	readings := make([]*[]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)

		history := make([]int, len(fields))
		for j, field := range fields {
			val, err := strconv.Atoi(field)
			errorHandling.Check(err)

			history[j] = val
		}

		readings[i] = &history
	}

	return &readings
}

func extraPolateValue(reading *[]int, reverse bool) int {
	readingVal := *reading
	differences := make([]int, len(readingVal)-1)

	for i := 0; i < len(differences); i++ {
		differences[i] = readingVal[i+1] - readingVal[i]
	}

	// If we only have zeroes, extrapolate
	if !slices.ContainsFunc(differences, func(el int) bool { return el != 0 }) {

		if reverse {
			return readingVal[0] - differences[0]
		}

		return readingVal[len(readingVal)-1] + differences[len(differences)-1]
	}

	// Otherwise recursively find the next value on these differences
	nextValue := extraPolateValue(&differences, reverse)

	if reverse {
		return readingVal[0] - nextValue
	}

	return readingVal[len(readingVal)-1] + nextValue
}

func easy() {
	readings := parseInput("input.txt")

	sum := 0
	for _, r := range *readings {
		sum += extraPolateValue(r, false)
	}

	fmt.Printf("The total sum is %d\n", sum)
}

func hard() {
	readings := parseInput("input.txt")

	sum := 0
	for _, r := range *readings {
		sum += extraPolateValue(r, true)
	}

	fmt.Printf("The total sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
