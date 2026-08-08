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

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		name		string
		input		string
		expected	VersionStatus
	} {
		{
			name: "accept HTTP/1.1 version",
			input: "HTTP/1.1",
			expected: Supported,
		},
		{
			name: "reject HTTP1.1 version",
			input: "HTTP1.1",
			expected: Invalid,
		},
		{
			name: "reject HTTP/2 version",
			input: "HTTP/2",
			expected: Unsupported,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status := validateVersion([]byte(test.input))
			if status != test.expected {
				t.Errorf("validateVersion() failed; expected %v, got %v", test.expected, status)
			}
		})
	}

}
