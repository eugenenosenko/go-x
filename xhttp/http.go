// Contains helper methods to facilitate writing JSON http responses to the http.ResponseWriter.

package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteResponse replies to the request with the specified JSON body and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func WriteResponse(w http.ResponseWriter, code int, body any) {
	if body == nil {
		w.WriteHeader(code)
		return
	}
	b, err := json.Marshal(body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to serialize body: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(b)
}

// Created replies to the request with an HTTP 201 StatusCreated and a supplied body (if body was provided).
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func Created(w http.ResponseWriter, body any) {
	WriteResponse(w, http.StatusCreated, body)
}

// OK replies to the request with an HTTP 200 StatusOK and a supplied body (if body was provided).
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func OK(w http.ResponseWriter, body any) {
	WriteResponse(w, http.StatusOK, body)
}

// UnprocessableEntity replies to the request with an HTTP 422 StatusUnprocessableEntity and a
// supplied error body (if body was provided).
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func UnprocessableEntity(w http.ResponseWriter, body any) {
	WriteResponse(w, http.StatusUnprocessableEntity, body)
}

// BadRequest replies to the request with an HTTP 400 StatusBadRequest and a supplied body (if body was provided).
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func BadRequest(w http.ResponseWriter, body any) {
	WriteResponse(w, http.StatusBadRequest, body)
}

// InternalServerError replies to the request with an HTTP 500 StatusInternalServerError and a
// supplied body (if body was provided).
//
// if body is not nil, it should be a value that can be serialized using json.Marshal.
func InternalServerError(w http.ResponseWriter, body any) {
	WriteResponse(w, http.StatusInternalServerError, body)
}
