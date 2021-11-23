package core

import (
	"net/http"
)

type Response struct {
	RawResponse *http.Response
}

func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode() > 199 && r.StatusCode() <= 300
}

func (r *Response) IsError() bool {
	return r.StatusCode() > 399
}
