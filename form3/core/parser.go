package core

import (
	"encoding/json"
	"io/ioutil"

	Model "github.com/rohan-sharma92/accountapi-client-lib/form3/models"
	log "github.com/sirupsen/logrus"
)

func ParseSingleResponse(response *Response) (Model.AccountData, error) {
	defer response.RawResponse.Body.Close()
	content, err := ioutil.ReadAll(response.RawResponse.Body)
	if err != nil {
		log.Error(err)
		return Model.AccountData{}, err
	}
	serverResponse := &Model.Response{}
	err = json.Unmarshal(content, &serverResponse)
	if err != nil {
		log.Error(err)
		return Model.AccountData{}, err
	}

	return serverResponse.Data, nil
}
