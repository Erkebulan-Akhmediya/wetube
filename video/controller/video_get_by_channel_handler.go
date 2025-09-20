package controller

import (
	"net/http"
)

func NewGetByChannelHandler() http.Handler {
	return &getByChannelHandler{}
}

type getByChannelHandler struct{}

func (g *getByChannelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
