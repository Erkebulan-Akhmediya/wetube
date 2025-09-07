package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
	"wetube/users/service"

	"golang.org/x/crypto/bcrypt"
)

func NewUpdateByIdHandler() http.Handler {
	return &updateByIdHandler{}
}

type updateByIdHandler struct{}

func (uh *updateByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("urlUser").(*service.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updateById(w, r, user)
}

func updateById(w http.ResponseWriter, r *http.Request, user *service.User) {
	var dto userDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if dto.Password != "" {
		if ok := updatePassword(w, user, &dto); !ok {
			return
		}
	}

	user.Username = dto.Username

	if dto.DeletedAt != "" {
		if ok := updateDeletedAt(w, user, &dto); !ok {
			return
		}
	}

	if err := service.Update(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedDto := newUserDto(user)
	if err := json.NewEncoder(w).Encode(updatedDto); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func updatePassword(w http.ResponseWriter, user *service.User, dto *userDto) bool {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		http.Error(w, "Password too long", http.StatusBadRequest)
		return false
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	user.Password = string(newPassword)
	return true
}

func updateDeletedAt(w http.ResponseWriter, user *service.User, dto *userDto) bool {
	deletedAt, err := time.Parse(time.DateOnly, dto.DeletedAt)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid date for deleted_at: "+dto.DeletedAt, http.StatusBadRequest)
		return false
	}
	user.DeletedAt.Valid = true
	user.DeletedAt.Time = deletedAt
	return true
}
