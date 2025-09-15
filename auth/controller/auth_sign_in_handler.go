package controller

import (
	"encoding/json"
	"log"
	"net/http"
	authService "wetube/auth/service"
	"wetube/users/service"

	"golang.org/x/crypto/bcrypt"
)

func NewSignInHandler() http.Handler {
	return &signInHandler{}
}

type signInHandler struct{}

func (sh *signInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, code, err := getSignInDto(r)
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

	if user.DeletedAt.Valid {
		http.Error(w, "Your account has been deleted", http.StatusBadRequest)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		log.Println(err)
		http.Error(w, "Username or password incorrect", http.StatusBadRequest)
		return
	}

	token, err := authService.Create(user.Id, user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(jwtDto{Token: token, Id: user.Id}); err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
