package reduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	var input = []int{1, 2}
	got := Reduce[int, int](input, func(acc int, i int, j int, input []int) int {
		return acc + i
	}, 0)
	var expected = 3
	assert.Equal(t, expected, got)
}
