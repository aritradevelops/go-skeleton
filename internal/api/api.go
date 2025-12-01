package api

import (
	"context"
	"log"
	"os"
	"os/signal"
	"skeleton-test/internal/config"
	"skeleton-test/internal/db"
	"skeleton-test/internal/http"
	"skeleton-test/internal/services"
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

	// build services
	conn, _ := db.Conn()
	srv := services.New(conn)

	// build server
	httpServer := http.NewServer(config, db, srv)

	// run server in separate goroutine
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

	if err := db.Disconnect(); err != nil {
		return err
	}
	return nil
}
