package xmaps

import (
	"fmt"
)

func ExampleDifference() {
	m1 := map[string]string{"a": "1", "b": "2"}
	m2 := map[string]string{"b": "2", "c": "3", "d": "4"}

	res1 := Difference(m1, m2)
	res2 := Difference(m2, m1)
	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// map[a:1 c:3 d:4]
	// map[a:1 c:3 d:4]
}

func ExampleEqual() {
	m1 := map[string]string{"a": "1", "b": "2"}
	m2 := map[string]string{"b": "2", "c": "3", "d": "4"}

	eq := Equal(m1, m2)
	fmt.Println(eq)

	m1 = map[string]string{"a": "1"}
	m2 = map[string]string{"a": "1"}

	eq = Equal(m1, m2)
	fmt.Println(eq)

	m1 = map[string]string{}
	m2 = map[string]string{}

	eq = Equal(m1, m2)
	fmt.Println(eq)

	// Output:
	// false
	// true
	// true
}

func ExampleIntersect() {
	m1 := map[string]string{"a": "1", "b": "2"}
	m2 := map[string]string{"b": "2", "c": "3", "d": "4"}

	res := Intersect(m1, m2)
	fmt.Println(res)

	// Output:
	// map[b:2]
}

func ExampleMerge() {
	m1 := map[string]string{"a": "1", "b": "2222"}
	m2 := map[string]string{"b": "2", "c": "3", "d": "4"}
	m3 := map[string]string{"b": "2", "c": "3", "d": "4", "f": "7"}

	res := Merge(m1, m2, m3)
	fmt.Println(res)

	// Output:
	// map[a:1 b:2 c:3 d:4 f:7]
}
