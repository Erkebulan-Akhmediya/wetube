package middleware

import (
	"context"
	"net/http"
	"wetube/users/service"
)

// NewSelfOrAdminMiddleware returns a middleware.
// The middleware works only endpoints that modify a single user
// or in other words, on user router endpoints with {userId} path value.
func NewSelfOrAdminMiddleware(next http.Handler) http.Handler {
	return &selfOrAdminMiddleware{next: next}
}

type selfOrAdminMiddleware struct {
	next http.Handler
}

func (soam *selfOrAdminMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestingUser, ok := r.Context().Value("user").(*service.User)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	urlUser, ok := r.Context().Value("urlUser").(*service.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, role := range requestingUser.Roles {
		if (role == "admin") || (role == "user" && urlUser.Id == requestingUser.Id) {
			ctx := context.WithValue(r.Context(), "urlUser", urlUser)
			soam.next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}
	w.WriteHeader(http.StatusForbidden)
}
