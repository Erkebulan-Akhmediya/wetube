package users

import (
	"net/http"
	"wetube/auth/middleware"
	"wetube/users/controller"
)

func RegisterRoutes() {
	http.HandleFunc("/users/{userId}", middleware.Auth(controller.User))
	http.HandleFunc("/users/{userId}/restore", middleware.Auth(controller.Restore))
}
