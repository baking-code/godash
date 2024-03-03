package memoize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicMemoize(t *testing.T) {
	var i int = 0
	var theFunc = func(args ...int) any {
		a, b := args[0], args[1]
		i++
		return a * b
	}
	memed := BasicMemoize(theFunc)
	memed(1, 2)
	ans1 := memed(1, 2)
	ans2 := memed(3, 4)
	assert.EqualValues(t, 2, i)
	assert.EqualValues(t, 2, ans1)
	assert.EqualValues(t, 12, ans2)
}
func TestMemoize(t *testing.T) {
	var i int = 0
	var theFunc = func(a, b int) int {
		i++
		return a * b
	}
	memed, _ := Memoize(theFunc)
	memed(1, 2)
	ans1 := memed(1, 2)
	ans2 := memed(3, 4)
	assert.EqualValues(t, 2, i)
	assert.EqualValues(t, 2, ans1)
	assert.EqualValues(t, 12, ans2)
}

func BenchmarkBasicMemoize(b *testing.B) {
	var theFunc = func(args ...int) any {
		a, b := args[0], args[1]
		return a * b
	}
	for i := 0; i < b.N; i++ {
		memed := BasicMemoize(theFunc)
		memed(i, i+1)
	}
}

func BenchmarkMemoize(b *testing.B) {
	var theFunc = func(a, b int) int {
		return a * b
	}
	for i := 0; i < b.N; i++ {
		memed, _ := Memoize(theFunc)
		memed(i, i+1)
	}
}
