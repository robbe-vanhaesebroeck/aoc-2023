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

func getPeriod(instructions string, connections *map[string]([2]string), startNode string, endNodes []string, ch chan [2]int) {
	currentNode := startNode

	leading, period := 0, 0

	for i := 0; period == 0; i++ {
		connection, ok := (*connections)[currentNode]

		if !ok {
			panic("No connection found")
		}

		instruction := instructions[i%len(instructions)]
		idx := getInstructionIndex(instruction)

		currentNode = connection[idx]

		if !slices.Contains(endNodes, currentNode) {
			continue
		}

		if leading == 0 {
			leading = i + 1
			continue
		}

		period = i + 1 - leading
	}

	ch <- [2]int{leading, period}
}

func getPeriods(instructions string, connections map[string]([2]string), startNodes []string, endNodes []string) ([]int, []int) {
	ch := make(chan [2]int)
	for _, node := range startNodes {
		go getPeriod(instructions, &connections, node, endNodes, ch)
	}

	leading := make([]int, len(startNodes))
	periods := make([]int, len(startNodes))
	for i := range startNodes {
		p := <-ch

		leading[i] = p[0]
		periods[i] = p[1]
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
	instructions, connections := parseInput("input.txt")

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
