package remove

// Removes all elements from input that predicate returns truthy for and returns an slice of the removed elements.
// The predicate is invoked with three arguments: (value, index, input).

// Note: This mutates input
func Remove[T interface{}](input *[]T, predicate func(o T, i int, orig []T) bool) []T {
	index := 0
	remove := make([]T, 0)
	for {
		if predicate((*input)[index], index, *input) {
			remove = append(remove, (*input)[index])
			copy((*input)[index:], (*input)[index+1:])
			(*input) = (*input)[:len(*input)-1]
		} else {
			index++
		}
		if index >= len((*input)) {
			return remove
		}
	}
}
