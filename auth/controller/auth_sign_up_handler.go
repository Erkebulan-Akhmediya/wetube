package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"wetube/users/service"

	"golang.org/x/crypto/bcrypt"
)

func NewSignUpHandler() http.Handler {
	return &signUpHandler{}
}

type signUpHandler struct{}

func (sh *signUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, code, err := getAuthDto(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), code)
		return
	}

	user, err := service.GetByUsername(dto.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		http.Error(w, "Password is too long", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = service.Create(dto.Username, string(pwd), []string{"user"}); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
