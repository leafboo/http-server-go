package parser

import (
	"bytes"
	"errors"
	"http-server-go/internal/http1"
)

// NOTE: Implement at the validation of HTTP message semantics code
// According to the RFC 9112, if the request message lacks
// a Host header field, contain more than one Host header,
// or if the Host header has invalid value, return status 400

func parseFieldLine(b []byte) ([][]byte, error) {
	fieldLine := bytes.SplitN(b, []byte(":"), 2)
	fieldName := fieldLine[0]

	if len(fieldLine) != 2 {
		return nil, errors.New("Malformed field line")
	}
	if bytes.HasSuffix(fieldName, []byte(" ")) {
		return nil, errors.New("Malformed field name")
	}

	fieldValue := bytes.TrimSpace(fieldLine[1])
	return [][]byte{fieldName, fieldValue}, nil
}

// should I return a pointer here instead ???
func parseRequestHeaders(b []byte) (http1.RequestHeaders, error) {
	requestHeaders := http1.RequestHeaders{}
	crlf := []byte{'\r', '\n'}

	currLine := 0
	for {
		iCrlf := bytes.Index(b[currLine:], crlf)
		if iCrlf < 0 {
			break
		}
		// Host: example.com\r\n
		fieldLine, err := parseFieldLine(b[currLine : currLine+iCrlf])
		if err != nil {
			return requestHeaders, err
		}

		fieldName := string(fieldLine[0])
		fieldValue := string(fieldLine[1])
		requestHeaders[fieldName] = fieldValue
		currLine += iCrlf + 2
	}

	return requestHeaders, nil
}
