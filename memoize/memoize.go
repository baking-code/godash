package memoize

import (
	"errors"
	"fmt"
	"reflect"
)

type cacheKey interface {
	comparable
}

// BasicMemoize creates a function that memoizes the result of func. If resolver is provided, it determines
// the cache key for storing the result based on the arguments provided to the memoized function.
// By default, the first argument provided to the memoized function is used as the map cache key.
//
// Uses generics, so a limitation of fn is that it's arguments are of the same type
func BasicMemoize[T cacheKey](fn func(args ...T) any) func(args ...T) any {
	cache := map[string]any{}
	var params []T
	wrap := func(p ...T) any {
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

// Memoize creates a function that memoizes the result of func. If resolver is provided, it determines
// the cache key for storing the result based on the arguments provided to the memoized function.
// By default, the first argument provided to the memoized function is used as the map cache key.
//
// This uses reflection so is relatively slow - only memoize stateless, long running/expensive functions
func Memoize[T any](fn T) (T, error) {
	typeOfFn := reflect.TypeOf(fn)
	if typeOfFn.Kind() != reflect.Func {
		return fn, errors.New("provided argument not a function")
	}
	valueOfFn := reflect.ValueOf(fn)
	cache := map[string][]reflect.Value{}
	wrapperF := reflect.MakeFunc(typeOfFn, func(params []reflect.Value) []reflect.Value {
		s := ""
		for _, arg := range params {
			s += fmt.Sprint(arg)
		}
		v, ok := cache[s]
		if ok {
			return v
		} else {
			u := valueOfFn.Call(params)
			cache[s] = u
			return u
		}
	})
	return wrapperF.Interface().(T), nil
}
