package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"math"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func parseInput(fileName string) *[]string {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	spaceMap := strings.Split(string(data), "\n")

	return &spaceMap
}

func findGalaxies(spaceMap *[]string) []Coord {
	galaxies := make([]Coord, 0)

	for i, line := range *spaceMap {
		for j, r := range line {
			if r == rune('#') {
				galaxies = append(galaxies, Coord{j, i})
			}
		}
	}

	return galaxies
}

func findEmptyRows(spaceMap *[]string) []int {
	emptyRows := make([]int, 0)

	for i, line := range *spaceMap {
		if !strings.ContainsRune(line, rune('#')) {
			emptyRows = append(emptyRows, i)
		}
	}

	return emptyRows
}

func findEmptyColumns(spaceMap *[]string) []int {
	emptyColumns := make([]int, 0)

	for colIdx := range (*spaceMap)[0] {
		isEmpty := true
		for _, line := range *spaceMap {
			if line[colIdx] != byte('.') {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyColumns = append(emptyColumns, colIdx)
		}
	}

	return emptyColumns
}

func isBetween(n, lim1, lim2 int) bool {
	if lim1 < lim2 {
		return lim1 <= n && n <= lim2
	}

	return lim2 <= n && n <= lim1
}

func galaxyDistance(galaxy1, galaxy2 Coord, emptyRows, emptyColumns []int) int {
	// Compute the regular manhattan distance
	dist := int(math.Abs(float64(galaxy1.x)-float64(galaxy2.x)) + math.Abs(float64(galaxy1.y)-float64(galaxy2.y)))

	for _, row := range emptyRows {
		if isBetween(row, galaxy1.y, galaxy2.y) {
			dist++
		}
	}

	for _, col := range emptyColumns {
		if isBetween(col, galaxy1.x, galaxy2.x) {
			dist++
		}
	}

	return dist
}

func easy() {
	spaceMap := parseInput("input.txt")

	galaxies := findGalaxies(spaceMap)
	emptyRows := findEmptyRows(spaceMap)
	emptyColumns := findEmptyColumns(spaceMap)

	sum := 0
	for i, g1 := range galaxies[:len(galaxies)-1] {
		for _, g2 := range galaxies[i+1:] {
			dist := galaxyDistance(g1, g2, emptyRows, emptyColumns)

			sum += dist
		}
	}

	fmt.Printf("The sum of distances is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()
}
