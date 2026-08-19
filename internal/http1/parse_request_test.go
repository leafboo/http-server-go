package http1

import (
	"reflect"
	"testing"
)

func TestParseRequest(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   HTTPMessage
		shouldFail bool
	}{
		{
			name:  "accept simple GET request",
			input: "GET / HTTP/1.1\r\nHost: example.com\r\nUser-Agent: test-client/1.0\r\nAccept: text/html\r\nConnection: close\r\n\r\n",
			expected: HTTPMessage{
				RequestLine: RequestLine{
					Method:        "GET",
					RequestTarget: "/",
					HttpVersion:   "HTTP/1.1",
				},
				RequestHeaders: RequestHeaders{
					"host":       "example.com",
					"user-agent": "test-client/1.0",
					"accept":     "text/html",
					"connection": "close",
				},
				Body: nil,
			},
		},
		{
			name:  "accept simple POST request",
			input: "POST /submit HTTP/1.1\r\nHost: example.com\r\nContent-Type: application/x-www-form-urlencoded\r\nContent-Length: 11\r\n\r\nhello world",
			expected: HTTPMessage{
				RequestLine: RequestLine{
					Method:        "POST",
					RequestTarget: "/submit",
					HttpVersion:   "HTTP/1.1",
				},
				RequestHeaders: RequestHeaders{
					"host":           "example.com",
					"content-type":   "application/x-www-form-urlencoded",
					"content-length": "11",
				},
				Body: []byte("hello world"),
			},
		},
		{
			name:       "reject malformed request line",
			input:      "GET / HTTP/1.1 EXTRA\r\nHost: example.com\r\n\r\n",
			shouldFail: true,
		},
		{
			name:       "reject malformed POST request",
			input:      "POST /submit HTTP/1.1\r\nHost: example.com\r\nContent-Length: abc\r\n\r\nshould return error",
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpMessage, err := ParseRequest([]byte(test.input))
			if test.shouldFail {
				if err == nil {
					t.Error("ParseRequest() should fail but didn't")
				}
				return
			}

			if err != nil {
				t.Error(err)
			} else if !reflect.DeepEqual(httpMessage, test.expected) {
				t.Error("ParseRequest() failed; the returned value is different from the expected value")
			}
		})
	}
}
