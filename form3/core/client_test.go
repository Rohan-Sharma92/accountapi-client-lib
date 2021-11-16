package core

import (
	"net/http"
	"testing"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	"github.com/stretchr/testify/assert"
)

func Test_ClientCreated(t *testing.T) {
	c := CreateClient("test")
	assert.Equal(t, c.Header[http.CanonicalHeaderKey("content-type")][0], Common.JSON_TYPE)
	assert.Equal(t, c.HostURL, "test")
}
