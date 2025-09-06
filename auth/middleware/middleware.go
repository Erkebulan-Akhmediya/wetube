package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"wetube/auth/service"
	userService "wetube/users/service"
)

func NewAuthMiddleware(next http.Handler) http.Handler {
	return &authMiddleware{next: next}
}

type authMiddleware struct {
	next http.Handler
}

func (a *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := service.Validate(tokenStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Access denied", http.StatusUnauthorized)
		return
	}

	user, err := userService.GetById(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Access denied", http.StatusUnauthorized)
		return
	}
	if user.DeletedAt.Valid {
		http.Error(w, "Access denied", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "user", user)
	a.next.ServeHTTP(w, r.WithContext(ctx))
}

func AuthFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := service.Validate(tokenStr)
		if err != nil {
			log.Println(err)
			http.Error(w, "Access denied", http.StatusUnauthorized)
			return
		}

		user, err := userService.GetById(userId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Access denied", http.StatusUnauthorized)
			return
		}
		if user.DeletedAt.Valid {
			http.Error(w, "Access denied", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}
