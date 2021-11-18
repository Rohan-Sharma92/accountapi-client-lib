package functions

import (
	"net/http"
	"testing"
	"time"

	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Mocks "github.com/rohan-sharma92/accountapi-client-lib/form3/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClient_Delete_Invalid_Account(t *testing.T) {
	body := `{"data":""}`
	mock := Mocks.MockClient(404, []byte(body))
	c := &Core.Client{HttpClient: mock}
	f3 := &Form3{Client: *c}
	response, err := f3.Delete("Rohan", 1)
	assert.NotNil(t, err)
	assert.Equal(t, response.StatusCode, 404)
	assert.Equal(t, err.Error(), "Invalid Request 404")
}

func TestClient_Delete_Valid_Account(t *testing.T) {
	body := "Rohan"
	mock := Mocks.MockClient(200, []byte(body))
	c := &Core.Client{HttpClient: mock}
	f3 := &Form3{Client: *c}
	response, err := f3.Delete("Rohan", 1)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestClient_Delete_Malformed_URL(t *testing.T) {
	c := &Core.Client{HttpClient: &http.Client{Timeout: time.Duration(1 * time.Second)},
		HostURL: "http://wrong-address"}
	f3 := &Form3{Client: *c}
	_, err := f3.Delete("Rohan", 1)
	assert.NotNil(t, err)
}
