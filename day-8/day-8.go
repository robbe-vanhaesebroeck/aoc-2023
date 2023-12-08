package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"slices"
	"strings"
)

func parseInput(fileName string) (string, map[string]([2]string)) {
	data, err := os.ReadFile(fileName)

	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	instructions := lines[0]
	connections := make(map[string]([2]string))

	// There's a white space so skip the second line
	for _, l := range lines[2:] {
		split := strings.Split(l, " = ")

		src := split[0]

		dest := strings.Split(split[1], ", ")

		first := strings.Trim(dest[0], "()")
		second := strings.Trim(dest[1], "()")

		destPair := [2]string{first, second}

		connections[src] = destPair
	}

	return instructions, connections
}

const instructionSet = "LR"

func getInstructionIndex(instruction byte) int {
	return strings.IndexByte(instructionSet, instruction)
}

func getPeriods(instructions string, connections map[string]([2]string), startNodes []string, endNodes []string) ([]int, []int) {
	currentNodes := startNodes

	leading := make([]int, len(currentNodes))
	periods := make([]int, len(currentNodes))

	for i := 0; slices.Contains(leading, 0) || slices.Contains(periods, 0); i++ {
		for nodeIdx, currentNode := range currentNodes {
			if periods[nodeIdx] != 0 {
				// If we already have the period, skip
				continue
			}

			connection, ok := connections[currentNode]

			if !ok {
				panic("No connection found")
			}

			instruction := instructions[i%len(instructions)]
			idx := getInstructionIndex(instruction)

			// Travel node
			// We can change it in place because length of array doesn't change
			currentNodes[nodeIdx] = connection[idx]

			// If we have an end node now, fill in the bookkeeping
			if !slices.Contains(endNodes, currentNodes[nodeIdx]) {
				continue
			}

			if leading[nodeIdx] == 0 {
				// 0 based indexing. Actual number of steps should be + 1
				leading[nodeIdx] = i + 1
				continue
			}

			// Subtract the leading period
			periods[nodeIdx] = i + 1 - leading[nodeIdx]
		}
	}

	return leading, periods
}

func easy() {
	instructions, connections := parseInput("input.txt")

	_, periods := getPeriods(instructions, connections, []string{"AAA"}, []string{"ZZZ"})

	numSteps := periods[0]

	fmt.Printf("The number of steps is %d\n", numSteps)
}

func hard() {
	instructions, connections := parseInput("example-input3.txt")

	startNodes := make([]string, 0)
	endNodes := make([]string, 0)
	for key := range connections {
		if strings.HasSuffix(key, "A") {
			startNodes = append(startNodes, key)
		} else if strings.HasSuffix(key, "Z") {
			endNodes = append(endNodes, key)
		}
	}

	leading, periods := getPeriods(instructions, connections, startNodes, endNodes)

	// Printing out to manually verify that periods play nicely
	fmt.Println(leading, periods)

	commonPeriod := common.LCM(periods...)
	fmt.Printf("The number of steps is %d\n", commonPeriod)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
