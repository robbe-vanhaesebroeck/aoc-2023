package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Coord [2]int

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
			return Coord{i, j}
		}
	}

	panic("No start node found")
}

func step(graph *[]string, prevNode, currentNode Coord) Coord {
	// Takes the previous node and current node to determine the next node
	currentSymbol := (*graph)[currentNode[0]][currentNode[1]]

	verticalDiff := currentNode[0] - prevNode[0]
	horizontalDiff := currentNode[1] - prevNode[1]

	// north to south
	if currentSymbol == byte('|') && verticalDiff != 0 {
		return Coord{currentNode[0] + verticalDiff, currentNode[1]}
	}

	// east to west
	if currentSymbol == byte('-') && horizontalDiff != 0 {
		return Coord{currentNode[0], currentNode[1] + horizontalDiff}
	}

	// north to east
	if currentSymbol == byte('L') {
		return Coord{currentNode[0] - int(math.Abs(float64(horizontalDiff))), currentNode[1] + int(math.Abs(float64(verticalDiff)))}
	}

	// north to west
	if currentSymbol == byte('J') {
		return Coord{currentNode[0] - int(math.Abs(float64(horizontalDiff))), currentNode[1] - int(math.Abs(float64(verticalDiff)))}
	}

	// south to west
	if currentSymbol == byte('7') {
		return Coord{currentNode[0] + int(math.Abs(float64(horizontalDiff))), currentNode[1] - int(math.Abs(float64(verticalDiff)))}
	}

	// south to east
	if currentSymbol == byte('F') {
		return Coord{currentNode[0] + int(math.Abs(float64(horizontalDiff))), currentNode[1] + int(math.Abs(float64(verticalDiff)))}
	}

	return currentNode
}

func findCycles(graph *[]string, startNode Coord) [][]Coord {
	diffs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	cycles := make([][]Coord, 0, 4)

	for _, diff := range diffs {
		currentNode := startNode
		nextNode := Coord{startNode[0] + diff[0], startNode[1] + diff[1]}

		// Node is not a valid node to traverse
		if nextNode[0] < 0 || nextNode[0] >= len(*graph) ||
			nextNode[1] < 0 || nextNode[1] >= len((*graph)[0]) ||
			(*graph)[nextNode[0]][nextNode[1]] == byte('.') {
			continue
		}

		nodes := make([]Coord, 2)
		nodes[0] = startNode
		nodes[1] = nextNode

		for nextNode[0] != startNode[0] || nextNode[1] != startNode[1] {
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

func easy() {
	graph := parseInput("input.txt")

	startNode := getStartNode(graph)

	paths := findCycles(graph, startNode)

	maxPath := slices.MaxFunc(paths, func(p1, p2 []Coord) int { return len(p1) - len(p2) })

	maxDistanceFromStart := len(maxPath) / 2

	fmt.Printf("The maximal distance from start is %d\n", maxDistanceFromStart)
}

func main() {
	fmt.Println("Part one")
	easy()
}
