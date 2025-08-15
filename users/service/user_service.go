package service

import (
	"database/sql"
	"time"
	"wetube/database"
)

const baseSelect = `select id, username, password, created_at, deleted_at from "user"`

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
	DeletedAt sql.NullTime
}

func GetById(id int) (*User, error) {
	query := baseSelect + " where id = $1"
	return get(query, id)
}

func GetByUsername(username string) (*User, error) {
	query := baseSelect + " where username = $1"
	return get(query, username)
}

func get(query string, args ...any) (*User, error) {
	var user User
	row := database.Db().QueryRow(query, args...)
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.DeletedAt); err != nil {
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
	query := `UPDATE "user" SET username = $1, password = $2, deleted_at = $3 where id = $4`
	_, err := database.Db().Exec(query, user.Username, user.Password, user.DeletedAt, user.Id)
	return err
}
