package controller

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	} else if r.Method == "PUT" {
		update(w, r, user)
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

func update(w http.ResponseWriter, r *http.Request, user *service.User) {
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
