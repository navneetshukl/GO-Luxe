package luxe

import (
	"strings"
)

// GetMethod return HTTP method
func (c *LTX) GetMethod() string {
	return c.Request.Method
}

// GetQuery return request path
func (c *LTX) GetQuery() string {
	return c.Request.Query
}

// GetParam return particular query
func (c *LTX) GetParam(key string) string {
	return c.Request.Params[key]
}

// GetAllParams return all query parameter
func (c *LTX) GetAllParams() map[string]string {
	return c.Request.Params
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

// SetData set the data for use between handlers
func (c *LTX) SetData(key string, val interface{}) {
	c.dataStore[key] = val
}

// Get retrieves data as interface{}
func (c *LTX) GetData(key string) (interface{}, bool) {
	val, exists := c.dataStore[key]
	return val, exists
}

// DeleteKey delete key from datastore
func (c *LTX) DeleteKey(key string) {
	delete(c.dataStore, key)
}
