package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"wetube/jwt"
	"wetube/users/service"

	"golang.org/x/crypto/bcrypt"
)

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtDto struct {
	Token string `json:"token"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
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

	if err = service.Create(dto.Username, string(pwd)); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	dto, code, err := getAuthDto(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), code)
		return
	}
	user, err := service.GetByUsername(dto.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Username or password incorrect", http.StatusBadRequest)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		log.Println(err)
		http.Error(w, "Username or password incorrect", http.StatusBadRequest)
		return
	}

	token, err := jwt.Create(user.Id, user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(jwtDto{Token: token}); err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
