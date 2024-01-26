package concat

// Concat creates a new slice concatenating each input slice
func Concat[T any](inputs ...[]T) []T {
	collected := make([]T, 0)
	for _, input := range inputs {
		collected = append(collected, input...)
	}
	return collected
}
