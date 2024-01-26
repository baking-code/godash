package fill

import "math"

// Fills elements of input with value from opts[0] up to, but not including, opts[1].
// (Defaults: opts[0] = 0, opts[1] = len(input)])
// Note: This method mutates input.
func Fill[T any](input []T, filler T, opts ...int) []T {
	start := 0
	end := len(input)
	if len(opts) == 1 {
		start = opts[0]
	} else if len(opts) == 2 {
		start = opts[0]
		end = int(math.Min(float64(opts[1]), float64(len(input))))
	}

	slice := input[start:end]
	for i := range slice {
		slice[i] = filler
	}
	return input
}
