package http1

import (
	"bytes"
	"errors"
	"strconv"
)

func validateHeaderSemantics(h RequestHeaders) error {
	// should exist and should be the first line in the header section
	if _, hasHost := h["host"]; !hasHost {
		return errors.New("The header doesn't have a host field'")
	}

	return nil
}

func isMessageBodyPresent(h RequestHeaders) bool {
	// NOTE: we'll ignore the `Transfer-encoding` header for now because that header is mainly used in the
	// response message even though technically it is allowed to have it in the request message as well
	if _, hasContentLength := h["content-length"]; hasContentLength {
		return true
	}
	return false
}

func ParseRequest(b []byte) (HTTPMessage, error) {
	message := HTTPMessage{}

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
	// NOTE: before adding the headers, validate the semantics
	if validateHeaderSemantics(headers) != nil {
		return message, err
	}
	message.RequestHeaders = headers

	// Handle only request with POST methods for now
	if isMessageBodyPresent(headers) && message.RequestLine.Method == "POST" {
		sBytesToSend, _ := headers["content-length"]
		iBytesToSend, err := strconv.Atoi(sBytesToSend)
		if err != nil {
			return message, errors.New("the value of the `content-length` header is not a number")
		}

		beginningOfBody := j + 4
		if iBytesToSend > len(b[beginningOfBody:]) {
			return message, errors.New("value of the `content-length` header is greater than the length(in bytes) of the body")
		}
		message.Body = b[beginningOfBody : beginningOfBody+iBytesToSend]
	}

	return message, nil
}
