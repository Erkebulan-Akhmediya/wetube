package controller

import (
	"log"
	"net/http"
	"strconv"
	"wetube/users/service"
)

func NewRestoreHandler() http.Handler {
	return &restoreHandler{}
}

type restoreHandler struct{}

func (rh *restoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	user, err := service.GetById(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.DeletedAt.Valid = false
	if err = service.Update(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
