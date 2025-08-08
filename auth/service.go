package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

func getAuthDto(r *http.Request) (*authDto, int, error) {
	if r.Method != "POST" {
		return nil, http.StatusMethodNotAllowed, errors.New("method not allowed")
	}

	var dto authDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &dto, http.StatusOK, nil
}
