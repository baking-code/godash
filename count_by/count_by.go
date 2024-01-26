package count_by

// CountBy creates a struct composed of keys generated from the results of running each element of collection thru iteratee. The corresponding value of each key is the number of times the key was returned by iteratee. The iteratee is invoked with one argument: (value).
func CountBy[T any](collection []T, iteratee func(e T) string) map[string]int {
	result := make(map[string]int)
	for _, element := range collection {
		value := iteratee(element)
		result[value]++
	}
	return result
}
