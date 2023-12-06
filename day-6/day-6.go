package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Race struct {
	time, distance int
}

func parseInput(fileName string) []Race {
	f, err := os.Open(fileName)

	errorHandling.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var times, distances []int

	if scanner.Scan() {
		timeLine := scanner.Text()
		timesStr := strings.Split(timeLine, ":")[1]

		splitTimeStr := strings.Fields(timesStr)

		times = common.StringListToInt(splitTimeStr)
	} else {
		panic("Could not read time")
	}

	if scanner.Scan() {
		distanceLine := scanner.Text()
		distancesStr := strings.Split(distanceLine, ":")[1]

		splitDistancesStr := strings.Fields(distancesStr)

		distances = common.StringListToInt(splitDistancesStr)
	} else {
		panic("Could not read time")
	}

	races := make([]Race, len(times))

	for i := range times {
		races[i] = Race{time: times[i], distance: distances[i]}
	}

	return races
}

func clamp(n, min, max float64) float64 {
	if n < min {
		return min
	}

	if n > max {
		return max
	}

	return n
}

func getWins(time, distance float64) int {
	// t needs to be positive
	disc := time*time - 4*distance

	if disc < 0 {
		panic("No real solutions")
	}

	t1 := clamp((time-math.Sqrt(disc))/2, 0, time)
	t2 := clamp((time+math.Sqrt(disc))/2, 0, time)

	solutions := []float64{t1, t2}
	slices.Sort(solutions)

	solutionMin := math.Floor(solutions[0])
	solutionMax := math.Ceil(solutions[1])

	numWins := 0
	for i := solutionMin; i <= solutionMax; i++ {
		if i*i-time*i+distance < 0 {
			numWins++
		}
	}

	return numWins
}

func easy() {
	races := parseInput("input.txt")

	numWins := make([]int, len(races))
	for i, r := range races {
		numWins[i] = getWins(float64(r.time), float64(r.distance))
	}

	finalRes := 1
	for _, w := range numWins {
		finalRes *= w
	}

	fmt.Printf("The final number is %d\n", finalRes)
}

func main() {
	fmt.Println("Part one")
	easy()
}
