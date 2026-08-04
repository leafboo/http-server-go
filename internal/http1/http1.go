package http1

type RequestLine struct {
	method         string
	targetResource string
	httpVersion    string
}

func (r *RequestLine) isMethodValid() bool {
	switch string(r.method) {
	case "GET", "HEAD", "POST", "PUT", "DELETE", "CONNEC", "OPTIONS", "TRACE", "PATCH":
		return true
	default:
		return false
	}
}

type HTTPMessage struct {
	requestLine HTTPRequestLine
}
