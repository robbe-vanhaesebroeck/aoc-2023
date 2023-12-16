package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Grid []string

type Coord struct {
	x, y int
}

func parseInput(fileName string) Grid {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	return strings.Split(string(data), "\n")
}

func nextSteps(grid *Grid, prevNode Coord, currentNode Coord) []Coord {
	symbol := (*grid)[currentNode.y][currentNode.x]
	xDiff := currentNode.x - prevNode.x
	yDiff := currentNode.y - prevNode.y

	switch symbol {
	case byte('/'):
		{
			return []Coord{{currentNode.x - yDiff, currentNode.y - xDiff}}
		}
	case byte('\\'):
		{
			return []Coord{{currentNode.x + yDiff, currentNode.y + xDiff}}
		}
	case byte('|'):
		{
			if yDiff != 0 {
				return []Coord{{currentNode.x + xDiff, currentNode.y + yDiff}}
			}

			return []Coord{{currentNode.x, currentNode.y + 1}, {currentNode.x, currentNode.y - 1}}
		}
	case byte('-'):
		{
			if xDiff != 0 {
				return []Coord{{currentNode.x + xDiff, currentNode.y + yDiff}}
			}

			return []Coord{{currentNode.x + 1, currentNode.y}, {currentNode.x - 1, currentNode.y}}
		}
	default:
		{
			return []Coord{{currentNode.x + xDiff, currentNode.y + yDiff}}
		}
	}
}

func tracePath(grid *Grid, pathMap *map[Coord]([]Coord), prevNode Coord, currentNode Coord) {
	if _, ok := (*pathMap)[prevNode]; !ok {
		(*pathMap)[prevNode] = []Coord{currentNode}
	}

	// Iterate as long as our current node is within bounds
	for currentNode.x < len((*grid)[0]) && currentNode.x >= 0 &&
		currentNode.y < len(*grid) && currentNode.y >= 0 {
		nextNodes := nextSteps(grid, prevNode, currentNode)

		newNextNodes := make([]Coord, 0)
		if _, ok := (*pathMap)[currentNode]; ok {
			for _, n := range nextNodes {
				if !slices.Contains((*pathMap)[currentNode], n) {
					(*pathMap)[currentNode] = append((*pathMap)[currentNode], n)
					newNextNodes = append(newNextNodes, n)
				}
			}
		} else {
			(*pathMap)[currentNode] = nextNodes
			newNextNodes = nextNodes
		}

		if len(newNextNodes) > 1 {
			// We have to split
			for _, n := range newNextNodes {
				tracePath(grid, pathMap, currentNode, n)
			}
		} else if len(newNextNodes) == 1 {
			prevNode = currentNode
			currentNode = newNextNodes[0]
		} else {
			// No new nodes
			break
		}
	}
}

func easy() {
	grid := parseInput("input.txt")

	pathMap := make(map[Coord]([]Coord))
	tracePath(&grid, &pathMap, Coord{0, 0}, Coord{1, 0})

	fmt.Printf("The total number of keys is %d\n", len(pathMap))
}

func main() {
	fmt.Println("Part one")
	easy()
}
