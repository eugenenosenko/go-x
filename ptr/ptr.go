// Package ptr provides utilities for converting literal type values to/from pointers inline.

package ptr

// To returns a pointer to val
// this func should substitute aws.* functions and entire "pointer" package
// Should be used only for expression when "&" operator is not suitable
func To[T any](val T) *T {
	return &val
}

// Value returns the value of the pointer passed in or type's zero value.
func Value[V any](v *V) (value V) {
	if v != nil {
		value = *v
	}
	return value
}
