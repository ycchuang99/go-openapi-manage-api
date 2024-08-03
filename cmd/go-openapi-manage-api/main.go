package main

import (
	"fmt"
	"net/http"

	"github.com/ycchuang99/go-openapi-manage-api/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
