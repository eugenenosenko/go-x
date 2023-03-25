package xslices

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	type typ struct {
		Value string
	}
	type TypList []*typ

	want := TypList{{Value: "a"}, {Value: "b"}, {Value: "c"}}
	if got := Map[[]string, TypList]([]string{"a", "b", "c"}, func(s string) *typ {
		return &typ{Value: s}
	}); !reflect.DeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapWithError(t *testing.T) {
	t.Run("should map correctly and collect all items and combined error", func(t *testing.T) {
		failfast := false
		got, err := MapWithError[[]string, []int]([]string{"aaa", "bbb", "cc", "dd", "ee"}, func(t string) (int, error) {
			if len(t) < 3 {
				return 0, fmt.Errorf("%s is shorter than expected length of 3", t)
			}
			return len(t), nil
		}, failfast)

		want := []int{3, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapWithError() = got %v, want %v", got, want)
		}

		msg := `cc is shorter than expected length of 3
dd is shorter than expected length of 3
ee is shorter than expected length of 3`
		if got := err.Error(); got != msg {
			t.Errorf("MapWithError() = got error %v, want error %v", got, msg)
		}

	})
	t.Run("should map correctly if no errors occurred", func(t *testing.T) {
		got, err := MapWithError[[]string, []int]([]string{"aaa", "bbb"}, func(t string) (int, error) {
			if len(t) < 3 {
				return 0, fmt.Errorf("%s is shorter than expected length of 3", t)
			}
			return len(t), nil
		}, false)
		if err != nil {
			t.Errorf("MapWithError() = got error %v, want error %v ", got, nil)
		}

		want := []int{3, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapWithError() = got %v, want %v", got, want)
		}
	})
	t.Run("should map and fail fast on first error", func(t *testing.T) {
		failfast := true
		got, err := MapWithError[[]string, []int]([]string{"aaa", "cc", "bbb", "dd", "ee"}, func(t string) (int, error) {
			if len(t) < 3 {
				return 0, fmt.Errorf("%s is shorter than expected length of 3", t)
			}
			return len(t), nil
		}, failfast)

		want := []int{3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapWithError() = got %v, want %v", got, want)
		}
		msg := "cc is shorter than expected length of 3"
		if gotErr := err.Error(); gotErr != msg {
			t.Errorf("MapWithError() = got error %v, want error %v ", gotErr, msg)
		}
	})

}

func TestFlatten(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 5}
	if got := Flatten([][]int{{1}, {2}, {3}, {4}, {5}, {5}}); !reflect.DeepEqual(got, want) {
		t.Errorf("Flatten() = %v, want %v", got, want)
	}
}

func TestToSet(t *testing.T) {
	want := map[string]struct{}{"a": {}, "b": {}, "c": {}}
	if got := ToSet([]string{"a", "b", "c", "c", "b", "a"}); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSet() = %v, want %v", got, want)
	}
}

func TestToSetFunc(t *testing.T) {
	want := map[string]struct{}{"A": {}, "B": {}, "C": {}}
	if got := ToSetFunc([]string{"a", "b", "c", "c", "b", "a"}, strings.ToUpper); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSetFunc() = %v, want %v", got, want)
	}
}

func TestAssociate(t *testing.T) {
	want := map[string]string{"A": "a", "B": "b", "C": "c"}
	if got := Associate(
		[]string{"a", "b", "c", "c", "b", "a"}, func(item string) (string, string) {
			return strings.ToUpper(item), item
		},
	); !reflect.DeepEqual(got, want) {
		t.Errorf("Associate() = %v, want %v", got, want)
	}
}

func TestFirst(t *testing.T) {
	t.Run("should return first element of the list", func(t *testing.T) {
		got, ok := First([]string{"a", "b", "c", "c", "b", "a"})
		if !ok {
			t.Error("First() got = false, want true")
		}
		want := "a"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("First() got = %v, want %v", got, want)
		}
	})

	t.Run("should return zero value and false for an empty list", func(t *testing.T) {
		got, ok := First([]string{})
		if ok {
			t.Error("First() got = true, want false")
		}
		want := ""
		if !reflect.DeepEqual(got, want) {
			t.Errorf("First() got = %v, want %v", got, want)
		}
	})
}

func TestDifference(t *testing.T) {
	want := []string{"a", "b"}
	got := Difference([]string{"a", "b", "f", "d"}, []string{"f", "d", "c"})
	sort.Strings(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	t.Run("should correctly filter out items", func(t *testing.T) {
		want := []string{"abc"}
		if got := Filter([]string{"abc", "de", "f", ""}, func(s string) bool {
			return len(s) > 2
		}); !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})

	t.Run("should return empty slice if none of the items match predicate", func(t *testing.T) {
		want := make([]string, 0, 4)
		if got := Filter([]string{"abc", "de", "f", ""}, func(s string) bool {
			return len(s) > 3
		}); !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})
}

func TestFindFirst(t *testing.T) {
	t.Run("should correctly return matched item", func(t *testing.T) {
		want := "b"
		got, got1 := FindFirst([]string{"a", "b", "c"}, func(s string) bool {
			return s == "b"
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindFirst() got = %v, want %v", got, want)
		}
		if got1 != true {
			t.Errorf("FindFirst() got1 = %v, want %v", got1, true)
		}
	})

	t.Run("should return zero value if none of the items match predicate", func(t *testing.T) {
		want := ""
		got, got1 := FindFirst([]string{"a", "b", "c"}, func(s string) bool {
			return s == "f"
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindFirst() got = %v, want %v", got, want)
		}
		if got1 != false {
			t.Errorf("FindFirst() got1 = %v, want %v", got1, false)
		}
	})
}

func TestParallelize(t *testing.T) {
	want := []int{2, 4, 6, 8}
	got := MapParallel[[]int, []int](-1, []int{1, 2, 3, 4}, func(x int) int { return x * 2 })
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapParallel() = %v, want %v", got, want)
	}
}

func TestParallelizeWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	want := []int{2, 4, 6, 8}
	got := MapParallelWithContext[[]int, []int](ctx, 2, []int{1, 2, 3, 4}, func(x int) int { return 2 * x })
	sort.Ints(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapParallelWithContext() = %v, want %v", got, want)
	}
}

func TestMax(t *testing.T) {
	t.Run("should return max item from the slice", func(t *testing.T) {
		input := []int{2, 4, 6, 8}
		got, ok := Max(input)
		if !ok {
			t.Errorf("Max() = %v want %v", ok, false)
		}

		want := 8
		if got != want {
			t.Errorf("Max() = %v want %v", got, want)
		}
	})
	t.Run("should return zero value and false from the slice if slice is empty", func(t *testing.T) {
		var input []int
		got, ok := Max(input)
		if ok {
			t.Errorf("Max() = %v want %v", ok, true)
		}

		want := 0
		if got != want {
			t.Errorf("Max() = %v want %v", got, want)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("should return min item from the slice", func(t *testing.T) {
		input := []int{2, 4, 6, 8}
		got, ok := Min(input)
		if !ok {
			t.Errorf("Min() = %v want %v", ok, false)
		}

		want := 2
		if got != want {
			t.Errorf("Min() = %v want %v", got, want)
		}
	})
	t.Run("should return zero value and false from the slice if slice is empty", func(t *testing.T) {
		var input []int
		got, ok := Min(input)
		if ok {
			t.Errorf("Min() = %v want %v", ok, true)
		}

		want := 0
		if got != want {
			t.Errorf("Min() = %v want %v", got, want)
		}
	})
}
