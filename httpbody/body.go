// Package httpbody provides utilities to create/bind http.Request body.

package httpbody

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// FromJSON takes in JSON serializable input and returns either io.ReadCloser or error if the
// operation failed.
//
// Provided input, should support json.Marshal serialization.
func FromJSON(input any) (io.ReadCloser, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("marshaling input: %w", err)
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

// BindJSON binds provided io.ReadCloser body to a T type or returns an error in case operation fails.
//
// Provided generic T type, should support json.Unmarshal.
//
// Doesn't close underlying io.ReadCloser.
func BindJSON[T any](body io.ReadCloser) (zero T, err error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return zero, fmt.Errorf("reading body: %w", err)
	}
	var t T
	if err = json.Unmarshal(data, &t); err != nil {
		return zero, fmt.Errorf("unmarshaling body: %w", err)
	}
	return t, err
}
