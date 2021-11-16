package core

import (
	"bytes"
	"net/http"
	"net/url"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
)

type Request struct {
	URL        string
	Method     string
	QueryParam url.Values
	Header     http.Header
	Body       *bytes.Buffer
}

func r() *Request {
	r := &Request{
		QueryParam: url.Values{},
		Header:     http.Header{},
	}
	return r
}

func (r *Request) SetQueryParam(param, value string) *Request {
	r.QueryParam.Set(param, value)
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for p, v := range params {
		r.SetQueryParam(p, v)
	}
	return r
}

func (r *Request) SetBody(body *bytes.Buffer) *Request {
	r.Body = body
	return r
}

func Get(url string) *Request {
	return create(Common.MethodGet, url)
}

func Post(url string) *Request {
	return create(Common.MethodPost, url)
}

func Put(url string) *Request {
	return create(Common.MethodPut, url)
}

func Delete(url string) *Request {
	return create(Common.MethodDelete, url)
}

func create(method, url string) *Request {
	request := r()
	request.Method = method
	request.URL = url
	return request
}
