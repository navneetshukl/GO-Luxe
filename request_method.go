package luxe

import (
	"net/url"
	"strings"
)

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
func (c *LTX) GetBody() []byte {
	return c.Request.Body
}

// BodyString returns request body as string
func (c *LTX) GetBodyString() string {
	return string(c.Request.Body)
}
