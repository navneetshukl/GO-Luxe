package luxe

import (
	"net"
	"net/url"
	"strings"
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

// GetMethod return HTTP method
func (c *LTX) GetMethod() string {
	return c.Request.Method
}

// GetPath return request path
func (c *LTX) GetPath() string {
	return c.Request.Path
}

// GetQuery return particular query 
func (c *LTX) GetQuery(key string) string {
	return c.Request.Query.Get(key)
}

// GetAllQueries return all query parameter
func (c *LTX) GetAllQueries() url.Values {
	return c.Request.Query
}

// GetHeader return request header
func (c *LTX) GetHeader(key string) string {
	keyToLower := strings.ToLower(key)
	return c.Request.Headers[keyToLower]
}

// GetBody return body
func(c *LTX)GetBody()[]byte{
	return c.Request.Body
}
