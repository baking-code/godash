package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {
	var input = []func(int) int{
		func(a int) int {
			return a + 1
		},
		func(a int) int {
			return a + 2
		},
		func(int) int {
			return +3
		},
	}
	funcs := Flow[int](input)
	result := funcs(2)
	assert.EqualValues(t, 8, result)
}
