package utils

import "net/http"

type MethodHandler map[string]http.Handler

func (mh MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := mh[r.Method]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
