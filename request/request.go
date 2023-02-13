package request

import (
	"strings"
	"net/http"
)

type Request struct {
	Params  map[string]string
	R 		*http.Request
}

func (r *Request) GetParam(key string) string {
	value, ok := r.Params[key]
	if !ok {
		return ""
	}
	return value
}

func (r *Request) GetMethod() string {
	// the method is automatically converted to uppercase for consistency
	return strings.ToUpper(r.R.Method)
}