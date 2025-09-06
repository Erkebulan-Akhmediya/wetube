package controller

import (
	"net/http"
	"strconv"
	"wetube/users/service"
)

func getUserFromUrl(r *http.Request) (*service.User, error) {
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}
	return service.GetById(userId)
}
