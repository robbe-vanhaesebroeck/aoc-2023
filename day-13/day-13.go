package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"os"
)

type Grid []string

func parseInput(fileName string) []Grid {
	f, err := os.Open(fileName)
	errorHandling.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grids := make([]Grid, 0)
	currentGrid := make(Grid, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Add the finished grid to the list
			grids = append(grids, currentGrid)
			currentGrid = make(Grid, 0)

		} else {
			currentGrid = append(currentGrid, line)
		}
	}
	grids = append(grids, currentGrid)

	return grids
}

func getDifference(r1, r2 string) int {
	dist := 0
	for i := range r1 {
		if r1[i] != r2[i] {
			dist++
		}
	}

	return dist
}

func findRowSymmetry(grid Grid, smudges int) int {
	// Finds the number of rows above a symmetry line
	// Returns 0 if there is none

	for i := 1; i < len(grid); i++ {
		totalSmudges := 0
		for j := 1; i-j >= 0 || i+j-1 < len(grid); j++ {
			// One of the rows falls out of the boundary
			if i-j < 0 || i+j-1 >= len(grid) {
				continue
			}

			// Both the rows fall in the boundary
			totalSmudges += getDifference(grid[i-j], grid[i+j-1])

			if totalSmudges > smudges {
				break
			}
		}

		if totalSmudges == smudges {
			// Return the number of rows above the current one
			return i
		}
	}

	return 0
}

func findColSymmetry(grid Grid, smudges int) int {
	// Finds the number of columns before the symmetry line
	// Returns 0 if there is none
	return findRowSymmetry(common.Transpose(grid), smudges)
}

func easy() {
	grids := parseInput("input.txt")

	sum := 0
	for _, grid := range grids {
		var rows, cols int

		rows = findRowSymmetry(grid, 0)

		if rows == 0 {
			cols = findColSymmetry(grid, 0)
		}
		sum += rows*100 + cols
	}

	fmt.Printf("The sum is %d\n", sum)
}

func hard() {
	grids := parseInput("input.txt")

	sum := 0
	for _, grid := range grids {

		var rows, cols int

		rows = findRowSymmetry(grid, 1)

		if rows == 0 {
			cols = findColSymmetry(grid, 1)
		}
		sum += rows*100 + cols
	}

	fmt.Printf("The sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
