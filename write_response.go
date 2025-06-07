package luxe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type HTTPResponse struct {
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Headers    map[string]string `json:"headers"`
	Body       []byte            `json:"body"`
}

// NewHTTPResponse creates a new HTTP response with default values
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 200,
		Status:     "OK",
		Headers:    make(map[string]string),
		Body:       []byte{},
	}
}

// SetStatus sets the status code and status text
func (r *HTTPResponse) SetStatus(code int, status string) *HTTPResponse {
	r.StatusCode = code
	r.Status = status
	return r
}

// SetHeader sets a header value
func (r *HTTPResponse) SetHeader(key, value string) *HTTPResponse {
	r.Headers[key] = value
	return r
}

// SetBody sets the response body
func (r *HTTPResponse) SetBody(body []byte) *HTTPResponse {
	r.Body = body
	r.Headers["Content-Length"] = strconv.Itoa(len(body))
	return r
}

// SetJSONBody sets the response body as JSON
func (r *HTTPResponse) SetJSONBody(data interface{}) *HTTPResponse {
	jsonData, err := json.Marshal(data)
	if err != nil {
		r.SetStatus(500, STATUSINTERNALSERVERERROR.ToString())
		r.SetHeader("Context-Type", "text/plain")
		r.SetBody([]byte(JSONEncoding.ToString()))
		return r
	}
	r.SetHeader("Content-Type", "application/json")
	r.SetBody(jsonData)
	return r
}

// ToString converts the response to HTTP response string
func (r *HTTPResponse) ToString() string {
	var buffer bytes.Buffer

	// Status line
	buffer.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.Status))

	// Headers
	for key, value := range r.Headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	// Empty line between headers and body
	buffer.WriteString("\r\n")

	// Body
	buffer.Write(r.Body)

	return buffer.String()
}

// ToBytes converts the response to bytes
func (r *HTTPResponse) ToBytes() []byte {
    return []byte(r.ToString())
}
func (r *HTTPResponse) SetTextBody(text string) *HTTPResponse {
    return r.SetBody([]byte(text))
}
