package controller

import (
	"errors"
	"log"
	"net/http"
	"time"
	"wetube/users/service"
)

func NewDeleteByIdHandler() http.Handler {
	return &deleteByIdHandler{}
}

type deleteByIdHandler struct{}

func (dh *deleteByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	deleteById(w, user)
}

func deleteById(w http.ResponseWriter, user *service.User) {
	user.DeletedAt.Time = time.Now()
	user.DeletedAt.Valid = true
	if err := service.Update(user); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
