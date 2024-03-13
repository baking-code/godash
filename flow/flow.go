package flow

// Flow creates a function that returns the result of invoking the given functions, where each successive
// invocation is supplied the return value of the previous.
func Flow[T any](inputFunctions []func(T) T) func(arg T) T {
	wrapper := func(input T) T {
		var output T = input
		for _, fn := range inputFunctions {
			output = fn(output)
		}
		return output
	}
	return wrapper
}
