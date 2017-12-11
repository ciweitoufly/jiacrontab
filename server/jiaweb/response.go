package jiaweb

import (
	"net/http"
)

type (
	Response struct {
		http.ResponseWriter
	}
)

func (res *Response) reset(rw http.ResponseWriter) {
	res.ResponseWriter = rw
}
