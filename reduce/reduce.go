package reduce

// Reduces collection to a value which is the accumulated result of running each element in collection thru iteratee, where each successive invocation is supplied the return value of the previous. The iteratee is invoked with four arguments:
// (accumulator, value, index|key, collection).
func Reduce[T interface{}, V interface{}](collection []T, iteratee func(accumulator V, e T, index int, collection []T) V, accumulator V) V {
	for i, element := range collection {
		accumulator = iteratee(accumulator, element, i, collection)
	}
	return accumulator
}
