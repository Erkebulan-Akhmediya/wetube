package controller

import (
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
	user, ok := r.Context().Value("urlUser").(*service.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.DeletedAt.Time = time.Now()
	user.DeletedAt.Valid = true
	if err := service.Update(user); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
