package users

import (
	"time"
	"wetube/database"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func GetById(id int) (*User, error) {
	query := `select id, username, password, created_at from "user" where id = $1`
	return get(query, id)
}

func GetByUsername(username string) (*User, error) {
	query := `select id, username, password, created_at from "user" where username = $1`
	return get(query, username)
}

func get(query string, args ...any) (*User, error) {
	var user User
	row := database.Db().QueryRow(query, args...)
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(username string, password string) error {
	query := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
	_, err := database.Db().Exec(query, username, password)
	return err
}

func Update(user *User) error {
	query := `UPDATE "user" SET username = $1, password = $2 where id = $3`
	_, err := database.Db().Exec(query, user.Username, user.Password, user.Id)
	return err
}
