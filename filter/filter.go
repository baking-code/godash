package filter

import (
	"godash/partition"
)

/*
Iterates over elements of collection, returning an array of all elements predicate returns truthy for. The predicate is invoked with three arguments: (value, index|key, collection).

Note: this method returns a new array.
*/
func Filter[T interface{}](collection []T, predicate func(e T, index int, collection []T) bool) []T {
	left, _ := partition.Partition(collection, predicate)
	return left
}
