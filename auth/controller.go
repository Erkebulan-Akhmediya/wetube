package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"wetube/database"
)

type signUpDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto signUpDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
	if _, err := database.Db().Exec(query, dto.Username, dto.Password); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
