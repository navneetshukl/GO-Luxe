package luxe

import "errors"

var (
	errReadMaxSize = errors.New("request size exceeds maximum allowed")
)
