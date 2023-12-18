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

// var directions = []string{"U", "D", "L", "R"}
// var deltas = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

var directions = []string{"R", "D", "L", "U"}
var deltas = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

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

func getHexInstruction(hex string) DigInstruction {
	lenStr := hex[:5]
	dirStr := hex[len(hex)-1]

	dirIdx := int(dirStr - '0')

	length, err := strconv.ParseInt(lenStr, 16, 64)
	errorHandling.Check(err)

	return DigInstruction{dirIdx, int(length), hex}
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

func getLavaVolume(instructions []DigInstruction) int {
	coords := getCoords(instructions)
	area := getArea(coords)

	// We end up where we began, don't count twice
	boundaryPoints := -1
	for _, ins := range instructions {
		boundaryPoints += ins.length
	}

	// Using pick's theorem
	innerPoints := int(math.Ceil(area - float64(boundaryPoints)/2 + 1))

	return innerPoints + boundaryPoints
}

func easy() {
	instructions := parseInput("input.txt")

	lavaVolume := getLavaVolume(instructions)

	fmt.Printf("Total cubic meters of lava is %d\n", lavaVolume)
}

func hard() {
	instructions := parseInput("input.txt")

	hexInstructions := make([]DigInstruction, len(instructions))
	for i, ins := range instructions {
		hexInstructions[i] = getHexInstruction(ins.hex)
	}

	lavaVolume := getLavaVolume(hexInstructions)

	fmt.Printf("Total cubic meters of lava is %d\n", lavaVolume)

}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
