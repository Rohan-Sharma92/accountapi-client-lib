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

func Create(a Core.Client, account Model.AccountData) (Model.AccountData, error) {
	accountRequest := &Model.AccountRequest{Data: account}
	payload, _ := json.Marshal(accountRequest)
	uri := fmt.Sprintf("%s/%s/%s", a.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE)
	request := Core.Post(uri)
	request.SetBody(bytes.NewBuffer(payload))
	response, err := a.Execute(request)
	if err != nil {
		return Model.AccountData{}, err
	}
	if response.IsError() {
		errorMessage := fmt.Sprintf("Invalid Request %d", response.StatusCode())
		return Model.AccountData{}, errors.New(errorMessage)
	}
	return Core.ParseSingleResponse(response)
}
