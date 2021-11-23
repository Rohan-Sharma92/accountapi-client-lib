package core

import (
	"bytes"
	"net/http"
	"testing"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Mocks "github.com/rohan-sharma92/accountapi-client-lib/form3/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_ClientCreated(t *testing.T) {
	c := CreateClient("test")
	assert.Equal(t, c.Header[http.CanonicalHeaderKey("content-type")][0], Common.JSON_TYPE)
	assert.Equal(t, c.HostURL, "test")
}

func Test_SetHeaders(t *testing.T) {
	c := CreateClient("test")
	headers := make(map[string]string)
	headers["Authorization"] = "SHA-1"
	c.SetHeaders(headers)
	assert.Equal(t, c.Header[http.CanonicalHeaderKey("content-type")][0], Common.JSON_TYPE)
	assert.Equal(t, c.Header["Authorization"][0], "SHA-1")
	assert.Equal(t, c.HostURL, "test")
}

func Test_SetHeader(t *testing.T) {
	c := CreateClient("test")
	c.SetHeader("Authorization", "SHA-1")
	assert.Equal(t, c.Header[http.CanonicalHeaderKey("content-type")][0], Common.JSON_TYPE)
	assert.Equal(t, c.Header["Authorization"][0], "SHA-1")
	assert.Equal(t, c.HostURL, "test")
}

func Test_Execute(t *testing.T) {
	body := `{"data":{"attributes":{"alternative_names":null,"country":"GB","name":["Rohan"]},"created_on":"2021-11-10T19:12:49.286Z","id":"37185cc2-4257-11ec-81d3-0242ac130003","modified_on":"2021-11-10T19:12:49.286Z","organisation_id":"6223ff16-4257-11ec-81d3-0242ac130003","type":"accounts","version":0},"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	bodyByteBuffer := []byte(body)
	mock := Mocks.MockClient(200, bodyByteBuffer)
	c := &Client{HttpClient: mock}
	req := &Request{}
	req.SetBody(bytes.NewBuffer(bodyByteBuffer))
	resp, err := c.Execute(req)
	assert.True(t, resp.IsSuccess())
	assert.Equal(t, err, nil)
}

func Test_ExecuteWithRequestHeaders(t *testing.T) {
	body := `{"data":{"attributes":{"alternative_names":null,"country":"GB","name":["Rohan"]},"created_on":"2021-11-10T19:12:49.286Z","id":"37185cc2-4257-11ec-81d3-0242ac130003","modified_on":"2021-11-10T19:12:49.286Z","organisation_id":"6223ff16-4257-11ec-81d3-0242ac130003","type":"accounts","version":0},"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	bodyByteBuffer := []byte(body)
	mock := Mocks.MockClient(200, bodyByteBuffer)
	c := &Client{HttpClient: mock}
	req := &Request{Header: http.Header{}}
	req.Header.Set("Content-Type", "application/text")
	req.SetBody(bytes.NewBuffer(bodyByteBuffer))
	resp, err := c.Execute(req)
	assert.True(t, resp.IsSuccess())
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.RawResponse.Request.Header.Get("Content-Type"), "application/text")
}
