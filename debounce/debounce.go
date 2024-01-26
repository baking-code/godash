package debounce

import (
	"fmt"
	"time"
)

type Options struct {
	leading  bool // default false
	trailing bool // default true
}

// Debounce creates a debounced function that delays invoking `fn` until after wait milliseconds have elapsed since the last time the debounced function was invoked.
// The debounced function comes with a cancel method to cancel delayed `fn` invocations and a flush method to immediately invoke them.
// Provide options to indicate whether `fn` should be invoked on the leading and/or trailing edge of the wait timeout.
// The `fn` is invoked with the last arguments provided to the debounced function.
// Subsequent calls to the debounced function return the result of the last func invocation.
func Debounce[T interface{}](fn func(args ...T), wait int, options Options) (func(args ...T), func(), func()) {
	invoked := false
	cancelled := false
	var params []T

	var leading, trailing bool
	if options.leading {
		leading = true
	}
	if !options.trailing {
		trailing = false
	} else {
		trailing = true
	}

	f := func() {
		if leading && !cancelled {
			fn(params...)
		}
		time.Sleep(time.Duration(wait) * time.Millisecond)
		if trailing && !cancelled {
			fn(params...)
		}
		invoked = false
	}

	wrap := func(p ...T) {
		params = p
		fmt.Println(p)
		if !invoked {
			invoked = true
			fmt.Println("go", p)
			go f()
		}
		return
	}

	cancel := func() {
		cancelled = true
	}
	flush := func() {
		fn(params...)
		cancel()
	}
	return wrap, cancel, flush
}
