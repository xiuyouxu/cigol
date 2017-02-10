package entity

import (
	"net/http"
)

// define Restful Handler
type RestHandler struct {
	Url     string `json:"url"`
	Handler func(http.ResponseWriter,
		*http.Request) `json:"handler"`
	Methods []string `json:"methods"`
}

func NewRestHandler(url string, handler func(http.ResponseWriter,
	*http.Request), methods ...string) RestHandler {
	return RestHandler{Url: url, Handler: handler, Methods: methods}
}
