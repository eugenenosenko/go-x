package ptr

import (
	"reflect"
	"testing"
)

func TestTo(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		want := new(int)
		*want = 123

		if got := To(123); !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})
}

func TestValue(t *testing.T) {
	t.Run("should return empty string for nil string pointer", func(t *testing.T) {
		want := ""
		got := Value(new(string))

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})

	t.Run("should return empty string for nil int pointer", func(t *testing.T) {
		want := 0
		got := Value(new(int))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})

	t.Run("should return empty string for nil slice pointer", func(t *testing.T) {
		want := []string(nil)
		got := Value(new([]string))

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})

	t.Run("should return empty string for nil bool pointer", func(t *testing.T) {
		got := Value(new(bool))
		if !reflect.DeepEqual(got, false) {
			t.Errorf("Merge() = %v, want %v", got, false)
		}
	})

	t.Run("should return empty string for nil map pointer", func(t *testing.T) {
		want := map[string]string(nil)
		got := Value(new(map[string]string))

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})
}
