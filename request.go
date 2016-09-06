package giny

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ReadStringFromBody read string from http body
func ReadStringFromBody(r *http.Request) (string, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//ReadJSONFromBody read json from http body
func ReadJSONFromBody(r *http.Request, obj interface{}) error {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	return nil
}
