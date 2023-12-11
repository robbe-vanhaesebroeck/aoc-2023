package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x, y int
}

func parseInput(fileName string) *[]string {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	return &lines
}

func getStartNode(graph *[]string) Coord {
	for i, line := range *graph {
		j := strings.Index(line, "S")

		if j != -1 {
			return Coord{x: j, y: i}
		}
	}

	panic("No start node found")
}

func step(graph *[]string, prevNode, currentNode Coord) Coord {
	// Takes the previous node and current node to determine the next node
	currentSymbol := (*graph)[currentNode.y][currentNode.x]

	verticalDiff := currentNode.y - prevNode.y
	horizontalDiff := currentNode.x - prevNode.x

	// north to south
	if currentSymbol == byte('|') && verticalDiff != 0 {
		return Coord{currentNode.x, currentNode.y + verticalDiff}
	}

	// east to west
	if currentSymbol == byte('-') && horizontalDiff != 0 {
		return Coord{currentNode.x + horizontalDiff, currentNode.y}
	}

	// north to east
	if currentSymbol == byte('L') {
		return Coord{currentNode.x + int(math.Abs(float64(verticalDiff))), currentNode.y - int(math.Abs(float64(horizontalDiff)))}
	}

	// north to west
	if currentSymbol == byte('J') {
		return Coord{currentNode.x - int(math.Abs(float64(verticalDiff))), currentNode.y - int(math.Abs(float64(horizontalDiff)))}
	}

	// south to west
	if currentSymbol == byte('7') {
		return Coord{currentNode.x - int(math.Abs(float64(verticalDiff))), currentNode.y + int(math.Abs(float64(horizontalDiff)))}
	}

	// south to east
	if currentSymbol == byte('F') {
		return Coord{currentNode.x + int(math.Abs(float64(verticalDiff))), currentNode.y + int(math.Abs(float64(horizontalDiff)))}
	}

	return currentNode
}

func findCycles(graph *[]string, startNode Coord) [][]Coord {
	diffs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	cycles := make([][]Coord, 0, 4)

	for _, diff := range diffs {
		currentNode := startNode
		nextNode := Coord{startNode.x + diff[0], startNode.y + diff[1]}

		// Node is not a valid node to traverse
		if nextNode.y < 0 || nextNode.y >= len(*graph) ||
			nextNode.x < 0 || nextNode.x >= len((*graph)[0]) ||
			(*graph)[nextNode.y][nextNode.x] == byte('.') {
			continue
		}

		nodes := make([]Coord, 2)
		nodes[0] = startNode
		nodes[1] = nextNode

		for nextNode.y != startNode.y || nextNode.x != startNode.x {
			nodes = append(nodes, nextNode)

			newNode := step(graph, currentNode, nextNode)

			// Dead end, this is not a valid cycle
			if newNode == currentNode {
				break
			}

			currentNode = nextNode
			nextNode = newNode
		}

		cycles = append(cycles, nodes)
	}

	return cycles
}

func findInsideTiles(graph *[]string, path []Coord) []Coord {
	// This is using the even odd rule
	// For each dot, check horizontally and vertically how many times the path intersects

	insideTiles := make([]Coord, 0)

	for row, line := range *graph {
		isInside := false
		for col, ch := range line {
			inPath := slices.Contains(path, Coord{x: col, y: row})
			if !inPath && isInside {
				insideTiles = append(insideTiles, Coord{x: col, y: row})
			} else if inPath && strings.ContainsRune("|LJ", ch) {
				// Only consider north-facing pipes because they intersect the 'ray'
				isInside = !isInside
			}
		}
	}

	return insideTiles
}

func easy() {
	graph := parseInput("input.txt")

	startNode := getStartNode(graph)

	paths := findCycles(graph, startNode)

	maxPath := slices.MaxFunc(paths, func(p1, p2 []Coord) int { return len(p1) - len(p2) })

	maxDistanceFromStart := len(maxPath) / 2

	fmt.Printf("The maximal distance from start is %d\n", maxDistanceFromStart)
}

func hard() {
	graph := parseInput("input.txt")

	startNode := getStartNode(graph)

	paths := findCycles(graph, startNode)
	maxPath := slices.MaxFunc(paths, func(p1, p2 []Coord) int { return len(p1) - len(p2) })

	insideTiles := findInsideTiles(graph, maxPath)

	fmt.Printf("There are %d inside tiles\n", len(insideTiles))
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
