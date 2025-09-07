package service

import (
	"database/sql"
	"fmt"
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

func DeleteAllByUserId(tx *sql.Tx, userId int) error {
	query := `DELETE FROM users_roles WHERE user_id = $1`
	_, err := tx.Exec(query, userId)
	return err
}

func AddUserRoles(tx *sql.Tx, userId int, roles []string) error {
	query := "INSERT INTO users_roles (user_id, role_name) VALUES "
	args := make([]interface{}, len(roles)+1)
	args[0] = userId
	for i, role := range roles {
		if i > 0 {
			query += ", "
		}
		args[i+1] = role
		query += fmt.Sprintf("($1, $%d)", i+2)
	}
	_, err := tx.Exec(query, args...)
	return err
}

func UpdateUserRoles(userId int, roles []string) error {
	tx, err := database.Db().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = DeleteAllByUserId(tx, userId); err != nil {
		return err
	}

	if err = AddUserRoles(tx, userId, roles); err != nil {
		return err
	}

	return tx.Commit()
}
