package luxe

import (
	"net"
	"net/url"
)

// LTX represents the context of HTTP request
type LTX struct {
	conn     net.Conn
	Request  *Request
	Response *Response
	luxe     *Luxe
}

// Request holds the request data for HTTP
type Request struct {
	Method      string
	Path        string
	Query       url.Values
	Headers     map[string]string
	Body        []byte
	ContentType string
}

// Response represents the response of HTTP
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

// NewLTX creates new context
func NewLTX(conn net.Conn, luxe *Luxe) *LTX {
	return &LTX{
		conn: conn,
		Request: &Request{
			Headers: make(map[string]string),
			Query:   make(url.Values),
		},
		Response: &Response{
			StatusCode: 200,
			Headers:    make(map[string]string),
		},
		luxe: luxe,
	}
}
