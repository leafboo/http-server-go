package parser

import (
	"bytes"
	"errors"
	"http-server-go/internal/http1"
)

func ParseRequest(b []byte) (http1.HTTPMessage, error) {
	message := http1.HTTPMessage{}

	separator := []byte{'\r', '\n'}
	i := bytes.Index(b, separator)
	if i < 0 {
		return message, errors.New("<CRLF> not found in the http message")
	}

	requestLine, err := parseRequestLine(b[:i])
	if err != nil {
		return message, err
	}
	message.RequestLine = requestLine

	whitespace := []byte{'\r', '\n', '\r', '\n'}
	j := bytes.Index(b, whitespace)
	if j < 0 {
		return message, errors.New("whitespace at the end of the header section not found")
	}

	headers, err := parseRequestHeaders(b[i+2 : j+2])
	if err != nil {
		return message, err
	}
	message.RequestHeaders = headers

	// body, err := parseRequestBody(b[j+4:])

	return message, nil
}
