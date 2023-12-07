package common

import "slices"

func RemoveFromSlice[T comparable](slice []T, el T) []T {
	idx := slices.Index(slice, el)

	if idx == -1 {
		return slice
	}

	return append(slice[:idx], slice[idx+1:]...)
}
