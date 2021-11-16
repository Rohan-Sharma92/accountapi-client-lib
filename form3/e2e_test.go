package accountapiclientlib

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	"github.com/rohan-sharma92/accountapi-client-lib/form3/functions"
	"github.com/rohan-sharma92/accountapi-client-lib/form3/models"
	"github.com/stretchr/testify/assert"
)

func Test_E2E_Fetch_Fresh_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := core.CreateClient(apiURL)
	accountID := uuid.New().String()
	acct := models.AccountData{
		ID:             accountID,
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country: "GB",
			Name:    []string{"Hello"},
		},
	}

	acc, err := functions.Create(*f3, acct)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, accountID)
	result, err := functions.Fetch(*f3, accountID)

	assert.Nil(t, err)
	assert.Equal(t, acc.ID, result.ID)
}

func Test_E2E_Delete_Fresh_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := core.CreateClient(apiURL)
	accountID := uuid.New().String()
	acct := models.AccountData{
		ID:             accountID,
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country: "GB",
			Name:    []string{"Hello"},
		},
	}

	acc, err := functions.Create(*f3, acct)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, accountID)
	response, err := functions.Delete(*f3, acc.ID, *acc.Version)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, http.StatusOK)
}
