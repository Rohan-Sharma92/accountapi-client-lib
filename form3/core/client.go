package core

import (
	"net/http"
	"net/url"
	"strings"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
)

type Client struct {
	HostURL    string
	QueryParam url.Values
	Header     http.Header
	HttpClient *http.Client
}

func (c *Client) SetBaseURL(url string) *Client {
	c.HostURL = strings.TrimRight(url, "/")
	return c
}

func (c *Client) SetHeader(header, value string) *Client {
	c.Header.Set(header, value)
	return c
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	for h, v := range headers {
		c.Header.Set(h, v)
	}
	return c
}

func (c *Client) Execute(req *Request) (*Response, error) {
	var err error
	var rawReq *http.Request
	rawReq, err = createHTTPRequest(req, c)
	addClientHeaders(rawReq, c)
	addRequestHeaders(rawReq, req)
	if err != nil {
		return &Response{}, err
	}
	resp, err := c.HttpClient.Do(rawReq)
	response := &Response{
		RawResponse: resp,
	}
	return response, err
}

func CreateClient(baseURL string) *Client {
	c := &Client{
		QueryParam: url.Values{},
		Header:     http.Header{},
		HttpClient: &http.Client{},
		HostURL:    baseURL,
	}
	//Default headers, which can be overridden if required
	c.SetHeader("Content-Type", Common.JSON_TYPE)
	c.SetHeader("Accept", Common.JSON_TYPE)
	return c
}

func createHTTPRequest(req *Request, c *Client) (*http.Request, error) {
	if req.Body != nil {
		return http.NewRequest(req.Method, req.URL, req.Body)
	}
	return http.NewRequest(req.Method, req.URL, http.NoBody)
}

func addClientHeaders(req *http.Request, c *Client) {
	for k, v := range c.Header {
		req.Header[k] = v
	}
}

func addRequestHeaders(rawReq *http.Request, req *Request) {
	for k, v := range req.Header {
		rawReq.Header[k] = v
	}
}
