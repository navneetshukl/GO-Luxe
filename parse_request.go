package luxe

import (
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
	c.Request.Query = urlParts[0]

	// handling 
	if len(urlParts) > 1 {
		queryParams := urlParts[1]

		allQuery := strings.Split(queryParams, "&")
		for _, val := range allQuery {
			k := strings.Split(val, "=")[0]
			v := strings.Split(val, "=")[1]

			c.Request.Params[k] = v
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
		c.Request.Body = []byte(body)
	}
	return nil
}
