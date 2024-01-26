package partition

// Partition creates a slice of elements split into two groups, _left_ which contains elements predicate returns truthy for,
// _right_ which contains elements predicate returns falsey for. The predicate is invoked with three arguments: (value, index, collection).
//
// Note this returns new slices.
func Partition[T interface{}](collection []T, predicate func(e T, index int, collection []T) bool) ([]T, []T) {
	left := make([]T, 0)
	right := make([]T, 0)
	for i, element := range collection {
		if predicate(element, i, collection) {
			left = append(left, element)
		} else {
			right = append(right, element)

		}
	}
	return left, right
}
