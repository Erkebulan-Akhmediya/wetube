package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"wetube/users/service"
)

func NewGetByIdHandler() http.Handler {
	return &getByIdHandler{}
}

type getByIdHandler struct{}

func (gh *getByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("urlUser").(*service.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dto := newUserDto(user)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
