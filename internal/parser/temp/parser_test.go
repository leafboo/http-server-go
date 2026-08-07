package parser

import (
	"http-server-go/internal/http1"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		name		string
		input		string
		expected	http1.RequestLine
		shouldFail	bool
	}{
		{
			name:  "GET request",
			input: "GET / HTTP/1.1",
			expected: http1.RequestLine{
				Method:         "GET",
				TargetResource: "/",
				HttpVersion:    "HTTP/1.1",
			},
		},
		{
			name:  "GET request with a target resource",
			input: "GET /home/homepage.html HTTP/1.1",
			expected: http1.RequestLine{
				Method:         "GET",
				TargetResource: "/home/homepage.html",
				HttpVersion:    "HTTP/1.1",
			},
		},
		{
			name:  "HTTP request with invalid method",
			input: "GO /home/homepage.html HTTP/1.1",
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseRequestLine([]byte(test.input))
			if err != nil {
				t.Error(err)
			}
			if got != test.expected {
				t.Errorf("ParseRequestLine(input) failed; return value does not match with the expected result")
			}
			if test.shouldFail {
				if err != nil {
					t.Error(err)
				}
			}
		})
	}
}
