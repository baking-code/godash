package concat

/*
*
Creates a new slice concatenating each input slice
*/
func Concat[T interface{}](inputs ...[]T) []T {
	collected := make([]T, 0)
	for _, input := range inputs {
		for _, element := range input {
			collected = append(collected, element)
		}
	}
	return collected
}
