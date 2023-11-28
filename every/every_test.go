package every

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	type User struct {
		user   string
		active bool
	}
	input := []User{{user: "barney", active: true}, {user: "fred", active: true}}
	one := Every[User](input, func(e User, index int, collection []User) bool {
		return e.user == "barney" && e.active == false
	})
	assert.Equal(t, false, one)

	count := 0
	two := Every[User](input, func(e User, index int, collection []User) bool {
		count++
		return e.active == false
	})
	assert.Equal(t, false, two)
	assert.Equal(t, 1, count)

	count = 0
	three := Every[User](input, func(e User, index int, collection []User) bool {
		count++
		return e.active == true
	})
	assert.Equal(t, true, three)
	assert.Equal(t, 2, count)
}
