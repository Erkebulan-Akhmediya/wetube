package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"wetube/users/service"
)

func NewGetByIdHandler() http.Handler {
	return &getByIdHandler{}
}

type getByIdHandler struct{}

func (gh *getByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromUrl(r)
	if err != nil {
		log.Println(err)
		if errors.As(err, &err) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	getById(w, user)
}

func getById(w http.ResponseWriter, user *service.User) {
	dto := newUserDto(user)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
