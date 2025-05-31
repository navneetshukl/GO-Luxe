package luxe

type H map[string]interface{}

type HandlerFunc func(*LTX)

type HTTPMethod string

const (
	METHODGET    HTTPMethod = "GET"
	METHODPOST   HTTPMethod = "POST"
	METHODPUT    HTTPMethod = "PUT"
	METHODDELETE HTTPMethod = "DELETE"
	METHODPATCH  HTTPMethod = "PATCH"
)

func (m HTTPMethod) ToString() string {
	return string(m)
}
