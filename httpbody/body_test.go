package httpbody

import (
	"bytes"
	"io"
	"math"
	"net/http"
	"reflect"
	"testing"
)

func Test_FromJSON(t *testing.T) {
	type args struct {
		input   any
		want    []byte
		wantErr bool
	}

	scenario := func(tt args) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			r, err := FromJSON(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Create() wanted error but error was nil")
					return
				}
				return
			}
			got, err := io.ReadAll(r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() wanted = %v, got = %v", string(got), string(tt.want))
			}

		}
	}
	t.Run("should correctly serialize value", scenario(args{
		input:   "nil",
		want:    []byte(`"nil"`),
		wantErr: false,
	}))

	t.Run("should correctly serialize struct", scenario(args{
		input: struct {
			A string `json:"a"`
			B string `json:"b,omitempty"`
		}{"abcd", ""},
		want:    []byte(`{"a":"abcd"}`),
		wantErr: false,
	}))

	t.Run("should return error if serialization fails", scenario(args{
		input:   math.NaN(),
		wantErr: true,
	}))
}

func TestBindJSON(t *testing.T) {
	type payload struct {
		Value string `json:"value"`
	}

	type args struct {
		body io.ReadCloser
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[payload]{
		{
			name: "should return error if the reader is empty",
			args: args{
				body: io.NopCloser(&bytes.Reader{}),
			},
			want:    payload{},
			wantErr: true,
		},
		{
			name: "should correctly bind body to the provided struct",
			args: args{
				body: io.NopCloser(bytes.NewReader([]byte(`{"value":"1234"}`))),
			},
			want:    payload{Value: "1234"},
			wantErr: false,
		},
		{
			name: "should not error if the payload doesn't match expected structure",
			args: args{
				body: io.NopCloser(bytes.NewReader([]byte(`{"abc":"1234"}`))),
			},
			want:    payload{},
			wantErr: false,
		},
		{
			name: "should return error if the http.NoBody is returned",
			args: args{
				body: http.NoBody,
			},
			want:    payload{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotZero, err := BindJSON[payload](tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotZero, tt.want) {
				t.Errorf("BindJSON() gotZero = %v, want %v", gotZero, tt.want)
			}
		})
	}
}
