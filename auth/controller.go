package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
	"wetube/users"
)

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
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

	user, err := users.GetByUsername(dto.Username)
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

	if err = users.Create(dto.Username, string(pwd)); err != nil {
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
	user, err := users.GetByUsername(dto.Username)
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

	token, err := createJwt(user)
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

func getAuthDto(r *http.Request) (*authDto, int, error) {
	if r.Method != "POST" {
		return nil, http.StatusMethodNotAllowed, errors.New("method not allowed")
	}

	var dto authDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &dto, http.StatusOK, nil
}

func createJwt(user *users.User) (string, error) {
	claims := &jwtClaims{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "wetube",
			Subject:   user.Username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secretKey))
}
