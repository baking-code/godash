package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type User struct {
		user   string
		active bool
	}
	input := []User{{user: "barney", active: true}, {user: "fred", active: true}}
	one := Filter[User](input, func(e User, index int, collection []User) bool {
		return e.user == "barney" && e.active == false
	})
	assert.EqualValues(t, []User{}, one)

	two := Filter[User](input, func(e User, index int, collection []User) bool {
		return e.user == "barney"
	})
	assert.EqualValues(t, []User{{user: "barney", active: true}}, two)

	three := Filter[User](input, func(e User, index int, collection []User) bool {
		return e.active == true
	})
	assert.EqualValues(t, input, three)
}
