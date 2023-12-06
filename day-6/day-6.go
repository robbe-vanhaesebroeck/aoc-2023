package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Race struct {
	time, distance int
}

func parseInput(fileName string) (string, string) {
	f, err := os.Open(fileName)

	errorHandling.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var timesStr, distancesStr string

	if scanner.Scan() {
		timeLine := scanner.Text()
		timesStr = strings.Split(timeLine, ":")[1]

	} else {
		panic("Could not read time")
	}

	if scanner.Scan() {
		distanceLine := scanner.Text()
		distancesStr = strings.Split(distanceLine, ":")[1]
	} else {
		panic("Could not read distance")
	}

	return timesStr, distancesStr
}

func parseInputEasy(fileName string) []Race {
	timesStr, distancesStr := parseInput(fileName)

	splitTimeStr := strings.Fields(timesStr)
	times := common.StringListToInt(splitTimeStr)

	splitDistancesStr := strings.Fields(distancesStr)
	distances := common.StringListToInt(splitDistancesStr)

	races := make([]Race, len(times))

	for i := range times {
		races[i] = Race{time: times[i], distance: distances[i]}
	}

	return races
}

func parseInputHard(fileName string) Race {
	timesStr, distancesStr := parseInput(fileName)

	time, err := strconv.Atoi(strings.Join(strings.Split(timesStr, " "), ""))
	errorHandling.Check(err)

	distance, err := strconv.Atoi(strings.Join(strings.Split(distancesStr, " "), ""))
	errorHandling.Check(err)

	return Race{time, distance}
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

func determineValue(t, time, distance float64) float64 {
	return t*t - time*t + distance
}

func getWins(time, distance float64) int {
	// We know that time (t) has an acceleration and a travel part
	// t = t_a + t_v
	// The time spent accelerating (t_a) is equal to the velocity
	// v = t_a
	// We also know that the distance (d) travelled is
	// d = v * t_v
	//	 = t_a * t_v
	//   = t_a * (t - t_a)
	// <=>
	// t_a^2 - t * t_a + d = 0
	// Our distance d should be greater than D (the current record)
	// So, this means for a total time T
	// v * t_v > D
	// <=>
	// t_a * (T - t_a) > D
	// <=>
	// t_a^2 - T * t_a + D < 0
	// We then solve for t_a under the constraints that
	// - t_a >= 0
	// - t_a <= T

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

	numWins := int(solutionMax - solutionMin + 1)

	// Check that the boundaries are actually solutions
	if determineValue(solutionMax, time, distance) >= 0 {
		numWins--
	}

	if determineValue(solutionMin, time, distance) >= 0 {
		numWins--
	}

	return numWins
}

func easy() {
	races := parseInputEasy("input.txt")

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

func hard() {
	race := parseInputHard("input.txt")

	numWins := getWins(float64(race.time), float64(race.distance))

	fmt.Printf("The final number is %d\n", numWins)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
