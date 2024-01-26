package memoize

import "fmt"

type cacheKey interface {
	comparable
}

// Memoize creates a function that memoizes the result of func. If resolver is provided, it determines
// the cache key for storing the result based on the arguments provided to the memoized function.
// By default, the first argument provided to the memoized function is used as the map cache key.
//
// Note: The cache is exposed as the cache property on the memoized function. Its creation may be customized by replacing the _.memoize.Cache constructor with one whose instances implement the Map method interface of clear, delete, get, has, and set.
func Memoize[T cacheKey, V any](fn func(args ...T) V) func(args ...T) V {
	cache := map[string]V{}
	var params []T
	wrap := func(p ...T) V {
		params = p
		s := ""
		for _, arg := range params {
			s += fmt.Sprint(arg)
		}
		v, ok := cache[s]
		if ok {
			return v
		} else {
			u := fn(params...)
			cache[s] = u
			return u
		}
	}
	return wrap
}
