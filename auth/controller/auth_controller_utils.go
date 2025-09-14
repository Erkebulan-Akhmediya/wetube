package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func getSignInDto(r *http.Request) (*signInDto, int, error) {
	var dto signInDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &dto, http.StatusOK, nil
}

func getSignUpDto(w http.ResponseWriter, r *http.Request) (*signUpDto, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<21)
	if err := r.ParseMultipartForm(10 << 21); err != nil {
		log.Println("Error parsing multipart form:", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return nil, false
	}

	f, h, err := r.FormFile("pfp")
	if err != nil {
		log.Println("Error getting pfp file:", err)
		http.Error(w, "Unable to parse form file", http.StatusBadRequest)
		return nil, false
	}

	dto := signUpDto{
		username:  r.FormValue("username"),
		password:  r.FormValue("password"),
		pfp:       f,
		pfpHeader: h,
	}
	return &dto, true
}
