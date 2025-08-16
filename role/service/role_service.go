package service

import (
	"log"
	"wetube/database"
)

type Role struct {
	Name string
}

func GetAll() ([]*Role, error) {
	query := `SELECT name FROM role`
	rows, err := database.Db().Query(query)
	if err != nil {
		return nil, err
	}
	var roles []*Role
	for rows.Next() {
		if err = rows.Err(); err != nil {
			log.Println(err)
			continue
		}
		var role Role
		if err = rows.Scan(&role.Name); err != nil {
			log.Println(err)
			continue
		}
		roles = append(roles, &role)
	}
	return roles, nil
}

func CreateAll() error {
	query := `INSERT INTO role (name) VALUES ('admin'), ('user')`
	_, err := database.Db().Exec(query)
	return err
}
