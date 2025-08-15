package controller

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"wetube/users/service"
)

type updatePasswordDto struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type userDto struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto updatePasswordDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if dto.OldPassword == dto.NewPassword {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "Passwords are identical")
		return
	}

	userId := r.Context().Value("userId").(int)
	user, err := service.GetById(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.OldPassword)); err != nil {
		log.Println(err)
		http.Error(w, "Password incorrect", http.StatusBadRequest)
		return
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		http.Error(w, "Password too long", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.Password = string(newPassword)
	if err = service.Update(user); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func User(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := service.GetById(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method == "GET" {
		getById(w, user)
	} else if r.Method == "DELETE" {
		deleteUser(w, user)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getById(w http.ResponseWriter, user *service.User) {
	dto := newUserDto(user)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

func deleteUser(w http.ResponseWriter, user *service.User) {
	user.DeletedAt.Time = time.Now()
	user.DeletedAt.Valid = true
	if err := service.Update(user); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
