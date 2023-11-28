package fill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	var input = []string{"a", "b", "c"}

	got := Fill[string](input, "z")
	var expected = []string{"z", "z", "z"}
	assert.EqualValues(t, expected, got)
}

func TestFillStart(t *testing.T) {
	var input = []string{"a", "b", "c"}

	got := Fill[string](input, "z", 1)
	var expected = []string{"a", "z", "z"}
	assert.EqualValues(t, expected, got)
}

func TestFillEnd(t *testing.T) {
	var input = []string{"a", "b", "c"}

	got := Fill[string](input, "z", 0, 1)
	var expected = []string{"z", "b", "c"}
	assert.EqualValues(t, expected, got)
}
