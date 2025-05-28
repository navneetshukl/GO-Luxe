package luxe

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	splitPart string = "\r\n"
)

// ParseRequest parse the request string to valid request format
func (c *LTX) ParseRequest(reqData string) error {
	lines := strings.Split(reqData, splitPart)
	if len(lines) == 0 {
		return errInvalidRequest
	}
	requestLine := strings.Split(lines[0], " ")
	if len(requestLine) != 3 {
		return errInvalidRequest
	}
	c.Request.Method = requestLine[0]

	urlParts := strings.Split(requestLine[1], "?")
	c.Request.Path = urlParts[0]

	if len(urlParts) > 1 {
		queryParams, err := url.ParseQuery(urlParts[1])
		if err == nil {
			c.Request.Query = queryParams
		}
	}

	headerIdx := -1
	for i, line := range lines[1:] {
		if line == "" {
			headerIdx = i + 1
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.ToLower(strings.TrimSpace(headerParts[0]))
			val := strings.TrimSpace(headerParts[1])
			c.Request.Headers[key] = val
		}
	}

	// set content type
	c.Request.ContentType = c.Request.Headers["content-type"]

	// parse body
	if headerIdx >= 0 && headerIdx < len(lines)-1 {
		bodyLines := lines[headerIdx+1:]
		body := strings.Join(bodyLines, splitPart)

		fmt.Println("Body is inside ",body)
		c.Request.Body = []byte(body)
	}
	return nil
}
