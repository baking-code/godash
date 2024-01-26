package some

// Some checks if predicate returns truthy for any element of collection. Iteration is stopped once predicate returns truthy. The predicate is invoked with three arguments: (value, index|key, collection).
func Some[T interface{}](collection []T, predicate func(e T, index int, collection []T) bool) bool {
	for i, element := range collection {
		if predicate(element, i, collection) {
			return true
		}
	}
	return false
}
