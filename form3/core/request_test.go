package core

import (
	"testing"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	"github.com/stretchr/testify/assert"
)

func Test_Request_Method_Get(t *testing.T) {
	r := Get("test")
	assert.Equal(t, r.Method, Common.MethodGet)
	assert.Equal(t, r.URL, "test")
}

func Test_Request_Method_Post(t *testing.T) {
	r := Post("test")
	assert.Equal(t, r.Method, Common.MethodPost)
	assert.Equal(t, r.URL, "test")
}

func Test_Request_Method_Delete(t *testing.T) {
	r := Delete("test")
	assert.Equal(t, r.Method, Common.MethodDelete)
	assert.Equal(t, r.URL, "test")
}
