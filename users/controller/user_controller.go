package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wetube/users/service"
)

type userDto struct {
	Id        int    `json:"id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func newUserDto(user *service.User) *userDto {
	dto := userDto{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format(time.DateOnly),
	}
	if user.DeletedAt.Valid {
		dto.DeletedAt = user.DeletedAt.Time.Format(time.DateOnly)
	}
	return &dto
}

func Restore(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
