package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"strings"
)

func parseInput(fileName string) []string {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	return strings.Split(string(data), ",")
}

const multiplier = 17
const modulus = 256

func hash(s string) int {
	currentValue := 0

	// Rune is ASCII value
	for _, r := range s {
		currentValue += int(r)
		currentValue *= multiplier
		currentValue = currentValue % modulus
	}

	return currentValue
}

func easy() {
	strs := parseInput("input.txt")

	total := 0
	for _, s := range strs {
		total += hash(s)
	}

	fmt.Printf("The sum of hashes is %d\n", total)
}

func main() {
	fmt.Println("Part one")
	easy()
}
