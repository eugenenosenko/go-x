package xjson

import (
	"fmt"
)

func ExampleUnmarshal() {
	type exampleType struct {
		Value string `json:"value"`
	}

	t, err := Unmarshal[exampleType]([]byte(`{"value":"1234"}`))
	fmt.Println(t.Value, err)

	// Output:
	// 1234 <nil>
}
