package main

import (
	"log"
	"net/http"
	"wetube/auth"
	"wetube/database"
	"wetube/users"
)

func main() {
	if err := database.Open(); err != nil {
		log.Fatal("Error opening db: " + err.Error())
	}
	defer func() {
		if err := database.Db().Close(); err != nil {
			log.Fatal("Error closing db: " + err.Error())
		}
	}()
	registerRoutes()
	log.Fatal(http.ListenAndServe(":2121", nil))
}

func registerRoutes() {
	auth.RegisterRoutes()
	users.RegisterRoutes()
}
