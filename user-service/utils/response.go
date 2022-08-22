package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ResponseJsonBody(response *http.Response) interface{} {
	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var body interface{}
	json.Unmarshal(rawBody, &body)
	return body
}

func ResponseJson(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}