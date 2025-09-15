package controller

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
)

func getSignInDto(r *http.Request) (*jsonDto, int, error) {
	var dto jsonDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &dto, http.StatusOK, nil
}

func getSignUpDto(w http.ResponseWriter, r *http.Request) (*formDataDto, bool) {
	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Println(err)
		http.Error(w, "Content type unspecified", http.StatusUnsupportedMediaType)
		return nil, false
	}
	switch mediaType {
	case "multipart/form-data":
		return handleFormData(w, r)
	case "application/json":
		return handleJSON(w, r)
	default:
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
		return nil, false
	}
}

func handleJSON(w http.ResponseWriter, r *http.Request) (*formDataDto, bool) {
	var jsonDto jsonDto
	if err := json.NewDecoder(r.Body).Decode(&jsonDto); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}
	formData := formDataDto{
		username:  jsonDto.Username,
		password:  jsonDto.Password,
		pfp:       nil,
		pfpHeader: nil,
	}
	return &formData, true
}

func handleFormData(w http.ResponseWriter, r *http.Request) (*formDataDto, bool) {
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

	dto := formDataDto{
		username:  r.FormValue("username"),
		password:  r.FormValue("password"),
		pfp:       f,
		pfpHeader: h,
	}
	return &dto, true
}
