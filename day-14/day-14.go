package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Grid []string

var directions = [...]string{"north", "west", "south", "east"}

func parseInput(fileName string) Grid {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	return lines
}

func reverse(s string) string {
	runeLine := []rune(s)
	for i, j := 0, len(runeLine)-1; i < len(runeLine)/2; i, j = i+1, j-1 {
		runeLine[i], runeLine[j] = runeLine[j], runeLine[i]
	}
	return string(runeLine)
}

func reverseRows(grid Grid) Grid {
	reversedGrid := make(Grid, len(grid))

	for row, line := range grid {
		reversedGrid[row] = reverse(line)
	}

	return reversedGrid
}

func tiltRowsWest(grid Grid) Grid {
	tiltedGrid := make(Grid, len(grid))

	var builder strings.Builder
	for row, line := range grid {
		partitioned := strings.Split(line, "#")

		// Sort each partition so O comes first
		for _, part := range partitioned {
			runes := []rune(part)
			// Sort descending
			slices.SortFunc(runes, func(r1, r2 rune) int { return int(r2 - r1) })

			builder.WriteString(string(runes) + "#")
		}

		tiltedRow := builder.String()
		tiltedGrid[row] = tiltedRow[:len(tiltedRow)-1]
		builder.Reset()
	}
	return tiltedGrid
}

func tilt(grid Grid, direction string) Grid {
	switch direction {
	case "north":
		{
			// Reduce to west case by transposing
			return common.Transpose(tilt(common.Transpose(grid), "west"))
		}
	case "south":
		{
			// Reduce to east case by transposing
			return common.Transpose(tilt(common.Transpose(grid), "east"))
		}
	case "east":
		{
			// Reduce to west case by reversing each row
			return reverseRows(tilt(reverseRows(grid), "west"))
		}
	default:
		{
			return tiltRowsWest(grid)
		}
	}
}

const maxIterations = 1000000000

func spinCycle(grid Grid) ([]Grid, int) {
	grids := make(map[string]int)

	gridPeriod := make([]Grid, 0)

	currentGrid := grid
	for i := 0; i < maxIterations; i++ {
		// Do the transformation
		intermediateGrid := currentGrid
		for _, d := range directions {
			intermediateGrid = tilt(intermediateGrid, d)
		}
		currentGrid = intermediateGrid

		// If we already have the grid either set leading or return
		if previousOccurrence, hasGrid := grids[strings.Join(currentGrid, " ")]; hasGrid {
			// Add 1 because we only started the repetition on the next iteration
			return gridPeriod[previousOccurrence:], previousOccurrence + 1
		}

		// Record the grid in the cache
		grids[strings.Join(currentGrid, " ")] = i
		gridPeriod = append(gridPeriod, currentGrid)

	}

	return gridPeriod, maxIterations
}

func calculateNorthLoad(grid Grid) int {
	load := 0
	total_rows := len(grid)
	for col := range grid[0] {
		for row := range grid {
			if grid[row][col] == byte('O') {
				load += total_rows - row
			}
		}
	}

	return load
}

func easy() {
	grid := parseInput("input.txt")

	total_load := 0
	total_rows := len(grid)
	for col := range grid[0] {
		cube_row := 0
		num_round := 0
		for row := range grid {
			if grid[row][col] == byte('O') {
				total_load += total_rows - cube_row - num_round
				num_round++
			} else if grid[row][col] == byte('#') {
				// Reset the count as if this is the top
				// Add 1 because we can't count the current row for the next round boulders
				cube_row = row + 1
				num_round = 0
			}
		}
	}

	fmt.Printf("The total load is %d\n", total_load)
}

func hard() {
	grid := parseInput("input.txt")

	grids, leading := spinCycle(grid)

	finalGrid := grids[(maxIterations-leading)%len(grids)]
	finalLoad := calculateNorthLoad(finalGrid)

	fmt.Printf("The final load is %d\n", finalLoad)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
