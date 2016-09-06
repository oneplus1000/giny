package giny

import "net/http"

//view model
type ViewModel struct {
	R *http.Request
	D interface{}
}
