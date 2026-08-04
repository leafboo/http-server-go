package parser

import (
	"bytes"
	"errors"
	"http-server-go/internal/http1"
	"log"
)

var sampleRequest = `GET / HTTP/1.1` + "\r\n" +
					`Host: localhost:8080` + "\r\n" +
					`User-Agent: curl/8.14.1` + "\r\n" +
					"\r\n" // another <CRLF> for whitespace to indicate end of header section

// should I return error here or just exit
func parseRequestLine(b []byte) (http1.RequestLine, error) {
	requestLine := http1.RequestLine{}

	whitespace := ' '

	first := bytes.IndexByte(b, whitespace)
	if first < 0 {
		return nil, errors.New("whitespace not found")
	}

	second := bytes.IndexByte(b[first+1:], whitespace)
	if first < 0 {
		return nil, errors.New("whitespace not found")
	}

	third := bytes.IndexByte(b[second+1:], whitespace)
	if third < 0 {
		return nil, errors.New("whitespace not found")
	}

	requestLine.method := string(b[:first])
	requestLine.targetResource := string(b[first+1 : second])
	requestLine.httpVersion := string(b[second+1:])

	return requestLine, nil
}

func parseRequestHeaders() {
	// find the end of the header section(whitespace)
}

func parseRequestBody() {
}

func ParseRequest(b []byte) (http1.HTTPRequest, error) {
	message := http1.HTTPMessage{}
	var restOfMessage []byte

	separator := []byte{'\r', '\n'}
	i := bytes.Index(b, separator)
	if i < 0 {
		log.Fatal("<CRLF> not found in the http message")
	}

	// store the request line and the rest of the message in separate variables
	restOfMessage := b[i+2:]

	message.requestLine, err := ParseRequestLine(b[:i])
	if err != nil {
		return nil, err
	}
	ParseRequestHeaders()
	ParseRequestBody()

	return message, nil
}
