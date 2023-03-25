package xjson

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	t.Run("should correctly unmarshal into struct", func(t *testing.T) {
		type result struct {
			A, B, C string
		}
		got, err := Unmarshal[*result]([]byte(`{"a": "a", "b": "b", "c": "c"}`))
		if err != nil {
			t.Errorf("Unmarshal() error = %v", err)
			return
		}
		want := &result{
			A: "a",
			B: "b",
			C: "c",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unmarshal() got = %v, want %v", got, want)
		}
	})
	t.Run("should return error if unmarshaling fails", func(t *testing.T) {
		type result struct {
			A, B, C string
		}
		_, err := Unmarshal[*result]([]byte(`{"a": ["a"], "b": ["b"], "c": ["c"]}`))
		if err == nil {
			t.Errorf("Unmarshal() expected error")

		}
		want := "json: cannot unmarshal array into Go struct field result.A of type string"
		if got := err.Error(); got != want {
			t.Errorf("Unmarshal() got = %q, want %q", got, want)
		}
		return
	})
	t.Run("should support unmarshaling nested structs", func(t *testing.T) {
		type value struct {
			V string `json:"v"`
		}
		type result struct {
			A, B, C string
			Value   *value   `json:"value"`
			Values  []*value `json:"values"`
		}
		got, err := Unmarshal[*result]([]byte(
			`{"a": "a", "b": "b", "c": "c","value":{"v":"test1"},"values":[{"v":"test2"}]}`,
		))
		if err != nil {
			t.Errorf("Unmarshal() error = %v", err)
			return
		}
		want := &result{
			A:     "a",
			B:     "b",
			C:     "c",
			Value: &value{V: "test1"},
			Values: []*value{
				{
					V: "test2",
				},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unmarshal() got = %v, want %v", got, want)
		}
	})
}
