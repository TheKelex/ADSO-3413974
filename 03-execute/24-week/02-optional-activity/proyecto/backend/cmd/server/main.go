package main

import (
	"context"
	"log"
	"time"

	"workshop/internal/config"
	"workshop/internal/repository"
	transport "workshop/internal/transport/http"
)

// main bootstraps the process: it reads the configuration, opens the database
// pool and starts the HTTP server. It holds no business rule and runs no query.
func main() {
	settings, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}
	database, err := repository.Open(context.Background(), settings.DatabaseDSN, settings.DatabaseTimeout)
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
	defer func() { _ = database.Close() }()

	server := transport.NewServer(transport.Dependency{
		Database: database,
		Config:   settings,
		Now:      time.Now,
	})
	log.Printf("backend listening on port %s", settings.HTTPPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
