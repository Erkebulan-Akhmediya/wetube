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
		newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			http.Error(w, "Password too long", http.StatusBadRequest)
			return
		}
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.Password = string(newPassword)
	}

	user.Username = dto.Username

	if dto.DeletedAt != "" {
		deletedAt, err := time.Parse(time.DateOnly, dto.DeletedAt)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid date for deleted_at: "+dto.DeletedAt, http.StatusBadRequest)
			return
		}
		user.DeletedAt.Valid = true
		user.DeletedAt.Time = deletedAt
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
