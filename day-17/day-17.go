package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Grid []string
type Coord struct {
	x, y int
}

// Up, right, down, left
var delta = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type PathElement struct {
	coord        Coord
	direction    int
	streakLength int
}

type PathCache map[Coord]PathElement

func parseInput(fileName string) Grid {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	return strings.Split(string(data), "\n")
}

func getCost(gridPtr *Grid, current Coord) int {
	grid := *gridPtr

	c := grid[current.y][current.x]

	val := int(c - '0')

	return val
}

func manhattan(node1, node2 Coord) int {
	// Manhattan distance
	return int(math.Abs(float64(node2.x-node1.x)) + math.Abs(float64(node2.y-node1.y)))
}

func reconstructPath(cameFrom PathCache, current Coord) []Coord {
	path := make([]Coord, 1)

	path[0] = current

	prevNode := current
	prevEntry, hasPrev := cameFrom[prevNode]
	for hasPrev {
		prevNode = prevEntry.coord
		path = append(path, prevNode)

		prevEntry, hasPrev = cameFrom[prevNode]
	}

	return path
}

func getNeighbours(gridPtr *Grid, currentElement PathElement, cameFrom *PathCache) []PathElement {
	// Determine potential next nodes for the current node
	// We can't go back
	// We can't go in a straight line for more than 3 turns

	grid := *gridPtr
	maxX := len(grid[0])
	maxY := len(grid)

	neighbours := make([]PathElement, 0, 3)

	currentNode := currentElement.coord

	for direction, offset := range delta {
		newNode := Coord{currentNode.x + offset[1], currentNode.y + offset[0]}

		// If the node is out of bounds, discard
		if newNode.x < 0 || newNode.x >= maxX || newNode.y < 0 || newNode.y >= maxY {
			continue
		}

		if direction == currentElement.direction && currentElement.streakLength >= 3 {
			continue
		}

		prevEl, hasPrev := (*cameFrom)[currentNode]

		if hasPrev && prevEl.coord == newNode {
			continue
		}

		if currentElement.direction == -1 || direction == currentElement.direction {
			neighbours = append(neighbours, PathElement{newNode, direction, currentElement.streakLength + 1})
		} else {
			neighbours = append(neighbours, PathElement{newNode, direction, 1})
		}
	}

	return neighbours
}

func findPath(gridPtr *Grid, startNode Coord, endNode Coord) []Coord {
	pq := make(common.PriorityQueue[PathElement], 0)
	heap.Init(&pq)

	heap.Push(&pq, &common.Item[PathElement]{Value: PathElement{startNode, -1, 0}, Priority: 1})

	cameFrom := make(PathCache)

	gScore := make(map[Coord]int)

	gScore[startNode] = 0

	fScore := make(map[Coord]int)
	fScore[startNode] = manhattan(startNode, endNode)

	for len(pq) > 0 {
		popped := heap.Pop(&pq)

		currentElement := popped.(*common.Item[PathElement]).Value
		currentNode := currentElement.coord

		if currentNode == endNode {
			return reconstructPath(cameFrom, currentNode)
		}

		neighbours := getNeighbours(gridPtr, currentElement, &cameFrom)
		for _, next := range neighbours {
			nextNode := next.coord

			currentGScore, ok := gScore[currentNode]

			if !ok {
				panic("No g score for current node")
			}

			tentativeGScore := currentGScore + getCost(gridPtr, nextNode)
			neighbourGScore, hasScore := gScore[nextNode]

			if !hasScore || tentativeGScore < neighbourGScore {
				cameFrom[nextNode] = PathElement{currentNode, currentElement.direction, currentElement.streakLength}
				gScore[nextNode] = tentativeGScore
				fScore[nextNode] = tentativeGScore + manhattan(nextNode, endNode)

				if !slices.ContainsFunc(pq, func(el *common.Item[PathElement]) bool {
					return el.Value.coord == next.coord && el.Value.direction == next.direction
				}) {
					heap.Push(&pq, &common.Item[PathElement]{Value: next, Priority: fScore[nextNode]})
				}
			}
		}
	}

	return []Coord{}
}

func easy() {
	grid := parseInput("example-input2.txt")

	startNode := Coord{0, 0}
	endNode := Coord{x: len(grid[0]) - 1, y: len(grid) - 1}

	path := findPath(&grid, startNode, endNode)

	fmt.Println(path)

	sum := 0
	for _, c := range path[:len(path)-1] {
		sum += getCost(&grid, c)
	}

	fmt.Println(sum)
}

func main() {
	fmt.Println("Part one")
	easy()
}
