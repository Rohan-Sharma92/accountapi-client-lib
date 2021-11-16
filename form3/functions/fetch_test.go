package functions

import (
	"testing"

	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Mocks "github.com/rohan-sharma92/accountapi-client-lib/form3/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClient_Fetch_Invalid_Account(t *testing.T) {
	body := `{"data":""}`
	mock := Mocks.MockClient(404, []byte(body))
	c := &Core.Client{HttpClient: mock}
	response, err := Fetch(*c, "Rohan")
	assert.NotNil(t, err)
	assert.Equal(t, response.ID, "")
	assert.Equal(t, err.Error(), "Invalid Request 404")
}

func TestClient_Fetch_Valid_Account(t *testing.T) {
	body := `{"data":{"attributes":{"alternative_names":null,"country":"GB","name":["Rohan"]},"created_on":"2021-11-10T19:12:49.286Z","id":"37185cc2-4257-11ec-81d3-0242ac130003","modified_on":"2021-11-10T19:12:49.286Z","organisation_id":"6223ff16-4257-11ec-81d3-0242ac130003","type":"accounts","version":0},"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	mock := Mocks.MockClient(200, []byte(body))
	c := &Core.Client{HttpClient: mock}
	response, err := Fetch(*c, "Rohan")
	assert.Nil(t, err)
	assert.Equal(t, response.Attributes.Name[0], "Rohan")
}
