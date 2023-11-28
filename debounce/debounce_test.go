package debounce

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDebounce(t *testing.T) {
	var now = time.Now()
	i := 0
	var input = func(args ...int) {
		i++
		fmt.Printf("Printing after %d millis\n", time.Since(now).Milliseconds())
	}
	dbncd, _, _ := Debounce[int](input, 100, Options{leading: false, trailing: true})
	dbncd()
	dbncd()
	time.Sleep(time.Duration(200) * time.Millisecond)
	assert.EqualValues(t, 1, i)
}

func TestDebounceLast(t *testing.T) {
	var now = time.Now()
	i := 0
	var input = func(args ...int) {
		i = args[0]
		fmt.Printf("Printing %v after %d millis\n", args, time.Since(now).Milliseconds())
	}
	dbncd, _, _ := Debounce[int](input, 100, Options{leading: false, trailing: true})
	dbncd(54)
	dbncd(79)
	time.Sleep(time.Duration(200) * time.Millisecond)
	assert.EqualValues(t, 79, i)
}

func TestDebounceCancel(t *testing.T) {
	var now = time.Now()
	i := 0
	var input = func(args ...int) {
		i = args[0]
		fmt.Printf("Printing %v after %d millis\n", args, time.Since(now).Milliseconds())
	}
	dbncd, cancel, _ := Debounce[int](input, 100, Options{leading: false, trailing: true})
	dbncd(54)
	dbncd(79)
	cancel()
	time.Sleep(time.Duration(200) * time.Millisecond)
	assert.EqualValues(t, 0, i)
}

func TestDebounceFlush(t *testing.T) {
	var now = time.Now()
	i := 0
	var input = func(args ...int) {
		i = args[0]
		fmt.Printf("Printing %v after %d millis\n", args, time.Since(now).Milliseconds())
	}
	dbncd, _, flush := Debounce[int](input, 100, Options{leading: false, trailing: true})
	dbncd(54)
	dbncd(79)
	flush()
	assert.EqualValues(t, 79, i)
}

func TestDebounceLeading(t *testing.T) {
	var now = time.Now()
	i := 0
	var input = func(args ...int) {
		i = args[0]
		fmt.Printf("Printing %v after %d millis\n", args, time.Since(now).Milliseconds())
	}
	dbncd, _, _ := Debounce[int](input, 100, Options{leading: true, trailing: false})
	dbncd(54)
	dbncd(79)
	time.Sleep(time.Duration(50) * time.Millisecond)
	assert.EqualValues(t, 79, i)
}
