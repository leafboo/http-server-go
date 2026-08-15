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
}
