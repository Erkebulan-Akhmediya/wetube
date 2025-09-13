package controller

import (
	"encoding/json"
	"net/http"
)

func getAuthDto(r *http.Request) (*authDto, int, error) {
	var dto authDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &dto, http.StatusOK, nil
}
