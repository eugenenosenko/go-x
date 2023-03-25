package httpbody

import (
	"fmt"
	"io"
	"strings"
)

func ExampleBindJSON() {
	type payload struct {
		ID string `json:"id"`
	}
	httpBody := io.NopCloser(strings.NewReader(`{"id":"1234"}`))

	p, err := BindJSON[payload](httpBody)

	fmt.Println(p.ID)
	fmt.Println(err)

	// Output:
	// 1234
	// <nil>

}

func ExampleFromJSON() {
	type payload struct {
		ID string `json:"id"`
	}

	body, err := FromJSON(&payload{ID: "1234"})

	bytes, _ := io.ReadAll(body)
	fmt.Println(string(bytes))
	fmt.Println(err)

	// Output:
	// {"id":"1234"}
	// <nil>
}
