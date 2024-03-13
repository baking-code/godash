package xor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {
	got := Xor([]int{2, 1}, []int{2, 3})
	expected := []int{1, 3}
	assert.EqualValues(t, expected, got)

	got = Xor([]int{2, 1}, []int{2, 3}, []int{3, 4, 5})
	expected = []int{1, 4, 5}
	assert.EqualValues(t, expected, got)

	got = Xor([]int{2, 1}, []int{2, 1, 4}, []int{1, 4, 2})
	expected = []int{}
	assert.EqualValues(t, expected, got)
}
