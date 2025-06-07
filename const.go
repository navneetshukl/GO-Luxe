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

type HTTPStatus string

const (
	STATUSINTERNALSERVERERROR HTTPMethod = "Internal Server Error"
	STATUSOK                  HTTPMethod = "Status OK"
)

func (m HTTPStatus) ToString() string {
	return string(m)
}

type Error string

const (
	JSONEncoding Error = "json encoding error"
)

func (e Error) ToString() string {
	return string(e)
}
