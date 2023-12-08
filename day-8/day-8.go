package main

import (
	"aoc-2023/errorHandling"
	"fmt"
	"os"
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
const startNode = "AAA"
const endNode = "ZZZ"

func getInstructionIndex(instruction byte) int {
	return strings.IndexByte(instructionSet, instruction)
}

func walkNetwork(instructions string, connections map[string]([2]string)) int {
	currentNode := startNode

	numSteps := 0
	for ; currentNode != endNode; numSteps++ {
		connection, ok := connections[currentNode]

		if !ok {
			panic("No connection found")
		}

		instruction := instructions[numSteps%len(instructions)]
		idx := getInstructionIndex(instruction)

		// Travel node
		currentNode = connection[idx]
	}

	return numSteps
}

func easy() {
	instructions, connections := parseInput("input.txt")

	numSteps := walkNetwork(instructions, connections)

	fmt.Printf("The number of steps is %d\n", numSteps)
}

func main() {
	fmt.Println("Part one")
	easy()
}
