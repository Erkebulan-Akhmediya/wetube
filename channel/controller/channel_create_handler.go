package controller

import (
	"encoding/json"
	"log"
	"net/http"
	channelService "wetube/channel/service"
	userService "wetube/users/service"
)

type createChannelDto struct {
	Name string `json:"name"`
}

func NewCreateHandler() http.Handler {
	return &createHandler{}
}

type createHandler struct{}

func (ch *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := r.Context().Value("user").(userService.User)

	var dto createChannelDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := channelService.Create(dto.Name, user.Id); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
