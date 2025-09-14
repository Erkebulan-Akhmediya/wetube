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
	dto, ok := getSignUpDto(w, r)
	if !ok {
		return
	}

	user, err := service.GetByUsername(dto.username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(dto.password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		http.Error(w, "Password is too long", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user = &service.User{
		Username: dto.username,
		Password: string(pwd),
		Roles:    []string{"user"},
	}
	if err = service.Create(user, dto.pfp, dto.pfpHeader); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
