package functions

import (
	"testing"

	"github.com/google/uuid"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Mocks "github.com/rohan-sharma92/accountapi-client-lib/form3/mocks"
	"github.com/rohan-sharma92/accountapi-client-lib/form3/models"
	"github.com/stretchr/testify/assert"
)

func TestClient_Create_Empty_Account(t *testing.T) {
	body := `{"data":""}`
	mock := Mocks.MockClient(404, []byte(body))
	c := &Core.Client{HttpClient: mock}
	f3 := &Form3{Client: *c}
	acct := models.AccountData{
		ID:             "",
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country: "GB",
			Name:    []string{"Hello"},
		},
	}

	response, err := f3.Create(acct)
	assert.NotNil(t, err)
	assert.Equal(t, response.ID, "")
	assert.Equal(t, err.Error(), "Invalid Request 404")
}

func TestClient_Create_Valid_Account(t *testing.T) {
	body := `{"data":{"attributes":{"alternative_names":null,"country":"GB","name":["Rohan"]},"created_on":"2021-11-10T19:12:49.286Z","id":"37185cc2-4257-11ec-81d3-0242ac130003","modified_on":"2021-11-10T19:12:49.286Z","organisation_id":"6223ff16-4257-11ec-81d3-0242ac130003","type":"accounts","version":0},"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	mock := Mocks.MockClient(200, []byte(body))
	c := &Core.Client{HttpClient: mock}
	acct := models.AccountData{
		ID:             "37185cc2-4257-11ec-81d3-0242ac130003",
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country: "GB",
			Name:    []string{"Rohan"},
		},
	}
	f3 := &Form3{Client: *c}
	response, err := f3.Create(acct)
	assert.Nil(t, err)
	assert.Equal(t, response.Attributes.Name[0], "Rohan")
}
