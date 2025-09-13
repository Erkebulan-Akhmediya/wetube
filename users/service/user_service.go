package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"wetube/database"
	roleService "wetube/role/service"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
	DeletedAt sql.NullTime
	Roles     []string
}

func Create(user *User) error {
	if user.Roles == nil || len(user.Roles) == 0 {
		return fmt.Errorf("no roles were provided for user with username %s", user.Username)
	}

	tx, err := database.Db().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var userId int
	query := `INSERT INTO "user" (username, password) VALUES ($1, $2) RETURNING id`
	if err = tx.QueryRow(query, user.Username, user.Password).Scan(&userId); err != nil {
		return err
	}

	if err = roleService.AddUserRoles(tx, userId, user.Roles); err != nil {
		return err
	}

	return tx.Commit()
}

func Update(user *User) error {
	query := `UPDATE "user" SET username = $1, password = $2, deleted_at = $3 where id = $4`
	_, err := database.Db().Exec(query, user.Username, user.Password, user.DeletedAt, user.Id)
	return err
}

func CheckForDeletes(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping periodic check for deletes")
			return
		case <-ticker.C:
			log.Println("Running periodic check for deletes...")
			if err := permanentlyDelete(); err != nil {
				log.Printf("Error checking for deletes: %v\n", err)
			}
		}
	}
}

func permanentlyDelete() error {
	query := `DELETE FROM "user" WHERE deleted_at < CURRENT_DATE - INTERVAL '30 days'`
	_, err := database.Db().Exec(query)
	return err
}
