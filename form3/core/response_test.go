package core

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Response_Success(t *testing.T) {
	res := Response{RawResponse: &http.Response{StatusCode: 200}}
	assert.True(t, res.IsSuccess())
	assert.Equal(t, res.StatusCode(), 200)
}

func Test_Response_Error(t *testing.T) {
	res := Response{RawResponse: &http.Response{StatusCode: 400}}
	assert.True(t, res.IsError())
	assert.Equal(t, res.StatusCode(), 400)
}
