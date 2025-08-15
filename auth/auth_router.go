package auth

import (
	"net/http"
	"wetube/auth/controller"
)

func RegisterRoutes() {
	http.HandleFunc("/sign-up", controller.SignUp)
	http.HandleFunc("/sign-in", controller.SignIn)
}
