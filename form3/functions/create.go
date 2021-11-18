package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Model "github.com/rohan-sharma92/accountapi-client-lib/form3/models"
)

type ClientInterface interface {
	Create(account Model.AccountData) (Model.AccountData, error)
	Delete(accountID string, version int) (bool, error)
	Fetch(accountID string) (Model.AccountData, error)
}

func (f *Form3) Create(account Model.AccountData) (Model.AccountData, error) {
	accountRequest := &Model.AccountRequest{Data: account}
	payload, _ := json.Marshal(accountRequest)
	client := f.Client
	uri := fmt.Sprintf("%s/%s/%s", client.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE)
	request := Core.Post(uri)
	request.SetBody(bytes.NewBuffer(payload))
	response, err := client.Execute(request)
	if err != nil {
		return Model.AccountData{}, err
	}
	if response.IsError() {
		errorMessage := fmt.Sprintf("Invalid Request %d", response.StatusCode())
		return Model.AccountData{}, errors.New(errorMessage)
	}
	return Core.ParseSingleResponse(response)
}
