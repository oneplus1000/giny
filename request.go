package giny

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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

//ReadJSON read json from http body
func ReadJSON(ctx gin.Context, obj interface{}) error {
	return ReadJSONFromBody(ctx.Request, obj)
}

//ReadString read string fomr http body
func ReadString(ctx gin.Context) (string, error) {
	return ReadStringFromBody(ctx.Request)
}
