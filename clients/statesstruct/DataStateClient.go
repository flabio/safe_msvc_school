package statesstruct

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/safe_msvc_user/clients"
)

func DataStateClient(data *http.Request, err error) (StateResponse, string) {
	var stateResponse StateResponse
	var dataMessage clients.MessageClient
	var msg string
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	if err != nil {
		msg = err.Error()
	}
	clientHttp := &http.Client{}
	resp, err := clientHttp.Do(data)
	if err != nil {
		msg = err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg = err.Error()
	}
	err = json.Unmarshal(body, &stateResponse)
	if err != nil {
		msg = err.Error()
	}
	if stateResponse.Id == 0 {
		json.Unmarshal(body, &dataMessage)
		msg = dataMessage.Message
	}
	return stateResponse, msg
}
