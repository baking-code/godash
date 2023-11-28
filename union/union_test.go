package union

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	var input = []string{"a", "b"}
	var input2 = []string{"b", "c"}
	got := Union[string](input, input2)
	var expected = []string{"a", "b", "c"}
	assert.EqualValues(t, expected, got)
}
