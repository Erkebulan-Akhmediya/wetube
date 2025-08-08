package users

import (
	"io"
	"net/http"
)

func updatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Update password successfully")
}
