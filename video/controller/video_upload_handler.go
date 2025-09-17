package controller

import (
	"log"
	"net/http"
	"wetube/video/service"
)

func NewUploadHandler() http.Handler {
	return &uploadHandler{}
}

type uploadHandler struct{}

func (uh *uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, ok := getVideoDto(w, r)
	if !ok {
		return
	}

	if err := service.Create(dto); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
