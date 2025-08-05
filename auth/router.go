package auth

import "net/http"

func Router() {
	http.HandleFunc("/sign-up", signUp)
}
