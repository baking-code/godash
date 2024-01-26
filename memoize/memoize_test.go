package memoize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoize(t *testing.T) {
	var i int = 0
	var theFunc = func(args ...int) int {
		a, b := args[0], args[1]
		i++
		return a * b
	}
	memed := Memoize[int, int](theFunc)
	memed(1, 2)
	ans1 := memed(1, 2)
	ans2 := memed(3, 4)
	assert.EqualValues(t, 2, i)
	assert.EqualValues(t, 2, ans1)
	assert.EqualValues(t, 12, ans2)
}
