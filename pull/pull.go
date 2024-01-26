package pull

import (
	"slices"
)

// Pull removes all given values from array using comparable.
// Note: This method mutates input. Use _.remove to remove elements from an slice by predicate.
func Pull[T comparable](input *[]T, values ...T) []T {
	// deal with pointer to input as slicing references a copy
	index := 0
	for {
		if slices.Contains(values, (*input)[index]) {
			copy((*input)[index:], (*input)[index+1:]) // Shift left one index.
			(*input) = (*input)[:len(*input)-1]        // Truncate slice.
		} else {
			index++
		}
		if index >= len((*input)) {
			return *input
		}
	}
}
