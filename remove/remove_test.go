package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	var input = []string{"a", "b", "c", "d"}
	got := Remove[string](&input, func(o string, i int, orig []string) bool {
		return o == "b"
	})
	var expected = []string{"b"}

	assert.EqualValues(t, input, []string{"a", "c", "d"})
	assert.EqualValues(t, expected, got)
}
