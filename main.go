package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wetube/auth"
	"wetube/channel"
	"wetube/database"
	"wetube/users"
	"wetube/users/service"
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

	server := http.Server{Addr: ":2121"}

	cleanupCtx, cleanupCancel := context.WithCancel(context.Background())

	go service.CheckForDeletes(cleanupCtx, time.Hour*24*30)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server error: %v\n", err)
			cleanupCancel()
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalChan:
		log.Println("Shut down signal received")
	case <-cleanupCtx.Done():
		log.Println("Context done")
	}

	cleanupCancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func registerRoutes() {
	auth.RegisterRoutes()
	users.RegisterRoutes()
	channel.RegisterRoutes()
}
