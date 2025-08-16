package service

import (
	"database/sql"
	"time"
	"wetube/database"
	"wetube/users/service"
)

type Channel struct {
	Name      string
	Author    *service.User
	CreatedAt time.Time
	DeletedAt sql.NullTime
}

func Create(name string, authorId int) error {
	query := `INSERT INTO channel (name, author_id) VALUES ($1, $2)`
	_, err := database.Db().Exec(query, name, authorId)
	return err
}
