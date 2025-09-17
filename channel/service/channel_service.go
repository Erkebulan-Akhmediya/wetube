package service

import (
	"database/sql"
	"time"
	"wetube/database"
	userService "wetube/users/service"
)

type Channel struct {
	Id        int
	Name      string
	Author    *userService.User
	CreatedAt time.Time
	DeletedAt sql.NullTime
}

func Create(name string, authorId int) error {
	query := `INSERT INTO channel (name, author_id) VALUES ($1, $2)`
	_, err := database.Db().Exec(query, name, authorId)
	return err
}

func GetById(id int) (*Channel, error) {
	var channel Channel
	var authorId int
	query := `SELECT id, name, author_id, created_at, deleted_at FROM channel WHERE id = $1`
	err := database.Db().
		QueryRow(query, id).
		Scan(&channel.Id, &channel.Name, &authorId, &channel.CreatedAt, &channel.DeletedAt)
	if err != nil {
		return nil, err
	}

	author, err := userService.GetById(authorId)
	if err != nil {
		return nil, err
	}
	channel.Author = author
	return &channel, nil
}
