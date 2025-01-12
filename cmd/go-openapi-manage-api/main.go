package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourusername/go-openapi-manage-api/internal/config"
	"github.com/yourusername/go-openapi-manage-api/internal/handlers"
	"github.com/yourusername/go-openapi-manage-api/internal/routes"
	"github.com/yourusername/go-openapi-manage-api/internal/storage/postgres"
)

func main() {
	// Initialize database connection
	db, err := config.NewDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize storage
	storage := postgres.NewPostgresStorage(db)

	// Initialize handlers with storage
	appHandlers := handlers.NewHandlers(storage)

	// Initialize router with handlers
	router := routes.NewRouter(appHandlers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
