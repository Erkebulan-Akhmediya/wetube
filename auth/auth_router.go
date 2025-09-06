package auth

import (
	"net/http"
	"wetube/auth/controller"
)

func RegisterRoutes() {
	http.Handle("/sign-up", controller.NewSignUpHandler())
	http.Handle("/sign-in", controller.NewSignInHandler())
}
