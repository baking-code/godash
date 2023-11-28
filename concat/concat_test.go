package concat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	var input = []string{"a", "b"}

	var compare = []string{"c", "d"}

	got := Concat[string](input, compare)
	var expected = []string{"a", "b", "c", "d"}
	assert.EqualValues(t, expected, got)
}
