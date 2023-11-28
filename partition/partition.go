package partition

/*
*
 */
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
