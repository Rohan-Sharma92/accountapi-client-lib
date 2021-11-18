package functions

import (
	"errors"
	"fmt"
	"net/http"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
)

func (f *Form3) Delete(accountID string, version int64) (http.Response, error) {
	var err error

	err = Common.ValidateAccountId(accountID)
	if err != nil {
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}
	err = Common.ValidateVersion(int(version))
	if err != nil {
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}
	c := f.Client
	uri := fmt.Sprintf("%s/%s/%s/%s?version=%d", c.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID, version)
	request := Core.Delete(uri)
	response, err := c.Execute(request)
	if err != nil {
		return http.Response{}, err
	}
	if response.IsError() {
		errorMessage := fmt.Sprintf("Invalid Request %d", response.StatusCode())
		return http.Response{StatusCode: response.StatusCode()}, errors.New(errorMessage)
	}
	return http.Response{StatusCode: http.StatusOK, Status: "Ok"}, nil
}
