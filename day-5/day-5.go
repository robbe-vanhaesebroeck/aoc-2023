package main

import (
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Seeds []int

type ConvertRange struct {
	sourceStart, destinationStart, length int
}

type ConvertMap struct {
	source, destination string
	ranges              []ConvertRange
}

var mapHeaderRegex = regexp.MustCompile(`(?P<source>\w+)\-to\-(?P<destination>\w+) map:`)

func parseSeeds(line string) []int {
	seedStrings := strings.Fields(line)[1:]

	seeds := make([]int, len(seedStrings))

	for i, s := range seedStrings {
		val, err := strconv.Atoi(s)

		errorHandling.Check(err)

		seeds[i] = val
	}

	return seeds
}

func parseConvertRange(line string) ConvertRange {
	rangeStrs := strings.Fields(line)

	rangeNums := make([]int, len(rangeStrs))
	for i, s := range rangeStrs {
		val, err := strconv.Atoi(s)

		errorHandling.Check(err)

		rangeNums[i] = val
	}

	return ConvertRange{rangeNums[1], rangeNums[0], rangeNums[2]}
}

func parseInput(fileName string) (Seeds, map[string]*ConvertMap) {
	f, err := os.Open(fileName)

	errorHandling.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	conversionMaps := make(map[string]*ConvertMap)
	var seeds Seeds
	var currentMap = &ConvertMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds = parseSeeds(line)
		} else if match := mapHeaderRegex.FindStringSubmatch(line); len(match) > 0 {
			sourceIdx := mapHeaderRegex.SubexpIndex("source")
			destIdx := mapHeaderRegex.SubexpIndex("destination")

			source := match[sourceIdx]
			dest := match[destIdx]

			currentMap = &ConvertMap{source: source, destination: dest}
			conversionMaps[source] = currentMap
		} else if line == "" && currentMap.source != "" && currentMap.destination != "" {

		} else if line != "" {
			convertRange := parseConvertRange(line)

			currentMap.ranges = append(currentMap.ranges, convertRange)
		}
	}

	return seeds, conversionMaps
}

func convertNumber(seed int, convertMap *ConvertMap) int {
	for _, r := range convertMap.ranges {
		if v := seed - r.sourceStart; v >= 0 && v < r.length {
			return r.destinationStart + v
		}
	}

	return seed
}

func convertTuples(seedTuples [][]int, convertMap *ConvertMap) [][]int {
	convertedTuples := make([][]int, 0)

	for _, tuple := range seedTuples {
		tuples := make([][]int, 0)

		start, length := tuple[0], tuple[1]
		end := start + length

		for _, r := range convertMap.ranges {
			if start >= r.sourceStart && end <= r.sourceStart+r.length {
				// Range falls completely in mapped range
				newStart := r.destinationStart + start - r.sourceStart

				tuples = append(tuples, []int{newStart, length})
				break
			} else if start >= r.sourceStart && start < r.sourceStart+r.length {
				// Range starts after the mapped start but ends after the range ends
				newStart := r.destinationStart + start - r.sourceStart
				newLength := r.sourceStart + r.length - start

				tuples = append(tuples, []int{newStart, newLength})
				tuples = append(tuples, convertTuples([][]int{{r.sourceStart + r.length, length - newLength}}, convertMap)...)
				break
			} else if end >= r.sourceStart && end < r.sourceStart+length {
				// Range starts before start but ends within the range end
				newStart := r.destinationStart
				newLength := end - r.sourceStart

				tuples = append(tuples, []int{newStart, newLength})
				tuples = append(tuples, convertTuples([][]int{{r.sourceStart, length - newLength}}, convertMap)...)
			}
		}

		if len(tuples) == 0 {
			convertedTuples = append(convertedTuples, tuple)
		} else {
			convertedTuples = append(convertedTuples, tuples...)
		}

	}

	return convertedTuples
}

func getLocationNumbers(seeds Seeds, conversionMaps map[string]*ConvertMap) []int {
	locationNumbers := make([]int, len(seeds))

	for i, seed := range seeds {
		currentMap, hasMap := conversionMaps["seed"]
		convertedNumber := seed

		for hasMap {
			convertedNumber = convertNumber(convertedNumber, currentMap)
			currentMap, hasMap = conversionMaps[currentMap.destination]
		}

		locationNumbers[i] = convertedNumber
	}

	return locationNumbers
}

func mapTuples(seedTuples []int, convertMaps map[string]*ConvertMap) []int {
	locationNumbers := make([]int, len(seedTuples)/2)
	for i := 0; i < len(seedTuples); i += 2 {
		currentMap, hasMap := convertMaps["seed"]
		convertedTuples := [][]int{seedTuples[i : i+2]}

		for hasMap {
			convertedTuples = convertTuples(convertedTuples, currentMap)
			currentMap, hasMap = convertMaps[currentMap.destination]
		}

		locationNumbers[i/2] = slices.MinFunc(convertedTuples, func(t1, t2 []int) int { return t1[0] - t2[0] })[0]
	}

	return locationNumbers
}

func easy() {
	seeds, conversionMaps := parseInput("input.txt")

	locationNumbers := getLocationNumbers(seeds, conversionMaps)

	minLocation := slices.Min(locationNumbers)

	fmt.Printf("The minimal location is %d\n", minLocation)
}

func hard() {
	seedTuples, convertMaps := parseInput("input.txt")

	locationNumbers := mapTuples(seedTuples, convertMaps)

	minLocation := slices.Min(locationNumbers)

	fmt.Printf("The minimal location is %d\n", minLocation)
}

func main() {
	fmt.Println("Part one")
	easy()

	fmt.Println("Part two")
	hard()
}
