package functions

import (
	"errors"
	"fmt"
	"net/http"

	Common "github.com/rohan-sharma92/accountapi-client-lib/form3/common"
	Core "github.com/rohan-sharma92/accountapi-client-lib/form3/core"
)

func Delete(a Core.Client, accountID string, version int64) (http.Response, error) {
	var err error

	if accountID == "" {
		return http.Response{StatusCode: http.StatusBadRequest, Status: "Bad Request"}, nil
	}

	uri := fmt.Sprintf("%s/%s/%s/%s?version=%d", a.HostURL, Common.API_VERSION, Common.ACCOUNT_ROUTE, accountID, version)
	request := Core.Delete(uri)
	response, err := a.Execute(request)
	if err != nil {
		return http.Response{}, err
	}
	if response.IsError() {
		errorMessage := fmt.Sprintf("Invalid Request %d", response.StatusCode())
		return http.Response{StatusCode: response.StatusCode()}, errors.New(errorMessage)
	}
	return http.Response{StatusCode: http.StatusOK, Status: "Ok"}, nil
}
