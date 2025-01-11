package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourusername/go-openapi-manage-api/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.NewRouter()

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
