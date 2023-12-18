package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type DigInstruction struct {
	direction int
	length    int
	hex       string
}

type Coord struct {
	x, y int
}

const (
	Up = iota
	Down
	Left
	Right
)

var directions = []string{"U", "D", "L", "R"}
var deltas = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

var digRegex = regexp.MustCompile(`(?P<direction>\w) (?P<length>\d+) \(#(?P<hex>.{6})\)`)

func parseInput(fileName string) []DigInstruction {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	digInstructions := make([]DigInstruction, len(lines))

	for i, line := range lines {
		match := digRegex.FindStringSubmatch(line)

		directionIdx := digRegex.SubexpIndex("direction")
		lengthIdx := digRegex.SubexpIndex("length")
		hexIdx := digRegex.SubexpIndex("hex")

		direction := match[directionIdx]
		lengthStr := match[lengthIdx]
		hex := match[hexIdx]

		directionNum := slices.Index(directions, direction)
		length, err := strconv.Atoi(lengthStr)

		errorHandling.Check(err)

		digInstructions[i] = DigInstruction{directionNum, length, hex}
	}

	return digInstructions
}

func getCoords(instructions []DigInstruction) []Coord {
	coords := make([]Coord, 1)
	coords[0] = Coord{0, 0}

	currentX, currentY := 0, 0
	for _, instruction := range instructions {
		delta := deltas[instruction.direction]

		currentX += delta[0] * instruction.length
		currentY += delta[1] * instruction.length

		coords = append(coords, Coord{currentX, currentY})
	}

	return coords
}

func getArea(coords []Coord) float64 {
	// Shoelace formula
	area := 0
	for i, c1 := range coords {
		c2 := coords[(i+1)%len(coords)]
		area += c1.x*c2.y - c1.y*c2.x
	}

	return float64(area) / 2
}

func easy() {
	instructions := parseInput("input.txt")

	coords := getCoords(instructions)
	area := getArea(coords)

	// We end up where we began, don't count twice
	boundaryPoints := -1
	for _, ins := range instructions {
		boundaryPoints += ins.length
	}

	// Using pick's theorem
	innerPoints := int(math.Ceil(area - float64(boundaryPoints)/2 + 1))

	fmt.Println(innerPoints + boundaryPoints)
}

func main() {
	fmt.Println("Part one")
	easy()
}
