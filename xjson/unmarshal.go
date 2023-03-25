package xjson

import (
	"encoding/json"
)

// Unmarshal parses the JSON-encoded data and returns the result [T any] or error that
// occurred during json.Unmarshal operation.
func Unmarshal[T any](data []byte) (t T, err error) {
	err = json.Unmarshal(data, &t)
	return
}
