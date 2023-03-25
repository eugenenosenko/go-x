package xtesting

import (
	"fmt"
	"testing"
)

func ExampleLoadFixture() {
	t := &testing.T{}

	fixture := LoadFixture(t, "testdata/payload.json") // content: {"id":"12345"}

	fmt.Println(string(fixture))

	// Output:
	// {"id":"12345"}
}

func ExampleLoadFixtureJSON() {
	t := &testing.T{}
	type dummyStruct struct {
		ID string `json:"id"`
	}

	out := LoadFixtureJSON[dummyStruct](t, "testdata/payload.json") // content: {"id":"12345"}

	fmt.Println(out.ID)

	// Output:
	// 12345
}
