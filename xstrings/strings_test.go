package xstrings

import (
	"testing"
)

func TestCapitalizeFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should capitalize first letter",
			args: args{
				str: "abcd",
			},
			want: "Abcd",
		},
		{
			name: "should not cause error if already capitalized",
			args: args{
				str: "Abcd",
			},
			want: "Abcd",
		},
		{
			name: "should not cause error if string is empty",
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.str); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUncapitalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly un-capitalize string",
			args: args{
				s: "ABC",
			},
			want: "aBC",
		},
		{
			name: "should not panic if string is empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "should not change anything if the string is lowercase",
			args: args{
				s: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uncapitalize(tt.args.s); got != tt.want {
				t.Errorf("Uncapitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	type args struct {
		s          string
		prefix     string
		ignorecase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if the string starts with a specific suffix",
			args: args{
				s:          "abc",
				prefix:     "a",
				ignorecase: false,
			},
			want: true,
		},
		{
			name: "should return false if the string does not start with a specific suffix",
			args: args{
				s:          "bbc",
				prefix:     "a",
				ignorecase: false,
			},
			want: false,
		},
		{
			name: "should return true if the string starts with a specific suffix ignore case",
			args: args{
				s:          "Abbc",
				prefix:     "ab",
				ignorecase: true,
			},
			want: true,
		},
		{
			name: "should return false if strings are different length",
			args: args{
				s:          "",
				prefix:     "ab",
				ignorecase: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartsWith(tt.args.s, tt.args.prefix, tt.args.ignorecase); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	type args struct {
		s          string
		suffix     string
		ignorecase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if the string ends with a specific suffix",
			args: args{
				s:          "abc",
				suffix:     "bc",
				ignorecase: false,
			},
			want: true,
		},
		{
			name: "should return false if the string does not end with a specific suffix",
			args: args{
				s:          "bbc",
				suffix:     "a",
				ignorecase: false,
			},
			want: false,
		},
		{
			name: "should return true if the string ends with a specific suffix ignore case",
			args: args{
				s:          "Abbc",
				suffix:     "BC",
				ignorecase: true,
			},
			want: true,
		},
		{
			name: "should return false if strings are different length",
			args: args{
				s:          "",
				suffix:     "ab",
				ignorecase: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWith(tt.args.s, tt.args.suffix, tt.args.ignorecase); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if the string contains only alpha characters",
			args: args{
				s: "ABCsas",
			},
			want: true,
		},
		{
			name: "should return false if the string contains non alpha characters",
			args: args{
				s: "ABCsa12s",
			},
			want: false,
		},
		{
			name: "should not panic if string is empty",
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.args.s); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphanumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if the string contains only alphanumeric characters",
			args: args{
				s: "ABCsas",
			},
			want: true,
		},
		{
			name: "should return false if the string contains non alphanumeric characters",
			args: args{
				s: "ABCsa12s//",
			},
			want: false,
		},
		{
			name: "should not panic if string is empty",
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphanumeric(tt.args.s); got != tt.want {
				t.Errorf("IsAlphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly reverse provided string",
			args: args{
				s: "abc",
			},
			want: "cba",
		},
		{
			name: "should not panic if string is empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "should reverse strings with empty values",
			args: args{
				s: "   aaa",
			},
			want: "aaa   ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly repeat string n times",
			args: args{
				s: "a",
				n: 3,
			},
			want: "aaa",
		},
		{
			name: "should correctly repeat empty string n times",
			args: args{
				s: " ",
				n: 3,
			},
			want: "   ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastChars(t *testing.T) {
	type args struct {
		s    string
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should correctly return n last chars of a string",
			args: args{
				s:    "abcdef",
				size: 3,
			},
			want: "def",
		},
		{
			name: "should correctly entire string if size is larger than string itself",
			args: args{
				s:    "abcdef",
				size: 7,
			},
			want: "abcdef",
		},
		{
			name: "should not panic if string is empty",
			args: args{
				s:    "",
				size: 7,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastChars(tt.args.s, tt.args.size); got != tt.want {
				t.Errorf("LastChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
