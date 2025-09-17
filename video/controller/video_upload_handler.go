package controller

import (
	"net/http"
)

func NewUploadHandler() http.Handler {
	return &uploadHandler{}
}

type uploadHandler struct{}

func (uh *uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
