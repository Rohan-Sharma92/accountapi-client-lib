package accountapiclientlib

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	form3 "github.com/rohan-sharma92/accountapi-client-lib/form3/functions"
	"github.com/rohan-sharma92/accountapi-client-lib/form3/models"
	"github.com/stretchr/testify/assert"
)

func Test_E2E_Fetch_Fresh_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := form3.Init(apiURL)
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

	acc, err := f3.Create(acct)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, accountID)
	result, err := f3.Fetch(accountID)

	assert.Nil(t, err)
	assert.Equal(t, acc.ID, result.ID)
}

func Test_E2E_Delete_Fresh_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := form3.Init(apiURL)
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

	acc, err := f3.Create(acct)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, accountID)
	response, err := f3.Delete(acc.ID, *acc.Version)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, http.StatusOK)
}

func Test_E2E_Delete_Non_Existing_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := form3.Init(apiURL)
	accountID := uuid.New().String()
	response, err := f3.Delete(accountID, 0)
	assert.Equal(t, err.Error(), "Invalid Request 404")
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
}

func Test_E2E_Fetch_Non_Existing_User(t *testing.T) {
	apiURL := os.Getenv("API_URL")
	//apiURL := "http://localhost:8080"
	if apiURL == "" {
		t.Skip("API URL not configured")
	}
	f3 := form3.Init(apiURL)
	accountID := uuid.New().String()
	response, err := f3.Fetch(accountID)
	assert.Equal(t, err.Error(), "Invalid Request 404")
	assert.Equal(t, response.ID, "")
}
