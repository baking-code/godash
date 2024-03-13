package xor

// Xor creates an array of unique values that is the symmetric
// difference of the given arrays. The order of result values
// is determined by the order they occur in the arrays.
func Xor[T comparable](arrays ...[]T) []T {
	exclusive := map[T]int{}
	// collect exclusive elements in map
	for _, arr := range arrays {
		for _, v := range arr {
			exclusive[v]++
		}
	}
	result := make([]T, 0)
	// loop again to collect the exclusive elements in order
	for _, arr := range arrays {
		for _, v := range arr {
			count := exclusive[v]
			if count == 1 {
				result = append(result, v)
			}
		}
	}

	return result
}
