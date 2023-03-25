// Package xslices defines various utility functions useful with slices of any type.

package xslices

import (
	"context"
	"errors"
	"sync"
)

// Map applies the given function mapper to each element of slice s and returns a new slice O with the mapped elements.
func Map[S ~[]T, O ~[]R, T, R any](s S, mapper func(T) R) O {
	out := make(O, 0, len(s))
	for _, t := range s {
		out = append(out, mapper(t))
	}
	return out
}

// MapUntilError applies the given function mapper to each element of slice s and
// returns a new slice with the mapped elements.
//
// If the error occurs during mapping, function will short-circuit and return already
// mapped elements and error.
func MapUntilError[S ~[]T, O ~[]R, T, R any](s S, mapper func(T) (R, error)) (O, error) {
	return MapWithError[S, O](s, mapper, true)
}

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Min returns the smallest element in a slice.
// If slice is empty it returns zero value and false.
func Min[T ordered](s []T) (zero T, ok bool) {
	curr, ok := First(s)
	if !ok {
		return
	}

	for _, v := range s[1:] {
		if v < curr {
			curr = v
		}
	}
	return curr, true
}

// Max returns the largest element in a slice.
// If slice is empty it returns zero value and false.
func Max[T ordered](s []T) (zero T, ok bool) {
	curr, ok := First(s)
	if !ok {
		return
	}

	for _, v := range s[1:] {
		if v > curr {
			curr = v
		}
	}
	return curr, true
}

// MapWithError applies the given function mapper to each element of slice s and
// returns a new slice with the mapped elements. Function takes in additional bool parameter ff (fail-fast).
//
// If ff is set to true, in that case the mapping will stop after first error.
//
// If ff is set to false, mapping will continue until all items have been mapped or yielded error, errors
// from mapping function are joined using errors.Join function.
func MapWithError[S ~[]T, O ~[]R, T, R any](s S, mapper func(T) (R, error), ff bool) (O, error) {
	out := make(O, 0, len(s))
	var e error
	for _, t := range s {
		r, err := mapper(t)
		if err != nil {
			if ff {
				return out, err
			}
			e = errors.Join(e, err)
			continue
		}
		out = append(out, r)
	}
	return out, e
}

// Flatten takes a slice of slices s and returns a new slice that contains all the elements from the nested slices.
func Flatten[S ~[]T, T any](s []S) []T {
	target := make(S, 0)
	for _, ts := range s {
		target = append(target, ts...)
	}
	return target
}

// ToSet takes a slice s and returns a set in the form of a map where the keys are the elements
// in s and the values are empty structs
func ToSet[T comparable](s []T) map[T]struct{} {
	return ToSetFunc(s, func(t T) T { return t })
}

// ToSetFunc takes a slice s and a mapping function mapper that maps the elements in s to keys of a set.
// It returns a set in the form of a map where the keys are the mapped values and the values are empty structs.
func ToSetFunc[T, K comparable](s []T, mapper func(item T) K) map[K]struct{} {
	return Associate(s, func(item T) (K, struct{}) {
		return mapper(item), struct{}{}
	})
}

// Associate returns a map containing key-value pairs provided by mapper function
// applied to elements of the given collection.
//
// If any of two pairs would have the same key the last one gets added to the map.
//
// The returned map preserves the entry iteration order of the original collection.
func Associate[K comparable, T, V any](s []T, mapper func(item T) (K, V)) map[K]V {
	res := make(map[K]V, len(s))
	for _, item := range s {
		k, v := mapper(item)
		res[k] = v
	}
	return res
}

// First if provided slice has any elements the function will return first element and true.
// If the provided slice is empty in that case the function will return zero value of [E] and false
func First[E any](s []E) (zero E, ok bool) {
	for _, i := range s {
		return i, true
	}
	return
}

// Difference returns a new slice of unique [T] items that are present in s1 but not in s2
func Difference[S ~[]T, T comparable](s1, s2 S) S {
	set := make(map[T]struct{}, len(s1))
	for _, t := range s2 {
		set[t] = struct{}{}
	}
	var diff []T
	for _, x := range s1 {
		if _, found := set[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// Filter returns a new slice of [T] items that match the supplied predicate.
func Filter[T any](input []T, pred func(T) bool) []T {
	res := make([]T, 0, len(input))
	for _, item := range input {
		if pred(item) {
			res = append(res, item)
		}
	}
	return res
}

// FindFirst finds the first [T] element in the slice that matches the predicate, or
// it returns zero value and false.
func FindFirst[T any](input []T, pred func(T) bool) (zero T, ok bool) {
	for _, item := range input {
		if pred(item) {
			return item, true
		}
	}
	return
}

// MapParallelWithContext applies the given function `f` to each element in the input
// slice inputs  in parallel, using the n number of goroutines.
//
// A negative n value indicates that n will be set to inputs items count.
//
// See also MapParallel.
func MapParallelWithContext[S ~[]T, O ~[]R, T, R any](ctx context.Context, n int, inputs S, f func(T) R) O {
	// create a channel for final results
	out := make(chan R, len(inputs))

	// creates a channel to send inputs
	in := make(chan T, len(inputs))

	// set number of workers (n)
	// if the n < 0 or n > len(inputs) we set n to items count
	if max := len(inputs); n < 0 || n > max {
		n = max
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ { // spawn the worker goroutines.
		wg.Add(1)
		go func() {
			defer wg.Done()
			for input := range in {
				select {
				case <-ctx.Done():
					return
				case out <- f(input):
				}
			}
		}()
	}

	for _, input := range inputs {
		in <- input // send inputs to the worker goroutines.
	}
	close(in) // no more inputs to be expected

	wg.Wait() // wait for goroutines to finish work

	close(out) // close out channel since goroutines are done
	results := make([]R, len(inputs))
	for i := 0; i < len(inputs); i++ {
		results[i] = <-out
	}
	return results
}

// MapParallel applies the given function `f` to each element in the input
// slice `inputs` in parallel, using the  n number of goroutines. It returns a slice of results,
// one for each input element.
//
// See also MapParallelWithContext.
func MapParallel[S ~[]T, O ~[]R, T, R any](n int, inputs S, f func(T) R) O {
	return MapParallelWithContext[S, O](context.Background(), n, inputs, f)
}
