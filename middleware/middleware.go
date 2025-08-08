package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"wetube/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := jwt.Validate(tokenStr)
		if err != nil {
			log.Println(err)
			http.Error(w, "Access denied", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		next(w, r.WithContext(ctx))
	}
}
