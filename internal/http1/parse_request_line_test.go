package http1

import (
	"testing"
)

func TestIsMethodValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "accept GET method",
			input:    "GET",
			expected: true,
		},
		{
			name:     "reject Get method",
			input:    "Get",
			expected: false,
		},
		{
			name:     "accept POST method",
			input:    "POST",
			expected: true,
		},
		{
			name:     "reject invalid method",
			input:    "FETCH",
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
		name     string
		input    string
		expected VersionStatus
	}{
		{
			name:     "accept HTTP/1.1 version",
			input:    "HTTP/1.1",
			expected: Supported,
		},
		{
			name:     "reject HTTP1.1 version",
			input:    "HTTP1.1",
			expected: Invalid,
		},
		{
			name:     "reject HTTP/2 version",
			input:    "HTTP/2",
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

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   RequestLine
		shouldFail bool
	}{
		{
			name:  "accept GET /homepage.html HTTP/1.1",
			input: "GET /homepage.html HTTP/1.1",
			expected: RequestLine{
				Method:        "GET",
				RequestTarget: "/homepage.html",
				HttpVersion:   "HTTP/1.1",
			},
		},
		{
			name:       "reject GET /homepage.html HTTP/1.1 with trailing whitespace",
			input:      "GET /homepage.html HTTP/1.1 ",
			shouldFail: true,
		},
		{
			name:  "accept GET http://example.com/ HTTP/1.1",
			input: "GET http://example.com/ HTTP/1.1",
			expected: RequestLine{
				Method:        "GET",
				RequestTarget: "http://example.com/",
				HttpVersion:   "HTTP/1.1",
			},
		},
		{
			name:       "reject GET HTTP/1.1",
			input:      "GET HTTP/1.1",
			shouldFail: true,
		},
		{
			name:       "reject GET  / HTTP/1.1",
			input:      "GET  / HTTP/1.1",
			shouldFail: true,
		},
		{
			name:       "reject GET /  HTTP/1.1",
			input:      "GET /  HTTP/1.1",
			shouldFail: true,
		},
		{
			name:       "reject GET / HTTP/1.1 extra",
			input:      "GET / HTTP/1.1 extra",
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpMessage, err := parseRequestLine([]byte(test.input))
			if test.shouldFail {
				if err == nil {
					t.Error("parseRequestLine() should fail but did not.")
				}
				return
			}

			if err != nil {
				t.Error(err)
			} else if httpMessage != test.expected {
				t.Errorf("parseRequestLine() failed\nexpected: %v\ngot: %v", test.expected, httpMessage)
			}
		})
	}
}
