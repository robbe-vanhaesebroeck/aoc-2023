package main

import (
	"aoc-2023/common"
	"aoc-2023/errorHandling"
	"fmt"
	"os"
	"strings"
)

type SpringConfig struct {
	config string
	groups []int
}

func parseInput(fileName string) []SpringConfig {
	data, err := os.ReadFile(fileName)
	errorHandling.Check(err)

	lines := strings.Split(string(data), "\n")

	configs := make([]SpringConfig, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")

		config := split[0]

		strGroups := strings.Split(split[1], ",")
		groups := common.StringListToInt(strGroups)

		configs[i] = SpringConfig{config, groups}
	}

	return configs
}

func findPossibleConfigs(config SpringConfig) int {
	totalConfigs := 0

	// Exit conditions
	if len(config.config) == 0 && len(config.groups) == 0 {
		return 1
	} else if len(config.config) == 0 {
		return 0
	}

	if len(config.config) == 1 {
		if (config.config[0] == byte('.') || config.config[0] == byte('?')) &&
			len(config.groups) == 0 {
			return 1
		}

		if (config.config[0] == byte('#') || config.config[0] == byte('?')) &&
			len(config.groups) == 1 && config.groups[0] == 1 {
			return 1
		}

		return 0
	}

	if len(config.groups) == 0 {
		if strings.ContainsRune(config.config, rune('#')) {
			return 0
		}

		return 1
	}

	// Reducing the input
	switch config.config[0] {
	case byte('?'):
		{
			workingConfig := strings.Replace(config.config, "?", ".", 1)
			totalConfigs += findPossibleConfigs(SpringConfig{workingConfig, config.groups})

			brokenConfig := strings.Replace(config.config, "?", "#", 1)
			totalConfigs += findPossibleConfigs(SpringConfig{brokenConfig, config.groups})
		}
	case byte('.'):
		{
			// A working spring does not alter the groups
			totalConfigs += findPossibleConfigs(SpringConfig{config.config[1:], config.groups})
		}
	case byte('#'):
		{
			groupLength := config.groups[0]

			if len(config.config) < groupLength {
				// Not a valid config
				return 0
			} else if strings.ContainsRune(config.config[:groupLength], rune('.')) {
				// No dots are allowed
				return 0
			} else if len(config.config) == groupLength && len(config.groups) == 1 {
				return 1
			} else if len(config.config) == groupLength {
				// Not all groups are used
				return 0
			}

			// The sequence is longer than the current group, the next element should be a '.'
			nextEl := config.config[groupLength]

			if nextEl == byte('#') {
				return 0
			} else if nextEl == byte('.') {
				totalConfigs += findPossibleConfigs(SpringConfig{config.config[groupLength+1:], config.groups[1:]})
			} else {
				newConfig := strings.Replace(config.config[groupLength:], "?", ".", 1)
				totalConfigs += findPossibleConfigs(SpringConfig{newConfig, config.groups[1:]})
			}
		}
	}

	return totalConfigs
}

func easy() {
	configs := parseInput("input.txt")

	sum := 0
	for _, config := range configs {
		sum += findPossibleConfigs(config)
	}

	fmt.Printf("The total number of configs is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()
}
