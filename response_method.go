package luxe

import (
	"encoding/json"
	"fmt"
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
func (l *LTX) SendString(data string) error {
	l.Response.Body = []byte(data)
	if l.Request.Headers["Content-Type"] == "" {
		l.SetHeader("Content-Type", "text/plain")
	}
	return l.writeResponse()
}

// SendBytes method send response back in bytes
func (l *LTX) SendBytes(data []byte) error {
	l.Response.Body = data
	if l.Request.Headers["Content-Type"] == "" {
		l.SetHeader("Content-Type", "application/octet-stream")
	}
	return l.writeResponse()
}

func(l *LTX)SendJSON(data interface{})error{
	jsonData,err:=json.Marshal(data)
	if err!=nil{
		return err
	}

	l.Response.Body=jsonData
	l.SetHeader("Content-Type","application/json")
	return l.writeResponse()
}

// writeResponse method to write the response back
func (l *LTX) writeResponse() error {

	statusStr := getStatusText(l.Response.StatusCode)
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", l.Response.StatusCode, statusStr)

	// add headers
	for key, value := range l.Response.Headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	// add Content-Length
	response += fmt.Sprintf("Content-Length: %d\r\n", len(l.Response.Body))
	response += "Connection: close\r\n"
	response += "\r\n"

	// write response
	_, err := l.conn.Write([]byte(response))
	if err != nil {
		return err
	}

	// write body
	if len(l.Response.Body) > 0 {
		_, err = l.conn.Write(l.Response.Body)
	}

	return err
}
