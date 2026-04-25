package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		{input: "aaф0b__________", expected: "aab__________"},
		{input: "__________0", expected: "_________"},
		{input: "%^%#$", expected: "%^%#$"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidStringExtended(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
	}{
		{
			name:        "digit at beginning of string",
			inputString: "3abc",
		},
		{
			name:        "two digits",
			inputString: "a45bc",
		},
		{
			name:        "zero at beginning of string",
			inputString: "0abc",
		},
		{
			name:        "three digits with special characters",
			inputString: "*jhjkh$$453@_",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Unpack(tc.inputString)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
