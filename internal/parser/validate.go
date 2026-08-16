package parser

import (
	"errors"
	"http-server-go/internal/http1"
)

func isHeaderSemanticsValid(h http1.RequestHeaders) error {
	// should exist and should be the first line in the header section
	if _, hasHost := h["host"]; !hasHost {
		return errors.New("The header doesn't have a host field'")
	}

	return nil
}
