package xtesting

import (
	"fmt"
	"reflect"
	"testing"
)

type mockT struct {
	*testing.T

	MockFatalf func(format string, args ...any)
}

func (t *mockT) Fatalf(format string, args ...any) {
	t.MockFatalf(format, args...)
}

var _ testing.TB = (*mockT)(nil)

type dummyStruct struct {
	ID string `json:"id"`
}

func TestLoadFixtureJSON(t *testing.T) {
	mock := func(f func(format string, args ...any)) *mockT { return &mockT{T: t, MockFatalf: f} }
	mockNoInteractions := mock(func(_ string, _ ...any) {
		t.Fatalf("Fatalf() not expecting function to be called")
	})

	t.Run("should return error if payload is empty", func(t *testing.T) {
		mockInteraction := mock(func(format string, args ...any) {
			want := "unmarshaling file data testdata/empty_payload.json: unexpected end of JSON input"
			if got := fmt.Sprintf(format, args...); got != want {
				t.Fatalf("Fatalf() error message = %v, want %v", got, want)
			}
		})

		out := LoadFixtureJSON[dummyStruct](mockInteraction, "testdata/empty_payload.json")
		want := ""
		if got := out.ID; got != want {
			t.Errorf("LoadFixtureJSON[dummyStruct]() = %v, want %v", got, want)
		}
	})

	t.Run("should correctly marshal payload into struct", func(t *testing.T) {
		out := LoadFixtureJSON[dummyStruct](mockNoInteractions, "testdata/payload.json")
		want := "12345"
		if got := out.ID; got != want {
			t.Errorf("LoadFixtureJSON[dummyStruct]() = %v, want %v", got, want)
		}
	})

	t.Run("should return error when trying to load non-existent data", func(t *testing.T) {
		mockInteraction := mock(func(format string, args ...any) {
			wants := []string{
				"reading file: open testdata/non_existent: no such file or directory",
				"unmarshaling file data testdata/non_existent: unexpected end of JSON input",
			}
			got := fmt.Sprintf(format, args...)
			var found bool
			for _, want := range wants {
				if got == want {
					found = true
				}
			}
			if !found {
				t.Fatalf("Fatalf() error message = %v doesn't match expected", got)
			}
		})

		got := LoadFixtureJSON[*dummyStruct](mockInteraction, "testdata/non_existent")
		if got != nil {
			t.Errorf("LoadFixtureJSON[dummyStruct]() = %v, want %v", got, nil)
		}
	})

}

func TestLoadFixture(t *testing.T) {
	t.Run("should return error when trying to load non-existent data", func(t *testing.T) {
		m := &mockT{MockFatalf: func(format string, args ...any) {
			want := "reading file: open testdata/non_existent: no such file or directory"
			if got := fmt.Sprintf(format, args...); got != want {
				t.Fatalf("Fatalf() error message = %v, want %v", got, want)
			}
		}}

		LoadFixture(m, "testdata/non_existent")
	})

	t.Run("should a correct byte slice after loading testdata", func(t *testing.T) {
		want := []byte("{\"id\":\"12345\"}\n")
		if got := LoadFixture(t, "testdata/payload.json"); !reflect.DeepEqual(got, want) {
			t.Fatalf("LoadFixture() = %v, got %v", string(got), string(want))
		}
	})
}
