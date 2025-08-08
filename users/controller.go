package users

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type updatePasswordDto struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
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
	user, err := GetById(userId)
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
	if err = Update(user); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, "Update password successfully"); err != nil {
		log.Println(err)
	}
}
