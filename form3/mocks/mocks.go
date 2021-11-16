package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func MockResponse(name string) *http.Response {
	body := []byte(name)
	raw := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}
	return raw
}

func MockRequest(body string) *http.Request {
	return &http.Request{Body: ioutil.NopCloser(bytes.NewReader([]byte(body)))}
}
