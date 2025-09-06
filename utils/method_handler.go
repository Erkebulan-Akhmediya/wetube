package utils

import "net/http"

type MethodHandler map[string]http.HandlerFunc

func (mh MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := mh[r.Method]; ok {
		handler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
