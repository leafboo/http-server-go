package parser

import (
	"http-server-go/internal/http1"
)

func validateHeaderSemantics(h http1.RequestHeaders) (http1.RequestHeaders, error) {
	return h, nil
}
