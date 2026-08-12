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

type HTTPMessage struct {
	RequestLine RequestLine
	RequestHeaders RequestHeaders
}
