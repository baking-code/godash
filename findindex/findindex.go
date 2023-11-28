package findindex

func iterate[T interface{}](input []T, predicate func(o T) bool, returnEarly bool) []int {
	collect := make([]int, 0)
	for i, element := range input {
		if predicate(element) {
			collect = append(collect, i)
			if returnEarly {
				break
			}
		}
	}
	return collect
}

/*
* Find first index of element in input slice which satisfies predicate function
 */
func FindIndex[T interface{}](input []T, predicate func(o T) bool) int {
	collect := iterate[T](input, predicate, true)
	if len(collect) > 0 {
		return collect[0]
	}
	return -1
}

/*
* Find last index of element in input slice which satisfies predicate function
 */
func FindLastIndex[T interface{}](input []T, predicate func(o T) bool) int {
	collect := iterate[T](input, predicate, false)
	if len(collect) > 0 {
		return collect[len(collect)-1]
	}
	return -1
}
