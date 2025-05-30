package luxe

import "errors"

var (
	errReadMaxSize error = errors.New("request size exceeds maximum allowed")
	errInvalidRequest error=errors.New("invalid request")
	errMarshallingJson error=errors.New("error converting to json")
)
