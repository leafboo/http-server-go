package parser

import (
	"bytes"
	"errors"
	"http-server-go/internal/http1"
)

var sampleRequest = `GET / HTTP/1.1` + "\r\n" +
	`Host: localhost:8080` + "\r\n" +
	`User-Agent: curl/8.14.1` + "\r\n" +
	"\r\n" // another <CRLF> for whitespace to indicate end of header section

func parseRequestLine(b []byte) (http1.RequestLine, error) {
	requestLine := http1.RequestLine{}

	whitespace := byte(' ')

	first := bytes.IndexByte(b, whitespace)
	if first < 0 {
		return requestLine, errors.New("whitespace not found")
	}

	second := bytes.IndexByte(b[first+1:], whitespace)
	if first < 0 {
		return requestLine, errors.New("whitespace not found")
	}

	third := bytes.IndexByte(b[first+1+second+1:], whitespace)
	if third > 0 {
		return requestLine, errors.New("HTTP Message is malformed; There is a whitespace at the end of the Request Line")
	}

	requestLine.Method = string(b[:first])
	requestLine.TargetResource = string(b[first+1 : first+1+second])
	requestLine.HttpVersion = string(b[first+1+second+1:])

	return requestLine, nil
}

func parseRequestHeaders() {
	// find the end of the header section(whitespace)
}

func parseRequestBody() {
}

func ParseRequest(b []byte) (http1.HTTPMessage, error) {
	message := http1.HTTPMessage{}
	var restOfMessage []byte

	separator := []byte{'\r', '\n'}
	i := bytes.Index(b, separator)
	if i < 0 {
		return message, errors.New("<CRLF> not found in the http message")
	}

	// store the request line and the rest of the message in separate variables
	restOfMessage = b[i+2:]
	_ = restOfMessage

	requestLine, err := parseRequestLine(b[:i])
	message.RequestLine = requestLine
	if err != nil {
		return message, err
	}
	parseRequestHeaders()
	parseRequestBody()

	return message, nil
}
