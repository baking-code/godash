package findindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	var input = []string{"a", "b", "b"}
	got := FindIndex[string](input, func(str string) bool {
		return str == "b"
	})
	var expected = 1
	assert.EqualValues(t, expected, got)
}
func TestFindIndexNone(t *testing.T) {
	var input = []string{"a", "b", "b"}
	got := FindIndex[string](input, func(str string) bool {
		return str == "c"
	})
	var expected = -1
	assert.EqualValues(t, expected, got)
}

func TestLastIndex(t *testing.T) {
	var input = []string{"a", "b", "b"}
	got := FindLastIndex[string](input, func(str string) bool {
		return str == "b"
	})
	var expected = 2
	assert.EqualValues(t, expected, got)
}

func TestLastIndexNone(t *testing.T) {
	var input = []string{"a", "b", "b"}
	got := FindLastIndex[string](input, func(str string) bool {
		return str == "c"
	})
	var expected = -1
	assert.EqualValues(t, expected, got)
}
