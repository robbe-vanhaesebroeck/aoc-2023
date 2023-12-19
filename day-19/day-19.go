package main

import (
	"aoc-2023/errorHandling"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	category rune
	smaller  bool
	target   int
	dest     string
}

type Part map[rune]int

var ruleRegex = regexp.MustCompile(`(?P<workflow>\w+)\{(?P<rules>.*)\}`)

func parseRules(scanner *bufio.Scanner) map[string]([]Rule) {
	rulesMap := make(map[string]([]Rule))

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Break to the next part
			break
		}

		match := ruleRegex.FindStringSubmatch(line)

		if match == nil {
			log.Fatalf("No match found for %s\n", line)
		}

		workflowIdx := ruleRegex.SubexpIndex("workflow")
		rulesIdx := ruleRegex.SubexpIndex("rules")

		workflowName := match[workflowIdx]
		rawRules := match[rulesIdx]

		splitRules := strings.Split(rawRules, ",")

		rules := make([]Rule, len(splitRules))

		for i, splitRule := range splitRules {
			split := strings.Split(splitRule, ":")

			if len(split) == 1 {
				// No condition, set target to -1 to signify this
				rules[i] = Rule{target: -1, dest: split[0]}
				continue
			}

			dest := split[1]

			condition := split[0]
			category := rune(condition[0])
			smaller := condition[1] == byte('<')
			target, err := strconv.Atoi(condition[2:])
			errorHandling.Check(err)

			rules[i] = Rule{category, smaller, target, dest}
		}

		rulesMap[workflowName] = rules
	}

	return rulesMap
}

func parseParts(scanner *bufio.Scanner) []Part {
	parts := make([]Part, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// Removes "{" and "}"
		trimmed := line[1 : len(line)-1]
		split := strings.Split(trimmed, ",")

		part := make(Part)

		for _, s := range split {
			category := rune(s[0])
			val, err := strconv.Atoi(s[2:])
			errorHandling.Check(err)

			part[category] = val
		}

		parts = append(parts, part)
	}

	return parts
}

func parseInput(fileName string) (map[string]([]Rule), []Part) {
	f, err := os.Open(fileName)

	errorHandling.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rulesMap := parseRules(scanner)
	parts := parseParts(scanner)

	return rulesMap, parts
}

func acceptPart(part Part, rulesMapPtr *map[string][]Rule, startRule string) bool {
	rulesMap := *rulesMapPtr

	currentName := startRule
	for currentName != "R" && currentName != "A" {
		rules, ok := rulesMap[currentName]

		if !ok {
			log.Fatalf("No rule found with the name %s\n", currentName)
		}

		for _, rule := range rules {
			if rule.target == -1 {
				// Go to next immediately
				currentName = rule.dest
				break
			}

			val := part[rule.category]

			if (rule.smaller && val < rule.target) ||
				(!rule.smaller && val > rule.target) {
				currentName = rule.dest
				break
			}
		}
	}

	return currentName == "A"
}

func easy() {
	rulesMap, parts := parseInput("input.txt")

	accepted := make([]Part, 0)
	for _, part := range parts {
		if acceptPart(part, &rulesMap, "in") {
			accepted = append(accepted, part)
		}
	}

	sum := 0
	for _, part := range accepted {
		sum += part[rune('x')] + part[rune('m')] + part[rune('a')] + part[rune('s')]
	}

	fmt.Printf("The final sum is %d\n", sum)
}

func main() {
	fmt.Println("Part one")
	easy()
}
