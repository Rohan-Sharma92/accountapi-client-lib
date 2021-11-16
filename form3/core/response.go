package core

import (
	"io"
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

func (r *Response) Body() io.ReadCloser {
	defer r.RawResponse.Body.Close()
	if r.RawResponse == nil {
		return nil
	}
	return r.RawResponse.Body
}
