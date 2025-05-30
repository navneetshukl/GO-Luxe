package luxe

import (
	"encoding/json"
	"os"
	"strconv"
)

// SetStatusCode set response status code
func (l *LTX) SetStatusCode(code int) *LTX {
	l.Response.StatusCode = code
	return l
}

// SetHeader set the response header
func (l *LTX) SetHeader(key, val string) *LTX {
	l.Response.Headers[key] = val
	return l
}

// SendString method to send response back in string
func (l *LTX) SendString(code int, data string) error {
	l.Response.Body = []byte(data)
	l.Response.StatusCode = code
	if l.Request.Headers["Content-Type"] == "" {
		l.SetHeader("Content-Type", "text/plain")
	}
	return l.writeResponse()
}

// SendBytes method send response back in bytes
func (l *LTX) SendBytes(code int, data []byte) error {
	l.Response.Body = data
	l.Response.StatusCode = code
	if l.Request.Headers["Content-Type"] == "" {
		l.SetHeader("Content-Type", "application/octet-stream")
	}
	return l.writeResponse()
}

func (l *LTX) SendJSON(code int, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		l.sendError(500,"something went wrong")
	}

	// Set response data
	l.Response.StatusCode = code
	l.Response.Headers["Content-Type"] = "application/json; charset=utf-8"
	l.Response.Headers["Content-Length"] = strconv.Itoa(len(jsonData))
	l.Response.Body = jsonData

	// Write response
	return l.writeResponse()
}

// SendHTML parse html file
func (l *LTX) SendHTML(code int, fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		l.sendError(500, "something went wrong")
	}
	l.Response.StatusCode = code
	l.Response.Body = data
	l.Response.Headers["Content-Type"] = "text/html; charset=utf-8"
	l.Response.Headers["Content-Length"] = strconv.Itoa(len(data))
	return l.writeResponse()

}

// writeResponse method to write the response back
func (l *LTX) writeResponse() error {

	statusText := getStatusText(l.Response.StatusCode)
	statusLine := "HTTP/1.1 " + strconv.Itoa(l.Response.StatusCode) + " " + statusText + "\r\n"
	l.conn.Write([]byte(statusLine))

	// Write headers
	for key, value := range l.Response.Headers {
		headerLine := key + ": " + value + "\r\n"
		l.conn.Write([]byte(headerLine))
	}

	// End headers
	l.conn.Write([]byte("\r\n"))

	// Write body
	if len(l.Response.Body) > 0 {
		l.conn.Write(l.Response.Body)
	}

	return nil
}

// sendError handle internal error
func (l *LTX) sendError(code int, message string) error {
	l.Response.StatusCode = code
	l.Response.Headers["Content-Type"] = "text/plain; charset=utf-8"
	l.Response.Body = []byte(message)
	l.Response.Headers["Content-Length"] = strconv.Itoa(len(message))
	return l.writeResponse()
}
