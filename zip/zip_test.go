package zip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	var input = []string{"a", "b"}
	var input2 = []string{"c", "d"}
	got := Zip[string](input, input2)
	var expected = [][]string{{"a", "c"}, {"b", "d"}}
	assert.EqualValues(t, expected, got)
}

func TestUnzip(t *testing.T) {
	var input = [][]string{{"a", "c"}, {"b", "d"}}
	got := Unzip[string](input)
	var expected = [][]string{{"a", "b"}, {"c", "d"}}
	assert.EqualValues(t, expected, got)
}

func TestBoth(t *testing.T) {
	var input = []string{"a", "b"}
	var input2 = []string{"c", "d"}
	got := Zip[string](input, input2)
	var actual = Unzip[string](got)
	assert.EqualValues(t, [][]string{{"a", "b"}, {"c", "d"}}, actual)
}
