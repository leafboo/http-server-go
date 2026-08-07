package parser

import (
	"testing"
)

func TestIsMethodValid(t *testing.T) {
	tests := []struct {
		name		string
		input		string
		expected	bool
	} {
		{
			name: "GET request",
			input: "GET",
			expected: true,
		},
		{
			name: "Get request",
			input: "Get",
			expected: false,
		},
		{
			name: "POST request",
			input: "POST",
			expected: true,
		},
		{
			name: "Invalid request",
			input: "FETCH",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := isMethodValid([]byte(test.input))
			if val != test.expected {
				t.Errorf("isMethodValid() failed; expected %t, got %t", test.expected, val)
			}
		})
	}


}
