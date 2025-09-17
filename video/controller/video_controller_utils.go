package controller

import (
	"log"
	"net/http"
	"strconv"
	"wetube/video/controller/dto"
)

func getVideoDto(w http.ResponseWriter, r *http.Request) (*dto.VideoDto, bool) {
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

	channelIdStr := r.PathValue("channelId")
	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		log.Println("Error converting channelId to int:", err)
		http.Error(w, "Invalid channelId", http.StatusBadRequest)
		return nil, false
	}

	videoDto := dto.VideoDto{
		ChannelId:   channelId,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		File:        f,
		FileHeader:  h,
	}
	return &videoDto, true
}
