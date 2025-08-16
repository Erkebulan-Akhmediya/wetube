package controller

import (
	"encoding/json"
	"net/http"
	"wetube/role/service"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	roles, err := service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var arr []string
	for _, role := range roles {
		arr = append(arr, role.Name)
	}
	if err = json.NewEncoder(w).Encode(arr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
