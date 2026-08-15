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

// NOTE: figure out how to deal with list based fields(field names that allow multiple field values)

func isToken(b []byte) bool {
	for _, byte := range b {
		if byte >= 'A' && byte <= 'Z' ||
			byte >= 'a' && byte <= 'z' ||
			byte >= 0 && byte <= 9 ||
			byte == '!' || byte == '#' ||
			byte == '$' || byte == '%' ||
			byte == '&' || byte == '\'' ||
			byte == '*' || byte == '+' ||
			byte == '-' || byte == '.' ||
			byte == '^' || byte == '_' ||
			byte == '`' || byte == '|' || byte == '~' {
			continue
		}
		return false
	}
	return true
}

func parseFieldLine(b []byte) ([][]byte, error) {
	fieldLine := bytes.SplitN(b, []byte(":"), 2)
	fieldName := fieldLine[0]

	if len(fieldLine) != 2 {
		return nil, errors.New("Malformed field line")
	}
	// As per the RFC 9110, field name must only contain tokens
	if !isToken(fieldName) {
		return nil, errors.New("Malformed field name")
	}

	fieldValue := bytes.TrimSpace(fieldLine[1])
	return [][]byte{fieldName, fieldValue}, nil
}

func parseRequestHeaders(b []byte) (http1.RequestHeaders, error) {
	requestHeaders := make(http1.RequestHeaders)
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

		fieldName := string(bytes.ToLower(fieldLine[0]))
		fieldValue := string(fieldLine[1])
		requestHeaders.SetRequestHeaders(fieldName, fieldValue)
		currLine += iCrlf + 2
	}

	return requestHeaders, nil
}
