package xmaps

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	want := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	if got := Merge(
		map[string]string{"a": "1"},
		map[string]string{"a": "1"},
		map[string]string{"b": "2"},
		map[string]string{"c": "3"},
	); !reflect.DeepEqual(got, want) {
		t.Errorf("Merge() = %v, want %v", got, want)
	}
}

func TestDifference(t *testing.T) {
	want := map[string]string{
		"a": "1",
		"b": "2",
	}
	if got := Difference(
		map[string]string{"a": "1"},
		map[string]string{"b": "2"},
	); !reflect.DeepEqual(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}

func TestIntersect(t *testing.T) {
	want := map[string]string{
		"b": "2",
		"c": "3",
	}
	if got := Intersect(
		map[string]string{"b": "2", "c": "3", "d": "4"},
		map[string]string{"c": "3", "b": "2", "a": "5"},
	); !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect() = %v, want %v", got, want)
	}
}

func TestEqual(t *testing.T) {
	if equal := Equal(
		map[string]string{"c": "2"},
		map[string]string{"c": "2"},
	); !equal {
		t.Errorf("Equal() = %v, want %v", equal, true)
	}
}
