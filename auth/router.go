package auth

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/sign-in", signIn)
}
