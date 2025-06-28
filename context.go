package luxe

import (
	"net"
)

// LTX represents the context of HTTP request
type LTX struct {
	dataStore map[string]interface{}
	conn      net.Conn
	Request   *Request
	Response  *Response
	luxe      *Luxe
}

// Request holds the request data for HTTP
type Request struct {
	Method      string
	Query        string
	Params       map[string]string
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
		conn:      conn,
		dataStore: make(map[string]interface{}),
		Request: &Request{
			Headers: make(map[string]string),
			Params:   make(map[string]string),
		},
		Response: &Response{
			StatusCode: 200,
			Headers:    make(map[string]string),
		},
		luxe: luxe,
	}
}
