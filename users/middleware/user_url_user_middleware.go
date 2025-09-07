package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"wetube/users/service"
)

func NewURLUserMiddleware(next http.Handler) http.Handler {
	return &urlUserMiddleware{next: next}
}

type urlUserMiddleware struct {
	next http.Handler
}

func (uum *urlUserMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlUser, err := getUserFromUrl(r)
	if err != nil {
		log.Println(err)
		if errors.As(err, &err) {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := context.WithValue(r.Context(), "urlUser", urlUser)
	uum.next.ServeHTTP(w, r.WithContext(ctx))
}

func getUserFromUrl(r *http.Request) (*service.User, error) {
	userIdStr := r.PathValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}
	return service.GetById(userId)
}
