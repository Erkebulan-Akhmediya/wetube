package users

import (
	"wetube/database"
)

type User struct {
	Id       int
	Username string
	Password string
}

func GetByUsername(username string) (*User, error) {
	var user User
	query := `select id, username, password from "user" where username = $1`
	row := database.Db().QueryRow(query, username)
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(username string, password string) error {
	query := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
	_, err := database.Db().Exec(query, username, password)
	return err
}
