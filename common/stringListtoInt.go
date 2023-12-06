package common

import (
	"aoc-2023/errorHandling"
	"strconv"
)

func StringListToInt(strs []string) []int {
	res := make([]int, len(strs))
	for i, s := range strs {
		val, err := strconv.Atoi(s)

		errorHandling.Check(err)

		res[i] = val
	}

	return res
}
