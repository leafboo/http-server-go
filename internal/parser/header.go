package parser

// NOTE:
// According to the RFC 9112, if the request message lacks
// a Host header field, contain more than one Host header,
// or if the Host header has invalid value, return status 400
