package main

import (
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"os"
	"strings"
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

func findRowSymmetry(grid Grid) int {
	// Finds the number of rows above a symmetry line
	// Returns 0 if there is none

	for i := 1; i < len(grid); i++ {
		isSymmetric := true
		for j := 1; i-j >= 0 || i+j-1 < len(grid); j++ {
			// One of the rows falls out of the boundary
			if i-j < 0 || i+j-1 >= len(grid) {
				continue
			}

			// Both the rows fall in the boundary
			if grid[i-j] != grid[i+j-1] {
				isSymmetric = false
				break
			}
		}

		if isSymmetric {
			// Return the number of rows above the current one
			return i
		}
	}

	return 0
}

func transpose(grid Grid) Grid {
	transposed := make(Grid, len(grid[0]))

	for j := 0; j < len(grid[0]); j++ {
		var transposedStrBuilder strings.Builder
		for i := 0; i < len(grid); i++ {
			transposedStrBuilder.WriteByte(grid[i][j])
		}

		transposed[j] = transposedStrBuilder.String()
	}

	return transposed
}

func findColSymmetry(grid Grid) int {
	// Finds the number of columns before the symmetry line
	// Returns 0 if there is none
	return findRowSymmetry(transpose(grid))
}

func easy() {
	grids := parseInput("input.txt")

	sum := 0
	for _, grid := range grids {
		rows := findRowSymmetry(grid)

		cols := findColSymmetry(grid)

		sum += rows*100 + cols
	}

	fmt.Printf("The sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()
}
