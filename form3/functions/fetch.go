package functions

import (
	"errors"
	"fmt"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Model "github.com/rohan-sharma92/accountapi-client-lib/form3/models"
)

func Fetch(a Core.Client, accountID string) (Model.AccountData, error) {
	var err error

	if accountID == "" {
		return Model.AccountData{}, errors.New("cannot process empty accountID")
	}

	uri := fmt.Sprintf("%s/%s/%s/%s", a.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID)
	request := Core.Get(uri)
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
