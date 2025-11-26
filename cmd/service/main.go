package main

import (
	"log"
	"os"
	"skeleton-test/internal/api"
)

func main() {
	// wire up the service
	err := api.Run()
	if err != nil {
		log.Printf("failed to start the service : %v", err)
		os.Exit(1)
	}
}
