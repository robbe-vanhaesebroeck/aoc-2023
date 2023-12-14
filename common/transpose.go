package common

import "strings"

func Transpose(grid []string) []string {
	transposed := make([]string, len(grid[0]))

	for j := 0; j < len(grid[0]); j++ {
		var transposedStrBuilder strings.Builder
		for i := 0; i < len(grid); i++ {
			transposedStrBuilder.WriteByte(grid[i][j])
		}

		transposed[j] = transposedStrBuilder.String()
	}

	return transposed
}
