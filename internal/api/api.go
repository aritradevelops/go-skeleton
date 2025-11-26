package api

import (
	"context"
	"log"
	"os"
	"os/signal"
	"skeleton-test/internal/config"
	"skeleton-test/internal/db"
	"skeleton-test/internal/http"
	"syscall"
	"time"
)

func Run() error {

	// load config
	config, err := config.Load(true)
	if err != nil {
		return err
	}

	// connect to database
	db := db.NewPostgres(config.Db.Url)
	err = db.Connect()
	if err != nil {
		return err
	}
	// TODO: connect to database

	httpServer := http.NewServer(config, db)

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatalf("failed to start the http server: %v", err)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quitCh
	log.Printf("Bye! shutting down gracefully...")
	shutdownContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownContext); err != nil {
		return err
	}
	// TODO: close the database connection
	return nil
}
