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

func easy() {
	seeds, conversionMaps := parseInput("input.txt")

	locationNumbers := getLocationNumbers(seeds, conversionMaps)

	minLocation := slices.Min(locationNumbers)

	fmt.Printf("The minimal location is %d\n", minLocation)
}

func main() {
	fmt.Println("Part one")
	easy()
}
