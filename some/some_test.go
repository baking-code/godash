package some

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	type User struct {
		user   string
		active bool
	}
	input := []User{{user: "barney", active: true}, {user: "fred", active: false}}
	one := Some[User](input, func(e User, index int, collection []User) bool {
		return e.user == "barney" && e.active == false
	})
	assert.Equal(t, false, one)

	two := Some[User](input, func(e User, index int, collection []User) bool {
		return e.active == false
	})
	assert.Equal(t, true, two)

	count := 0
	three := Some[User](input, func(e User, index int, collection []User) bool {
		count++
		return e.active == true
	})
	assert.Equal(t, true, three)
	assert.Equal(t, 1, count)
}
