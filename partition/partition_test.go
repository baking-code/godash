package partition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	type User struct {
		user   string
		active bool
	}
	input := []User{{user: "barney", active: true}, {user: "fred", active: true}}
	one, two := Partition[User](input, func(e User, index int, collection []User) bool {
		return e.user == "barney" && e.active == false
	})
	assert.EqualValues(t, []User{}, one)
	assert.EqualValues(t, input, two)
}
