package every

/*
Checks if predicate returns truthy for all elements of collection. Iteration is stopped once predicate returns falsey. The predicate is invoked with three arguments: (value, index|key, collection).

Note: This method returns true for empty collections because everything is true of elements of empty collections.
*/
func Every[T interface{}](collection []T, predicate func(e T, index int, collection []T) bool) bool {
	for i, element := range collection {
		if !predicate(element, i, collection) {
			return false
		}
	}
	return true
}
