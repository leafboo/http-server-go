package http1

type RequestLine struct {
	Method         string
	RequestTarget  string
	HttpVersion    string
}

type HTTPMessage struct {
	RequestLine RequestLine
}
