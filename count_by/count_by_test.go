package count_by

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBy(t *testing.T) {
	var input = []float64{6.1, 4.2, 6.3}
	got := CountBy[float64](input, func(e float64) string {
		return fmt.Sprintf("%d", int(math.Floor(e)))
	})
	var expected = map[string]int{
		"4": 1,
		"6": 2,
	}
	assert.EqualValues(t, expected, got)
}
