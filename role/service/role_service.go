package service

import (
	"wetube/database"
)

type Role struct {
	Name string
}

func CreateAll() error {
	query := `INSERT INTO role (name) VALUES ('admin'), ('user')`
	_, err := database.Db().Exec(query)
	return err
}
