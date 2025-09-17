package controller

import (
	"log"
	"net/http"
)

func getVideoDto(w http.ResponseWriter, r *http.Request) (*videoDto, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<21)
	if err := r.ParseMultipartForm(10 << 21); err != nil {
		log.Println("Error parsing multipart form:", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return nil, false
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		log.Println("Error getting file:", err)
		http.Error(w, "Unable to parse form file", http.StatusBadRequest)
		return nil, false
	}

	dto := videoDto{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		File:        f,
		FileHeader:  h,
	}
	return &dto, true
}
