package main

import (
	"log"
	"strconv"

	"github.com/pusrenk/auth-service/cmd/server"
	"github.com/pusrenk/auth-service/configs"
)

func main() {
	// Load configuration
	cfg := configs.GetConfig()

	// Create server
	srv := server.NewServer(cfg)

	// Setup dependencies
	if err := srv.SetupDependencies(); err != nil {
		log.Fatalf("Failed to setup dependencies: %v", err)
	}

	// Start server
	if err := srv.Start(strconv.Itoa(cfg.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
