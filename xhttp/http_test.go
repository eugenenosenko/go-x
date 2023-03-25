package xhttp

import (
	"math"
	"net/http"
	"reflect"
	"testing"
)

type writerMock struct {
	WriteMock       func(bytes []byte) (int, error)
	WriteHeaderMock func(int)
	headers         map[string][]string
}

func (w *writerMock) Header() http.Header             { return w.headers }
func (w *writerMock) Write(bytes []byte) (int, error) { return w.WriteMock(bytes) }
func (w *writerMock) WriteHeader(statusCode int)      { w.WriteHeaderMock(statusCode) }

var _ http.ResponseWriter = (*writerMock)(nil)

type dummyResponse struct {
	Value string `json:"value"`
}

type errorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func TestOK(t *testing.T) {
	w := &writerMock{
		headers: map[string][]string{},
		WriteMock: func(bytes []byte) (i int, err error) {
			want := `{"value":"test"}`
			if got := string(bytes); got != want {
				t.Errorf("OK() = body got %s, want %s", got, want)
			}
			return
		},
		WriteHeaderMock: func(code int) {
			if code != http.StatusOK {
				t.Errorf("OK() = status code %d, want %d", code, http.StatusOK)
			}
		},
	}
	OK(w, &dummyResponse{Value: "test"})

	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("UnprocessableEntity() = body got %q, want %q", got, "application/json")
	}
	if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("UnprocessableEntity() = body got %q, want %q", got, "nosniff")
	}
}

func TestUnprocessableEntity(t *testing.T) {
	t.Run("", func(t *testing.T) {
		w := &writerMock{
			headers: map[string][]string{},
			WriteMock: func(bytes []byte) (i int, e error) {
				want := `{"message":"failed to do something","code":"error-code"}`
				if got := string(bytes); got != want {
					t.Errorf("UnprocessableEntity() = body got %s, want %s", got, want)
				}
				return
			},
			WriteHeaderMock: func(code int) {
				if code != http.StatusUnprocessableEntity {
					t.Errorf("UnprocessableEntity() = status code %d, want %d", code, http.StatusUnprocessableEntity)
				}
			},
		}
		UnprocessableEntity(w, &errorResponse{
			Message: "failed to do something",
			Code:    "error-code",
		})

		if got := w.Header().Get("Content-Type"); got != "application/json" {
			t.Errorf("UnprocessableEntity() = body got %q, want %q", got, "application/json")
		}
		if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
			t.Errorf("UnprocessableEntity() = body got %q, want %q", got, "nosniff")
		}
	})
}

func TestBadRequest(t *testing.T) {
	w := &writerMock{
		headers: map[string][]string{},
		WriteMock: func(bytes []byte) (i int, e error) {
			want := `{"message":"failed to do something","code":"error-code"}`
			if got := string(bytes); got != want {
				t.Errorf("BadRequest() = body got %s, want %s", got, want)
			}
			return
		},
		WriteHeaderMock: func(code int) {
			if code != http.StatusBadRequest {
				t.Errorf("BadRequest() = status code %d, want %d", code, http.StatusBadRequest)
			}
		},
	}
	BadRequest(w, &errorResponse{
		Message: "failed to do something",
		Code:    "error-code",
	})

	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("BadRequest() = body got %q, want %q", got, "application/json")
	}
	if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("BadRequest() = body got %q, want %q", got, "nosniff")
	}
}

func TestInternalError(t *testing.T) {
	w := &writerMock{
		headers: map[string][]string{},
		WriteMock: func(bytes []byte) (i int, e error) {
			want := `{"message":"failed to do something","code":"error-code"}`
			if got := string(bytes); got != want {
				t.Errorf("InternalServerError() = body got %s, want %s", got, want)
			}
			return
		},
		WriteHeaderMock: func(code int) {
			if code != http.StatusInternalServerError {
				t.Errorf("InternalServerError() = status code %d, want %d", code, http.StatusInternalServerError)
			}
		},
	}
	InternalServerError(w, &errorResponse{
		Message: "failed to do something",
		Code:    "error-code",
	})

	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("InternalServerError() = body got %q, want %q", got, "application/json")
	}
	if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("InternalServerError() = body got %q, want %q", got, "nosniff")
	}
}

func TestResponse(t *testing.T) {
	t.Run("should correctly respond to upstream", func(t *testing.T) {
		w := &writerMock{
			headers: map[string][]string{},
			WriteMock: func(bytes []byte) (i int, e error) {
				want := `{"message":"failed to do something","code":"error-code"}`
				if got := string(bytes); got != want {
					t.Errorf("WriteResponse() = body got %s, want %s", got, want)
				}
				return
			},
			WriteHeaderMock: func(code int) {
				if code != http.StatusForbidden {
					t.Errorf("WriteResponse() = status code %d, want %d", code, http.StatusForbidden)
				}
			},
		}
		WriteResponse(w, http.StatusForbidden, &errorResponse{
			Message: "failed to do something",
			Code:    "error-code",
		})

		if got := w.Header().Get("Content-Type"); got != "application/json" {
			t.Errorf("WriteResponse() = body got %q, want %q", got, "application/json")
		}
		if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
			t.Errorf("WriteResponse() = body got %q, want %q", got, "nosniff")
		}
	})

	t.Run("should only respond with status if body is not present", func(t *testing.T) {
		status := http.StatusInternalServerError
		w := &writerMock{
			WriteMock: func(got []byte) (i int, e error) {
				if got != nil {
					t.Errorf("WriteResponse() = expected body to be empty got %v", got)
				}
				return
			},
			WriteHeaderMock: func(s int) {
				if got := s; got != status {
					t.Errorf("WriteResponse() = status got %q, want %q", got, status)
				}
			},
			headers: map[string][]string{},
		}
		WriteResponse(w, http.StatusInternalServerError, nil)
		if got := w.headers; len(got) > 0 {
			t.Errorf("WriteResponse() = expected headers to be empty got %v", got)
		}
	})

	t.Run("should return error if body marshal failed", func(t *testing.T) {
		status := http.StatusInternalServerError
		message := "failed to serialize body: json: unsupported value: NaN\n"
		w := &writerMock{
			WriteMock: func(body []byte) (i int, e error) {
				if got := string(body); got != message {
					t.Errorf("WriteResponse() = body got %q, want %q", got, message)
				}
				return
			},
			WriteHeaderMock: func(s int) {
				if got := s; got != status {
					t.Errorf("WriteResponse() = status got %q, want %q", got, status)
				}
			},
			headers: map[string][]string{},
		}
		WriteResponse(w, status, math.NaN())
		headers := map[string][]string{
			"Content-Type":           {"text/plain; charset=utf-8"},
			"X-Content-Type-Options": {"nosniff"},
		}
		if got := w.headers; !reflect.DeepEqual(got, headers) {
			t.Errorf("WriteResponse() = headers got %v, want %v", got, headers)
		}
	})

}

func TestCreated(t *testing.T) {
	w := &writerMock{
		headers: map[string][]string{},
		WriteMock: func(bytes []byte) (i int, e error) {
			want := `{"value":"test"}`
			if got := string(bytes); got != want {
				t.Errorf("Created() = body got %s, want %s", got, want)
			}
			return
		},
		WriteHeaderMock: func(code int) {
			if code != http.StatusCreated {
				t.Errorf("Created() = status code %d, want %d", code, http.StatusCreated)
			}
		},
	}
	Created(w, &dummyResponse{
		Value: "test",
	})

	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Created() = body got %q, want %q", got, "application/json")
	}
	if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("Created() = body got %q, want %q", got, "nosniff")
	}
}
