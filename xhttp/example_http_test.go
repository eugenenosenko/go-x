package xhttp

import (
	"net/http"
)

func ExampleBadRequest() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// respond with HTTP 400 BadRequest
		BadRequest(w, &errorResponse{
			Message: "id is required",
			Code:    "1234",
		})
	})
}

func ExampleCreated() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// handler logic here

		// respond with HTTP 201 Created
		Created(w, &dummyResponse{
			Value: "entity-value",
		})
	})
}

func ExampleOK() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// handler logic here

		// respond with HTTP 20 OK
		OK(w, &dummyResponse{
			Value: "success",
		})
	})
}

func ExampleInternalServerError() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// respond with HTTP 500 InternalServerError
		InternalServerError(w, &errorResponse{
			Message: "failed to connect to db",
			Code:    "1234",
		})
	})
}

func ExampleWriteResponse() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// write response to upstream
		WriteResponse(w, http.StatusConflict, &dummyResponse{
			Value: "entity 123 state conflict",
		})
	})
}

func ExampleUnprocessableEntity() {
	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		// respond with HTTP 422 UnprocessableEntity
		UnprocessableEntity(w, &errorResponse{
			Message: "name field is empty",
			Code:    "1234",
		})
	})
}
