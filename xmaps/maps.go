// Package xmaps defines various utility functions useful with maps of any type.

package xmaps

// Merge merges provided maps into a single map. If duplicates occur, values will be overwritten.
func Merge[M ~map[K]V, K comparable, V any](maps ...M) M {
	res := make(M, 0)
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// Equal determines if two maps are equal to each other. If they have the same cardinality
// and contain the same elements, they are considered equal.
// The order in which the elements were added is irrelevant.
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

// Intersect returns a new map containing only the key-value pairs that are present and equal in both maps,
func Intersect[M ~map[K]V, K, V comparable](m1, m2 M) M {
	res := make(M, 0)
	for k, v1 := range m1 {
		if v2, ok := m2[k]; ok && v1 == v2 {
			res[k] = v1
		}
	}
	return res
}

// Difference returns a new map containing the difference between two provided maps.
func Difference[M ~map[K]V, K, V comparable](m1, m2 M) M {
	diff := make(map[K]V)
	for k, v := range m1 {
		if e, ok := m2[k]; ok {
			if e != v {
				diff[k] = e
			}
		} else {
			diff[k] = v
		}
	}
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			diff[k] = v
		}
	}
	return diff
}
