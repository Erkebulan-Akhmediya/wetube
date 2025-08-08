package users

import (
	"net/http"
	"wetube/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/users/update/password", middleware.Auth(updatePassword))
}
