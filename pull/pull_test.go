package pull

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPull(t *testing.T) {
	var input = []string{"a", "b", "c", "d"}
	got := Pull[string](&input, "a", "b")
	var expected = []string{"c", "d"}

	assert.EqualValues(t, expected, input)
	assert.EqualValues(t, expected, got)
}

func TestPullTwo(t *testing.T) {
	var input = []string{"a", "b", "c", "d"}
	got := Pull[string](&input, "c", "b")
	var expected = []string{"a", "d"}

	assert.EqualValues(t, expected, input)
	assert.EqualValues(t, expected, got)
}
