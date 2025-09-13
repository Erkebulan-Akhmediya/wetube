package auth

import (
	"net/http"
	"wetube/auth/controller"
	"wetube/utils"
)

func RegisterRoutes() {
	signUpHandler := controller.NewSignUpHandler()
	signUpHandler = utils.MethodHandler{
		http.MethodPost: signUpHandler,
	}
	http.Handle("/sign-up", signUpHandler)

	signInHandler := controller.NewSignInHandler()
	signInHandler = utils.MethodHandler{
		http.MethodPost: signInHandler,
	}
	http.Handle("/sign-in", signInHandler)
}
