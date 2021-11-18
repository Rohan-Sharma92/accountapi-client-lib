package functions

import (
	"errors"
	"fmt"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	Model "github.com/rohan-sharma92/accountapi-client-lib/form3/models"
)

func (f *Form3) Fetch(accountID string) (Model.AccountData, error) {
	var err error
	err = Common.ValidateAccountId(accountID)
	if err != nil {
		return Model.AccountData{}, err
	}
	c := f.Client
	uri := fmt.Sprintf("%s/%s/%s/%s", c.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID)
	request := Core.Get(uri)
	response, err := c.Execute(request)
	if err != nil {
		return Model.AccountData{}, err
	}
	if response.IsError() {
		errorMessage := fmt.Sprintf("Invalid Request %d", response.StatusCode())
		return Model.AccountData{}, errors.New(errorMessage)
	}
	return Core.ParseSingleResponse(response)
}
