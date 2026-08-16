package http1

import "fmt"

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

func (r RequestLine) String() string {
	return fmt.Sprintf("%v %v %v", r.Method, r.RequestTarget, r.HttpVersion)
}

type RequestHeaders map[string]string

// NOTE: this can be a value receiver and still change the data because the `RequestHeaders` type is a map and
// under the hood, maps in go are pointers that point to the actual map data structures
func (r RequestHeaders) SetRequestHeaders(fName, fVal string) {
	if _, ok := r[fName]; !ok {
		r[fName] = fVal
		return
	}
	r[fName] += fmt.Sprint(", " + fVal)
}

type HTTPMessage struct {
	RequestLine    RequestLine
	RequestHeaders RequestHeaders
	Body           []byte
}

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
	var responseByteStream []byte

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
	responseByteStream = append(responseByteStream, []byte(statusLine))

	return responseByteStream, nil
}
