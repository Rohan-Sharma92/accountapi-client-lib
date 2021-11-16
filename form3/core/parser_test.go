package core

import (
	"testing"

	Mocks "github.com/rohan-sharma92/accountapi-client-lib/form3/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Valid_Response_Parsing(t *testing.T) {
	body := `{"data":{"attributes":{"alternative_names":null,"country":"GB","name":["Rohan"]},"created_on":"2021-11-10T19:12:49.286Z","id":"37185cc2-4257-11ec-81d3-0242ac130003","modified_on":"2021-11-10T19:12:49.286Z","organisation_id":"6223ff16-4257-11ec-81d3-0242ac130003","type":"accounts","version":0},"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	r := Mocks.MockResponse(body)
	response := &Response{RawResponse: r}
	account, err := ParseSingleResponse(response)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, account.ID, "37185cc2-4257-11ec-81d3-0242ac130003")
}

func Test_Invalid_Response_Parsing(t *testing.T) {
	body := `test`
	r := Mocks.MockResponse(body)
	response := &Response{RawResponse: r}
	account, err := ParseSingleResponse(response)
	assert.NotNil(t, err)
	assert.Equal(t, account.ID, "")
}
