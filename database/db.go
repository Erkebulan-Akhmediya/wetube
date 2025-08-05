package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Open() error {
	var err error
	connStr := "host=localhost port=2122 user=postgres password=123456 dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return db.Ping()
}

func Db() *sql.DB {
	return db
}
