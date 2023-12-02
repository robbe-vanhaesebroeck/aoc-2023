package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fileName string) []string {
	rawData, err := os.ReadFile(fileName)
	check(err)

	data := string(rawData)

	return strings.Split(data, "\n")
}

func parseDigitsEasy(line string) int {
	firstDigit := -1
	lastDigit := -1

	for _, ch := range line {
		if unicode.IsDigit(ch) {

			if firstDigit == -1 {
				firstDigit = int(ch - '0')
			}

			lastDigit = int(ch - '0')

		}

	}

	return firstDigit*10 + lastDigit

}

var strDigits = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func startsWithStringDigit(line string) int {
	for strIdx, digit := range strDigits {
		if strings.HasPrefix(line, digit) {
			return strIdx + 1
		}
	}

	return -1
}

func parseDigitsHard(line string) int {
	firstDigit := -1
	lastDigit := -1

	var _assignDigits = func(digit int) {
		if firstDigit == -1 {
			firstDigit = digit
		}

		lastDigit = digit
	}

	for idx, ch := range line {
		if unicode.IsDigit(ch) {
			_assignDigits(int(ch - '0'))
		} else if digit := startsWithStringDigit(line[idx:]); digit != -1 {
			_assignDigits(digit)
		}
	}

	return firstDigit*10 + lastDigit
}

func sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func easy() {
	lines := readLines("input.txt")

	numbers := make([]int, len(lines))

	for _, line := range lines {
		num := parseDigitsEasy(line)
		numbers = append(numbers, num)
	}

	fmt.Printf("The sum is %d\n", sum(numbers))
}

func hard() {
	lines := readLines("input.txt")

	numbers := make([]int, len(lines))

	for _, line := range lines {
		num := parseDigitsHard(line)
		numbers = append(numbers, num)
	}

	fmt.Printf("The sum is %d\n", sum(numbers))
}

func main() {
	fmt.Println("Easy part")
	easy()

	fmt.Println()
	fmt.Println("Hard part")

	hard()
}
