// Package main implements the application configuration and initializes the HTTP server.
package main

import (
	"log"

	"github.com/blevels/weatherAPI/config"
	"github.com/blevels/weatherAPI/infrastructure"
)

func main() {
	// Init configuration options
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run HTTP Server
	infrastructure.NewHTTPServer().Start(cfg)
}
