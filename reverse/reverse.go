package reverse

// Reverses input so that the first element becomes the last, the second element becomes the second to last, and so on.
//
// Note: Mutates input
func Reverse[T interface{}](input *[]T) []T {
	temp := make([]T, len(*input))
	for index := range *input {
		temp[index] = (*input)[len(*input)-(index+1)]
	}
	*input = temp
	return *input
}
