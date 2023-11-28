package union

import "slices"

/*
Creates a slice of unique values, in order, from all given slices
*/
func Union[T comparable](inputs ...[]T) []T {
	result := make([]T, 0)
	for _, input := range inputs {
		for _, element := range input {
			if !slices.Contains(result, element) {
				result = append(result, element)
			}
		}
	}
	return result
}
