package parser

import (
	"http-server-go/internal/http1"
	"maps"
	"testing"
)

func TestParseRequestHeaders(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   http1.RequestHeaders
		shouldFail bool
	}{
		{
			name:  "accept 2 headers",
			input: "Host: example.com\r\nContent-Length: 35\r\n",
			expected: http1.RequestHeaders{
				"Host":           "example.com",
				"Content-Length": "35",
			},
		},
		{
			name:       "reject malformed header",
			input:      "Host example.com\r\nContent-Length: 35\r\n",
			shouldFail: true,
		},
		{
			name:  "acccept localhost:8000 field value",
			input: "Host:   localhost:8000\r\n",
			expected: http1.RequestHeaders{
				"Host": "localhost:8000",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers, err := parseRequestHeaders([]byte(test.input))
			if test.shouldFail {
				if err == nil {
					t.Error("parseRequestHeaders() should fail but did not")
				}
				return
			}

			if err != nil {
				t.Error(err)
			} else if !maps.Equal(headers, test.expected) {
				t.Error("parseRequestHeaders() returned a value different from the expected value")
			}
		})
	}
}
