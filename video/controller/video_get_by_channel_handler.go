package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wetube/video/controller/dto"
	"wetube/video/service"
)

func NewGetByChannelHandler() http.Handler {
	return &getByChannelHandler{}
}

type getByChannelHandler struct{}

func (g *getByChannelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	channelIdStr := r.PathValue("channelId")
	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid channel id", http.StatusBadRequest)
		return
	}

	videos, err := service.GetByChannelId(channelId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get videos", http.StatusInternalServerError)
		return
	}

	var videoDtos []dto.GetVideoDto
	for _, video := range videos {
		videoDto := dto.GetVideoDto{
			Id:          video.Id,
			Name:        video.Name,
			Description: video.Description,
			File:        video.File,
			ChannelId:   video.ChannelId,
		}
		videoDtos = append(videoDtos, videoDto)
	}

	if err = json.NewEncoder(w).Encode(videoDtos); err != nil {
		log.Println(err)
		http.Error(w, "Failed to send videos", http.StatusInternalServerError)
		return
	}
}
