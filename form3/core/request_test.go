package core

import (
	"testing"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	"github.com/stretchr/testify/assert"
)

func Test_Request_Method(t *testing.T) {
	r := Get("test")
	assert.Equal(t, r.Method, Common.MethodGet)
	assert.Equal(t, r.URL, "test")
}

func Test_Request_Method_WithQueryParam(t *testing.T) {
	r := Get("test")
	r.SetQueryParam("Id", "Id1")
	assert.Equal(t, r.Method, Common.MethodGet)
	assert.Equal(t, r.URL, "test")
	assert.Equal(t, r.queryParam.Get("Id"), "Id1")
}

func Test_Request_Method_WithQueryParams(t *testing.T) {
	r := Get("test")
	queryParams := make(map[string]string)
	queryParams["Id"] = "Id1"
	r.SetQueryParams(queryParams)
	assert.Equal(t, r.Method, Common.MethodGet)
	assert.Equal(t, r.URL, "test")
	assert.Equal(t, r.queryParam.Get("Id"), "Id1")
}

func Test_Request_Method_Post(t *testing.T) {
	r := Post("test")
	assert.Equal(t, r.Method, Common.MethodPost)
	assert.Equal(t, r.URL, "test")
}

func Test_Request_Method_Put(t *testing.T) {
	r := Put("test")
	assert.Equal(t, r.Method, Common.MethodPut)
	assert.Equal(t, r.URL, "test")
}

func Test_Request_Method_Delete(t *testing.T) {
	r := Delete("test")
	assert.Equal(t, r.Method, Common.MethodDelete)
	assert.Equal(t, r.URL, "test")
}
