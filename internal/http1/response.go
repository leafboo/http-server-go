package http1

import (
	"fmt"
)

type StatusLine struct {
	HttpVersion  string
	StatusCode   int
	ReasonPhrase string
}

type ResponseHeaders map[string]string

type HTTPResponse struct {
	StatusLine      StatusLine
	ResponseHeaders ResponseHeaders
	Body            []byte
}

// NOTE: the content/body is defined by the: request method, response status, and response fields
func CreateResponse(h HTTPMessage) ([]byte, error) {
	httpResponse := HTTPResponse{}
	responseByteStream := make([]byte, 0)

	// check request method
	method := h.RequestLine.Method
	switch method {
	case "GET":
		httpResponse.StatusLine = StatusLine{
			HttpVersion:  "HTTP/1.1",
			StatusCode:   200,
			ReasonPhrase: "OK",
		}
	}

	// put the contents of the target resource to the response body/content
	httpVersion := httpResponse.StatusLine.HttpVersion
	statusCode := httpResponse.StatusLine.StatusCode
	reasonPhrase := httpResponse.StatusLine.ReasonPhrase

	statusLine := fmt.Sprintf("%s %d %s \r\n", httpVersion, statusCode, reasonPhrase)
	responseByteStream = append(responseByteStream, statusLine...)

	return responseByteStream, nil
}
