package xslices

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

func ExampleMap() {
	input := []string{"a", "b", "c"}

	res := Map[[]string, []int](input, func(s string) int {
		return len(s) + 1
	})
	fmt.Println(res)

	// Output:
	// [2 2 2]
}

func ExampleFindFirst() {
	input := []string{"a", "b", "c"}

	first, ok := FindFirst(input, func(s string) bool {
		return s == "a"
	})
	fmt.Println(first, ok)

	// Output:
	// a true
}

func ExampleAssociate() {
	input := []string{"a", "bb", "ccc"}

	res := Associate(input, func(item string) (string, int) {
		return item, len(item) + 1
	})
	fmt.Println(res)

	// Output:
	// map[a:2 bb:3 ccc:4]
}

func ExampleDifference() {
	input1 := []string{"a", "bb", "ddd", "fff"}
	input2 := []string{"a", "bb", "ccc"}

	difference := Difference(input1, input2)
	fmt.Println(difference)

	// Output:
	// [ddd fff]
}

func ExampleFilter() {
	input := []string{"a", "b", "c"}

	res := Filter(input, func(s string) bool {
		return s == "a" || s == "b"
	})
	fmt.Println(res)

	// Output:
	// [a b]
}

func ExampleFirst() {
	input1 := []string{"a", "b", "c"}
	first, ok := First(input1)
	fmt.Println(first, ok)

	var input2 []string
	first, ok = First(input2)
	fmt.Println(first, ok)

	// Output:
	// a true
	//  false
}

func ExampleFlatten() {
	input := [][]string{{"a", "aa"}, {"b", "bb"}, {"c", "cc"}}

	res := Flatten(input)
	fmt.Println(res)

	// Output:
	// [a aa b bb c cc]
}

func ExampleToSetFunc() {
	input := []string{"a", "b", "c"}

	set := ToSetFunc(input, func(item string) int {
		return len(item)
	})
	fmt.Println(set)

	// Output:
	// map[1:{}]
}

func ExampleToSet() {
	input := []string{"a", "b", "c", "a", "b"}

	set := ToSet(input)
	fmt.Println(set)

	// Output:
	// map[a:{} b:{} c:{}]
}

func ExampleMapParallel() {
	input := []string{"a", "b", "c", "a", "b"}

	res := MapParallel[[]string, []int](2, input, func(i string) int {
		return len(i)
	})
	fmt.Println(res)

	// Output:
	// [1 1 1 1 1]

}

func ExampleMapParallelWithContext() {
	input := []string{"a", "b", "c", "a", "b"}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	res := MapParallelWithContext[[]string, []int](ctx, 2, input, func(i string) int {
		return len(i)
	})
	fmt.Println(res)

	// Output:
	// [1 1 1 1 1]
}

func ExampleMapUntilError() {
	type A string
	s := []A{"a", "bb", "", "ccc"}

	res, err := MapUntilError[[]A, []int](s, func(t A) (int, error) {
		if t > "" {
			return len(t), nil
		}
		return 0, errors.New("ups")
	})
	fmt.Println(res, err)

	// Output:
	// [1 2] ups
}

func ExampleMapWithError() {
	type A string
	s := []A{"a", "bb", "", "ccc"}

	res, err := MapWithError[[]A, []string](s, func(a A) (string, error) {
		if a > "" {
			return strings.ToUpper(string(a)), nil
		}
		return "", errors.New("empty string")
	}, true)

	fmt.Println(res, err)

	// Output:
	// [A BB] empty string
}
