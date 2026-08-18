package http1

import (
	"maps"
	"testing"
)

func TestParseRequestHeaders(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   RequestHeaders
		shouldFail bool
	}{
		{
			name:  "accept 2 headers",
			input: "Host: example.com\r\nContent-Length: 35\r\n",
			expected: RequestHeaders{
				"host":           "example.com",
				"content-length": "35",
			},
		},
		{
			name:       "reject malformed header",
			input:      "Host example.com\r\nContent-Length: 35\r\n",
			shouldFail: true,
		},
		{
			name:  "accept localhost:8000",
			input: "Host:   localhost:8000    \r\n",
			expected: RequestHeaders{
				"host": "localhost:8000",
			},
		},
		{
			name:  "accept multi-line identical field-names",
			input: "Host: localhost:8000\r\naccept: text/html\r\nAccept: application/xhtml+xml\r\n",
			expected: RequestHeaders{
				"host":   "localhost:8000",
				"accept": "text/html, application/xhtml+xml",
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
