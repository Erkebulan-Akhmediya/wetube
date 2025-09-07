package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wetube/role/service"
)

func NewUpdateByUserIdHandler() http.Handler {
	return &updateByUserIdHandler{}
}

type updateByUserIdHandler struct{}

func (u *updateByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var newRoles []string
	if err := json.NewDecoder(r.Body).Decode(&newRoles); err != nil {
		log.Println(err)
		http.Error(w, "Please provide list of roles", http.StatusBadRequest)
		return
	}

	rawUserId := r.PathValue("userId")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Please provide a valid userId", http.StatusBadRequest)
		return
	}

	if err = service.UpdateUserRoles(userId, newRoles); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
