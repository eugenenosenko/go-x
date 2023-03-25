// Package xstrings provides utility function to perform various operations on the strings.

package xstrings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Capitalize returns string with first and only first rune (rune, not byte) capitalized.
func Capitalize(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Uncapitalize un-capitalizes a string, changing the first letter to lower case.
func Uncapitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	if unicode.IsUpper(r[0]) {
		return string(unicode.ToLower(r[0])) + string(r[1:])
	}
	return s
}

// StartsWith check if a string starts with a specified prefix.
func StartsWith(s string, prefix string, ignorecase bool) bool {
	if s == "" || prefix == "" {
		return s == "" && prefix == ""
	}
	if utf8.RuneCountInString(prefix) > utf8.RuneCountInString(s) {
		return false
	}
	if ignorecase {
		return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
	}
	return strings.HasPrefix(s, prefix)
}

// EndsWith check if a string ends with a specified suffix.
func EndsWith(s string, suffix string, ignorecase bool) bool {
	if s == "" || suffix == "" {
		return s == "" && suffix == ""
	}
	if utf8.RuneCountInString(suffix) > utf8.RuneCountInString(s) {
		return false
	}
	if ignorecase {
		return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
	}
	return strings.HasSuffix(s, suffix)
}

// IsAlpha checks if the string contains only Unicode letters.
func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphanumeric checks if the string contains only Unicode letters and digits.
func IsAlphanumeric(s string) bool {
	if s == "" {
		return true
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Reverse reverses a string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Repeat repeats a string n times to form a new string.
func Repeat(s string, n int) string {
	buff := ""
	for n > 0 {
		n = n - 1
		buff += s
	}
	return buff
}

// LastChars gets the rightmost len characters of a string.
func LastChars(s string, size int) string {
	if s == "" || size < 0 {
		return ""
	}
	if len(s) <= size {
		return s
	}
	return s[len(s)-size:]
}
