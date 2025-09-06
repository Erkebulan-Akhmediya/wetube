package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"wetube/users/service"

	"golang.org/x/crypto/bcrypt"
)

func getUserFromUrl(r *http.Request) (*service.User, error) {
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}
	return service.GetById(userId)
}

func GetById(w http.ResponseWriter, r *http.Request) {
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
	getById(w, user)
}

func getById(w http.ResponseWriter, user *service.User) {
	dto := newUserDto(user)
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	deleteUser(w, user)
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

func Update(w http.ResponseWriter, r *http.Request) {
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
	update(w, r, user)
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
