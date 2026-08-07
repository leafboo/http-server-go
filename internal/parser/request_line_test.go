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
			name: "accept GET method",
			input: "GET",
			expected: true,
		},
		{
			name: "reject Get method",
			input: "Get",
			expected: false,
		},
		{
			name: "accept POST method",
			input: "POST",
			expected: true,
		},
		{
			name: "reject invalid method",
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
