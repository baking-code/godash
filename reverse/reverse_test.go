package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	var input = []string{"a", "b", "c"}
	got := Reverse[string](&input)
	var expected = []string{"c", "b", "a"}
	assert.EqualValues(t, expected, got)
	assert.EqualValues(t, input, got)
}
