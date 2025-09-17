package middleware

import (
	"log"
	"net/http"
	"strconv"
	channelService "wetube/channel/service"
	userService "wetube/users/service"
)

func NewIsOwnerMiddleware(next http.Handler) http.Handler {
	return &isOwnerMiddleware{
		next: next,
	}
}

type isOwnerMiddleware struct {
	next http.Handler
}

func (iom *isOwnerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*userService.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	channelIdStr := r.PathValue("channelId")
	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "channelId must be an integer", http.StatusBadRequest)
		return
	}

	channel, err := channelService.GetById(channelId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if channel.Author.Id != user.Id {
		http.Error(w, "You are not the owner of this channel", http.StatusForbidden)
		return
	}
	iom.next.ServeHTTP(w, r)
}
