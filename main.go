package main

import (
	"log"
	"net/http"
	"wetube/auth"
	"wetube/database"
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
	auth.Router()
	log.Fatal(http.ListenAndServe(":2121", nil))
}
