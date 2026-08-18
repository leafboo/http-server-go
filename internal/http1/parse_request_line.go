package http1

import (
	"bytes"
	"errors"
)

type VersionStatus int

const (
	Supported VersionStatus = iota
	Unsupported
	Invalid
)

var VersionStatusEnumName = map[VersionStatus]string{
	Supported:   "Supported",
	Unsupported: "Unsupported",
	Invalid:     "Invalid",
}

func (v VersionStatus) String() string {
	return VersionStatusEnumName[v]
}

func isMethodValid(b []byte) bool {
	switch string(b) {
	case "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH":
		return true
	default:
		return false
	}
}

func validateVersion(b []byte) VersionStatus {
	httpVersion := string(b)

	switch httpVersion {
	case "HTTP/1.1":
		return Supported
	case "HTTP/0.9", "HTTP/1.0", "HTTP/2", "HTTP/3":
		return Unsupported
	default:
		return Invalid
	}
}

// NOTE: According to the RFC 9112, invalid request line should return either status 400 or 301
func parseRequestLine(b []byte) (RequestLine, error) {
	requestLine := RequestLine{}

	whitespace := byte(' ')

	first := bytes.IndexByte(b, whitespace)
	if first < 0 {
		return requestLine, errors.New("whitespace not found")
	}
	method := b[:first]
	if !isMethodValid(method) {
		return requestLine, errors.New("Method invalid")
	}

	second := bytes.IndexByte(b[first+1:], whitespace)
	if second < 0 {
		return requestLine, errors.New("whitespace not found")
	}
	requestTarget := b[first+1 : first+1+second]
	if string(requestTarget) == "" {
		return requestLine, errors.New("request target empty")
	}

	third := bytes.IndexByte(b[first+1+second+1:], whitespace)
	if third > 0 {
		return requestLine, errors.New("HTTP Message is malformed; There is a whitespace at the end of the Request Line")
	}
	httpVersion := b[first+1+second+1:]
	switch validateVersion(httpVersion) {
	case Unsupported: // return status 505 (HTTP Version Not Supported)
		return requestLine, errors.New("HTTP version is not supported")
	case Invalid: // return status 400
		return requestLine, errors.New("HTTP version is invalid")
	}

	requestLine.Method = string(method)
	requestLine.RequestTarget = string(requestTarget)
	requestLine.HttpVersion = string(b[first+1+second+1:])

	return requestLine, nil
}
