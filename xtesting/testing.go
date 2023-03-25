package xtesting

import (
	"encoding/json"
	"os"
	"testing"
)

// LoadFixture reads provided file into memory and return a slice of bytes representing
// the contents of that file.
//
// If an error occurs during the operation, it will trigger the test to fail right away.
func LoadFixture(t testing.TB, filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	return content
}

// LoadFixtureJSON reads provided file into memory and will try to json.Unmarshal the
// loaded bytes as the provided generic type [T].
//
// If an error occurs during the operation, it will trigger the test to fail right away.
func LoadFixtureJSON[T any](t testing.TB, filename string) (out T) {
	content := LoadFixture(t, filename)
	if err := json.Unmarshal(content, &out); err != nil {
		t.Fatalf("unmarshaling file data %s: %v", filename, err)
	}
	return
}
