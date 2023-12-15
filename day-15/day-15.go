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

type Lens struct {
	label string
	value int
}

type Box []Lens

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

var lensRegex = regexp.MustCompile(`(?P<label>\w+)-?(?:=(?P<value>\d+))?`)

func hashMap(lensBoxes *[256]Box, s string) {
	match := lensRegex.FindStringSubmatch(s)

	if len(match) == 0 {
		panic("Lens does not follow correct format")
	}

	labelIdx := lensRegex.SubexpIndex("label")
	valueIdx := lensRegex.SubexpIndex("value")

	label := match[labelIdx]
	labelHash := hash(label)

	value := match[valueIdx]

	currentBox := lensBoxes[labelHash]

	if value != "" {
		parsedValue, err := strconv.Atoi(value)
		errorHandling.Check(err)

		if labelIdx := slices.IndexFunc(currentBox, func(lens Lens) bool { return lens.label == label }); labelIdx >= 0 {
			lensBoxes[labelHash] = slices.Replace(currentBox, labelIdx, labelIdx+1, Lens{label, parsedValue})
		} else {
			lensBoxes[labelHash] = append(currentBox, Lens{label, parsedValue})
		}

	} else {
		updatedBox := slices.DeleteFunc(currentBox, func(el Lens) bool { return el.label == label })
		(*lensBoxes)[labelHash] = updatedBox
	}

}

func easy() {
	strs := parseInput("input.txt")

	total := 0
	for _, s := range strs {
		total += hash(s)
	}

	fmt.Printf("The sum of hashes is %d\n", total)
}

func hard() {
	instructions := parseInput("input.txt")

	var boxes [256]Box
	for _, instruction := range instructions {
		hashMap(&boxes, instruction)
	}

	focusingPower := 0
	for boxIdx, box := range boxes {
		for slot, lens := range box {
			focusingPower += (1 + boxIdx) * (slot + 1) * lens.value
		}
	}

	fmt.Printf("The total focusing power is %d\n", focusingPower)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
