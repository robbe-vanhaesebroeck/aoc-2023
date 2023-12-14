package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"strings"
)

type Grid []string

func parseInput(fileName string) Grid {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	return lines
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
				cube_row = row + 1
				num_round = 0
			}
		}
	}

	fmt.Printf("The total load is %d\n", total_load)
}

func main() {
	fmt.Println("Part one")
	easy()
}
